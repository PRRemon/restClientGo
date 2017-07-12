
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	
	
	"time"
    "io/ioutil"
    "bytes"
    "net/http"
    "log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func testPost(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var ok bool

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	contractId := args[0]
	attachmentName := args[1]
	documentBlob := args[2]

	ok, err = insertAttachmentDetails(stub, contractId, attachmentName, documentBlob)
	if !ok && err == nil {
		return nil, errors.New("Error in inserting attachment")
	}

	return nil, err
}

func testGet(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Need 0 arguments")
	}

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

	return person1, nil
}
