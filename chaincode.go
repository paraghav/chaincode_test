package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ChaincodeType is the chaincode type for this module
type ChaincodeType struct {
}

func (t *ChaincodeType) storeBlob(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("incorrect number of arguments, expecting 2")
	}

	blobID := args[0]
	blobContent := []byte(args[1])

	err := stub.PutState(blobID, blobContent)

	return nil, err
}

func (t *ChaincodeType) retrieveBlob(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("incorrect number of arguments, expecting 1")
	}

	blobID := args[0]

	return stub.GetState(blobID)
}

// =============================================================================
// Chaincode Interface Functions
// =============================================================================

// Init initializes the chaincode
func (t *ChaincodeType) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

// Invoke invokes the chaincode
func (t *ChaincodeType) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "storeBlob" {
		return t.storeBlob(stub, args)
	}
	return nil, errors.New("invalid invoke function name")
}

// Query queries the chaincode
func (t *ChaincodeType) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "retrieveBlob" {
		return t.retrieveBlob(stub, args)
	}
	return nil, errors.New("invalid query function name")
}

// =============================================================================
// Main
// =============================================================================
func main() {
	err := shim.Start(new(ChaincodeType))
	if err != nil {
		fmt.Printf("error starting test chaincode: %s", err)
	}
}
