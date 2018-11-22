/**
 *  Xooa Zapier Logger Smart Contract
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
 *  https://github.com/xooa/integrations
 *
 */

package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("zapierCC")

// SimpleAsset implements a simple smart contract to manage an asset
type SimpleAsset struct {
}

// Init is called during smart contract instantiation to initialize any
// data. Note that smart contract upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Init() called.")
	return shim.Success(nil)
}

// Invoke is called per transaction on the smart contract. Each transaction is
// either updating the state or retreiving the state created by Init function.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Invoke() called.")
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()

	if function == "saveNewEvent" {
		return t.saveNewEvent(stub, args)
	} else if function == "getKeyDetails" {
		return t.getKeyDetails(stub, args)
	} else if function == "getHistoryByKey" {
		return t.getHistoryByKey(stub, args)
	} else if function == "getVersion" {
		return t.getVersion(stub)
	}
	logger.Error("Function declaration not found for ", function)
	resp := shim.Error("Invalid function name: " + function)
	resp.Status = 404
	return resp
}

// getVersion retrieves the name and version of this smart contract
func (t *SimpleAsset) getVersion(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("getVersion called.")

	return shim.Success([]byte("Zapier:1.0.0"))
}

// saveNewEvent stores the event on the ledger. For each key,
// it will override the current state with the new one
func (t *SimpleAsset) saveNewEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("saveNewEvent() called.")

	// Essential check to verify number of arguments
	if len(args) != 2 {
		logger.Error("Incorrect number of arguments passed in saveNewEvent.")
		resp := shim.Error("Incorrect number of arguments. Expecting 2 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}
	key := args[0]
	eventAsString := args[1]
	logger.Debug("eventAsString: ", eventAsString)

	eventAsBytes := []byte(eventAsString)
	if key == "" {
		logger.Error("Empty key passed to saveNewEvent()")
		resp := shim.Error("Key must not be empty.")
		resp.Status = 400
		return resp
	} else {
		err := stub.PutState(key, eventAsBytes)
		if err != nil {
			logger.Error("Error occured while calling PutState(): ", err)
			return shim.Error("Error in updating ledger.")
		}

		err = stub.SetEvent("saveNewEvent", eventAsBytes)
		if err != nil {
			logger.Error("Error occured while calling SetEvent(): ", err)
		}
	}
	return shim.Success([]byte(key))
}

// main function starts up the smart contract in the container during instantiate
func main() {
	logger.Debug("main() called.")
	if err := shim.Start(new(SimpleAsset)); err != nil {
		logger.Error("Error starting SimpleAsset smart contract: ", err)
		fmt.Printf("Error starting SimpleAsset smart contract: %s", err)
	}
}

// getHistoryByKey queries the ledger using key.
// It retrieve all the changes to the value happened over time.
func (t *SimpleAsset) getHistoryByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("getHistoryByKey called.")

	// Essential check to verify number of arguments
	if len(args) != 1 {
		logger.Error("Incorrect number of arguments passed in getHistoryByKey.")
		resp := shim.Error("Incorrect number of arguments. Expecting 1 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	key := args[0]
	resultsIterator, err := stub.GetHistoryForKey(key)

	if err != nil {
		logger.Error("Error occured while calling GetHistoryForKey(): ", err)
		return shim.Error("Error occured while calling GetHistoryForKey: " + err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the event
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			logger.Error("Error occured while calling resultsIterator.Next(): ", err)
			return shim.Error("Error occured while calling GetHistoryByKey (resultsIterator): " + err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		buffer.WriteString(string(response.Value))
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

// getKeyDetails queries using key.
// It retrieves the latest state of the value.
func (t *SimpleAsset) getKeyDetails(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("getKeyDetails called.")

	// Essential check to verify number of arguments
	if len(args) != 1 {
		logger.Error("Incorrect number of arguments passed in getKeyDetails.")
		resp := shim.Error("Incorrect number of arguments. Expecting 1 argument: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	key := args[0]
	if key == "" {
		logger.Error("Empty key passed to getKeyDetails()")
		resp := shim.Error("Key must not be empty.")
		resp.Status = 400
		return resp
	} else {
		valueAsBytes, err := stub.GetState(key)
		if err != nil {
			logger.Error("Error occured while calling GetState(): ", err)
			return shim.Error("Failed to get state for id=" + key)
		}
		if valueAsBytes == nil {
			logger.Info("No data received for key : ", key)
			resp := shim.Error("No data received for key: " + key)
			resp.Status = 400
			return resp
		}
		return shim.Success(valueAsBytes)
	}
}
