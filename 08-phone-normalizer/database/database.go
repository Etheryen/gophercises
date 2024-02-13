package database

import (
	"08-phone-normalizer/env"
	"08-phone-normalizer/phone"
	"08-phone-normalizer/utils"
	"bytes"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type PhoneNumberRow struct {
	Id    int
	Value string
}

var dbinfo env.DbInfo

const tablename = "numbers"

func Init() error {
	info, err := env.GetDbInfo()
	if err != nil {
		return err
	}

	dbinfo = info

	db, err := sql.Open("postgres", getConnStr(false))
	if err != nil {
		return err
	}

	err = reset(db)
	db.Close()
	if err != nil {
		return err
	}

	db, err = sql.Open("postgres", getConnStr(true))
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	return createPhoneNumTable(db)
}

func NormalizeAll() error {
	db, err := sql.Open("postgres", getConnStr(true))
	if err != nil {
		return err
	}
	defer db.Close()

	return normalizeAllPhones(db)
}

func normalizeAllPhones(db *sql.DB) error {
	originalRows, err := readAllRows(db)
	if err != nil {
		return err
	}

	for _, ogRow := range originalRows {
		err = normalizePhone(db, ogRow)
		if err != nil {
			return err
		}

	}

	return nil
}

func normalizePhone(db *sql.DB, row PhoneNumberRow) error {
	normalized := phone.Normalize(row.Value)

	duplicateIds, err := getDuplicateIds(db, normalized, row.Id)
	if err != nil {
		return err
	}

	if len(duplicateIds) > 0 {
		fmt.Printf(
			"id: %v val: %v normalized is same as ids: %v\n",
			row.Id,
			row.Value,
			duplicateIds,
		)
		err = removePhones(db, duplicateIds)
		if err != nil {
			return err
		}
	}

	if row.Value == normalized {
		return nil
	}

	err = updatePhone(db, row.Id, normalized)
	if err != nil {
		return err
	}

	return nil
}

func updatePhone(db *sql.DB, id int, normalized string) error {
	query := fmt.Sprintf("UPDATE %v SET value = $1 WHERE id = $2", tablename)

	_, err := db.Exec(query, normalized, id)
	return err
}

// TODO: remove only the original, not an array of ids
func removePhones(db *sql.DB, idsToRemove []int) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id = ANY($1)", tablename)

	_, err := db.Exec(query, pq.Array(idsToRemove))
	return err
}

func getDuplicateIds(db *sql.DB, phoneVal string, phoneId int) ([]int, error) {
	query := fmt.Sprintf(
		"SELECT id FROM %v WHERE value = $1 AND id != $2",
		tablename,
	)

	rows, err := db.Query(
		query,
		phoneVal,
		phoneId,
	)
	if err != nil {
		return nil, err
	}

	var ids []int

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func ReadAll() ([]PhoneNumberRow, error) {
	db, err := sql.Open("postgres", getConnStr(true))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return readAllRows(db)
}

func readAllRows(db *sql.DB) ([]PhoneNumberRow, error) {
	query := "SELECT * FROM " + tablename

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []PhoneNumberRow

	for rows.Next() {
		var row PhoneNumberRow
		err = rows.Scan(&row.Id, &row.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}

func Populate(phoneNumbers []string) error {
	db, err := sql.Open("postgres", getConnStr(true))
	if err != nil {
		return err
	}
	defer db.Close()

	return insertPhoneNumbers(db, phoneNumbers)
}

func insertPhoneNumbers(db *sql.DB, phoneNumbers []string) error {
	baseQuery := fmt.Sprintf("INSERT INTO %v (value) VALUES\n", tablename)

	queryBuf := bytes.NewBufferString(baseQuery)

	for i := range phoneNumbers {
		isLast := i == len(phoneNumbers)-1

		pnQuery := fmt.Sprintf("($%v)", i+1)
		queryBuf.WriteString(pnQuery)

		if !isLast {
			queryBuf.WriteString(",\n")
		}
	}

	queryArgs := utils.StrToAnySlice(phoneNumbers)

	_, err := db.Exec(queryBuf.String(), queryArgs...)
	return err
}

func createPhoneNumTable(db *sql.DB) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %v (
			id SERIAL PRIMARY KEY,
			value VARCHAR(255)
		)`, tablename)

	_, err := db.Exec(query)
	return err
}

func reset(db *sql.DB) error {
	query := fmt.Sprintf("DROP DATABASE IF EXISTS %v", dbinfo.Dbname)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("CREATE DATABASE %v", dbinfo.Dbname)
	_, err = db.Exec(query)
	return err
}

func getConnStr(includeDbName bool) string {
	connStr := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v sslmode=disable",
		dbinfo.Host,
		dbinfo.Port,
		dbinfo.User,
		dbinfo.Password,
	)

	if includeDbName {
		connStr = fmt.Sprintf("%v dbname=%v", connStr, dbinfo.Dbname)
	}

	return connStr
}
