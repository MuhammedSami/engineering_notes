package semaphore

import (
	"fmt"
	"sync"
)

func FetchURLsWithSemaphore(urls []string, limit int) []string {
	results := make([]string, len(urls))

	var wg sync.WaitGroup

	sem := make(chan struct{}, limit)

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(wID int, url string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			results[wID] = "fetched" + url
		}(i, urls[i])
	}

	wg.Wait()

	return results
}

func main() {
	urls := []string{
		"test1",
		"test2",
		"test3",
		"test4",
	}

	fmt.Println(FetchURLsWithSemaphore(urls, 3))
}
