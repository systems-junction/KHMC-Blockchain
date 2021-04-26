/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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

type ItemSchema struct {
	ObjectType       string   `json:"Type"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	ItemCode         string   `json:"itemCode"`
	Form             string   `json:"form"`
	DrugAllergy      []string `json:"drugAllergy"`
	ReceiptUnit      string   `json:"receiptUnit"`
	IssueUnit        string   `json:"issueUnit"`
	VendorId         string   `json:"vendorId"`
	PurchasePrice    string   `json:"purchasePrice"`
	MinimumLevel     string   `json:"minimumLevel"`
	MaximumLevel     string   `json:"maximumLevel"`
	ReorderLevel     string   `json:"reorderLevel"`
	Cls              string   `json:"cls"`
	MedClass         string   `json:"medClass"`
	SubClass         string   `json:"subClass"`
	GrandSubClass    string   `json:"grandSubClass"`
	Comments         string   `json:"comments"`
	CreatedAt        string   `json:"createdAt"`
	UpdatedAt        string   `json:"updatedAt"`
	ReceiptUnitCost  string   `json:"receiptUnitCost"`
	IssueUnitCost    string   `json:"issueUnitCost"`
	ScientificName   string   `json:"scientificName"`
	TradeName        string   `json:"tradeName"`
	Temprature       string   `json:"temprature"`
	Humidity         string   `json:"humidity"`
	LightSensitive   string   `json:"lightSensitive"`
	ResuableItem     string   `json:"resuableItem"`
	StorageCondition string   `json:"storageCondition"`
	Expiration       string   `json:"expiration"`
	Tax              string   `json:"tax"`
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

type ReceiveItemFUSchema struct {
	ObjectType             string `json:"Type"`
	ItemId                 string `json:"itemId"`
	CurrentQty             string `json:"currentQty"`
	RequestedQty           string `json:"requestedQty"`
	ReceivedQty            string `json:"receivedQty"`
	BonusQty               string `json:"bonusQty"`
	BatchNumber            string `json:"batchNumber"`
	LotNumber              string `json:"lotNumber"`
	ExpiryDate             string `json:"expiryDate"`
	Unit                   string `json:"unit"`
	Discount               string `json:"discount"`
	UnitDiscount           string `json:"unitDiscount"`
	DiscountAmount         string `json:"discountAmount"`
	Tax                    string `json:"tax"`
	TaxAmount              string `json:"taxAmount"`
	FinalUnitPrice         string `json:"finalUnitPrice"`
	SubTotal               string `json:"subTotal"`
	DiscountAmount2        string `json:"discountAmount2"`
	TotalPrice             string `json:"totalPrice"`
	Invoice                string `json:"invoice"`
	DateInvoice            string `json:"dateInvoice"`
	DateReceived           string `json:"dateReceived"`
	Notes                  string `json:"notes"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
	ReplenishmentRequestId string `json:"replenishmentRequestId"`
	BatchArray
}

