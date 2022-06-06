package test

import (
	"net/url"
	"testing"

	testPackage "github.com/Github-Reneon/yt-tui/src/yt_parser"
)

func TestCanSearchTitles(t *testing.T) {
	ret, err := testPackage.SearchTitles("Luke Smith", 10)

	if err != nil {
		t.Errorf("SearchTitles error: %v", err)
		t.Fail()
	}

	// Length Check
	if len(ret) != 10 {
		t.Errorf("Ret pass return %v number of results", len(ret))
	} else {
		t.Logf("Ret pass return %v number of results", len(ret))
	}

	for _, title := range ret {
		if title != "" {
			t.Logf("Title of %v returned", title)
		} else {
			t.Errorf("Error title of %v returned", title)
		}
	}
}

func TestCanSearchLinks(t *testing.T) {
	ret, err := testPackage.SearchLinks("Luke Smith", 10)

	if err != nil {
		t.Errorf("SearchLinks error: %v", err)
		t.Fail()
	}

	// Check ret length
	if len(ret) > 0 {
		t.Logf("SearchLinks pass return %v number of results", len(ret))
	} else {
		t.Errorf("SearchLinks pass return %v number of results", len(ret))
	}

	for _, link := range ret {
		if link != "" {
			t.Logf("SearchLinks pass link returned %v", link)
			_, urlErr := url.ParseRequestURI(link)

			if urlErr != nil {
				t.Errorf(urlErr.Error())
				t.Fail()
			} else {
				t.Logf("Url of \"%v\" is a url", link)
			}
		} else {
			t.Errorf("SearchLinks pass link returned %v", link)
		}
	}
}

func TestCanSearchVideoData(t *testing.T) {

	ret, err := testPackage.SearchVideoData("Luke Smith", 10)

	if err != nil {
		t.Errorf("SearchVideoData Error: %v", err)
		t.Fail()
	}

	// ID check
	if ret[0].Id > "" {
		t.Logf("Id pass return %v", ret[0].Id)
	} else {
		t.Errorf("Id error return %v", ret[0].Id)
	}

	// Title check
	if ret[0].Title != "" {
		t.Logf("Title pass return %v", ret[0].Title)
	} else {
		t.Errorf("Title error return %v", ret[0].Title)
	}
}
