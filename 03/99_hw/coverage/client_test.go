package main

// тут писать код тестов
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// тут писать код тестов

func TestSearchClient_FindUsers(t *testing.T) {
	cases := []SearchRequest{
		{
			Limit:      20,
			Query:      "commodo",
			OrderField: "Name",
			OrderBy:    OrderByAsIs,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	for caseNum, item := range cases {
		sc := &SearchClient{
			AccessToken: "asdas",
			URL:         ts.URL,
		}
		result, err := sc.FindUsers(item)
		if err != nil {
			fmt.Println(caseNum, ": FUCK , error: ", err)
		}
		if err == nil {
			fmt.Println(result.Users)
		}
	}
	ts.Close()

}
