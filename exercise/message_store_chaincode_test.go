package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"testing"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	fmt.Println("check init")
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, objectID string, paramsMap map[string]string) {
	var err error
	var response map[string]interface{}

	bytes, err := stub.GetState(objectID)
	if bytes == nil {
		fmt.Println("State", objectID, "failed to get value")
		t.FailNow()
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		fmt.Println("Failed to unmarshal the stored value: " + err.Error())
	}

	// do a deep comparsion of the two maps
	if reflect.DeepEqual(paramsMap, response) {
		fmt.Println("params value", response["params"], "was not", paramsMap, "as expected")
		t.FailNow()
	}

}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}

}

func Test_Init(t *testing.T) {
	// TODO specify the chaincode package name
	scc := new()

	// TODO generate a new mock stub
	stub := shim.NewMockStub("chaincode_stub_name", scc)

	// TODO pass the stub interface to the checkInit function and pass "init" as argument
	checkInit(t, my_own_stub, [][]byte{
		[]byte("argument"),
	})
}

func Test_Store_Message(t *testing.T) {
	scc := new(MessageStore)
	stub := shim.NewMockStub("messagestorecc", scc)

	// TODO add random values to the parmas map
	paramsMap := map[string]string{}

	// TODO add the same values you added to the parmasMap in json representation
	// e.g. parmasJSON := `{"key_1":"val_1"}`
	paramsJSON := `{}`

	// init the chaincode
	checkInit(t, stub, [][]byte{
		[]byte("init"),
	})

	// TODO use the checkInvoke function to write a message to the ledger
	// create a message
	checkInvoke(t, stub, [][]byte{
		[]byte("method"),
		[]byte("key"),
		[]byte("values"),
	})

	// TODO test if the message has been written to the ledger
	// validate if the params have been written to the ledger in the same way as they have been specified
	checkState(t, stub, "key", paramsMap)

}

func Test_Write_Read_Message(t *testing.T) {
	scc := new(MessageStore)
	stub := shim.NewMockStub("messagestorecc", scc)

	// TODO add random values to the parmas map
	paramsMap := map[string]string{}

	// TODO add the same values you added to the parmasMap in json representation
	paramsJSON := `{}`

	// init the chaincode
	checkInit(t, stub, [][]byte{
		[]byte("init"),
	})

	// TODO use the checkInvoke function to write a message to the ledger
	// create a message
	checkInvoke(t, stub, [][]byte{
		[]byte("method"),
		[]byte("key"),
		[]byte("values"),
	})

	// TODO test if the message has been written to the ledger
	// validate if the params have been written to the ledger in the same way as they have been specified
	checkState(t, stub, "message_key", paramsMap)

	// TODO add an invoke of the 'read' function
	// read the message
	checkInvoke(t, stub, [][]byte{
		[]byte("method"),
		[]byte("key"),
	})

}
