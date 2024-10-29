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
	identifierId := 12345
	response, err := fetchDate(ctx, identifierId)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response: ", response)
	fmt.Println("execution time: ", time.Since(start))
}

func fetchDate(ctx context.Context, identifier int) (int, error) {
	reponseCode, err := fetchExternalApi()
	if err != nil {
		return 500, err
	}

	return reponseCode, nil
}

// simulates a slow api call
func fetchExternalApi() (int, error) {
	time.Sleep(500 * time.Millisecond)
	return 200, nil

}
