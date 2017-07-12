package main

import (  
    "fmt"
    "io/ioutil"
    "bytes"
    "net/http"
    "encoding/json"
    "log"
    "time"
)

type person struct {  
    Name string `json:"name"`
}

func main() {


	// Post Request Using net/http

 	url := "http://localhost:8080/api/login/test2"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"name":"raval"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))


	// Get Request Using net/http

    urlGet := "http://localhost:8080/api/login/test"

    spaceClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, urlGet, nil)
    if err != nil {
        log.Fatal(err)
    }

    //req.Header.Set("User-Agent", "spacecount-tutorial")

    res, getErr := spaceClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    person1 := person{}
    jsonErr := json.Unmarshal(body, &person1)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

	fmt.Printf("HTTP: %s\n", res.Status)
    fmt.Println(person1)

	
}