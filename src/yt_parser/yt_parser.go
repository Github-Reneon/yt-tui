package yt_parser

import (
	"fmt"

	youtube "github.com/adrg/youtube-go"
)

func SearchTitles(search string, numOfResults int) ([]string, error) {
	searchParams := &youtube.SearchParams{search, 1, numOfResults}
	videos, err := youtube.Search(searchParams)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ret []string

	for _, video := range videos {
		ret = append(ret, video.Title)
	}

	return ret, nil
	//	for _, video := range
}

func SearchLinks(search string, numOfResults int) ([]string, error) {
	return nil, nil
}

func SearchVideoData(search string, numOfResults int) ([]*youtube.Video, error) {
	return nil, nil
}
