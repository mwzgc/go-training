package main

import (
	"fmt"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestErrorGroup(t *testing.T) {
	g := new(errgroup.Group)

	urls := []string{
		"https://www.baidu.com",
		"https://www.taobao.com",
	}

	for _, url := range urls {

		url := url

		g.Go(func() error {
			resp, err := http.Get(url)

			if err == nil {
				// bodyRes, _ := ioutil.ReadAll(resp.Body)
				// fmt.Printf("%v", string(bodyRes))
				resp.Body.Close()
			}

			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("success fetched all urls")
	}
}
