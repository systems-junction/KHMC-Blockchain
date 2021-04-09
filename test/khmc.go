/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct {
}

type PurchaseOrder struct {
	ObjectType        string `json:"Type"`
	PurchaseOrderNo   string `json:"purchaseOrderNo"`
	PurchaseRequestId string `json:"purchaseRequestId"`
	Date              string `json:"date"`
	Generated         string `json:"generated"`
	GeneratedBy       string `json:"generatedBy"`
	VendorId          string `json:"vendorId"`
	Status            string `json:"status"`
	CommitteeStatus   string `json:"committeeStatus"`
}

type Item struct {
	ItemId      string `json:"itemId"`
	CurrQty     string `json:"currQty"`
	ReqQty      string `json:"reqQty"`
	Comments    string `json:"comments"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemCode    string `json:"itemCode"`
}

type PurchaseRequest struct {
	ObjectType      string `json:"Type"`
	RequestNo       string `json:"requestNo"`
	GeneratedBy     string `json:"generatedBy"`
	Status          string `json:"status"`
	CommitteeStatus string `json:"committeeStatus"`
	Reason          string `json:"reason"`
	VendorId        string `json:"vendorId"`
	Rr              string `json:"rr"`
	Item            Item
	RequesterName   string `json:"requesterName"`
	RejectionReason string `json:"rejectionReason"`
	Department      string `json:"department"`
	OrderType       string `json:"orderType"`
	Generated       string `json:"generated"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)
	if function == "addPurchaseOrder" {
		return t.addPurchaseOrder(stub, args)
	}
	if function == "addPurchaseRequest" {
		return t.addPurchaseRequest(stub, args)
	}
	if function == "queryPurchaseOrder" {
		return t.queryPurchaseOrder(stub, args)
	}
	if function == "queryPurchaseRequest" {
		return t.queryPurchaseRequest(stub, args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 8 {
		return shim.Error("Incorrect Number of Aruments. Expecting 8")
	}

	fmt.Println("Adding new PurchaseOrder")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}

	purchaseOrderNo := args[0]
	purchaseRequestId := args[1]
	date := args[2]
	generated := args[3]
	generatedBy := args[4]
	vendorId := args[5]
	status := args[6]
	committeeStatus := args[7]

	// ======Check if PurchaseOrder Already exists

	PurchaseOrderAsBytes, err := stub.GetState(purchaseOrderNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseOrderAsBytes != nil {
		return shim.Error("The Inserted PurchaseOrder ID already Exists: " + purchaseOrderNo)
	}

	// ===== Create PurchaseOrder Object and Marshal to JSON

	objectType := "PurchaseOrder"
	PurchaseOrder := &PurchaseOrder{objectType, purchaseOrderNo, purchaseRequestId, date, generated, generatedBy, vendorId, status, committeeStatus}
	PurchaseOrderJSONasBytes, err := json.Marshal(PurchaseOrder)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseOrder to State

	err = stub.PutState(purchaseOrderNo, PurchaseOrderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseOrder")
	return shim.Success(nil)
}

func (t *SmartContract) queryPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	purchaseOrderNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"PurchaseOrder\",\"purchaseOrderNo\":\"%s\"}}", purchaseOrderNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addPurchaseRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 19 {
		return shim.Error("Incorrect Number of Aruments. Expecting 19")
	}

	fmt.Println("Adding new PurchaseOrder")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}
	if len(args[8]) <= 0 {
		return shim.Error("9th Argument Must be a Non-Empty String")
	}
	if len(args[9]) <= 0 {
		return shim.Error("10th Argument Must be a Non-Empty String")
	}
	if len(args[10]) <= 0 {
		return shim.Error("11th Argument Must be a Non-Empty String")
	}
	if len(args[11]) <= 0 {
		return shim.Error("12th Argument Must be a Non-Empty String")
	}
	if len(args[12]) <= 0 {
		return shim.Error("13th Argument Must be a Non-Empty String")
	}
	if len(args[13]) <= 0 {
		return shim.Error("14th Argument Must be a Non-Empty String")
	}
	if len(args[14]) <= 0 {
		return shim.Error("15th Argument Must be a Non-Empty String")
	}
	if len(args[15]) <= 0 {
		return shim.Error("16th Argument Must be a Non-Empty String")
	}
	if len(args[16]) <= 0 {
		return shim.Error("17th Argument Must be a Non-Empty String")
	}
	if len(args[17]) <= 0 {
		return shim.Error("18th Argument Must be a Non-Empty String")
	}
	if len(args[18]) <= 0 {
		return shim.Error("19th Argument Must be a Non-Empty String")
	}

	requestNo := args[0]
	generatedBy := args[1]
	status := args[2]
	committeeStatus := args[3]
	reason := args[4]
	vendorId := args[5]
	rr := args[6]
	itemId := args[7]
	currQty := args[8]
	reqQty := args[9]
	comments := args[10]
	name := args[11]
	description := args[12]
	itemCode := args[13]
	requesterName := args[14]
	rejectionReason := args[15]
	department := args[16]
	orderType := args[17]
	generated := args[18]

	// ======Check if PurchaseRequest Already exists

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseRequestAsBytes != nil {
		return shim.Error("The Inserted PurchaseOrder ID already Exists: " + requestNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "PurchaseRequest"
	PurchaseRequest := &PurchaseRequest{objectType, requestNo, generatedBy, status, committeeStatus, reason, vendorId, rr, Item{itemId, currQty, reqQty, comments, name, description, itemCode}, requesterName, rejectionReason, department, orderType, generated}
	PurchaseRequestJSONasBytes, err := json.Marshal(PurchaseRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseRequest to State

	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseRequest")
	return shim.Success(nil)
}

func (t *SmartContract) queryPurchaseRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requestNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"PurchaseRequest\",\"requestNo\":\"%s\"}}", requestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}
