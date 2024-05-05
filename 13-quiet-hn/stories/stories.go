package stories

import (
	"13-quiet-hn/utils"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const storiesNumber int = 30

const baseUrl string = "https://hacker-news.firebaseio.com/v0"

type TopStoriesResponse []int

type ItemResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

type Story struct {
	ID     int
	Title  string
	URL    string
	Domain string
}

type StoriesData struct {
	Stories []Story
	Time    time.Duration
}

func GetStoriesData() (*StoriesData, error) {
	start := time.Now()

	ids, err := getTopStoriesIds()
	if err != nil {
		return nil, err
	}

	stories, err := getStories(*ids)
	if err != nil {
		return nil, err
	}

	return &StoriesData{
		Stories: stories,
		Time:    time.Since(start).Round(10 * time.Millisecond),
	}, nil
}

func getStories(ids []int) ([]Story, error) {
	var stories []Story

	for i := 0; len(stories) < storiesNumber; i++ {
		items, err := getItems(ids, i)
		if err != nil {
			return nil, err
		}

		newStories, err := filterOutNonStories(items)
		if err != nil {
			return nil, err
		}

		stories = append(stories, newStories...)
	}

	sortedStories := sortStories(stories, ids)

	return sortedStories, nil
}

func getItems(ids []int, attempt int) ([]*ItemResponse, error) {
	var items []*ItemResponse

	itemsChan := make(chan *ItemResponse, storiesNumber)
	errChan := make(chan error)

	for i := attempt * storiesNumber; i < storiesNumber*(attempt+1); i++ {
		go getItemFromWorker(ids[i], itemsChan, errChan)
	}

	for {
		if len(items) >= storiesNumber {
			return items, nil
		}
		select {
		case err := <-errChan:
			return nil, err
		case item := <-itemsChan:
			items = append(items, item)
		}
	}
}

func getItemFromWorker(
	id int,
	itemChan chan<- *ItemResponse,
	errChan chan<- error,
) {
	item, err := getItem(id)
	if err != nil {
		errChan <- err
	}
	itemChan <- item
}

func filterOutNonStories(items []*ItemResponse) ([]Story, error) {
	var stories []Story

	for _, item := range items {
		if isItemStory(item) {
			story, err := itemToStory(item)
			if err != nil {
				return nil, err
			}
			stories = append(stories, story)
		}
	}

	return stories, nil
}

func sortStories(stories []Story, ids []int) []Story {
	var sorted []Story

	for _, id := range ids {
		if len(sorted) >= storiesNumber {
			break
		}
		story := findStoryWithId(stories, id)
		if story != nil {
			sorted = append(sorted, *story)
		}
	}

	return sorted
}

func findStoryWithId(stories []Story, id int) *Story {
	for _, story := range stories {
		if story.ID == id {
			return &story
		}
	}
	return nil
}

func isItemStory(item *ItemResponse) bool {
	return item.Type == "story" && item.URL != ""
}

func itemToStory(item *ItemResponse) (Story, error) {
	parsedUrl, err := url.Parse(item.URL)
	if err != nil {
		return Story{}, err
	}

	domain := strings.TrimPrefix(parsedUrl.Hostname(), "www.")

	return Story{
		ID:     item.ID,
		Title:  item.Title,
		URL:    item.URL,
		Domain: domain,
	}, nil
}

func getItem(id int) (*ItemResponse, error) {
	url := fmt.Sprintf("%v/item/%v.json", baseUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return utils.ParseJSON[ItemResponse](bytes)
}

func getTopStoriesIds() (*TopStoriesResponse, error) {
	resp, err := http.Get(baseUrl + "/topstories.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return utils.ParseJSON[TopStoriesResponse](bytes)
}
