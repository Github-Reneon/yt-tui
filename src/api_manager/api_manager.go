package api_manager

import (
	"encoding/json"
	"fmt"

	ytsearch "github.com/raitonoberu/ytsearch"
)

type Video struct {
	Title string
	Id    string
	Link  string
}

func SearchTitle(title string) ([]Video, error) {

	var ret []Video

	searchInfo := ytsearch.Search("test")

	result, err := searchInfo.Next()

	if err != nil {
		return nil, err
	}

	for video := range result.Videos {

		jsontr, _ := json.Marshal(video)

		var mappedJson map[string]interface{}

		json.Unmarshal(jsontr, &mappedJson)

		videoData := Video{
			Title: fmt.Sprint(mappedJson["title"]),
			Id:    fmt.Sprint(mappedJson["id"]),
			Link:  fmt.Sprint(mappedJson["url"]),
		}
		ret = append(ret, videoData)

	}

	return ret, nil
}
