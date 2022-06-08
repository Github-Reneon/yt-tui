package yt_parser

import (
	youtube "github.com/Github-Reneon/yt-tui/src/api_manager"
)

func SearchTitles(search string, numOfResults int) (*[]string, error) {
	videos, err := youtube.SearchTitle(search)

	if err != nil {
		return nil, err
	}

	var ret []string

	for i, video := range *videos {
		if i >= 10 {
			break
		}
		ret = append(ret, video.Title)
	}

	return &ret, nil
	//	for _, video := range
}

func SearchLinks(search string, numOfResults int) (*[]string, error) {
	videos, err := youtube.SearchTitle(search)

	if err != nil {
		return nil, err
	}

	var ret []string

	for i, video := range *videos {
		if i >= 10 {
			break
		}
		ret = append(ret, video.Link)
	}

	return &ret, nil
}

func SearchVideoData(search string, numOfResults int) (*[]youtube.Video, error) {
	videos, err := youtube.SearchTitle(search)

	if err != nil {
		return nil, err
	}

	var ret []youtube.Video

	for i, video := range *videos {
		if i == 10 {
			break
		}
		ret = append(ret, video)
	}

	return &ret, nil
}
