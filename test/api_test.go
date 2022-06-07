package test

import (
	"testing"

	testPackage "github.com/Github-Reneon/yt-tui/src/api_manager"
)

func TestVideoSearch(t *testing.T) {
	test, err := testPackage.SearchTitle("test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if test != nil {
		t.Logf("Pass: %v", len(test))
	} else {
		t.Errorf("Error: test is nil")
	}
}
