package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Test_Init(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	response := simpleCC.Init(mockStub)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init test failed")
		t.FailNow()
	}
}

func Test_Invoke_noFunction(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"
	mockStub.MockTransactionStart(txId)
	args := []string{"noFunction", "key1", "value1"}
	argsAsBytes := make([][]byte, len(args))
	for i, v := range args {
		argsAsBytes[i] = []byte(v)
	}
	response := mockStub.MockInvoke(txId, argsAsBytes)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 404 {
		fmt.Println("Invoke_noFunction test failed")
		t.FailNow()
	}
}
func Test_Invoke_saveNewResponse(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"
	mockStub.MockTransactionStart(txId)
	args := []string{"saveNewResponse", "key1", "value1"}
	argsAsBytes := make([][]byte, len(args))
	for i, v := range args {
		argsAsBytes[i] = []byte(v)
	}
	response := mockStub.MockInvoke(txId, argsAsBytes)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Invoke_saveNewResponse test failed")
		t.FailNow()
	}
}
func Test_Invoke_getKeyDetails(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"
	mockStub.MockTransactionStart(txId)
	args := []string{"getKeyDetails", "key1"}
	argsAsBytes := make([][]byte, len(args))
	for i, v := range args {
		argsAsBytes[i] = []byte(v)
	}
	response := mockStub.MockInvoke(txId, argsAsBytes)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 400 {
		fmt.Println("Invoke_getKeyDetails test failed")
		t.FailNow()
	}
}

func Test_Main(t *testing.T) {
	main()
}

func Test_saveNewResponse(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key1", "value1"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewResponse(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 200 {
		fmt.Println("saveNewResponse test failed")
		t.FailNow()
	}
}

func Test_saveNewResponse_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key", "value1", "value2"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewResponse(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("saveNewResponse_incorrectArgs test failed")
		t.FailNow()
	}
}

func Test_saveNewResponse_emptyKey(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"", "value1"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewResponse(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("saveNewResponse_emptyKey test failed")
		t.FailNow()
	}
}

func Test_getKeyDetails(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	args := []string{"key1", "value1"}
	response := simpleCC.saveNewResponse(mockStub, args)

	args = []string{"key1"}
	response = simpleCC.getKeyDetails(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 200 {
		fmt.Println("getKeyDetails test failed")
		t.FailNow()
	}
}
func Test_getKeyDetails_nodata(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key1"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.getKeyDetails(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("getKeyDetails_nodata test failed")
		t.FailNow()
	}
}

func Test_getKeyDetails_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key1", "key2"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.getKeyDetails(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("getKeyDetails_incorrectArgs test failed")
		t.FailNow()
	}
}
func Test_getKeyDetails_emptyKey(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{""}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.getKeyDetails(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("getKeyDetails_emptyKey test failed")
		t.FailNow()
	}
}

func Test_getHistoryByKey(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	args := []string{"key1", "value1"}
	response := simpleCC.saveNewResponse(mockStub, args)

	args = []string{"key1"}
	response = simpleCC.getHistoryByKey(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	// change error code here once GetHistoryForKey is implemented in mockstub
	if s := response.GetStatus(); s != 500 {
		fmt.Println("getHistoryByKey test failed")
		t.FailNow()
	}
}

// func Test_getHistoryByKey_nodata(t *testing.T) {
// 	simpleCC := new(SimpleAsset)
// 	mockStub := shim.NewMockStub("mockstub", simpleCC)
// 	txId := "mockTxID"

// 	args := []string{"key1"}
// 	mockStub.MockTransactionStart(txId)
// 	response := simpleCC.getHistoryByKey(mockStub, args)
// 	mockStub.MockTransactionEnd(txId)
// 	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
// 	fmt.Println("Payload: " + string(response.GetPayload()))
// 	fmt.Println("Message: " + response.GetMessage())

// 	if s := response.GetStatus(); s != 400 {
// 		fmt.Println("getHistoryByKey_nodata test failed")
// 		t.FailNow()
// 	}
// }
func Test_getHistoryByKey_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key1", "key2"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.getHistoryByKey(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("getHistoryByKey_incorrectArgs test failed")
		t.FailNow()
	}
}
