package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func makeHTTPPostReq(temperature, numOfPpl, humidity int) {
	fmt.Println("Sending the Post Request")
	jsonData := map[string]int{"Temperature": temperature, "NumofPPL": numOfPpl, "Humidity": humidity}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "http://localhost:8000/local", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

}

func main() {
	for {
		temperature := random(0, 100)
		numOfPpl := random(0, 100)
		humidity := random(0, 100)
		makeHTTPPostReq(temperature, numOfPpl, humidity)

		fmt.Println("successfully sent")

		duration := (5 * time.Second)
		time.Sleep(duration)
	}
}
