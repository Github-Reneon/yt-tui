package test

import (
	"fmt"
	"testing"

	testPackage "github.com/Github-Reneon/yt-tui/src/api_manager"
)

func TestVideoSearch(t *testing.T) {
	test, err := testPackage.SearchTitle("testy")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if test != nil {
		fmt.Println(len(*test))
	} else {
		t.Errorf("Error: test is nil")
	}
}
