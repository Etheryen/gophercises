package cache

import (
	"13-quiet-hn/stories"
	"fmt"
	"time"
)

const expiresAfter = time.Minute

type Cache struct {
	data       *stories.StoriesData
	lastUpdate time.Time
}

type GetRootData struct {
	StoriesData *stories.StoriesData
	CacheTime   time.Duration
	UpdateIn    time.Duration
}

var cache Cache

func init() {
	if err := update(); err != nil {
		panic(err)
	}

	go updater()
}

func GetTemplateData() (*GetRootData, error) {
	return &GetRootData{
		StoriesData: cache.data,
		CacheTime:   time.Since(cache.lastUpdate).Round(10 * time.Millisecond),
		UpdateIn: cache.lastUpdate.Add(expiresAfter).Sub(time.Now()).
			Round(10 * time.Millisecond),
	}, nil
}

func updater() {
	ticker := time.NewTicker(expiresAfter)
	for range ticker.C {
		fmt.Println("Updated:", time.Now())
		if err := update(); err != nil {
			panic(err)
		}
	}
}

func update() error {
	newData, err := stories.GetStoriesData()
	if err != nil {
		return err
	}
	cache.data = newData
	cache.lastUpdate = time.Now()
	return nil
}
