package tests

import (
	"net/http"
	"sync"
	"testing"
)

func TestLoad(t *testing.T) {
	const requests = 100

	var wg sync.WaitGroup

	for i := 0; i < requests; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resp, err := http.Get("http://goboxd:8080/healthz")
			if err != nil {
				t.Error(err)
				return
			}

			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected 200 got %d", resp.StatusCode)
			}
		}()
	}

	wg.Wait()
}
