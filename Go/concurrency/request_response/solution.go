package request_response

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Log struct {
	LogType   string `json:"type"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

type Request struct {
	ClientID string `json:"client_id"`
	LogType  string `json:"log_type"`
}

type Response struct {
	ClientID string
	Count    int
}

type FileDB struct {
	Logs       []Log
	NumWorkers int
}

func (f *FileDB) ReadData() error {
	data, err := os.ReadFile("logs.json")
	if err != nil {
		return fmt.Errorf("failed to read logs")
	}

	var logs []Log
	err = json.Unmarshal(data, &logs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal")
	}

	f.Logs = logs

	f.NumWorkers = len(f.Logs) / 10

	return nil
}

func (f *FileDB) ReadInChunksWorker(req Request, resp chan<- Response, mainWG *sync.WaitGroup) {
	defer mainWG.Done()

	var wg sync.WaitGroup
	var count int

	for i := 0; i < f.NumWorkers; i++ {
		wg.Add(1)
		go func(index int, logType string) {
			defer wg.Done()

			startingIndex := index * 10

			for _, v := range f.Logs[startingIndex : startingIndex+10] {
				if v.LogType == logType {
					var mu sync.Mutex
					mu.Lock()
					count++
					mu.Unlock()
				}
			}
		}(i, req.LogType)
	}

	wg.Wait()
	resp <- Response{
		ClientID: req.ClientID,
		Count:    count,
	}
}

func server(reqs <-chan Request, resp chan<- Response) {
	db := FileDB{}
	err := db.ReadData()
	if err != nil {
		panic("received error reading files")
	}

	var wg sync.WaitGroup

	for req := range reqs {
		wg.Add(1)
		go db.ReadInChunksWorker(req, resp, &wg)
	}

	wg.Wait()
	close(resp)
}

func main() {
	requests := make(chan Request)
	responses := make(chan Response)

	go server(requests, responses)

	go func() {
		data, err := os.ReadFile("clients.json")
		if err != nil {
			panic("failed to read from file")
		}

		var clientReqs []Request
		err = json.Unmarshal(data, &clientReqs)
		if err != nil {
			panic("failed to unmarshal file")
		}

		for _, v := range clientReqs {
			requests <- Request{
				ClientID: v.ClientID,
				LogType:  v.LogType,
			}
		}
		close(requests)
	}()

	for resp := range responses {
		fmt.Println(resp)
	}
}
