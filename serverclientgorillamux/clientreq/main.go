package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i <= 100000; i++ {
		var ctx context.Context
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/raj", nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		// You can context separately for each request
		request = request.WithContext(ctx)

		client := &http.Client{}
		response, err := client.Do(request)
		// You will get an error "net/http: request canceled" when request timeout exceeds limits
		if err != nil {
			log.Println(err.Error())
			return
		}
		response.Body.Close()
		fmt.Println(response)
		//time.Sleep(1 * time.Second)
	}
}
