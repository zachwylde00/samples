/**
 *  Salesforce blockchain Logger
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
 * Modifications from: Vishal Mullur:
 *  https://github.com/xooa/integrations/salesforce
 *
 * Changes:
 *  Logs to Xooa blockchain platform from Salesforce dashboard
 */

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("SalesforceCC")

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either updating the state or retreiving the state created by Init function.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "storeData" {
		return t.storeData(stub, args)
	} else if function == "retrieveData" {
		return t.retrieveData(stub, args)
	} else if function == "deleteData" {
		return t.deleteData(stub, args)
	}

	logger.Error("Function declaration not found for ", function)
	response := shim.Error("Invalid function name " + function + " for 'invoke'")
	response.Status = 404
	return response
}

// storeData stores the event on the ledger. For each key,
// it will override the current state with the new one
func (t *SimpleAsset) storeData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("storeData() called.")

	// Function call with argument "Xooa Test" is used for testing function names entered on the Salesforce page.
	// We return success wihout committing the call to the ledger
	if len(args) == 1 && args[0] == "Xooa Test" {
		logger.Debug("Method test call. Nothing will be committed to ledger")
		response := shim.Success([]byte("Method test call. Nothing will be committed to ledger"))
		response.Status = 200
		return response
	}

	if len(args) != 3 {
		logger.Error("Incorrect number of arguments passed in storeData.")
		response := shim.Error("Expecting 3 arguments, found " + string(len(args)) + " arguments instead")
		response.Status = 400
		return response
	}

	sObjectType := args[0]
	key := args[1]
	arr := []string{sObjectType, key}
	compositeKey, err1 := stub.CreateCompositeKey("sObject~key", arr)

	// Check if there were any errors while creating composite key
	if err1 != nil {
		logger.Error("Error while creating composite key", err1.Error())
		response := shim.Error("Error while creating composite key " + err1.Error())
		response.Status = 400
		return response
	}

	eventJSONAsString := args[2]
	eventJSONAsBytes := []byte(eventJSONAsString)

	// Put value to the ledger
	err := stub.PutState(compositeKey, eventJSONAsBytes)

	if err != nil {
		logger.Error("Error occured while calling PutState()", err.Error())
		response := shim.Error("Error occured while calling PutState() " + err.Error())
		response.Status = 400
		return response
	}

	return shim.Success([]byte(compositeKey))
}

// deleteData stores the event on the ledger. For each key,
// it will override the current state with the new one
func (t *SimpleAsset) deleteData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("deleteData() called.")

	// Function call with argument "Xooa Test" is used for testing function names entered on the Salesforce page.
	// We return success wihout committing the call to the ledger
	if len(args) == 1 && args[0] == "Xooa Test" {
		logger.Debug("Method test call. Nothing will be committed to ledger")
		response := shim.Success([]byte("Method test call. Nothing will be committed to ledger"))
		response.Status = 200
		return response
	}

	if len(args) != 2 {
		logger.Error("Incorrect number of arguments passed in deleteData.")
		response := shim.Error("Expecting 2 arguments, found " + string(len(args)) + " arguments instead")
		response.Status = 400
		return response
	}

	sObjectType := args[0]
	key := args[1]
	arr := []string{sObjectType, key}
	compositeKey, err1 := stub.CreateCompositeKey("sObject~key", arr)

	// Check if there were any errors while creating composite key
	if err1 != nil {
		logger.Error("Error while creating composite key", err1.Error())
		response := shim.Error("Error while creating composite key " + err1.Error())
		response.Status = 400
		return response
	}

	// Delete value from the ledger
	err := stub.DelState(compositeKey)

	if err != nil {
		logger.Error("Error occured while calling DelState()", err.Error())
		response := shim.Error("Error occured while calling DelState() " + err.Error())
		response.Status = 400
		return response
	}

	return shim.Success([]byte(compositeKey))
}

// retrieveData queries using key.
// It retrieves the latest state of the value.
func (t *SimpleAsset) retrieveData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("retrieveData() called.")

	if len(args) != 2 {
		logger.Error("Incorrect number of arguments passed in retrieveData.")
		response := shim.Error("Expecting 2 arguments, found " + string(len(args)) + " arguments instead")
		response.Status = 400
		return response
	}

	sObjectType := args[0]
	key := args[1]
	arr := []string{sObjectType, key}
	compositeKey, err1 := stub.CreateCompositeKey("sObject~key", arr)

	// Check if there were any errors while creating composite key
	if err1 != nil {
		logger.Error("Error while creating composite key", err1.Error())
		response := shim.Error("Error while creating composite key " + err1.Error())
		response.Status = 400
		return response
	}

	// Get value form the ledger
	valuesAsBytes, err := stub.GetState(compositeKey)

	if err != nil {
		logger.Error("Error occured while calling GetState()", compositeKey)
		response := shim.Error("Error occured while calling GetState() " + compositeKey)
		response.Status = 400
		return response
	}
	if valuesAsBytes == nil {
		logger.Error("No value found for key", compositeKey)
		response := shim.Error("No value found for key " + compositeKey)
		response.Status = 400
		return response
	}

	return shim.Success(valuesAsBytes)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		logger.Error("Error starting SimpleAsset smartcontract: ", err)
		fmt.Printf("Error starting smartcontract: %s", err)
	}
}
