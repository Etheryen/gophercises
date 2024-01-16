package utils

import (
	"03-cyoa/types"
	"encoding/json"
)

func ParseStoryJSON(storyJson []byte) (map[string]types.StoryArc, error) {
	var storyMap map[string]types.StoryArc

	err := json.Unmarshal(storyJson, &storyMap)
	if err != nil {
		return nil, err
	}

	return storyMap, nil
}
