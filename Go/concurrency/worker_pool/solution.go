package worker_pool

import (
	"fmt"
	"sync"
	"time"
)

func Multiply(num int) int {
	return 2 * num
}

func WorkerW(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		time.Sleep(1 * time.Second)

		results <- Multiply(j)
	}
}

func NumWorkerPool(nums []int, workerCount int) []int {
	jobs := make(chan int, workerCount)
	results := make(chan int, workerCount)

	var wg sync.WaitGroup

	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go WorkerW(jobs, results, &wg)
	}

	for i := 0; i < len(nums); i++ {
		jobs <- nums[i]
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var out []int
	for r := 0; r < len(nums); r++ {
		out = append(out, <-results)
	}

	return out
}

func main() {
	fmt.Println(NumWorkerPool([]int{1, 3, 4}, 2))
}
