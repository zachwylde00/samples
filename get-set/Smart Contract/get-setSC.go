/**
 *  Xooa Zapier Logger Chaincode
 *
 *  Copyright 2018 Xooa
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at:
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 *  for the specific language governing permissions and limitations under the License.
 */
/*
 * Original source via IBM Corp:
 *  https://hyperledger-fabric.readthedocs.io/en/release-1.2/chaincode4ade.html#pulling-it-all-together
 *
 * Modifications from: Arisht Jain:
 *  https://github.com/xooa/samples
 *
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("get-setSC")

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// ...... checking if commit id gets updated
// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Init() called.")
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Invoke() called.")
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	//checking for account level access
	channelId := stub.GetChannelID()
	accountAssertError := cid.AssertAttributeValue(stub, "ChannelId", channelId)
	if accountAssertError != nil {
		return shim.Error(accountAssertError.Error())
	}

	// checking for access to app
	chaincodeId := os.Getenv("CORE_CHAINCODE_ID_NAME")
	pair := strings.Split(chaincodeId, ":")
	chaincodeName := pair[0]
	appAssertError := cid.AssertAttributeValue(stub, "AppId", chaincodeName)
	if appAssertError != nil {
		return shim.Error(appAssertError.Error())
	}

	fmt.Println("invoke is running " + fn)

	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else { // assume 'get' even if fn is nil
		result, err = get(stub, args)
	}
	if err != nil {
		logger.Error("Error occured in Invoke.")
		return shim.Error(err.Error())
	}
	fmt.Println("invoke returning " + result)
	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	logger.Debug("set() called.")
	fmt.Println("- start set value")
	if len(args) != 2 {
		logger.Error("Incorrect number of arguments passed in set.")
		return "", fmt.Errorf("Incorrect number of arguments. Expecting 2 arguments: " + strconv.Itoa(len(args)) + " given.")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		logger.Error("Error occured while calling PutState(): ", err)
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	fmt.Println("- end set value")
	return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	fmt.Println("- start get value")
	if len(args) != 1 {
		logger.Error("Incorrect number of arguments passed in get.")
		return "", fmt.Errorf("Incorrect number of arguments. Expecting 1 arguments: " + strconv.Itoa(len(args)) + " given.")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		logger.Error("Error occured while calling GetState(): ", err)
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		logger.Info("No data received for key : ", args[0])
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	fmt.Println("- end get value")
	return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		logger.Error("Error starting SimpleAsset chaincode: ", err)
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
