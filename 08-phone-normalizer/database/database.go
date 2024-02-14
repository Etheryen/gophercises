package database

import (
	"08-phone-normalizer/env"
	"08-phone-normalizer/phone"
	"08-phone-normalizer/utils"
	"bytes"
	"database/sql"
	"fmt"

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

	if row.Value == normalized {
		return nil
	}

	hasDuplicates, err := hasDuplicates(db, normalized, row.Id)
	if err != nil {
		return err
	}

	if hasDuplicates {
		fmt.Printf("%v normalized has duplicates\n", row)
		return removePhone(db, row.Id)
	}

	return updatePhone(db, row.Id, normalized)
}

func updatePhone(db *sql.DB, id int, normalized string) error {
	query := fmt.Sprintf("UPDATE %v SET value = $1 WHERE id = $2", tablename)

	_, err := db.Exec(query, normalized, id)
	return err
}

func removePhone(db *sql.DB, idToRemove int) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id = $1", tablename)

	_, err := db.Exec(query, idToRemove)
	return err
}

func hasDuplicates(db *sql.DB, phoneVal string, phoneId int) (bool, error) {
	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %v WHERE value = $1 AND id != $2",
		tablename,
	)

	row := db.QueryRow(
		query,
		phoneVal,
		phoneId,
	)

	var duplicates int
	err := row.Scan(&duplicates)

	return duplicates > 0, err
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
		if err = rows.Scan(&row.Id, &row.Value); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
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
