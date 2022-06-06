package test

import (
	"testing"

	testPackage "github.com/Github-Reneon/yt-tui/src/api_manager"
)

func TestUrlCreation(t *testing.T) {
	videos, err := testPackage.CreateUrl("test", 1, 10)

	if err != nil {
		t.Errorf(err)
	} else {
		t.Logf(videos.Title)
	}
}
