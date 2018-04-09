// Disclaimer
//
// THIS SAMPLE CODE MAY BE USED SOLELY AS PART OF THE TEST AND EVALUATION OF THE SAP CLOUD PLATFORM BLOCKCHAIN SERVICE (THE “SERVICE”)
// AND IN ACCORDANCE WITH THE TERMS OF THE TEST AND EVALUATION AGREEMENT FOR THE SERVICE. THIS SAMPLE CODE PROVIDED “AS IS”, WITHOUT
// ANY WARRANTY, ESCROW, TRAINING, MAINTENANCE, OR SERVICE OBLIGATIONS WHATSOEVER ON THE PART OF SAP.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strings"
)

type MessageStore struct{}

type message struct {
	ID    string `json:"ID"`
	Value string `json:"value"`
}

// Main function starts up the chaincode in the container during instantiate
//
func main() {
	if err := shim.Start(new(MessageStore)); err != nil {
		fmt.Printf("Main: Error starting MessageStore chaincode: %s", err)
	}
}

// Init is called during Instantiate transaction after the chaincode container
// has been established for the first time, allowing the chaincode to
// initialize its internal data. Note that chaincode upgrade also calls this
// function to reset or to migrate data, so be careful to avoid a scenario
// where you inadvertently clobber your ledger's data!
//
func (cc *MessageStore) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// Validate supplied init parameters, in this case zero arguments!

	// TODO implement the check for the init parameters

	// TODO add a return 'success' statement
	return pb.Response{}
}

// Invoke is called to update or query the ledger in a proposal transaction.
// Updated state variables are not committed to the ledger until the
// transaction is committed.
//
func (cc *MessageStore) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Which function is been called?
	// TODO add args variable
	function, _ := stub.GetFunctionAndParameters()

	// Turn arguments to lower case
	function = strings.ToLower(function)


	// Route call to the correct function
	switch function {

	// TODO implement a case statement for "write"
	case "test":
		return cc.write(,)

	// TODO implement a case statement for "read"

	default:

		// TODO add a informative error message
		msg := "Error"
		return shim.Error(msg)
	}
}

// Write an ID and string to the blockchain
//
func (cc *MessageStore) write(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// extract the ID and value from the arguments the content and object id
	// TODO make sure that just two values have been passed as arguments

	// key should be lowercase
	id := strings.ToLower(args[0])

	// create a message struct
	msg := &message{ID: id, Value: args[1]}

	// marshal message struct to json
	// can be returned directly
	// allows more sophisticated search using couchdb /mq
	// TODO add variable for json return value
	_, err := json.Marshal(msg)


	// Validate that this ID does not yet exist
	// TODO check if the key has already been assigned


	// Write the message
	// TODO if key has not been assigned yet write it to the ledger



	// TODO if successful return 'shim.Success(nil)'
	return pb.Response{}
}

// Read a string from the blockchain, given its ID
//
// TODO implement the header of the 'read' function
func () 'function_name' (input_arguments) response_type {

	// TODO make sure that there is just one argument passed.


	// TODO evaluate the argument check and respond with a proper error message if the number was wrong.


	// TODO convert the key to lowercase so that there is no confusion with lower and upper case characters  			id := strings.ToLower(args[0])


	// TODO read the key from the ledger


	// TODO evaluate the response from the read operation and return a success or error
}
