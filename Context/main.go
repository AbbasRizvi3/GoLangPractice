package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func assignJobs(id int, chl chan int) {
	wg.Add(1)
	chl <- id

}

func doJob(chl1 chan int, chl2 chan int, ctx context.Context) {
	for {
		select {
		case msg, ok := <-chl1:
			{
				if ok {
					chl2 <- msg
				}
			}
		case <-ctx.Done():
			{
				print("Cancelling")
				return
			}
		}

	}
}

func showpassedvalue(ctx context.Context) {
	data, ok := ctx.Value("KEY").(string)
	if ok {
		fmt.Printf("%v\n", data)
	}
}

func fetchURL(ctx context.Context, url string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Request creation error:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status for", url, ":", resp.Status)
}

func showResult(x int) {
	defer wg.Done()
	fmt.Printf("%d\n", x)

}
func main() {
	workers := 3
	jobs := 5
	url := "http://www.google.com"
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	data_channel := make(chan int)
	results := make(chan int)

	go func() {
		for val := range results {
			showResult(val)
		}
	}()

	for i := 0; i < workers; i++ {
		go doJob(data_channel, results, ctx)
	}

	for i := 0; i < jobs; i++ {
		assignJobs(i, data_channel)
	}

	// cancel()

	close(data_channel)
	wg.Wait()
	close(results)

	fetchURL(ctx2, url)

	ctx3 := context.WithValue(context.Background(), "KEY", "abc")
	showpassedvalue(ctx3)

}