type ReceiveItemBUSchema struct {
	ObjectType                 string `json:"Type"`
	ItemId                     string `json:"itemId"`
	CurrentQty                 string `json:"currentQty"`
	RequestedQty               string `json:"requestedQty"`
	ReceivedQty                string `json:"receivedQty"`
	BonusQty                   string `json:"bonusQty"`
	BatchNumber                string `json:"batchNumber"`
	LotNumber                  string `json:"lotNumber"`
	ExpiryDate                 string `json:"expiryDate"`
	Unit                       string `json:"unit"`
	Discount                   string `json:"discount"`
	UnitDiscount               string `json:"unitDiscount"`
	DiscountAmount             string `json:"discountAmount"`
	Tax                        string `json:"tax"`
	TaxAmount                  string `json:"taxAmount"`
	FinalUnitPrice             string `json:"finalUnitPrice"`
	SubTotal                   string `json:"subTotal"`
	DiscountAmount2            string `json:"discountAmount2"`
	TotalPrice                 string `json:"totalPrice"`
	Invoice                    string `json:"invoice"`
	DateInvoice                string `json:"dateInvoice"`
	DateReceived               string `json:"dateReceived"`
	Notes                      string `json:"notes"`
	CreatedAt                  string `json:"createdAt"`
	UpdatedAt                  string `json:"updatedAt"`
	ReplenishmentRequestId     string `json:"replenishmentRequestId"`
	ReplenishmentRequestItemId string `json:"replenishmentRequestItemId"`
	QualityRate                string `json:"qualityRate"`
	BatchArray
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

type Staff struct {
	ObjectType           string   `json:"Type"`
	StaffId              string   `json:"staffId"`
	StaffTypeId          string   `json:"staffTypeId"`
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	Designation          string   `json:"designation"`
	ContactNumber        string   `json:"contactNumber"`
	IdentificationNumber string   `json:"identificationNumber"`
	Email                string   `json:"email"`
	Password             string   `json:"password"`
	Gender               string   `json:"gender"`
	Dob                  string   `json:"dob"`
	Address              string   `json:"address"`
	FunctionalUnit       string   `json:"functionalUnit"`
	SystemAdminId        string   `json:"systemAdminId"`
	Status               string   `json:"status"`
	Routes               []string `json:"routes"`
}

type Vendor struct {
	ObjectType             string   `json:"Type"`
	Uuid                   string   `json:"uuid"`
	VendorNo               string   `json:"vendorNo"`
	EnglishName            string   `json:"englishName"`
	ArabicName             string   `json:"arabicName"`
	Telephone1             string   `json:"telephone1"`
	Telephone2             string   `json:"telephone2"`
	ContactEmail           string   `json:"contactEmail"`
	Address                string   `json:"address"`
	Country                string   `json:"country"`
	City                   string   `json:"city"`
	Zipcode                string   `json:"zipcode"`
	Faxno                  string   `json:"faxno"`
	Taxno                  string   `json:"taxno"`
	ContactPersonName      string   `json:"contactPersonName"`
	ContactPersonTelephone string   `json:"contactPersonTelephone"`
	ContactPersonEmail     string   `json:"contactPersonEmail"`
	PaymentTerms           string   `json:"paymentTerms"`
	ShippingTerms          string   `json:"shippingTerms"`
	Rating                 string   `json:"rating"`
	Status                 string   `json:"status"`
	Cls                    string   `json:"cls"`
	SubClass               []string `json:"subClass"`
	Compliance             string   `json:"compliance"`
	CreatedAt              string   `json:"createdAt"`
	UpdatedAt              string   `json:"updatedAt"`
}

type DamageReport struct {
	ObjectType      string `json:"Type"`
	CausedBy        string `json:"causedBy"`
	TotalDamageCost string `json:"totalDamageCost"`
	Date            string `json:"date"`
	ItemCostPerUnit string `json:"itemCostPerUnit"`
}

type ReturnBatch struct {
	BatchNumber         string `json:"batchNumber"`
	ExpiryDatePerBatch  string `json:"expiryDatePerBatch"`
	ReceivedQtyPerBatch string `json:"receivedQtyPerBatch"`
	ReturnedQtyPerBatch string `json:"returnedQtyPerBatch"`
	Price               string `json:"price"`
}

type ReturnBatchArray []ReturnBatch

type InternalReturnRequestSchema struct {
	ObjectType             string `json:"Type"`
	ReturnRequestNo        string `json:"returnRequestNo"`
	GeneratedBy            string `json:"generatedBy"`
	DateGenerated          string `json:"dateGenerated"`
	ExpiryDate             string `json:"expiryDate"`
	To                     string `json:"to"`
	From                   string `json:"from"`
	CurrentQty             string `json:"currentQty"`
	ReturnedQty            string `json:"returnedQty"`
	ItemId                 string `json:"itemId"`
	Description            string `json:"description"`
	FuId                   string `json:"fuId"`
	Reason                 string `json:"reason"`
	ReasonDetail           string `json:"reasonDetail"`
	BuId                   string `json:"buId"`
	DamageReport           DamageReport
	Status                 string `json:"status"`
	ReplenishmentRequestBU string `json:"replenishmentRequestBU"`
	ReplenishmentRequestFU string `json:"replenishmentRequestFU"`
	ApprovedBy             string `json:"approvedBy"`
	CommentNote            string `json:"commentNote"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
	BatchNo                string `json:"batchNo"`
	ReturnBatchArray
}

type ExternalReturnRequestSchema struct {
	ObjectType      string `json:"Type"`
	ReturnRequestNo string `json:"returnRequestNo"`
	GeneratedBy     string `json:"generatedBy"`
	Generated       string `json:"generated"`
	DateGenerated   string `json:"dateGenerated"`
	ExpiryDate      string `json:"expiryDate"`
	ReturnedQty     string `json:"returnedQty"`
	ItemId          string `json:"itemId"`
	PrId            string `json:"prId"`
	Description     string `json:"description"`
	Reason          string `json:"reason"`
	ReasonDetail    string `json:"reasonDetail"`
	DamageReport    DamageReport
	Status          string `json:"status"`
	ApprovedBy      string `json:"approvedBy"`
	CommentNote     string `json:"commentNote"`
	InProgressTime  string `json:"inProgressTime"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	BatchArray
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
	if function == "updatePurchaseOrder" {
		return t.updatePurchaseOrder(stub, args)
	}
	if function == "updatePurchaseRequest" {
		return t.updatePurchaseRequest(stub, args)
	}
	if function == "updateReceiveItem" {
		return t.updateReceiveItem(stub, args)
	}
	if function == "updateReplenishmentRequest" {
		return t.updateReplenishmentRequest(stub, args)
	}
	if function == "updateFuInventory" {
		return t.updateFuInventory(stub, args)
	}
	if function == "updateFunctionalUnit" {
		return t.updateFunctionalUnit(stub, args)
	}
	if function == "addStaff" {
		return t.addStaff(stub, args)
	}
	if function == "addVendor" {
		return t.addVendor(stub, args)
	}
	if function == "queryStaff" {
		return t.queryStaff(stub, args)
	}
	if function == "queryVendor" {
		return t.queryVendor(stub, args)
	}
	if function == "updateStaff" {
		return t.updateStaff(stub, args)
	}
	if function == "updateVendor" {
		return t.updateVendor(stub, args)
	}
	if function == "addItem" {
		return t.addItem(stub, args)
	}
	if function == "queryItem" {
		return t.queryItem(stub, args)
	}
	if function == "updateItem" {
		return t.updateItem(stub, args)
	}
	if function == "addReceiveItemBUSchema" {
		return t.addReceiveItemBUSchema(stub, args)
	}
	if function == "queryReceiveItemBU" {
		return t.queryReceiveItemBU(stub, args)
	}
	if function == "addReceiveItemFUSchema" {
		return t.addReceiveItemFUSchema(stub, args)
	}
	if function == "queryReceiveItemFU" {
		return t.queryReceiveItemFU(stub, args)
	}
	if function == "updateReceiveItemFU" {
		return t.updateReceiveItemFU(stub, args)
	}
	if function == "updateReceiveItemBU" {
		return t.updateReceiveItemBU(stub, args)
	}
	if function == "addExternalReturnRequestSchema" {
		return t.addExternalReturnRequestSchema(stub, args)
	}
	if function == "queryExternalReturnRequest" {
		return t.queryExternalReturnRequest(stub, args)
	}
	if function == "updateExternalReturnRequestSchema" {
		return t.updateExternalReturnRequestSchema(stub, args)
	}
	if function == "addInternalReturnRequestSchema" {
		return t.addInternalReturnRequestSchema(stub, args)
	}
	if function == "queryInternalReturnRequest" {
		return t.queryInternalReturnRequest(stub, args)
	}
	if function == "updateInternalReturnRequestSchema" {
		return t.updateInternalReturnRequestSchema(stub, args)
	}
	if function == "getHistory" {
		return t.getHistory(stub, args)
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

	err = stub.PutState("purchaseOrder", PurchaseOrderJSONasBytes) //rewrite the marble
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

	err = stub.PutState("PurchaseRequest", PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseRequest")
	return shim.Success(nil)
}

func (t *SmartContract) addItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 30 {
		return shim.Error("Incorrect Number of Arguments. Expecting 30")
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
	if len(args[27]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}
	if len(args[28]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}
	if len(args[29]) <= 0 {
		return shim.Error("27th Argument Must be a Non-Empty String")
	}

	name := args[0]
	description := args[1]
	itemCode := args[2]
	form := args[3]
	drugAllergy := strings.Split(args[4], ",")
	receiptUnit := args[5]
	issueUnit := args[6]
	vendorId := args[7]
	purchasePrice := args[8]
	minimumLevel := args[9]
	maximumLevel := args[10]
	reorderLevel := args[11]
	cls := args[12]
	medClass := args[13]
	subClass := args[14]
	grandSubClass := args[15]
	comments := args[16]
	createdAt := args[17]
	updatedAt := args[18]
	receiptUnitCost := args[19]
	issueUnitCost := args[20]
	scientificName := args[21]
	tradeName := args[22]
	temprature := args[23]
	humidity := args[24]
	lightSensitive := args[25]
	resuableItem := args[26]
	storageCondition := args[27]
	expiration := args[28]
	tax := args[29]

	// ======Check if PurchaseRequest Already exists

	PurchaseRequestAsBytes, err := stub.GetState(itemCode)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseRequestAsBytes != nil {
		return shim.Error("The Inserted itemCode ID already Exists: " + itemCode)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "ItemSchema"
	PurchaseRequest := &ItemSchema{objectType, name, description, itemCode, form, drugAllergy, receiptUnit, issueUnit, vendorId, purchasePrice, minimumLevel, maximumLevel, reorderLevel, cls, medClass, subClass, grandSubClass, comments, createdAt, updatedAt, receiptUnitCost, issueUnitCost, scientificName, tradeName, temprature, humidity, lightSensitive, resuableItem, storageCondition, expiration, tax}
	PurchaseRequestJSONasBytes, err := json.Marshal(PurchaseRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseRequest to State

	err = stub.PutState(itemCode, PurchaseRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ItemInfo", PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseRequest")
	return shim.Success(nil)
}

func (t *SmartContract) queryItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itemCode := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ItemSchema\",\"itemCode\":\"%s\"}}", itemCode)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addReplenishmentRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 32 {
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
	batchArray := args[18]
	tempBatchArray := args[19]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)
	// batchNumber := args[18]
	// expiryDate := args[19]
	// quantity := args[20]
	// tempbatchNumber := args[21]
	// tempexpiryDate := args[22]
	// tempquantity := args[23]
	status := args[20]
	secondStatus := args[21]
	rrB := args[22]
	approvedBy := args[23]
	requesterName := args[24]
	orderType := args[25]
	department := args[26]
	commentNote := args[27]
	inProgressTime := args[28]
	completedTime := args[29]
	createdAt := args[30]
	updatedAt := args[31]

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
		append(ReplenishmentRequest{}.RItem.BatchArray, batch...), append(ReplenishmentRequest{}.RItem.TempBatchArray, tempbatch...)}, status, secondStatus, rrB, approvedBy, requesterName, orderType, department, commentNote, inProgressTime, completedTime, createdAt, updatedAt}
	ReplenishmentRequestJSONasBytes, err := json.Marshal(ReplenishmentRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save replenishmentRequest to State

	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReplenishmentRequest", ReplenishmentRequestJSONasBytes) //rewrite the marble
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

	err = stub.PutState("FU", FunctionalUnitJSONasBytes) //rewrite the marble
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

	if len(args) != 10 {
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
	// if len(args[10]) <= 0 {
	// 	return shim.Error("11th Argument Must be a Non-Empty String")
	// }
	// if len(args[11]) <= 0 {
	// 	return shim.Error("12th Argument Must be a Non-Empty String")
	// }
	// if len(args[12]) <= 0 {
	// 	return shim.Error("13th Argument Must be a Non-Empty String")
	// }
	// if len(args[13]) <= 0 {
	// 	return shim.Error("14th Argument Must be a Non-Empty String")
	// }

	fuId := args[0]
	itemId := args[1]
	qty := args[2]
	maximumLevel := args[3]
	reorderLevel := args[4]
	minimumLevel := args[5]
	createdAt := args[6]
	updatedAt := args[7]
	batchArray := args[8]
	tempBatchArray := args[9]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)

	// batchNumber := args[8]
	// expiryDate := args[9]
	// quantity := args[10]
	// tempbatchNumber := args[11]
	// tempexpiryDate := args[12]
	// tempquantity := args[13]

	// ======Check if FuInventory Already exists

	FuInventoryAsBytes, err := stub.GetState(fuId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if FuInventoryAsBytes != nil {
		return shim.Error("The Inserted FuInventory ID already Exists: " + fuId)
	}

	// ===== Create FuInventory Object and Marshal to JSON

	objectType := "FuInventory"
	FuInventory := &FuInventory{objectType, fuId, itemId, qty, maximumLevel, reorderLevel, minimumLevel, createdAt, updatedAt, append(FuInventory{}.BatchArray, batch...), append(FuInventory{}.TempBatchArray, tempbatch...)}
	FuInventoryJSONasBytes, err := json.Marshal(FuInventory)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save FuInventory to State

	err = stub.PutState(fuId, FuInventoryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("FUInventory", FuInventoryJSONasBytes) //rewrite the marble
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

	if len(args) != 29 {
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
	batchArray := args[27]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	unitPrice := args[28]

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + itemId)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "ReceiveItem"
	ReceiveItem := &ReceiveItem{objectType, itemId, prId, status, currentQty, requestedQty, receivedQty, bonusQty, batchNumber, lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal, discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, returnedQty, append(ReceiveItem{}.BatchArray, batch...), unitPrice}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(itemId, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItem", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReceiveItem")
	return shim.Success(nil)
}

func (t *SmartContract) addReceiveItemFUSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 26 {
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

	itemId := args[0]
	currentQty := args[1]
	requestedQty := args[2]
	receivedQty := args[3]
	bonusQty := args[4]
	batchNumber := args[5]
	lotNumber := args[6]
	expiryDate := args[7]
	unit := args[8]
	discount := args[9]
	unitDiscount := args[10]
	discountAmount := args[11]
	tax := args[12]
	taxAmount := args[13]
	finalUnitPrice := args[14]
	subTotal := args[15]
	discountAmount2 := args[16]
	totalPrice := args[17]
	invoice := args[18]
	dateInvoice := args[19]
	dateReceived := args[20]
	notes := args[21]
	createdAt := args[22]
	updatedAt := args[23]
	replenishmentRequestId := args[24]
	batchArray := args[25]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + itemId)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "ReceiveItemFU"
	ReceiveItem := &ReceiveItemFUSchema{objectType, itemId, currentQty, requestedQty, receivedQty, bonusQty, batchNumber, lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal, discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, replenishmentRequestId, append(ReceiveItemFUSchema{}.BatchArray, batch...)}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(itemId, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItemFU", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReceiveItem")
	return shim.Success(nil)
}

func (t *SmartContract) addReceiveItemBUSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 28 {
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

	itemId := args[0]
	currentQty := args[1]
	requestedQty := args[2]
	receivedQty := args[3]
	bonusQty := args[4]
	batchNumber := args[5]
	lotNumber := args[6]
	expiryDate := args[7]
	unit := args[8]
	discount := args[9]
	unitDiscount := args[10]
	discountAmount := args[11]
	tax := args[12]
	taxAmount := args[13]
	finalUnitPrice := args[14]
	subTotal := args[15]
	discountAmount2 := args[16]
	totalPrice := args[17]
	invoice := args[18]
	dateInvoice := args[19]
	dateReceived := args[20]
	notes := args[21]
	createdAt := args[22]
	updatedAt := args[23]
	replenishmentRequestId := args[24]
	replenishmentRequestItemId := args[25]
	qualityRate := args[26]
	batchArray := args[27]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + itemId)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "ReceiveItemBU"
	ReceiveItem := &ReceiveItemBUSchema{objectType, itemId, currentQty, requestedQty, receivedQty, bonusQty, batchNumber, lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal, discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, replenishmentRequestId, replenishmentRequestItemId, qualityRate, append(ReceiveItemBUSchema{}.BatchArray, batch...)}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(itemId, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItemBU", ReceiveItemJSONasBytes) //rewrite the marble
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

func (t *SmartContract) queryReceiveItemFU(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itemId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ReceiveItemFU\",\"itemId\":\"%s\"}}", itemId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryReceiveItemBU(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itemId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ReceiveItemBU\",\"itemId\":\"%s\"}}", itemId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addWarehouseInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 9 {
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

	itemId := args[0]
	qty := args[1]
	maximumLevel := args[2]
	minimumLevel := args[3]
	reorderLevel := args[4]
	createdAt := args[5]
	updatedAt := args[6]
	batchArray := args[7]
	tempBatchArray := args[8]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)

	// ======Check if WarehouseInventory Already exists

	WarehouseInventoryAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if WarehouseInventoryAsBytes != nil {
		return shim.Error("The Inserted WarehouseInventory ID already Exists: " + itemId)
	}

	// ===== Create WarehouseInventory Object and Marshal to JSON

	objectType := "WarehouseInventory"
	WarehouseInventory := &WarehouseInventory{objectType, itemId, qty, maximumLevel, minimumLevel, reorderLevel, createdAt, updatedAt, append(WarehouseInventory{}.BatchArray, batch...), append(WarehouseInventory{}.TempBatchArray, tempbatch...)}
	WarehouseInventoryJSONasBytes, err := json.Marshal(WarehouseInventory)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save WarehouseInventory to State

	err = stub.PutState(itemId, WarehouseInventoryJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("WareHouse", WarehouseInventoryJSONasBytes) //rewrite
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

func (t *SmartContract) addStaff(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 16 {
		return shim.Error("Incorrect Number of Aruments. Expecting 14")
	}

	fmt.Println("Adding new Sttaff")

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
		return shim.Error("14th Argument Must be a Non-Empty String")
	}

	staffId := args[0]
	staffTypeId := args[1]
	firstName := args[2]
	lastName := args[3]
	designation := args[4]
	contactNumber := args[5]
	identificationNumber := args[6]
	email := args[7]
	password := args[8]
	gender := args[9]
	dob := args[10]
	address := args[11]
	functionalUnit := args[12]
	systemAdminId := args[13]
	status := args[14]
	routes := strings.Split(args[15], ",")

	// ======Check if PurchaseOrder Already exists

	PurchaseOrderAsBytes, err := stub.GetState(staffId)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseOrderAsBytes != nil {
		return shim.Error("The Inserted staffTypeId ID already Exists: " + staffId)
	}

	// ===== Create PurchaseOrder Object and Marshal to JSON

	objectType := "Staff"
	Staff := &Staff{objectType, staffId, staffTypeId, firstName, lastName, designation, contactNumber, identificationNumber, email, password, gender, dob, address, functionalUnit, systemAdminId, status, routes}
	PurchaseOrderJSONasBytes, err := json.Marshal(Staff)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseOrder to State

	err = stub.PutState(staffId, PurchaseOrderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("Staff", PurchaseOrderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseOrder")
	return shim.Success(nil)
}

func (t *SmartContract) queryStaff(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	staffId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Staff\",\"staffId\":\"%s\"}}", staffId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addVendor(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 25 {
		return shim.Error("Incorrect Number of Arguments. Expecting 27")
	}

	fmt.Println("Adding new Vendor")

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

	uuid := args[0]
	vendorNo := args[1]
	englishName := args[2]
	arabicName := args[3]
	telephone1 := args[4]
	telephone2 := args[5]
	contactEmail := args[6]
	address := args[7]
	country := args[8]
	city := args[9]
	zipcode := args[10]
	faxno := args[11]
	taxno := args[12]
	contactPersonName := args[13]
	contactPersonTelephone := args[14]
	contactPersonEmail := args[15]
	paymentTerms := args[16]
	shippingTerms := args[17]
	rating := args[18]
	status := args[19]
	cls := args[20]
	subClass := strings.Split(args[21], ",")
	compliance := args[22]
	createdAt := args[23]
	updatedAt := args[24]

	// ======Check if PurchaseRequest Already exists

	PurchaseRequestAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if PurchaseRequestAsBytes != nil {
		return shim.Error("The Inserted Vendor ID already Exists: " + uuid)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "Vendor"
	Vendor := &Vendor{objectType, uuid, vendorNo, englishName, arabicName, telephone1, telephone2, contactEmail, address, country, city, zipcode, faxno, taxno, contactPersonName, contactPersonTelephone, contactPersonEmail, paymentTerms, shippingTerms, rating, status, cls, subClass, compliance, createdAt, updatedAt}
	PurchaseRequestJSONasBytes, err := json.Marshal(Vendor)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseRequest to State

	err = stub.PutState(uuid, PurchaseRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("Vendor", PurchaseRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved PurchaseRequest")
	return shim.Success(nil)
}

func (t *SmartContract) queryVendor(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	uuid := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Vendor\",\"uuid\":\"%s\"}}", uuid)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addInternalReturnRequestSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 27 {
		return shim.Error("Incorrect Number of Arguments. Expecting 33")
	}

	fmt.Println("Adding new InternalReturnRequestSchema")

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

	returnRequestNo := args[0]
	generatedBy := args[1]
	dateGenerated := args[2]
	expiryDate := args[3]
	to := args[4]
	from := args[5]
	currentQty := args[6]
	returnedQty := args[7]
	itemId := args[8]
	description := args[9]
	fuId := args[10]
	reason := args[11]
	reasonDetail := args[12]
	buId := args[13]
	causedBy := args[14]
	totalDamageCost := args[15]
	date := args[16]
	itemCostPerUnit := args[17]
	status := args[18]
	replenishmentRequestBU := args[19]
	replenishmentRequestFU := args[20]
	approvedBy := args[21]
	commentNote := args[22]
	createdAt := args[23]
	updatedAt := args[24]
	batchNo := args[25]
	batchArray := args[26]

	var batch []ReturnBatch
	json.Unmarshal([]byte(batchArray), &batch)

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(returnRequestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + returnRequestNo)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "InternalReturnRequest"
	object := "DamageReport"
	ReceiveItem := &InternalReturnRequestSchema{objectType, returnRequestNo, generatedBy, dateGenerated, expiryDate, to, from, currentQty, returnedQty, itemId, description, fuId, reason, reasonDetail, buId, DamageReport{object, causedBy, totalDamageCost, date, itemCostPerUnit}, status, replenishmentRequestBU, replenishmentRequestFU, approvedBy, commentNote, createdAt, updatedAt, batchNo, append(InternalReturnRequestSchema{}.ReturnBatchArray, batch...)}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(returnRequestNo, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("InternalReturnRequest", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReceiveItem")
	return shim.Success(nil)
}

func (t *SmartContract) addExternalReturnRequestSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 22 {
		return shim.Error("Incorrect Number of Arguments. Expecting 33")
	}

	fmt.Println("Adding new ExternalReturnRequestSchema")

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

	returnRequestNo := args[0]
	generatedBy := args[1]
	generated := args[2]
	dateGenerated := args[3]
	expiryDate := args[4]
	returnedQty := args[5]
	itemId := args[6]
	prId := args[7]
	description := args[8]
	reason := args[9]
	reasonDetail := args[10]
	causedBy := args[11]
	totalDamageCost := args[12]
	date := args[13]
	itemCostPerUnit := args[14]
	status := args[15]
	approvedBy := args[16]
	commentNote := args[17]
	inProgressTime := args[18]
	createdAt := args[19]
	updatedAt := args[20]
	batchArray := args[21]

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	// ======Check if ReceiveItem Already exists

	ReceiveItemAsBytes, err := stub.GetState(returnRequestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if ReceiveItemAsBytes != nil {
		return shim.Error("The Inserted ReceiveItem ID already Exists: " + returnRequestNo)
	}

	// ===== Create ReceiveItem Object and Marshal to JSON

	objectType := "ExternalReturnRequest"
	object := "DamageReport"
	ReceiveItem := &ExternalReturnRequestSchema{objectType, returnRequestNo, generatedBy, generated, dateGenerated, expiryDate, returnedQty, itemId, prId, description, reason, reasonDetail, DamageReport{object, causedBy, totalDamageCost, date, itemCostPerUnit}, status, approvedBy, commentNote, inProgressTime, createdAt, updatedAt, append(ExternalReturnRequestSchema{}.BatchArray, batch...)}
	ReceiveItemJSONasBytes, err := json.Marshal(ReceiveItem)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save ReceiveItem to State

	err = stub.PutState(returnRequestNo, ReceiveItemJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ExternalReturnRequest", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved ReceiveItem")
	return shim.Success(nil)
}

func (t *SmartContract) queryInternalReturnRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	returnRequestNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"InternalReturnRequest\",\"returnRequestNo\":\"%s\"}}", returnRequestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryExternalReturnRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	returnRequestNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"ExternalReturnRequest\",\"returnRequestNo\":\"%s\"}}", returnRequestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) updateWarehouseInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 9 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	itemId := args[0]
	qty := args[1]
	maximumLevel := args[2]
	minimumLevel := args[3]
	reorderLevel := args[4]
	createdAt := args[5]
	updatedAt := args[6]
	batchArray := args[7]
	tempBatchArray := args[8]

	fmt.Println("- start  ", itemId, qty, maximumLevel, minimumLevel, reorderLevel, createdAt, updatedAt, batchArray, tempBatchArray)

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

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)

	responseToUpdate.Qty = qty
	responseToUpdate.MaximumLevel = maximumLevel
	responseToUpdate.MinimumLevel = minimumLevel
	responseToUpdate.ReorderLevel = reorderLevel
	responseToUpdate.CreatedAt = createdAt
	responseToUpdate.UpdatedAt = updatedAt

	responseToUpdate.BatchArray = append(responseToUpdate.BatchArray, batch...)
	responseToUpdate.TempBatchArray = append(responseToUpdate.TempBatchArray, tempbatch...) //change the status

	responseJSONasBytes, _ := json.Marshal(responseToUpdate)
	err = stub.PutState(itemId, responseJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("WareHouse", responseJSONasBytes) //rewrite
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

func (t *SmartContract) updatePurchaseOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 14 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
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
	fmt.Println("- start  ", purchaseOrderNo, purchaseRequestId, date, generated, generatedBy, commentNotes, approvedBy,
		vendorId, status, committeeStatus, inProgressTime, createdAt, sentAt, updatedAt)

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

	PurchaseOrderToUpdate.PurchaseRequestId = purchaseRequestId //change the status
	PurchaseOrderToUpdate.Date = date                           //change the status
	PurchaseOrderToUpdate.Generated = generated                 //change the status
	PurchaseOrderToUpdate.GeneratedBy = generatedBy             //change the status
	PurchaseOrderToUpdate.CommentNotes = commentNotes           //change the status
	PurchaseOrderToUpdate.ApprovedBy = approvedBy               //change the status
	PurchaseOrderToUpdate.VendorId = vendorId                   //change the status
	PurchaseOrderToUpdate.Status = status                       //change the status
	PurchaseOrderToUpdate.CommitteeStatus = committeeStatus     //change the status
	PurchaseOrderToUpdate.InProgressTime = inProgressTime       //change the status
	PurchaseOrderToUpdate.CreatedAt = createdAt                 //change the status
	PurchaseOrderToUpdate.SentAt = sentAt                       //change the status
	PurchaseOrderToUpdate.UpdatedAt = updatedAt                 //change the status

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(purchaseOrderNo, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("purchaseOrder", PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePurchaseRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 27 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
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
	fmt.Println("- start  ", requestNo, generatedBy, status, committeeStatus, availability, reason, vendorId, rr,
		itemId, currQty, reqQty, comments, name, description, itemCode, istatus, secondStatus, requesterName,
		rejectionReason, department, commentNotes, orderType, generated, approvedBy, inProgressTime, createdAt,
		updatedAt)

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
	PurchaseRequestToUpdate.GeneratedBy = generatedBy
	PurchaseRequestToUpdate.Status = status
	PurchaseRequestToUpdate.CommitteeStatus = committeeStatus
	PurchaseRequestToUpdate.Availability = availability
	PurchaseRequestToUpdate.Reason = reason
	PurchaseRequestToUpdate.VendorId = vendorId
	PurchaseRequestToUpdate.Rr = rr
	PurchaseRequestToUpdate.Item.ItemId = itemId
	PurchaseRequestToUpdate.Item.CurrQty = currQty
	PurchaseRequestToUpdate.Item.ReqQty = reqQty
	PurchaseRequestToUpdate.Item.Comments = comments
	PurchaseRequestToUpdate.Item.Name = name
	PurchaseRequestToUpdate.Item.Description = description
	PurchaseRequestToUpdate.Item.ItemCode = itemCode
	PurchaseRequestToUpdate.Item.IStatus = istatus
	PurchaseRequestToUpdate.Item.SecondStatus = secondStatus
	PurchaseRequestToUpdate.RequesterName = requesterName
	PurchaseRequestToUpdate.RejectionReason = rejectionReason
	PurchaseRequestToUpdate.Department = department
	PurchaseRequestToUpdate.CommentNotes = commentNotes
	PurchaseRequestToUpdate.OrderType = orderType
	PurchaseRequestToUpdate.Generated = generated
	PurchaseRequestToUpdate.ApprovedBy = approvedBy
	PurchaseRequestToUpdate.InProgressTime = inProgressTime
	PurchaseRequestToUpdate.CreatedAt = createdAt
	PurchaseRequestToUpdate.UpdatedAt = updatedAt

	PurchaseRequestJSONasBytes, _ := json.Marshal(PurchaseRequestToUpdate)
	err = stub.PutState(requestNo, PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("PurchaseRequest", PurchaseRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReplenishmentRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 32 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
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
	batchArray := args[18]
	tempBatchArray := args[19]
	status := args[20]
	secondStatus := args[21]
	rrB := args[22]
	approvedBy := args[23]
	requesterName := args[24]
	orderType := args[25]
	department := args[26]
	commentNote := args[27]
	inProgressTime := args[28]
	completedTime := args[29]
	createdAt := args[30]
	updatedAt := args[31]
	fmt.Println("- start  ", requestNo, generated, generatedBy, dateGenerated, reason, fuId, to, from, comments,
		itemId, currentQty, requestedQty, recieptUnit, issueUnit, fuItemCost, description, rstatus, rsecondStatus,
		batchArray, tempBatchArray, status, secondStatus,
		rrB, approvedBy, requesterName, orderType, department, commentNote, inProgressTime, completedTime, createdAt,
		updatedAt)

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

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)

	ReplenishmentRequestToUpdate.Generated = generated
	ReplenishmentRequestToUpdate.GeneratedBy = generatedBy
	ReplenishmentRequestToUpdate.DateGenerated = dateGenerated
	ReplenishmentRequestToUpdate.Reason = reason
	ReplenishmentRequestToUpdate.FuId = fuId
	ReplenishmentRequestToUpdate.To = to
	ReplenishmentRequestToUpdate.From = from
	ReplenishmentRequestToUpdate.Comments = comments
	ReplenishmentRequestToUpdate.RItem.ItemId = itemId
	ReplenishmentRequestToUpdate.RItem.CurrentQty = currentQty
	ReplenishmentRequestToUpdate.RItem.RequestedQty = requestedQty
	ReplenishmentRequestToUpdate.RItem.RecieptUnit = recieptUnit
	ReplenishmentRequestToUpdate.RItem.IssueUnit = issueUnit
	ReplenishmentRequestToUpdate.RItem.FuItemCost = fuItemCost
	ReplenishmentRequestToUpdate.RItem.Description = description
	ReplenishmentRequestToUpdate.RItem.RStatus = rstatus
	ReplenishmentRequestToUpdate.RItem.RSecondStatus = rsecondStatus
	ReplenishmentRequestToUpdate.RItem.BatchArray = append(ReplenishmentRequestToUpdate.RItem.BatchArray, batch...)
	ReplenishmentRequestToUpdate.RItem.TempBatchArray = append(ReplenishmentRequestToUpdate.RItem.TempBatchArray, tempbatch...)
	ReplenishmentRequestToUpdate.Status = status             //change the status
	ReplenishmentRequestToUpdate.SecondStatus = secondStatus //change the status
	ReplenishmentRequestToUpdate.RrB = rrB
	ReplenishmentRequestToUpdate.ApprovedBy = approvedBy
	ReplenishmentRequestToUpdate.RequesterName = requesterName
	ReplenishmentRequestToUpdate.OrderType = orderType
	ReplenishmentRequestToUpdate.Department = department
	ReplenishmentRequestToUpdate.CommentNote = commentNote
	ReplenishmentRequestToUpdate.InProgressTime = inProgressTime
	ReplenishmentRequestToUpdate.CompletedTime = completedTime
	ReplenishmentRequestToUpdate.CreatedAt = createdAt
	ReplenishmentRequestToUpdate.UpdatedAt = updatedAt

	ReplenishmentRequestJSONasBytes, _ := json.Marshal(ReplenishmentRequestToUpdate)
	err = stub.PutState(requestNo, ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReplenishmentRequest", ReplenishmentRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateFunctionalUnit(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 9 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
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
	fmt.Println("- start  ", uuid, fuName, description, fuHead, status, buId, fuLogId, createdAt, updatedAt)

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
	FunctionalUnitToUpdate.FuName = fuName           //change the status
	FunctionalUnitToUpdate.Description = description //change the status
	FunctionalUnitToUpdate.FuHead = fuHead           //change the status
	FunctionalUnitToUpdate.Status = status           //change the status
	FunctionalUnitToUpdate.BuId = buId               //change the status
	FunctionalUnitToUpdate.FuLogId = fuLogId         //change the status
	FunctionalUnitToUpdate.CreatedAt = createdAt     //change the status
	FunctionalUnitToUpdate.UpdatedAt = updatedAt     //change the status

	FunctionalUnitJSONasBytes, _ := json.Marshal(FunctionalUnitToUpdate)
	err = stub.PutState(uuid, FunctionalUnitJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("FU", FunctionalUnitJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateFuInventory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 10 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	fuId := args[0]
	itemId := args[1]
	qty := args[2]
	maximumLevel := args[3]
	reorderLevel := args[4]
	minimumLevel := args[5]
	createdAt := args[6]
	updatedAt := args[7]
	batchArray := args[8]
	tempBatchArray := args[9]

	fmt.Println("- start  ", fuId, itemId, qty, maximumLevel, reorderLevel, minimumLevel, createdAt, updatedAt, batchArray, tempBatchArray)

	FunctionalUnitAsBytes, err := stub.GetState(fuId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if FunctionalUnitAsBytes == nil {
		return shim.Error("FUID does not exist")
	}

	FunctionalUnitToUpdate := FuInventory{}
	err = json.Unmarshal(FunctionalUnitAsBytes, &FunctionalUnitToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	var tempbatch []TempBatch
	json.Unmarshal([]byte(tempBatchArray), &tempbatch)

	FunctionalUnitToUpdate.ItemId = itemId             //change the status
	FunctionalUnitToUpdate.Qty = qty                   //change the status
	FunctionalUnitToUpdate.MaximumLevel = maximumLevel //change the status
	FunctionalUnitToUpdate.ReorderLevel = reorderLevel //change the status
	FunctionalUnitToUpdate.MinimumLevel = minimumLevel //change the status
	FunctionalUnitToUpdate.CreatedAt = createdAt       //change the status
	FunctionalUnitToUpdate.UpdatedAt = updatedAt       //change the status
	FunctionalUnitToUpdate.BatchArray = append(FunctionalUnitToUpdate.BatchArray, batch...)
	FunctionalUnitToUpdate.TempBatchArray = append(FunctionalUnitToUpdate.TempBatchArray, tempbatch...)

	FunctionalUnitJSONasBytes, _ := json.Marshal(FunctionalUnitToUpdate)
	err = stub.PutState(fuId, FunctionalUnitJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("FUInventory", FunctionalUnitJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReceiveItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 29 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
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
	batchArray := args[27]
	unitPrice := args[28]

	fmt.Println("- start   ", itemId, prId, status, currentQty, requestedQty, receivedQty, bonusQty, batchNumber,
		lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal,
		discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, returnedQty,
		batchArray, unitPrice)

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

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	ReceiveItemToUpdate.PrId = prId
	ReceiveItemToUpdate.Status = status
	ReceiveItemToUpdate.CurrentQty = currentQty
	ReceiveItemToUpdate.RequestedQty = requestedQty
	ReceiveItemToUpdate.ReceivedQty = receivedQty
	ReceiveItemToUpdate.BonusQty = bonusQty
	ReceiveItemToUpdate.BatchNumber = batchNumber
	ReceiveItemToUpdate.LotNumber = lotNumber
	ReceiveItemToUpdate.ExpiryDate = expiryDate
	ReceiveItemToUpdate.Unit = unit
	ReceiveItemToUpdate.Discount = discount
	ReceiveItemToUpdate.UnitDiscount = unitDiscount
	ReceiveItemToUpdate.DiscountAmount = discountAmount
	ReceiveItemToUpdate.Tax = tax
	ReceiveItemToUpdate.TaxAmount = taxAmount
	ReceiveItemToUpdate.FinalUnitPrice = finalUnitPrice
	ReceiveItemToUpdate.SubTotal = subTotal
	ReceiveItemToUpdate.DiscountAmount2 = discountAmount2
	ReceiveItemToUpdate.TotalPrice = totalPrice
	ReceiveItemToUpdate.Invoice = invoice
	ReceiveItemToUpdate.DateInvoice = dateInvoice
	ReceiveItemToUpdate.DateReceived = dateReceived
	ReceiveItemToUpdate.Notes = notes
	ReceiveItemToUpdate.CreatedAt = createdAt
	ReceiveItemToUpdate.UpdatedAt = updatedAt
	ReceiveItemToUpdate.ReturnedQty = returnedQty
	ReceiveItemToUpdate.BatchArray = append(ReceiveItemToUpdate.BatchArray, batch...)
	ReceiveItemToUpdate.UnitPrice = unitPrice

	ReceiveItemJSONasBytes, _ := json.Marshal(ReceiveItemToUpdate)
	err = stub.PutState(itemId, ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItem", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReceiveItemFU(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 26 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	itemId := args[0]
	currentQty := args[1]
	requestedQty := args[2]
	receivedQty := args[3]
	bonusQty := args[4]
	batchNumber := args[5]
	lotNumber := args[6]
	expiryDate := args[7]
	unit := args[8]
	discount := args[9]
	unitDiscount := args[10]
	discountAmount := args[11]
	tax := args[12]
	taxAmount := args[13]
	finalUnitPrice := args[14]
	subTotal := args[15]
	discountAmount2 := args[16]
	totalPrice := args[17]
	invoice := args[18]
	dateInvoice := args[19]
	dateReceived := args[20]
	notes := args[21]
	createdAt := args[22]
	updatedAt := args[23]
	replenishmentRequestId := args[24]
	batchArray := args[25]

	fmt.Println("- start   ", itemId, currentQty, requestedQty, receivedQty, bonusQty, batchNumber,
		lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal,
		discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, replenishmentRequestId,
		batchArray)

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReceiveItemAsBytes == nil {
		return shim.Error("ReceiveItem does not exist")
	}

	ReceiveItemToUpdate := ReceiveItemFUSchema{}
	err = json.Unmarshal(ReceiveItemAsBytes, &ReceiveItemToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	ReceiveItemToUpdate.CurrentQty = currentQty
	ReceiveItemToUpdate.RequestedQty = requestedQty
	ReceiveItemToUpdate.ReceivedQty = receivedQty
	ReceiveItemToUpdate.BonusQty = bonusQty
	ReceiveItemToUpdate.BatchNumber = batchNumber
	ReceiveItemToUpdate.LotNumber = lotNumber
	ReceiveItemToUpdate.ExpiryDate = expiryDate
	ReceiveItemToUpdate.Unit = unit
	ReceiveItemToUpdate.Discount = discount
	ReceiveItemToUpdate.UnitDiscount = unitDiscount
	ReceiveItemToUpdate.DiscountAmount = discountAmount
	ReceiveItemToUpdate.Tax = tax
	ReceiveItemToUpdate.TaxAmount = taxAmount
	ReceiveItemToUpdate.FinalUnitPrice = finalUnitPrice
	ReceiveItemToUpdate.SubTotal = subTotal
	ReceiveItemToUpdate.DiscountAmount2 = discountAmount2
	ReceiveItemToUpdate.TotalPrice = totalPrice
	ReceiveItemToUpdate.Invoice = invoice
	ReceiveItemToUpdate.DateInvoice = dateInvoice
	ReceiveItemToUpdate.DateReceived = dateReceived
	ReceiveItemToUpdate.Notes = notes
	ReceiveItemToUpdate.CreatedAt = createdAt
	ReceiveItemToUpdate.UpdatedAt = updatedAt
	ReceiveItemToUpdate.ReplenishmentRequestId = replenishmentRequestId
	ReceiveItemToUpdate.BatchArray = append(ReceiveItemToUpdate.BatchArray, batch...)
	ReceiveItemJSONasBytes, _ := json.Marshal(ReceiveItemToUpdate)
	err = stub.PutState(itemId, ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItemFU", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateReceiveItemBU(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 28 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	itemId := args[0]
	currentQty := args[1]
	requestedQty := args[2]
	receivedQty := args[3]
	bonusQty := args[4]
	batchNumber := args[5]
	lotNumber := args[6]
	expiryDate := args[7]
	unit := args[8]
	discount := args[9]
	unitDiscount := args[10]
	discountAmount := args[11]
	tax := args[12]
	taxAmount := args[13]
	finalUnitPrice := args[14]
	subTotal := args[15]
	discountAmount2 := args[16]
	totalPrice := args[17]
	invoice := args[18]
	dateInvoice := args[19]
	dateReceived := args[20]
	notes := args[21]
	createdAt := args[22]
	updatedAt := args[23]
	replenishmentRequestId := args[24]
	replenishmentRequestItemId := args[25]
	qualityRate := args[26]
	batchArray := args[27]

	fmt.Println("- start   ", itemId, currentQty, requestedQty, receivedQty, bonusQty, batchNumber,
		lotNumber, expiryDate, unit, discount, unitDiscount, discountAmount, tax, taxAmount, finalUnitPrice, subTotal,
		discountAmount2, totalPrice, invoice, dateInvoice, dateReceived, notes, createdAt, updatedAt, replenishmentRequestId,
		replenishmentRequestItemId, qualityRate, batchArray)

	ReceiveItemAsBytes, err := stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ReceiveItemAsBytes == nil {
		return shim.Error("ReceiveItem does not exist")
	}

	ReceiveItemToUpdate := ReceiveItemBUSchema{}
	err = json.Unmarshal(ReceiveItemAsBytes, &ReceiveItemToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	ReceiveItemToUpdate.CurrentQty = currentQty
	ReceiveItemToUpdate.RequestedQty = requestedQty
	ReceiveItemToUpdate.ReceivedQty = receivedQty
	ReceiveItemToUpdate.BonusQty = bonusQty
	ReceiveItemToUpdate.BatchNumber = batchNumber
	ReceiveItemToUpdate.LotNumber = lotNumber
	ReceiveItemToUpdate.ExpiryDate = expiryDate
	ReceiveItemToUpdate.Unit = unit
	ReceiveItemToUpdate.Discount = discount
	ReceiveItemToUpdate.UnitDiscount = unitDiscount
	ReceiveItemToUpdate.DiscountAmount = discountAmount
	ReceiveItemToUpdate.Tax = tax
	ReceiveItemToUpdate.TaxAmount = taxAmount
	ReceiveItemToUpdate.FinalUnitPrice = finalUnitPrice
	ReceiveItemToUpdate.SubTotal = subTotal
	ReceiveItemToUpdate.DiscountAmount2 = discountAmount2
	ReceiveItemToUpdate.TotalPrice = totalPrice
	ReceiveItemToUpdate.Invoice = invoice
	ReceiveItemToUpdate.DateInvoice = dateInvoice
	ReceiveItemToUpdate.DateReceived = dateReceived
	ReceiveItemToUpdate.Notes = notes
	ReceiveItemToUpdate.CreatedAt = createdAt
	ReceiveItemToUpdate.UpdatedAt = updatedAt
	ReceiveItemToUpdate.ReplenishmentRequestId = replenishmentRequestId
	ReceiveItemToUpdate.ReplenishmentRequestItemId = replenishmentRequestItemId
	ReceiveItemToUpdate.QualityRate = qualityRate
	ReceiveItemToUpdate.BatchArray = append(ReceiveItemToUpdate.BatchArray, batch...)
	ReceiveItemJSONasBytes, _ := json.Marshal(ReceiveItemToUpdate)
	err = stub.PutState(itemId, ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ReceiveItemBU", ReceiveItemJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) getHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	info := args[0]

	fmt.Printf("- start getHistory: %s\n", info)

	resultsIterator, err := stub.GetHistoryForKey(info)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
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
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON )
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (t *SmartContract) updateStaff(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 15 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	staffId := args[0]
	staffTypeId := args[1]
	firstName := args[2]
	lastName := args[3]
	designation := args[4]
	contactNumber := args[5]
	identificationNumber := args[6]
	email := args[7]
	password := args[8]
	gender := args[9]
	dob := args[10]
	address := args[11]
	functionalUnit := args[12]
	systemAdminId := args[13]
	status := args[14]
	routes := strings.Split(args[15], ",")
	fmt.Println("- start  ", staffId, staffTypeId, firstName, lastName, designation, contactNumber, identificationNumber, email, password, gender, dob, address, functionalUnit, systemAdminId, status, routes)

	PurchaseOrderAsBytes, err := stub.GetState(staffId)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("staffId does not exist")
	}

	PurchaseOrderToUpdate := Staff{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	PurchaseOrderToUpdate.StaffTypeId = staffTypeId                   //change the status
	PurchaseOrderToUpdate.FirstName = firstName                       //change the status
	PurchaseOrderToUpdate.LastName = lastName                         //change the status
	PurchaseOrderToUpdate.Designation = designation                   //change the status
	PurchaseOrderToUpdate.ContactNumber = contactNumber               //change the status
	PurchaseOrderToUpdate.IdentificationNumber = identificationNumber //change the status
	PurchaseOrderToUpdate.Email = email                               //change the status
	PurchaseOrderToUpdate.Password = password                         //change the status
	PurchaseOrderToUpdate.Gender = gender                             //change the status
	PurchaseOrderToUpdate.Dob = dob                                   //change the status
	PurchaseOrderToUpdate.Address = address                           //change the status
	PurchaseOrderToUpdate.FunctionalUnit = functionalUnit             //change the status
	PurchaseOrderToUpdate.SystemAdminId = systemAdminId               //change the status
	PurchaseOrderToUpdate.Status = status
	PurchaseOrderToUpdate.Routes = append(PurchaseOrderToUpdate.Routes, routes...) //change the status

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(staffId, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("Staff", PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateVendor(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 25 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	uuid := args[0]
	vendorNo := args[1]
	englishName := args[2]
	arabicName := args[3]
	telephone1 := args[4]
	telephone2 := args[5]
	contactEmail := args[6]
	address := args[7]
	country := args[8]
	city := args[9]
	zipcode := args[10]
	faxno := args[11]
	taxno := args[12]
	contactPersonName := args[13]
	contactPersonTelephone := args[14]
	contactPersonEmail := args[15]
	paymentTerms := args[16]
	shippingTerms := args[17]
	rating := args[18]
	status := args[19]
	cls := args[20]
	subClass := strings.Split(args[21], ",")
	compliance := args[22]
	createdAt := args[23]
	updatedAt := args[24]
	fmt.Println("- start  ", uuid, vendorNo, englishName, arabicName, telephone1, telephone2, contactEmail, address,
		country, city, zipcode, faxno, taxno, contactPersonName, contactPersonTelephone, contactPersonEmail,
		paymentTerms, shippingTerms, rating, status, cls, subClass, compliance, createdAt, updatedAt)

	PurchaseOrderAsBytes, err := stub.GetState(uuid)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("Vendor does not exist")
	}

	PurchaseOrderToUpdate := Vendor{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	PurchaseOrderToUpdate.VendorNo = vendorNo         //change the status
	PurchaseOrderToUpdate.EnglishName = englishName   //change the status
	PurchaseOrderToUpdate.ArabicName = arabicName     //change the status
	PurchaseOrderToUpdate.Telephone1 = telephone1     //change the status
	PurchaseOrderToUpdate.Telephone2 = telephone2     //change the status
	PurchaseOrderToUpdate.ContactEmail = contactEmail //change the status
	PurchaseOrderToUpdate.Address = address           //change the status
	PurchaseOrderToUpdate.Country = country           //change the status
	PurchaseOrderToUpdate.City = city                 //change the status
	PurchaseOrderToUpdate.Zipcode = zipcode
	PurchaseOrderToUpdate.Faxno = faxno //change the status
	PurchaseOrderToUpdate.Taxno = taxno //change the status
	PurchaseOrderToUpdate.ContactPersonName = contactPersonName
	PurchaseOrderToUpdate.ContactPersonTelephone = contactPersonTelephone
	PurchaseOrderToUpdate.ContactPersonEmail = contactPersonEmail
	PurchaseOrderToUpdate.PaymentTerms = paymentTerms
	PurchaseOrderToUpdate.ShippingTerms = shippingTerms
	PurchaseOrderToUpdate.Rating = rating
	PurchaseOrderToUpdate.Status = status
	PurchaseOrderToUpdate.Cls = cls
	PurchaseOrderToUpdate.SubClass = append(PurchaseOrderToUpdate.SubClass, subClass...) //change the status
	PurchaseOrderToUpdate.Compliance = compliance
	PurchaseOrderToUpdate.CreatedAt = createdAt
	PurchaseOrderToUpdate.UpdatedAt = updatedAt

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(uuid, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("Vendor", PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateItem(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 30 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	name := args[0]
	description := args[1]
	itemCode := args[2]
	form := args[3]
	drugAllergy := strings.Split(args[4], ",")
	receiptUnit := args[5]
	issueUnit := args[6]
	vendorId := args[7]
	purchasePrice := args[8]
	minimumLevel := args[9]
	maximumLevel := args[10]
	reorderLevel := args[11]
	cls := args[12]
	medClass := args[13]
	subClass := args[14]
	grandSubClass := args[15]
	comments := args[16]
	createdAt := args[17]
	updatedAt := args[18]
	receiptUnitCost := args[19]
	issueUnitCost := args[20]
	scientificName := args[21]
	tradeName := args[22]
	temprature := args[23]
	humidity := args[24]
	lightSensitive := args[25]
	resuableItem := args[26]
	storageCondition := args[27]
	expiration := args[28]
	tax := args[29]
	fmt.Println("- start  ", name, description, itemCode, form, drugAllergy, receiptUnit, issueUnit, vendorId, purchasePrice,
		minimumLevel, maximumLevel, reorderLevel, cls, medClass, subClass, grandSubClass, comments, createdAt, updatedAt,
		receiptUnitCost, issueUnitCost, scientificName, tradeName, temprature, humidity, lightSensitive, resuableItem,
		storageCondition, expiration, tax)

	PurchaseOrderAsBytes, err := stub.GetState(itemCode)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("item does not exist")
	}

	PurchaseOrderToUpdate := ItemSchema{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	PurchaseOrderToUpdate.Name = name               //change the status
	PurchaseOrderToUpdate.Description = description //change the status
	PurchaseOrderToUpdate.Form = form               //change the status
	PurchaseOrderToUpdate.DrugAllergy = append(PurchaseOrderToUpdate.DrugAllergy, drugAllergy...)
	PurchaseOrderToUpdate.ReceiptUnit = receiptUnit     //change the status
	PurchaseOrderToUpdate.IssueUnit = issueUnit         //change the status
	PurchaseOrderToUpdate.VendorId = vendorId           //change the status
	PurchaseOrderToUpdate.PurchasePrice = purchasePrice //change the status
	PurchaseOrderToUpdate.MinimumLevel = minimumLevel   //change the status
	PurchaseOrderToUpdate.MaximumLevel = maximumLevel
	PurchaseOrderToUpdate.ReorderLevel = reorderLevel //change the status
	PurchaseOrderToUpdate.Cls = cls                   //change the status
	PurchaseOrderToUpdate.MedClass = medClass
	PurchaseOrderToUpdate.SubClass = subClass
	PurchaseOrderToUpdate.GrandSubClass = grandSubClass
	PurchaseOrderToUpdate.Comments = comments
	PurchaseOrderToUpdate.CreatedAt = createdAt
	PurchaseOrderToUpdate.UpdatedAt = updatedAt
	PurchaseOrderToUpdate.ReceiptUnitCost = receiptUnitCost
	PurchaseOrderToUpdate.IssueUnitCost = issueUnitCost
	PurchaseOrderToUpdate.ScientificName = scientificName //change the status
	PurchaseOrderToUpdate.TradeName = tradeName
	PurchaseOrderToUpdate.Temprature = temprature
	PurchaseOrderToUpdate.Humidity = humidity
	PurchaseOrderToUpdate.LightSensitive = lightSensitive
	PurchaseOrderToUpdate.ResuableItem = resuableItem
	PurchaseOrderToUpdate.StorageCondition = storageCondition
	PurchaseOrderToUpdate.Expiration = expiration
	PurchaseOrderToUpdate.Tax = tax

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(itemCode, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ItemInfo", PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateInternalReturnRequestSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 27 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	returnRequestNo := args[0]
	generatedBy := args[1]
	dateGenerated := args[2]
	expiryDate := args[3]
	to := args[4]
	from := args[5]
	currentQty := args[6]
	returnedQty := args[7]
	itemId := args[8]
	description := args[9]
	fuId := args[10]
	reason := args[11]
	reasonDetail := args[12]
	buId := args[13]
	causedBy := args[14]
	totalDamageCost := args[15]
	date := args[16]
	itemCostPerUnit := args[17]
	status := args[18]
	replenishmentRequestBU := args[19]
	replenishmentRequestFU := args[20]
	approvedBy := args[21]
	commentNote := args[22]
	createdAt := args[23]
	updatedAt := args[24]
	batchNo := args[25]
	returnBatchArray := args[26]
	fmt.Println("- start  ", returnRequestNo, generatedBy, dateGenerated, expiryDate, to, from, currentQty, returnedQty,
		itemId, description, fuId, reason, reasonDetail, buId, causedBy, totalDamageCost, date, itemCostPerUnit, status,
		replenishmentRequestBU, replenishmentRequestFU, approvedBy, commentNote, createdAt, updatedAt, batchNo, returnBatchArray)

	PurchaseOrderAsBytes, err := stub.GetState(returnRequestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("item does not exist")
	}

	PurchaseOrderToUpdate := InternalReturnRequestSchema{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	var batch []ReturnBatch
	json.Unmarshal([]byte(returnBatchArray), &batch)

	PurchaseOrderToUpdate.GeneratedBy = generatedBy     //change the status
	PurchaseOrderToUpdate.DateGenerated = dateGenerated //change the status
	PurchaseOrderToUpdate.ExpiryDate = expiryDate
	PurchaseOrderToUpdate.To = to                   //change the status
	PurchaseOrderToUpdate.From = from               //change the status
	PurchaseOrderToUpdate.CurrentQty = currentQty   //change the status
	PurchaseOrderToUpdate.ReturnedQty = returnedQty //change the status
	PurchaseOrderToUpdate.ItemId = itemId           //change the status
	PurchaseOrderToUpdate.Description = description
	PurchaseOrderToUpdate.FuId = fuId     //change the status
	PurchaseOrderToUpdate.Reason = reason //change the status
	PurchaseOrderToUpdate.ReasonDetail = reasonDetail
	PurchaseOrderToUpdate.BuId = buId
	PurchaseOrderToUpdate.DamageReport.CausedBy = causedBy
	PurchaseOrderToUpdate.DamageReport.TotalDamageCost = totalDamageCost
	PurchaseOrderToUpdate.DamageReport.Date = date
	PurchaseOrderToUpdate.DamageReport.ItemCostPerUnit = itemCostPerUnit
	PurchaseOrderToUpdate.Status = status
	PurchaseOrderToUpdate.ReplenishmentRequestBU = replenishmentRequestBU //change the status
	PurchaseOrderToUpdate.ReplenishmentRequestFU = replenishmentRequestFU
	PurchaseOrderToUpdate.ApprovedBy = approvedBy
	PurchaseOrderToUpdate.CommentNote = commentNote
	PurchaseOrderToUpdate.CreatedAt = createdAt
	PurchaseOrderToUpdate.UpdatedAt = updatedAt
	PurchaseOrderToUpdate.BatchNo = batchNo
	PurchaseOrderToUpdate.ReturnBatchArray = append(PurchaseOrderToUpdate.ReturnBatchArray, batch...)

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(returnRequestNo, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("InternalReturnRequest", PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateExternalReturnRequestSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 22 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}

	returnRequestNo := args[0]
	generatedBy := args[1]
	generated := args[2]
	dateGenerated := args[3]
	expiryDate := args[4]
	returnedQty := args[5]
	itemId := args[6]
	prId := args[7]
	description := args[8]
	reason := args[9]
	reasonDetail := args[10]
	causedBy := args[11]
	totalDamageCost := args[12]
	date := args[13]
	itemCostPerUnit := args[14]
	status := args[15]
	approvedBy := args[16]
	commentNote := args[17]
	inProgressTime := args[18]
	createdAt := args[19]
	updatedAt := args[20]
	batchArray := args[21]
	fmt.Println("- start  ", returnRequestNo, generatedBy, generated, dateGenerated, expiryDate, returnedQty,
		itemId, prId, description, reason, reasonDetail, causedBy, totalDamageCost, date, itemCostPerUnit, status,
		approvedBy, commentNote, inProgressTime, createdAt, updatedAt, batchArray)

	PurchaseOrderAsBytes, err := stub.GetState(returnRequestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PurchaseOrderAsBytes == nil {
		return shim.Error("item does not exist")
	}

	PurchaseOrderToUpdate := ExternalReturnRequestSchema{}
	err = json.Unmarshal(PurchaseOrderAsBytes, &PurchaseOrderToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	var batch []Batch
	json.Unmarshal([]byte(batchArray), &batch)

	PurchaseOrderToUpdate.GeneratedBy = generatedBy     //change the status
	PurchaseOrderToUpdate.Generated = generated         //change the status
	PurchaseOrderToUpdate.DateGenerated = dateGenerated //change the status
	PurchaseOrderToUpdate.ExpiryDate = expiryDate
	PurchaseOrderToUpdate.ReturnedQty = returnedQty //change the status
	PurchaseOrderToUpdate.ItemId = itemId
	PurchaseOrderToUpdate.PrId = prId //change the status
	PurchaseOrderToUpdate.Description = description
	PurchaseOrderToUpdate.Reason = reason //change the status
	PurchaseOrderToUpdate.ReasonDetail = reasonDetail
	PurchaseOrderToUpdate.DamageReport.CausedBy = causedBy
	PurchaseOrderToUpdate.DamageReport.TotalDamageCost = totalDamageCost
	PurchaseOrderToUpdate.DamageReport.Date = date
	PurchaseOrderToUpdate.DamageReport.ItemCostPerUnit = itemCostPerUnit
	PurchaseOrderToUpdate.Status = status
	PurchaseOrderToUpdate.ApprovedBy = approvedBy
	PurchaseOrderToUpdate.CommentNote = commentNote
	PurchaseOrderToUpdate.InProgressTime = inProgressTime
	PurchaseOrderToUpdate.CreatedAt = createdAt
	PurchaseOrderToUpdate.UpdatedAt = updatedAt
	PurchaseOrderToUpdate.BatchArray = append(PurchaseOrderToUpdate.BatchArray, batch...)

	PurchaseOrderJSONasBytes, _ := json.Marshal(PurchaseOrderToUpdate)
	err = stub.PutState(returnRequestNo, PurchaseOrderJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("ExternalReturnRequest", PurchaseOrderJSONasBytes) //rewrite the marble
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
