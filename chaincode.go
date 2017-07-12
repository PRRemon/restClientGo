
package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Region Chaincode implementation
type DTC_Chaincode struct {
}

func (t *DTC_Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	return nil, nil
}

func (t *DTC_Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "testPost" {
		// Initialize the User
		return testPost(stub, args)
	}
	return nil, nil
}

func (t *DTC_Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "testGet" {
		// Read contract details from blockchain
		return testGet(stub, args)
	} 

	return nil, nil
}

func main() {
	err := shim.Start(new(DTC_Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
