package main

import (
	"net/http"
	"fmt"

	"io/ioutil"

	"log"
)

func main(){
	resp, err := http.Get("http://localhost:8080")
	if err != nil{
		 log.Fatalf("some error %v",err)
	}
	defer resp.Body.Close()
	bytesOfData, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Errorf("some error reading the response body ", err.Error())
	}

	fmt.Println(string(bytesOfData))
}
