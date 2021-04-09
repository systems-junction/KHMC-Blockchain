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
	CommentNotes      string `json:"commentNotes"`
	ApprovedBy        string `json:"approvedBy"`
	VendorId          string `json:"vendorId"`
	Status            string `json:"status"`
	CommitteeStatus   string `json:"committeeStatus"`
	InProgressTime    string `json:"inProgressTime"`
	CreatedAt         string `json:"createdAt"`
	SentAt            string `json:"sentAt"`
	UpdatedAt         string `json:"updatedAt"`
}

type Item struct {
	ObjectType   string `json:"Type"`
	ItemId       string `json:"itemId"`
	CurrQty      string `json:"currQty"`
	ReqQty       string `json:"reqQty"`
	Comments     string `json:"comments"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ItemCode     string `json:"itemCode"`
	IStatus      string `json:"istatus"`
	SecondStatus string `json:"secondStatus"`
}

type RItem struct {
	ObjectType    string `json:"Type"`
	ItemId        string `json:"itemId"`
	CurrentQty    string `json:"currentQty"`
	RequestedQty  string `json:"requestedQty"`
	RecieptUnit   string `json:"recieptUnit"`
	IssueUnit     string `json:"issueUnit"`
	FuItemCost    string `json:"fuItemCost"`
	Description   string `json:"description"`
	RStatus       string `json:"rstatus"`
	RSecondStatus string `json:"rsecondStatus"`
	BatchArray
	TempBatchArray
}

type PurchaseRequest struct {
	ObjectType      string `json:"Type"`
	RequestNo       string `json:"requestNo"`
	GeneratedBy     string `json:"generatedBy"`
	Status          string `json:"status"`
	CommitteeStatus string `json:"committeeStatus"`
	Availability    string `json:"availability"`
	Reason          string `json:"reason"`
	VendorId        string `json:"vendorId"`
	Rr              string `json:"rr"`
	Item
	RequesterName   string `json:"requesterName"`
	RejectionReason string `json:"rejectionReason"`
	Department      string `json:"department"`
	CommentNotes    string `json:"commentNotes"`
	OrderType       string `json:"orderType"`
	Generated       string `json:"generated"`
	ApprovedBy      string `json:"approvedBy"`
	InProgressTime  string `json:"inProgressTime"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type Patient struct {
	ObjectType string `json:"Type"`
	PatientID  string `json:"patientID"`
	Name       string `json:"name"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
}

type ReplenishmentRequest struct {
	ObjectType    string `json:"Type"`
	RequestNo     string `json:"requestNo"`
	Generated     string `json:"generated"`
	GeneratedBy   string `json:"generatedBy"`
	DateGenerated string `json:"dateGenerated"`
	Reason        string `json:"reason"`
	FuId          string `json:"fuId"`
	To            string `json:"to"`
	From          string `json:"from"`
	Comments      string `json:"comments"`
	RItem
	Status         string `json:"status"`
	SecondStatus   string `json:"secondStatus"`
	RrB            string `json:"rrB"`
	ApprovedBy     string `json:"approvedBy"`
	RequesterName  string `json:"requesterName"`
	OrderType      string `json:"orderType"`
	Department     string `json:"department"`
	CommentNote    string `json:"commentNote"`
	InProgressTime string `json:"inProgressTime"`
	CompletedTime  string `json:"completedTime"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

type FunctionalUnit struct {
	ObjectType  string `json:"Type"`
	Uuid        string `json:"uuid"`
	FuName      string `json:"fuName"`
	Description string `json:"description"`
	FuHead      string `json:"fuHead"`
	Status      string `json:"status"`
	BuId        string `json:"buId"`
	FuLogId     string `json:"fuLogId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Batch struct {
	BatchNumber string `json:"batchNumber"`
	ExpiryDate  string `json:"expiryDate"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	QrCode      string `json:"qrCode"`
}

type RBatch struct {
	BatchNumber string `json:"batchNumber"`
	ExpiryDate  string `json:"expiryDate"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	QrCode      string `json:"qrCode"`
}

type TempBatch struct {
	BatchNumber string `json:"batchNumber"`
	ExpiryDate  string `json:"expiryDate"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	QrCode      string `json:"qrCode"`
}

type BatchArray []Batch

type RBatchArray []RBatch

type TempBatchArray []TempBatch

type FuInventory struct {
	ObjectType   string `json:"Type"`
	FuId         string `json:"fuId"`
	ItemId       string `json:"itemId"`
	Qty          string `json:"qty"`
	MaximumLevel string `json:"maximumLevel"`
	ReorderLevel string `json:"reorderLevel"`
	MinimumLevel string `json:"minimumLevel"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	BatchArray
	TempBatchArray
}

type ReceiveItem struct {
	ObjectType      string `json:"Type"`
	ItemId          string `json:"itemId"`
	PrId            string `json:"prId"`
	Status          string `json:"status"`
	CurrentQty      string `json:"currentQty"`
	RequestedQty    string `json:"requestedQty"`
	ReceivedQty     string `json:"receivedQty"`
	BonusQty        string `json:"bonusQty"`
	BatchNumber     string `json:"batchNumber"`
	LotNumber       string `json:"lotNumber"`
	ExpiryDate      string `json:"expiryDate"`
	Unit            string `json:"unit"`
	Discount        string `json:"discount"`
	UnitDiscount    string `json:"unitDiscount"`
	DiscountAmount  string `json:"discountAmount"`
	Tax             string `json:"tax"`
	TaxAmount       string `json:"taxAmount"`
	FinalUnitPrice  string `json:"finalUnitPrice"`
	SubTotal        string `json:"subTotal"`
	DiscountAmount2 string `json:"discountAmount2"`
	TotalPrice      string `json:"totalPrice"`
	Invoice         string `json:"invoice"`
	DateInvoice     string `json:"dateInvoice"`
	DateReceived    string `json:"dateReceived"`
	Notes           string `json:"notes"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	ReturnedQty     string `json:"returnedQty"`
	BatchArray
	UnitPrice string `json:"unitPrice"`
}

type WarehouseInventory struct {
	ObjectType   string `json:"Type"`
	ItemId       string `json:"itemId"`
	Qty          string `json:"qty"`
	MaximumLevel string `json:"maximumLevel"`
	MinimumLevel string `json:"minimumLevel"`
	ReorderLevel string `json:"reorderLevel"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	BatchArray
	TempBatchArray
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
	if function == "queryPatientByName" {
		return t.queryPatientByName(stub, args)
	}
	if function == "queryPurchaseOrder" {
		return t.queryPurchaseOrder(stub, args)
	}
	if function == "queryPurchaseRequest" {
		return t.queryPurchaseRequest(stub, args)
	}
	if function == "addPatient" {
		return t.addPatient(stub, args)
	}
	if function == "queryPatient" {
		return t.queryPatient(stub, args)
	}
	if function == "queryPatientByName" {
		return t.queryPatientByName(stub, args)
	}
	if function == "addReplenishmentRequest" {
		return t.addReplenishmentRequest(stub, args)
	}
	if function == "queryReplenishmentRequest" {
		return t.queryReplenishmentRequest(stub, args)
	}
	if function == "addFunctionalUnit" {
		return t.addFunctionalUnit(stub, args)
	}
	if function == "queryFunctionalUnit" {
		return t.queryFunctionalUnit(stub, args)
	}
	if function == "addFuInventory" {
		return t.addFuInventory(stub, args)
	}
	if function == "queryFuInventory" {
		return t.queryFuInventory(stub, args)
	}
	if function == "addReceiveItem" {
		return t.addReceiveItem(stub, args)
	}
	if function == "queryReceiveItem" {
		return t.queryReceiveItem(stub, args)
	}
	if function == "addWarehouseInventory" {
		return t.addWarehouseInventory(stub, args)
	}
	if function == "queryWarehouseInventory" {
		return t.queryWarehouseInventory(stub, args)
	}
	if function == "updateWarehouseInventory" {
		return t.updateWarehouseInventory(stub, args)
	}
	if function == "updatePurchaseOrderStatus" {
		return t.updatePurchaseOrderStatus(stub, args)
	}
	if function == "updatePurchaseOrderCommitteeStatus" {
		return t.updatePurchaseOrderCommitteeStatus(stub, args)
	}
	if function == "updatePurchaseRequestStatus" {
		return t.updatePurchaseRequestStatus(stub, args)
	}
	if function == "updatePurchaseRequestCommitteeStatus" {
		return t.updatePurchaseRequestCommitteeStatus(stub, args)
	}
	if function == "updatePurchaseRequestItemStatus" {
		return t.updatePurchaseRequestItemStatus(stub, args)
	}
	if function == "updatePurchaseRequestItemSecondStatus" {
		return t.updatePurchaseRequestItemSecondStatus(stub, args)
	}
	if function == "updateReplenishmentRequestStatus" {
		return t.updateReplenishmentRequestStatus(stub, args)
	}
	if function == "updateReplenishmentRequestSecondStatus" {
		return t.updateReplenishmentRequestSecondStatus(stub, args)
	}
	if function == "updateReplenishmentRequestItemStatus" {
		return t.updateReplenishmentRequestItemStatus(stub, args)
	}
	if function == "updateReplenishmentRequestItemSecondStatus" {
		return t.updateReplenishmentRequestItemSecondStatus(stub, args)
	}
	if function == "updateFunctionalUnitStatus" {
		return t.updateFunctionalUnitStatus(stub, args)
	}
	if function == "updateReceiveItemStatus" {
		return t.updateReceiveItemStatus(stub, args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addPurchaseOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 14 {
		return shim.Error("Incorrect Number of Aruments. Expecting 14")
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

	purchaseOrderNo := args[0]
	purchaseRequestId := args[1]
	date := args[2]
	generated := args[3]
	generatedBy := args[4]
	commentNotes := args[5]
	approvedBy := args[6]
	vendorId := args[7]
	status := args[8]
	committeeStatus := args[9]
	inProgressTime := args[10]
	createdAt := args[11]
	sentAt := args[12]
	updatedAt := args[13]

	// ======Check if PurchaseOrder Already exists

	PurchaseOrderAsBytes, err := stub.GetState(purchaseOrderNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseOrderAsBytes != nil {
		return shim.Error("The Inserted PurchaseOrder ID already Exists: " + purchaseOrderNo)
	}

	// ===== Create PurchaseOrder Object and Marshal to JSON

	objectType := "PurchaseOrder"
	PurchaseOrder := &PurchaseOrder{objectType, purchaseOrderNo, purchaseRequestId, date, generated, generatedBy, commentNotes, approvedBy, vendorId, status, committeeStatus, inProgressTime, createdAt, sentAt, updatedAt}
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

	if len(args) != 27 {
		return shim.Error("Incorrect Number of Arguments. Expecting 27")
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
	if len(args[19]) <= 0 {
		return shim.Error("20th Argument Must be a Non-Empty String")
	}
	if len(args[20]) <= 0 {
		return shim.Error("21th Argument Must be a Non-Empty String")
	}
	if len(args[21]) <= 0 {
		return shim.Error("22th Argument Must be a Non-Empty String")
	}
	if len(args[22]) <= 0 {
		return shim.Error("23th Argument Must be a Non-Empty String")
	}
	if len(args[23]) <= 0 {
		return shim.Error("24th Argument Must be a Non-Empty String")
	}
	if len(args[24]) <= 0 {
		return shim.Error("25th Argument Must be a Non-Empty String")
	}
	if len(args[25]) <= 0 {
		return shim.Error("26th Argument Must be a Non-Empty String")
	}
	if len(args[26]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}

	requestNo := args[0]
	generatedBy := args[1]
	status := args[2]
	committeeStatus := args[3]
	availability := args[4]
	reason := args[5]
	vendorId := args[6]
	rr := args[7]
	itemId := args[8]
	currQty := args[9]
	reqQty := args[10]
	comments := args[11]
	name := args[12]
	description := args[13]
	itemCode := args[14]
	istatus := args[15]
	secondStatus := args[16]
	requesterName := args[17]
	rejectionReason := args[18]
	department := args[19]
	commentNotes := args[20]
	orderType := args[21]
	generated := args[22]
	approvedBy := args[23]
	inProgressTime := args[24]
	createdAt := args[25]
	updatedAt := args[26]

	// ======Check if PurchaseRequest Already exists

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseRequestAsBytes != nil {
		return shim.Error("The Inserted PurchaseOrder ID already Exists: " + requestNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "PurchaseRequest"
	object := "Item"
	PurchaseRequest := &PurchaseRequest{objectType, requestNo, generatedBy, status, committeeStatus, availability, reason, vendorId, rr, Item{object, itemId, currQty, reqQty, comments, name, description, itemCode, istatus, secondStatus}, requesterName, rejectionReason, department, commentNotes, orderType, generated, approvedBy, inProgressTime, createdAt, updatedAt}
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

func (t *SmartContract) addReplenishmentRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 36 {
		return shim.Error("Incorrect Number of Arguments. Expecting 36")
	}

	fmt.Println("Adding new ReplenishmentRequest")

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
	if len(args[19]) <= 0 {
		return shim.Error("20th Argument Must be a Non-Empty String")
	}
	if len(args[20]) <= 0 {
		return shim.Error("21th Argument Must be a Non-Empty String")
	}
	if len(args[21]) <= 0 {
		return shim.Error("22th Argument Must be a Non-Empty String")
	}
	if len(args[22]) <= 0 {
		return shim.Error("23th Argument Must be a Non-Empty String")
	}
	if len(args[23]) <= 0 {
		return shim.Error("24th Argument Must be a Non-Empty String")
	}
	if len(args[24]) <= 0 {
		return shim.Error("25th Argument Must be a Non-Empty String")
	}
	if len(args[25]) <= 0 {
		return shim.Error("26th Argument Must be a Non-Empty String")
	}
	if len(args[26]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}
	if len(args[27]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[28]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[29]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[30]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[31]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[32]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[33]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[34]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[35]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}

	requestNo := args[0]
	generated := args[1]
	generatedBy := args[2]
	dateGenerated := args[3]
	reason := args[4]
	fuId := args[5]
	to := args[6]
	from := args[7]
	comments := args[8]
	itemId := args[9]
	currentQty := args[10]
	requestedQty := args[11]
	recieptUnit := args[12]
	issueUnit := args[13]
	fuItemCost := args[14]
	description := args[15]
	rstatus := args[16]
	rsecondStatus := args[17]
	batchNumber := args[18]
	expiryDate := args[19]
	quantity := args[20]
	tempbatchNumber := args[21]
	tempexpiryDate := args[22]
	tempquantity := args[23]
	status := args[24]
	secondStatus := args[25]
	rrB := args[26]
	approvedBy := args[27]
	requesterName := args[28]
	orderType := args[29]
	department := args[30]
	commentNote := args[31]
	inProgressTime := args[32]
	completedTime := args[33]
	createdAt := args[34]
	updatedAt := args[35]

	// ======Check if PurchaseRequest Already exists

	replenishmentRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if replenishmentRequestAsBytes != nil {
		return shim.Error("The Inserted replenishmentRequest ID already Exists: " + requestNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "ReplenishmentRequest"
	object := "RItem"
	ReplenishmentRequest := &ReplenishmentRequest{objectType, requestNo, generated, generatedBy, dateGenerated, reason, fuId, to, from, comments, RItem{object, itemId, currentQty, requestedQty, recieptUnit, issueUnit, fuItemCost, description, rstatus, rsecondStatus,
		BatchArray{Batch{
			BatchNumber: batchNumber,
			ExpiryDate:  expiryDate,
			Quantity:    quantity,
		}}, TempBatchArray{TempBatch{
			BatchNumber: tempbatchNumber,
			ExpiryDate:  tempexpiryDate,
			Quantity:    tempquantity,
		}}}, status, secondStatus, rrB, approvedBy, requesterName, orderType, department, commentNote, inProgressTime, completedTime, createdAt, updatedAt}
	ReplenishmentRequestJSONasBytes, err := json.Marshal(ReplenishmentRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save replenishmentRequest to State

	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReplenishmentRequest")
	return shim.Success(nil)
}

func (t *SmartContract) queryReplenishmentRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requestNo := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ReplenishmentRequest\",\"requestNo\":\"%s\"}}", requestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
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

func (t *SmartContract) addPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect Number of Aruments. Expecting 8")
	}

	fmt.Println("Adding new Patient")

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

	patientID := args[0]
	name := args[1]
	age := args[2]
	gender := args[3]

	// ======Check if Patient Already exists

	patientAsBytes, err := stub.GetState(patientID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if patientAsBytes != nil {
		return shim.Error("The Inserted Patient ID already Exists: " + patientID)
	}

	// ===== Create Patient Object and Marshal to JSON

	objectType := "Patient"
	Patient := &Patient{objectType, patientID, name, age, gender}
	PatientJSONasBytes, err := json.Marshal(Patient)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save Patient to State

	err = stub.PutState(patientID, PatientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Patient")
	return shim.Success(nil)
}

func (t *SmartContract) queryPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Patient\",\"patientID\":\"%s\"}}", patientID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryPatientByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	name := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Patient\",\"name\":\"%s\"}}", name)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addFunctionalUnit(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 9 {
		return shim.Error("Incorrect Number of Aruments. Expecting 8")
	}

	fmt.Println("Adding new FunctionalUnit")

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

	uuid := args[0]
	fuName := args[1]
	description := args[2]
	fuHead := args[3]
	status := args[4]
	buId := args[5]
	fuLogId := args[6]
	createdAt := args[7]
	updatedAt := args[8]

	// ======Check if FunctionalUnit Already exists

	FunctionalUnitAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if FunctionalUnitAsBytes != nil {
		return shim.Error("The Inserted FunctionalUnit ID already Exists: " + uuid)
	}

	// ===== Create FunctionalUnit Object and Marshal to JSON

	objectType := "FunctionalUnit"
	FunctionalUnit := &FunctionalUnit{objectType, uuid, fuName, description, fuHead, status, buId, fuLogId, createdAt, updatedAt}
	FunctionalUnitJSONasBytes, err := json.Marshal(FunctionalUnit)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save FunctionalUnit to State

	err = stub.PutState(uuid, FunctionalUnitJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved FunctionalUnit")
	return shim.Success(nil)
}

func (t *SmartContract) queryFunctionalUnit(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	uuid := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"FunctionalUnit\",\"uuid\":\"%s\"}}", uuid)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addFuInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 14 {
		return shim.Error("Incorrect Number of Aruments. Expecting 14")
	}

	fmt.Println("Adding new FuInventory")

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

	fuId := args[0]
	itemId := args[1]
	qty := args[2]
	maximumLevel := args[3]
	reorderLevel := args[4]
	minimumLevel := args[5]
	createdAt := args[6]
	updatedAt := args[7]
	batchNumber := args[8]
	expiryDate := args[9]
	quantity := args[10]
	tempbatchNumber := args[11]
	tempexpiryDate := args[12]
	tempquantity := args[13]

	// ======Check if FuInventory Already exists

	FuInventoryAsBytes, err := stub.GetState(fuId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if FuInventoryAsBytes != nil {
		return shim.Error("The Inserted FuInventory ID already Exists: " + fuId)
	}

	// ===== Create FuInventory Object and Marshal to JSON

	objectType := "FuInventory"
	FuInventory := &FuInventory{objectType, fuId, itemId, qty, maximumLevel, reorderLevel, minimumLevel, createdAt, updatedAt, BatchArray{
		{
			BatchNumber: batchNumber,
			ExpiryDate:  expiryDate,
			Quantity:    quantity,
		}}, TempBatchArray{TempBatch{
		BatchNumber: tempbatchNumber,
		ExpiryDate:  tempexpiryDate,
		Quantity:    tempquantity,
	}}}
	FuInventoryJSONasBytes, err := json.Marshal(FuInventory)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save FuInventory to State

	err = stub.PutState(fuId, FuInventoryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved FuInventory")
	return shim.Success(nil)
}

func (t *SmartContract) queryFuInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fuId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"FuInventory\",\"fuId\":\"%s\"}}", fuId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addReceiveItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 33 {
		return shim.Error("Incorrect Number of Arguments. Expecting 33")
	}

	fmt.Println("Adding new ReceiveItem")

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
	if len(args[19]) <= 0 {
		return shim.Error("20th Argument Must be a Non-Empty String")
	}
	if len(args[20]) <= 0 {
		return shim.Error("21th Argument Must be a Non-Empty String")
	}
	if len(args[21]) <= 0 {
		return shim.Error("22th Argument Must be a Non-Empty String")
	}
	if len(args[22]) <= 0 {
		return shim.Error("23th Argument Must be a Non-Empty String")
	}
	if len(args[23]) <= 0 {
		return shim.Error("24th Argument Must be a Non-Empty String")
	}
	if len(args[24]) <= 0 {
		return shim.Error("25th Argument Must be a Non-Empty String")
	}
	if len(args[25]) <= 0 {
		return shim.Error("26th Argument Must be a Non-Empty String")
	}
	if len(args[26]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}
	if len(args[27]) <= 0 {
		return shim.Error("28th Argument Must be a Non-Empty String")
	}
	if len(args[28]) <= 0 {
		return shim.Error("29th Argument Must be a Non-Empty String")
	}
	if len(args[29]) <= 0 {
		return shim.Error("29th Argument Must be a Non-Empty String")
	}
	if len(args[30]) <= 0 {
		return shim.Error("29th Argument Must be a Non-Empty String")
	}
	if len(args[31]) <= 0 {
		return shim.Error("29th Argument Must be a Non-Empty String")
	}
	if len(args[32]) <= 0 {
		return shim.Error("29th Argument Must be a Non-Empty String")
	}

	itemId := args[0]
	prId := args[1]
	status := args[2]
	currentQty := args[3]
	requestedQty := args[4]
	receivedQty := args[5]
	bonusQty := args[6]
	batchNumber := args[7]
	lotNumber := args[8]
	expiryDate := args[9]
	unit := args[10]
	discount := args[11]
	unitDiscount := args[12]
	discountAmount := args[13]
	tax := args[14]
	taxAmount := args[15]
	finalUnitPrice := args[16]
	subTotal := args[17]
	discountAmount2 := args[18]
	totalPrice := args[19]
	invoice := args[20]
	dateInvoice := args[21]
	dateReceived := args[22]
	notes := args[23]
	createdAt := args[24]
	updatedAt := args[25]
	returnedQty := args[26]
	batchNumberArr := args[27]
	expiryDateArr := args[28]
	quantity := args[29]
	price := args[30]
	qrCode := args[31]
	unitPrice := args[32]

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + itemId)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "ReceiveItem"
	ReceiveItem := &ReceiveItem{objectType, itemId, prId, status, currentQty, requestedQty, receivedQty, bonusQty, batchNumber, lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal, discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, returnedQty, BatchArray{
		Batch{
			BatchNumber: batchNumberArr,
			ExpiryDate:  expiryDateArr,
			Quantity:    quantity,
			Price:       price,
			QrCode:      qrCode,
		}}, unitPrice}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(itemId, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReceiveItem")
	return shim.Success(nil)
}

func (t *SmartContract) queryReceiveItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itemId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ReceiveItem\",\"itemId\":\"%s\"}}", itemId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addWarehouseInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 15 {
		return shim.Error("Incorrect Number of Aruments. Expecting 13")
	}

	fmt.Println("Adding new WarehouseInventory")

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
		return shim.Error("13th Argument Must be a Non-Empty String")
	}
	if len(args[14]) <= 0 {
		return shim.Error("13th Argument Must be a Non-Empty String")
	}

	itemId := args[0]
	qty := args[1]
	maximumLevel := args[2]
	minimumLevel := args[3]
	reorderLevel := args[4]
	createdAt := args[5]
	updatedAt := args[6]
	batchNumber := args[7]
	expiryDate := args[8]
	quantity := args[9]
	price := args[10]
	tempbatchNumber := args[11]
	tempexpiryDate := args[12]
	tempquantity := args[13]
	tempprice := args[14]

	// ======Check if WarehouseInventory Already exists

	WarehouseInventoryAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if WarehouseInventoryAsBytes != nil {
		return shim.Error("The Inserted WarehouseInventory ID already Exists: " + itemId)
	}

	// ===== Create WarehouseInventory Object and Marshal to JSON

	objectType := "WarehouseInventory"
	WarehouseInventory := &WarehouseInventory{objectType, itemId, qty, maximumLevel, minimumLevel, reorderLevel, createdAt, updatedAt, BatchArray{
		{
			BatchNumber: batchNumber,
			ExpiryDate:  expiryDate,
			Quantity:    quantity,
			Price:       price,
		}}, TempBatchArray{TempBatch{
		BatchNumber: tempbatchNumber,
		ExpiryDate:  tempexpiryDate,
		Quantity:    tempquantity,
		Price:       tempprice,
	}}}
	WarehouseInventoryJSONasBytes, err := json.Marshal(WarehouseInventory)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save WarehouseInventory to State

	err = stub.PutState(itemId, WarehouseInventoryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved WarehouseInventory")
	return shim.Success(nil)
}

func (t *SmartContract) queryWarehouseInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itemId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"WarehouseInventory\",\"itemId\":\"%s\"}}", itemId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) updateWarehouseInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 7 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	itemId := args[0]
	newBatchNumber := args[1]
	newExpiryDate := args[2]
	newQuantity := args[3]
	newTempBatchNumber := args[4]
	newTempExpiryDate := args[5]
	newTempQuantity := args[6]

	fmt.Println("- start  ", itemId, newBatchNumber, newExpiryDate, newQuantity, newTempBatchNumber, newTempExpiryDate, newTempQuantity)

	responseAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if responseAsBytes == nil {
		return shim.Error("response does not exist")
	}

	responseToUpdate := WarehouseInventory{}
	err = json.Unmarshal(responseAsBytes, &responseToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	responseToUpdate.BatchArray = append(responseToUpdate.BatchArray, Batch{
		BatchNumber: newBatchNumber,
		ExpiryDate:  newExpiryDate,
		Quantity:    newQuantity,
	}) //change the status
	responseToUpdate.TempBatchArray = append(responseToUpdate.TempBatchArray, TempBatch{
		BatchNumber: newTempBatchNumber,
		ExpiryDate:  newTempExpiryDate,
		Quantity:    newTempQuantity,
	}) //change the status

	responseJSONasBytes, _ := json.Marshal(responseToUpdate)
	err = stub.PutState(itemId, responseJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseOrderStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	purchaseOrderNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", purchaseOrderNo, newStatus)

	PurchaseOrderAsBytes, err := stub.GetState(purchaseOrderNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("PurchaseOrder does not exist")
	}

	PurchaseOrderToUpdate := PurchaseOrder{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseOrderToUpdate.Status = newStatus //change the status

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(purchaseOrderNo, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseOrderCommitteeStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	purchaseOrderNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", purchaseOrderNo, newStatus)

	PurchaseOrderAsBytes, err := stub.GetState(purchaseOrderNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("PurchaseOrder does not exist")
	}

	PurchaseOrderToUpdate := PurchaseOrder{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseOrderToUpdate.CommitteeStatus = newStatus //change the status

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(purchaseOrderNo, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseRequestStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseRequestAsBytes == nil {
		return shim.Error("PurchaseRequest does not exist")
	}

	PurchaseRequestToUpdate := PurchaseRequest{}
	err = json.Unmarshal(PurchaseRequestAsBytes, &PurchaseRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseRequestToUpdate.Status = newStatus //change the status

	PurchaseRequestJSONasBytes, _ := json.Marshal(PurchaseRequestToUpdate)
	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseRequestCommitteeStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseRequestAsBytes == nil {
		return shim.Error("PurchaseRequest does not exist")
	}

	PurchaseRequestToUpdate := PurchaseRequest{}
	err = json.Unmarshal(PurchaseRequestAsBytes, &PurchaseRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseRequestToUpdate.CommitteeStatus = newStatus //change the status

	PurchaseRequestJSONasBytes, _ := json.Marshal(PurchaseRequestToUpdate)
	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseRequestItemStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseRequestAsBytes == nil {
		return shim.Error("PurchaseRequest does not exist")
	}

	PurchaseRequestToUpdate := PurchaseRequest{}
	err = json.Unmarshal(PurchaseRequestAsBytes, &PurchaseRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseRequestToUpdate.Item.IStatus = newStatus //change the status

	PurchaseRequestJSONasBytes, _ := json.Marshal(PurchaseRequestToUpdate)
	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseRequestItemSecondStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	PurchaseRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseRequestAsBytes == nil {
		return shim.Error("PurchaseRequest does not exist")
	}

	PurchaseRequestToUpdate := PurchaseRequest{}
	err = json.Unmarshal(PurchaseRequestAsBytes, &PurchaseRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PurchaseRequestToUpdate.Item.SecondStatus = newStatus //change the status

	PurchaseRequestJSONasBytes, _ := json.Marshal(PurchaseRequestToUpdate)
	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReplenishmentRequestStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	ReplenishmentRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReplenishmentRequestAsBytes == nil {
		return shim.Error("ReplenishmentRequest does not exist")
	}

	ReplenishmentRequestToUpdate := ReplenishmentRequest{}
	err = json.Unmarshal(ReplenishmentRequestAsBytes, &ReplenishmentRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ReplenishmentRequestToUpdate.Status = newStatus //change the status

	ReplenishmentRequestJSONasBytes, _ := json.Marshal(ReplenishmentRequestToUpdate)
	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReplenishmentRequestSecondStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	ReplenishmentRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReplenishmentRequestAsBytes == nil {
		return shim.Error("ReplenishmentRequest does not exist")
	}

	ReplenishmentRequestToUpdate := ReplenishmentRequest{}
	err = json.Unmarshal(ReplenishmentRequestAsBytes, &ReplenishmentRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ReplenishmentRequestToUpdate.SecondStatus = newStatus //change the status

	ReplenishmentRequestJSONasBytes, _ := json.Marshal(ReplenishmentRequestToUpdate)
	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReplenishmentRequestItemStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	ReplenishmentRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReplenishmentRequestAsBytes == nil {
		return shim.Error("ReplenishmentRequest does not exist")
	}

	ReplenishmentRequestToUpdate := ReplenishmentRequest{}
	err = json.Unmarshal(ReplenishmentRequestAsBytes, &ReplenishmentRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ReplenishmentRequestToUpdate.RItem.RStatus = newStatus //change the status

	ReplenishmentRequestJSONasBytes, _ := json.Marshal(ReplenishmentRequestToUpdate)
	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReplenishmentRequestItemSecondStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	requestNo := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", requestNo, newStatus)

	ReplenishmentRequestAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReplenishmentRequestAsBytes == nil {
		return shim.Error("ReplenishmentRequest does not exist")
	}

	ReplenishmentRequestToUpdate := ReplenishmentRequest{}
	err = json.Unmarshal(ReplenishmentRequestAsBytes, &ReplenishmentRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ReplenishmentRequestToUpdate.RItem.RSecondStatus = newStatus //change the status

	ReplenishmentRequestJSONasBytes, _ := json.Marshal(ReplenishmentRequestToUpdate)
	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateFunctionalUnitStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	uuid := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", uuid, newStatus)

	FunctionalUnitAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if FunctionalUnitAsBytes == nil {
		return shim.Error("FunctionalUnit does not exist")
	}

	FunctionalUnitToUpdate := FunctionalUnit{}
	err = json.Unmarshal(FunctionalUnitAsBytes, &FunctionalUnitToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	FunctionalUnitToUpdate.Status = newStatus //change the status

	FunctionalUnitJSONasBytes, _ := json.Marshal(FunctionalUnitToUpdate)
	err = stub.PutState(uuid, FunctionalUnitJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReceiveItemStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	itemId := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", itemId, newStatus)

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReceiveItemAsBytes == nil {
		return shim.Error("ReceiveItem does not exist")
	}

	ReceiveItemToUpdate := ReceiveItem{}
	err = json.Unmarshal(ReceiveItemAsBytes, &ReceiveItemToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ReceiveItemToUpdate.Status = newStatus //change the status

	ReceiveItemJSONasBytes, _ := json.Marshal(ReceiveItemToUpdate)
	err = stub.PutState(itemId, ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
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
