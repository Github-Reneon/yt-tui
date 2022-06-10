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

func ParseInfo(info *ytsearch.SearchResult) *[]Video {
	var ret []Video

	for _, video := range info.Videos {

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

	return &ret
}

func SearchTitle(title string) (*[]Video, error) {

	searchInfo := ytsearch.Search(title)
	result, err := searchInfo.Next()

	if err != nil {
		return nil, err
	}

	ret := ParseInfo(result)

	return ret, nil
}
