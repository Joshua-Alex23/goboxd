package tests

import (
	"bytes"
	"net/http"
	"testing"
)

func TestCppHelloWorld(t *testing.T) {
	body := []byte(`{
		"language":"cpp",
		"source":"#include <iostream>\nint main(){std::cout<<\"hello\";}",
		"tests":[
			{
				"stdin":"",
				"expected_stdout":"hello"
			}
		]
	}`)

	resp, err := http.Post(
		"http://goboxd:8080/run",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
}
