/**
 *  Xooa Dropbox Logger Smart Contract
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
 * Modifications from: Kavi Sarna:
 *  https://github.com/xooa/integrations
 */

package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("dropboxCC")

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
	
	} else if function == "getEntityDetails" {

		return t.getEntityDetails(stub, args)

	} else if function == "getHistoryForEntity" {

		return t.getHistoryForEntity(stub, args)

	} else if function == "saveProfile" {

		return t.saveProfile(stub, args)

	} else if function == "getProfile" {

		return t.getProfile(stub, args)

	} else if function == "getVersion" {

		return t.getVersion(stub)
	}

	logger.Error("Function declaration not found for ", function)
	resp := shim.Error("Invalid function name : " + function)
	resp.Status = 404

	return resp
}

// getVersion retrieves the name and version of this smart contract
func (t *SimpleAsset) getVersion(stub shim.ChaincodeStubInterface) peer.Response {

	logger.Debug("getVersion called.")

	return shim.Success([]byte("dropbox:1.0.0"))
}

// saveNewEvent stores the event on the ledger. For each entity,
// it will override the current state with the new one.
func (t *SimpleAsset) saveNewEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	logger.Debug("saveNewEvent() called.")

	// Essential check to verify number of arguments
	if len(args) != 2 {

		logger.Error("Incorrect number of arguments passed in saveNewEvent.")
		resp := shim.Error("Incorrect number of arguments. Expecting 2 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	id := args[0]
	
	if id == "" {

		logger.Error("Empty key passed to saveNewEvent()")
		resp := shim.Error("Key must not be empty.")
		resp.Status = 400

		return resp

	} else {

		eventJSONasString := args[1]
		logger.Debug("eventJSONasString: ", eventJSONasString)
		eventJSONasBytes := []byte(eventJSONasString)

		err := stub.PutState(id, eventJSONasBytes)

		if err != nil {

			logger.Error("Error occured while calling PutState(): ", err)

			return shim.Error("Error in updating ledger.")
		}
	}

	return shim.Success([]byte(id))
}


// main function starts up the smart contract in the container during instantiate
func main() {

	logger.Debug("main() called.")
	if err := shim.Start(new(SimpleAsset)); err != nil {
		logger.Error("Error starting SimpleAsset chaincode: ", err)
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}


// getHistoryForEntity queries the entity using entity id.
// It retrieve all the changes to the entity happened over time.
func (t *SimpleAsset) getHistoryForEntity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	logger.Debug("getHistoryForEntity called.")

	// Essential check to verify number of arguments
	if len(args) != 1 {

		logger.Error("Incorrect number of arguments passed in getHistoryForEntity.")
		resp := shim.Error("Incorrect number of arguments. Expecting 1 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	id := args[0]

	resultsIterator, err := stub.GetHistoryForKey(id)

	if err != nil {

		logger.Error("Error occured while calling GetHistoryForKey(): ", err)
		return shim.Error("Failed to get history for entity with id = " + id)
	}

	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the event
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {

		response, err := resultsIterator.Next()

		if err != nil {

			return shim.Error("Error occured while calling getHistoryForEntity (resultsIterator): " + err.Error())
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

// getEntityDetails queries using id.
// It retrieves the latest state of the entity.
func (t *SimpleAsset) getEntityDetails(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	logger.Debug("getEntityDetails called.")

	var err error

	// Essential check to verify number of arguments
	if len(args) != 1 {

		logger.Error("Incorrect number of arguments passed in getEntityDetails.")
		resp := shim.Error("Incorrect number of arguments. Expecting 1 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	id := args[0]

	valueAsBytes, err := stub.GetState(id)

	if err != nil {

		logger.Error("Error occured while calling GetState(): ", err)

		return shim.Error("Failed to get state for entity with id = " + id)

	}

	if valueAsBytes == nil {

		logger.Info("No data received for entity with id = " + id)

		resp := shim.Error("nil result got for entity with id = " + id)
		resp.Status = 400
		return resp
	}

	return shim.Success(valueAsBytes)
}


// saveProfile stores the oauth2 token and cursor on the ledger. For each account id,
// it will override the current state with the new one.
func (t *SimpleAsset) saveProfile(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	logger.Debug("saveProfile() called.")

	if len(args) != 2 {

		logger.Error("Incorrect number of arguments passed in saveProfile().")
		resp := shim.Error("Incorrect number of arguments. Expecting 2 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	accountId := args[0]
	profile := args[1]

	logger.Debug("Profile: ", profile)

	profileAsBytes := []byte(profile)

	if accountId == "" {

		logger.Error("Empty key passed to saveProfile()")
		resp := shim.Error("Key must not be empty.")
		resp.Status = 400

		return resp

	} else {

		err := stub.PutState(accountId, profileAsBytes)

		if err != nil {

			logger.Error("Error occured while calling PutState(): ", err)
			return shim.Error("Error in updating ledger.")
		}
	}

	return shim.Success([]byte(accountId))
}


// getProfile queries using account id.
// It retrieves the latest stored oauth key and cursor for the account id.
func (t *SimpleAsset) getProfile(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	logger.Debug("getProfile() called.")

	var err error

	if len(args) != 1 {

		logger.Error("Incorrect number of arguments passed in getProfile().")
		resp := shim.Error("Incorrect number of arguments. Expecting 1 argument: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	accountId := args[0]

	valueAsBytes, err := stub.GetState(accountId)

	if err != nil {

		logger.Error("Error occured while calling getProfile(): ", err)
		return shim.Error("Failed to get oauth2 for " + accountId)
	}

	if valueAsBytes == nil {

		logger.Info("No data received for " + accountId + " accountId.")
		resp := shim.Error("nil result got for " + accountId + " accountId.")
		resp.Status = 400
		return resp
	}
	return shim.Success(valueAsBytes)
}
