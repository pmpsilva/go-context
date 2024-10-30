package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()
	// Background returns a non-nil, empty [Context]. It is never canceled, has no
	// values, and has no deadline.
	ctx := context.Background()
	slowId := 12345
	//id := 123

	response, err := fetchDate(ctx, slowId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response: ", response)
	fmt.Println("execution time: ", time.Since(start))
}

// structure to use on channel
type Response struct {
	statusCode int
	err        error
}

func fetchDate(ctx context.Context, identifier int) (int, error) {

	//The timeout is 200 milliseconds, but the fetchExternalApi() function has a 500 milliseconds sleep duration, causing a timeout error.
	childContext, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	//Defer to avoid context leaking - make sure after return will cancel the context
	defer cancel()

	respchan := make(chan Response)
	go func() {
		reponseCode, err := fetchExternalApi(identifier)
		respchan <- Response{
			statusCode: reponseCode,
			err:        err,
		}

	}()

	for {
		select {
		case <-childContext.Done():
			return 408, fmt.Errorf("request imeout")
		case resp := <-respchan:
			return resp.statusCode, resp.err
		}
	}

}

// simulates a slow api call
func fetchExternalApi(identifier int) (int, error) {
	if identifier == 12345 {
		time.Sleep(500 * time.Millisecond)
	}
	return 200, nil

}
