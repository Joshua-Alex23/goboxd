package tests

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPythonHelloWorld(t *testing.T) {
	body := []byte(`{
		"language":"py3",
		"source":"print(\"hello\")",
		"tests":[
			{
				"stdin":"",
				"expected_stdout":"hello\n"
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
