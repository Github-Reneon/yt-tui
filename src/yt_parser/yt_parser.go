package yt_parser

import (
	youtube "github.com/adrg/youtube-go"
)

func Check() string {
	youtube.GetVideo("test")

	return ""
}
