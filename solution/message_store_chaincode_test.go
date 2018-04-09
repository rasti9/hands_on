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
	scc := new(MessageStore)
	stub := shim.NewMockStub("messagestorecc", scc)

	checkInit(t, stub, [][]byte{
		[]byte("init"),
	})
}

func Test_Store_Message(t *testing.T) {
	scc := new(MessageStore)
	stub := shim.NewMockStub("messagestorecc", scc)

	paramsMap := map[string]string{
		"key_1": "val_1",
		"key_2": "val_2",
		"key_3": "val_3",
		"key_4": "val_4",
		"key_5": "val_5",
	}
	paramsJSON := `{"key_1":"val_1","key_2":"val_2","key_3":"val_3","key_4":"val_4","key_5":"val_5"}`

	// init the chaincode
	checkInit(t, stub, [][]byte{
		[]byte("init"),
	})

	// create a message
	checkInvoke(t, stub, [][]byte{
		[]byte("write"),
		[]byte("message_key"),
		[]byte(paramsJSON),
	})

	// validate if the params have been written to the ledger in the same way as they have been specified
	checkState(t, stub, "message_key", paramsMap)

}

func Test_Write_Read_Message(t *testing.T) {
	scc := new(MessageStore)
	stub := shim.NewMockStub("messagestorecc", scc)

	paramsMap := map[string]string{
		"key_1": "val_1",
		"key_2": "val_2",
		"key_3": "val_3",
		"key_4": "val_4",
		"key_5": "val_5",
	}
	paramsJSON := `{"key_1":"val_1","key_2":"val_2","key_3":"val_3","key_4":"val_4","key_5":"val_5"}`

	// init the chaincode
	checkInit(t, stub, [][]byte{
		[]byte("init"),
	})

	// create a message
	checkInvoke(t, stub, [][]byte{
		[]byte("write"),
		[]byte("message_key"),
		[]byte(paramsJSON),
	})

	// validate if the params have been written to the ledger in the same way as they have been specified
	checkState(t, stub, "message_key", paramsMap)

	// read the message
	checkInvoke(t, stub, [][]byte{
		[]byte("read"),
		[]byte("message_key"),
	})

}
