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

//Insurance info
type Insurance struct {
	ObjectType      string `json:"Type"`
	PatientsID      string `json:"patientsID"`
	InsuranceIDno   string `json:"insuranceIDNo"`
	PatientName     string `json:"patientName"`
	InsuranceStatus string `json:"insuranceStatus"`
	ClaimedBy       string `json:"claimedBy"`
	ItemFee         ItemFee
	Details         string `json:"details"`
	PrescriberSign  string `json:"prescriberSign"`
}

type Patient struct {
	ObjectType           string   `json:"Type"`
	ProfileNo            string   `json:"profileNo"` //MRN
	SIN                  string   `json:"SIN"`
	Title                string   `json:"title"`
	FirstName            string   `json:"firstName"` //name
	LastName             string   `json:"lastName"`  //name
	FullName             string   `json:"fullName"`
	Gender               string   `json:"gender"`
	Nationality          string   `json:"nationality"`
	Age                  string   `json:"age"`
	Height               string   `json:"height"`
	Weight               string   `json:"weight"`
	BloodGroup           string   `json:"bloodGroup"`
	Dob                  string   `json:"dob"`
	DrugAllergy          []string `json:"drugAllergy"`
	PhoneNumber          string   `json:"phoneNumber"`
	MobileNumber         string   `json:"mobileNumber"` //phone
	Email                string   `json:"email"`        //email
	Country              string   `json:"country"`
	City                 string   `json:"city"`
	Address              string   `json:"address"` //address
	OtherDetails         string   `json:"otherDetails"`
	PaymentMethod        string   `json:"paymentMethod"`
	DepositAmount        string   `json:"depositAmount"` //Amount
	AmountReceived       string   `json:"amountReceived"`
	BankName             string   `json:"bankName"`
	DepositorName        string   `json:"depositorName"` //Depositor
	DepositSlip          string   `json:"depositSlip"`
	InsuranceNo          string   `json:"insuranceNo"`
	InsuranceVendor      string   `json:"insuranceVendor"`
	CoverageDetails      string   `json:"coverageDetails"`
	CoverageTerms        string   `json:"coverageTerms"`
	Payment              string   `json:"payment"`
	RegisteredIn         string   `json:"registeredIn"`
	ReceivedBy           string   `json:"receivedBy"`
	EmergencyName        string   `json:"emergencyName"`
	EmergencyContactNo   string   `json:"emergencyContactNo"`
	EmergencyRelation    string   `json:"emergencyRelation"`
	CoveredFamilyMembers string   `json:"coveredFamilyMembers"`
	OtherCoverageDetails string   `json:"otherCoverageDetails"`
	OtherCity            string   `json:"otherCity"`
	PaymentDate          string   `json:"paymentDate"`
	CreatedAt            string   `json:"createdAt"` //Registered Date
	UpdatedAt            string   `json:"updatedAt"`
}

//ItemFee by Hospital
type ItemFee struct {
	TotalFee      string `json:"totalFee"`
	CoveredAmount string `json:"coveredAmount"`
}

type Consultation struct {
	ConsultationNo    string `json:"consultationNo"`
	Date              string `json:"date"`
	Description       string `json:"description"`
	ConsultationNotes string `json:"consultationNotes"`
	DoctorNotes       string `json:"doctorNotes"`
	AudioNotes        string `json:"audioNotes"`
	Status            string `json:"status"`
	Speciality        string `json:"speciality"`
	Specialist        string `json:"specialist"`
	Requester         string `json:"requester"`
	CompletedTime     string `json:"completedTime"`
}

type Resident struct {
	ResidentNoteNo string   `json:"residentNoteNo"`
	Date           string   `json:"date"`
	Description    string   `json:"description"`
	Doctor         string   `json:"doctor"`
	Note           string   `json:"note"`
	Section        string   `json:"section"`
	AudioNotes     string   `json:"audioNotes"`
	Code           []string `json:"code"`
}

type ResidentIPR struct {
	ResidentNoteNo string   `json:"residentNoteNo"`
	Date           string   `json:"date"`
	Description    string   `json:"description"`
	Doctor         string   `json:"doctor"`
	Note           string   `json:"note"`
	Status         string   `json:"status"`
	Section        string   `json:"section"`
	AudioNotes     string   `json:"audioNotes"`
	Code           []string `json:"code"`
}

type PharmacyReq struct {
	ReplenishmentRequestBuID string `json:"ReplenishmentRequestBuID"`
}

type LabReq struct {
	LRrequestNo   string `json:"lRrequestNo"`
	ServiceId     string `json:"serviceId"`
	Price         string `json:"price"`
	RequesterName string `json:"requesterName"`
	ServiceCode   string `json:"serviceCode"`
	ServiceName   string `json:"serviceName"`
	Status        string `json:"status"`
	Requester     string `json:"requester"`
	Results       string `json:"results"`
	SampleId      string `json:"sampleId"`
	Comments      string `json:"comments"`
	ServiceType   string `json:"serviceType"`
	ActiveDate    string `json:"activeDate"`
	CompletedDate string `json:"completedDate"`
	Date          string `json:"date"`
}

type RadiologyReq struct {
	RRrequestNo      string `json:"rRrequestNo"`
	ServiceId        string `json:"serviceId"`
	Price            string `json:"price"`
	ServiceCode      string `json:"serviceCode"`
	Status           string `json:"status"`
	RequesterName    string `json:"requesterName"`
	ServiceName      string `json:"serviceName"`
	Requester        string `json:"requester"`
	Results          string `json:"results"`
	Comments         string `json:"comments"`
	ServiceType      string `json:"serviceType"`
	ConsultationNote string `json:"consultationNote"`
	ActiveDate       string `json:"activeDate"`
	CompletedDate    string `json:"completedDate"`
	Date             string `json:"date"`
}

type Med struct {
	ItemId       string `json:"itemId"`
	Priority     string `json:"priority"`
	Schedule     string `json:"schedule"`
	Dosage       string `json:"dosage"`
	Frequency    string `json:"frequency"`
	Duration     string `json:"duration"`
	RequestedQty string `json:"requestedQty"`
	MedicineName string `json:"medicineName"`
	UnitPrice    string `json:"unitPrice"`
	TotalPrice   string `json:"totalPrice"`
	ItemType     string `json:"itemType"`
	Make_model   string `json:"make_model"`
	Size         string `json:"size"`
}

type Triage struct {
	Status            string   `json:"status"`
	Reason            string   `json:"reason"`
	TriageRequestNo   string   `json:"triageRequestNo"`
	HeartRate         string   `json:"heartRate"`
	BloodPressureSys  string   `json:"bloodPressureSys"`
	BloodPressureDia  string   `json:"bloodPressureDia"`
	RespiratoryRate   string   `json:"respiratoryRate"`
	Temperature       string   `json:"temperature"`
	FSBS              string   `json:"FSBS"`
	PainScale         string   `json:"painScale"`
	PulseOX           string   `json:"pulseOX"`
	TriageLevel       []string `json:"triageLevel"`
	GeneralAppearance []string `json:"generalAppearance"`
	HeadNeck          []string `json:"headNeck"`
	Respiratory       []string `json:"respiratory"`
	Cardiac           []string `json:"cardiac"`
	Abdomen           []string `json:"abdomen"`
	Neurological      []string `json:"neurological"`
	Requester         string   `json:"requester"`
	Date              string   `json:"date"`
}

type ConsultationNote []Consultation

type ResidentNotes []Resident

type ResidentNotesIPR []ResidentIPR

type PharmacyRequest []PharmacyReq

type LabRequest []LabReq

type RadiologyRequest []RadiologyReq

type Medicine []Med

type TriageAssessment []Triage

type DischargeSummary struct {
	DischargeNotes string `json:"dischargeNotes"`
	OtherNotes     string `json:"otherNotes"`
}

type DischargeMedication struct {
	Date      string `json:"date"`
	Status    string `json:"status"`
	Requester string `json:"requester"`
	Medicine
}

type DischargeRequest struct {
	ObjectType          string `json:"Type"`
	DischargeSummary    DischargeSummary
	DischargeMedication DischargeMedication
	Status              string `json:"status"`
	InProcessDate       string `json:"inProcessDate"`
	CompletionDate      string `json:"completionDate"`
}

type EDRSchema struct {
	ObjectType  string `json:"Type"`
	RequestNo   string `json:"requestNo"`
	PatientId   string `json:"patientId"`
	GeneratedBy string `json:"generatedBy"`
	ConsultationNote
	ResidentNotes
	PharmacyRequest
	LabRequest
	RadiologyRequest
	DischargeRequest DischargeRequest
	Status           string `json:"status"`
	TriageAssessment
	RequestType   string `json:"requestType"`
	Verified      string `json:"verified"`
	InsurerId     string `json:"insurerId"`
	PaymentMethod string `json:"paymentMethod"`
	Claimed       string `json:"claimed"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type NurseServ struct {
	NSrequestNo   string `json:"NSrequestNo"`
	ServiceId     string `json:"serviceId"`
	Price         string `json:"price"`
	RequesterName string `json:"requesterName"`
	ServiceCode   string `json:"serviceCode"`
	Status        string `json:"status"`
	ServiceName   string `json:"serviceName"`
	Comments      string `json:"comments"`
	Requester     string `json:"requester"`
	Date          string `json:"date"`
}

type Follow struct {
	Requester      string `json:"requester"`
	ApprovalNumber string `json:"approvalNumber"`
	ApprovalPerson string `json:"approvalPerson"`
	File           string `json:"file"`
	Description    string `json:"description"`
	Notes          string `json:"notes"`
	Status         string `json:"status"`
	DoctorName     string `json:"doctorName"`
	Doctor         string `json:"doctor"`
	Date           string `json:"date"`
}

type NurseService []NurseServ
type FollowUp []Follow

type IPRSchema struct {
	ObjectType  string `json:"Type"`
	RequestNo   string `json:"requestNo"`
	PatientId   string `json:"patientId"`
	GeneratedBy string `json:"generatedBy"`
	ConsultationNote
	ResidentNotesIPR
	PharmacyRequest
	LabRequest
	RadiologyRequest
	NurseService
	DischargeRequest DischargeRequest
	Status           string `json:"status"`
	TriageAssessment
	FollowUp
	RequestType    string `json:"requestType"`
	FunctionalUnit string `json:"functionalUnit"`
	Verified       string `json:"verified"`
	InsurerId      string `json:"insurerId"`
	PaymentMethod  string `json:"paymentMethod"`
	Claimed        string `json:"claimed"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

//Init method
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

//Invoke functions
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)
	if function == "addInsuranceInfo" {
		return t.addInsuranceInfo(stub, args)
	}
	if function == "addPatient" {
		return t.addPatient(stub, args)
	}
	if function == "queryInsuranceInfo" {
		return t.queryInsuranceInfo(stub, args)
	}
	if function == "queryPatient" {
		return t.queryPatient(stub, args)
	}
	if function == "updateDrugAllergy" {
		return t.updateDrugAllergy(stub, args)
	}
	if function == "updatePatient" {
		return t.updatePatient(stub, args)
	}
	if function == "getHistory" {
		return t.getHistory(stub, args)
	}
	if function == "addEDRSchema" {
		return t.addEDRSchema(stub, args)
	}
	if function == "queryEDRSchema" {
		return t.queryEDRSchema(stub, args)
	}
	if function == "addIPRSchema" {
		return t.addIPRSchema(stub, args)
	}
	if function == "queryIPRSchema" {
		return t.queryIPRSchema(stub, args)
	}
	if function == "updateEDRSchema" {
		return t.updateEDRSchema(stub, args)
	}
	if function == "updateIPRSchema" {
		return t.updateIPRSchema(stub, args)
	}
	if function == "getHistory" {
		return t.getHistory(stub, args)
	}
	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addInsuranceInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 9 {
		return shim.Error("Incorrect Number of Aruments. Expecting 19")
	}

	fmt.Println("Adding new Insurance Info")

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

	patientsID := args[0]
	insuranceIDNo := args[1]
	patientName := args[2]
	insuranceStatus := args[3]
	claimedBy := args[4]
	totalFee := args[5]
	coveredAmount := args[6]
	details := args[7]
	prescriberSign := args[8]

	// ======Check if PurchaseRequest Already exists

	InsuranceAsBytes, err := stub.GetState(patientsID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted Patient ID already Exists: " + patientsID)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "Insurance"
	Insurance := &Insurance{objectType, patientsID, insuranceIDNo, patientName, insuranceStatus, claimedBy, ItemFee{totalFee, coveredAmount}, details, prescriberSign}
	InsuranceJSONasBytes, err := json.Marshal(Insurance)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseRequest to State

	err = stub.PutState(patientsID, InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Insurance Info")
	return shim.Success(nil)
}

func (t *SmartContract) queryInsuranceInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientsID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Insurance\",\"patientsID\":\"%s\"}}", patientsID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 43 {
		return shim.Error("Incorrect Number of Aruments. Expecting 62")
	}

	fmt.Println("Adding new Patient Info")

	for i := 0; i < 42; i++ {
		if len(args[i]) <= 0 {
			argument0 := "Argument "
			argument := " Must be a Non-Empty String"
			concat := fmt.Sprint(argument0, i, argument)
			return shim.Error(concat)
		}

	}

	profileNo := args[0]
	SIN := args[1]
	title := args[2]
	firstName := args[3]
	lastName := args[4]
	fullName := args[5]
	gender := args[6]
	nationality := args[7]
	age := args[8]
	height := args[9]
	weight := args[10]
	bloodGroup := args[11]
	dob := args[12]
	drugAllergy := strings.Split(args[13], ",")
	//drugAllergy := args[13]
	// var drugAllergy []string
	// drugAllergy = append(drugAllergy, drugAllergytest)
	phoneNumber := args[14]
	mobileNumber := args[15]
	email := args[16]
	country := args[17]
	city := args[18]
	address := args[19]
	otherDetails := args[20]
	paymentMethod := args[21]
	depositAmount := args[22]
	amountReceived := args[23]
	bankName := args[24]
	depositorName := args[25]
	depositSlip := args[26]
	insuranceNo := args[27]
	insuranceVendor := args[28]
	coverageDetails := args[29]
	coverageTerms := args[30]
	payment := args[31]
	registeredIn := args[32]
	receivedBy := args[33]
	emergencyName := args[34]
	emergencyContactNo := args[35]
	emergencyRelation := args[36]
	coveredFamilyMembers := args[37]
	otherCoverageDetails := args[38]
	otherCity := args[39]
	paymentDate := args[40]
	createdAt := args[41]
	updatedAt := args[42]

	// ======Check if PatientINFO Already exists

	InsuranceAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted Patient ID already Exists: " + profileNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "Patient"
	Patient := &Patient{objectType, profileNo, SIN, title, firstName, lastName, fullName, gender, nationality,
		age, height, weight, bloodGroup, dob, drugAllergy, phoneNumber, mobileNumber, email, country, city, address, otherDetails, paymentMethod,
		depositAmount, amountReceived, bankName, depositorName, depositSlip, insuranceNo, insuranceVendor, coverageDetails, coverageTerms,
		payment, registeredIn, receivedBy, emergencyName, emergencyContactNo, emergencyRelation, coveredFamilyMembers, otherCoverageDetails,
		otherCity, paymentDate, createdAt, updatedAt}
	InsuranceJSONasBytes, err := json.Marshal(Patient)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PatientINFO to State

	err = stub.PutState(profileNo, InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("PatientInfo", InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Insurance Info")
	return shim.Success(nil)
}

func (t *SmartContract) queryPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	profileNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Patient\",\"profileNo\":\"%s\"}}", profileNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addEDRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 26 {
		return shim.Error("Incorrect Number of Aruments. Expecting 62")
	}

	fmt.Println("Adding new EDR Info")

	for i := 0; i < 26; i++ {
		if len(args[i]) <= 0 {
			argument0 := "Argument "
			argument := " Must be a Non-Empty String"
			concat := fmt.Sprint(argument0, i, argument)
			return shim.Error(concat)
		}

	}

	requestNo := args[0]
	patientId := args[1]
	generatedBy := args[2]
	consultationNote := args[3]

	var consult []Consultation
	json.Unmarshal([]byte(consultationNote), &consult)

	residentNotes := args[4]

	var resident []Resident
	json.Unmarshal([]byte(residentNotes), &resident)

	pharmacyRequest := args[5]

	var pharmacy []PharmacyReq
	json.Unmarshal([]byte(pharmacyRequest), &pharmacy)

	labRequest := args[6]

	var lab []LabReq
	json.Unmarshal([]byte(labRequest), &lab)

	radiologyRequest := args[7]

	var radiology []RadiologyReq
	json.Unmarshal([]byte(radiologyRequest), &radiology)

	dischargeNotes := args[8]
	otherNotes := args[9]
	dDate := args[10]
	dStatus := args[11]
	dRequester := args[12]

	medicine := args[13]

	var med []Med
	json.Unmarshal([]byte(medicine), &med)

	drStatus := args[14]
	inProcessDate := args[15]
	completionDate := args[16]
	status := args[17]
	triageAssessment := args[18]

	var triage []Triage
	json.Unmarshal([]byte(triageAssessment), &triage)

	requestType := args[19]
	verified := args[20]
	insurerId := args[21]
	paymentMethod := args[22]
	claimed := args[23]
	createdAt := args[24]
	updatedAt := args[25]

	// ======Check if PatientINFO Already exists

	InsuranceAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted EDR ID already Exists: " + requestNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "EDR"
	objectType1 := "DischargeRequest"
	EDRSchema := &EDRSchema{objectType, requestNo, patientId, generatedBy, append(EDRSchema{}.ConsultationNote, consult...),
		append(EDRSchema{}.ResidentNotes, resident...), append(EDRSchema{}.PharmacyRequest, pharmacy...), append(EDRSchema{}.LabRequest, lab...),
		append(EDRSchema{}.RadiologyRequest, radiology...), DischargeRequest{objectType1, DischargeSummary{dischargeNotes, otherNotes}, DischargeMedication{dDate, dStatus, dRequester,
			append(EDRSchema{}.DischargeRequest.DischargeMedication.Medicine, med...)}, drStatus, inProcessDate, completionDate}, status,
		append(EDRSchema{}.TriageAssessment, triage...), requestType, verified, insurerId, paymentMethod, claimed, createdAt, updatedAt}
	InsuranceJSONasBytes, err := json.Marshal(EDRSchema)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PatientINFO to State

	err = stub.PutState(requestNo, InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("EDR", InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Insurance Info")
	return shim.Success(nil)
}

func (t *SmartContract) queryEDRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requestNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"EDR\",\"requestNo\":\"%s\"}}", requestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addIPRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 29 {
		return shim.Error("Incorrect Number of Aruments. Expecting 62")
	}

	fmt.Println("Adding new EDR Info")

	for i := 0; i < 29; i++ {
		if len(args[i]) <= 0 {
			argument0 := "Argument "
			argument := " Must be a Non-Empty String"
			concat := fmt.Sprint(argument0, i, argument)
			return shim.Error(concat)
		}

	}

	requestNo := args[0]
	patientId := args[1]
	generatedBy := args[2]
	consultationNote := args[3]

	var consult []Consultation
	json.Unmarshal([]byte(consultationNote), &consult)

	residentNotes := args[4]

	var resident []ResidentIPR
	json.Unmarshal([]byte(residentNotes), &resident)

	pharmacyRequest := args[5]

	var pharmacy []PharmacyReq
	json.Unmarshal([]byte(pharmacyRequest), &pharmacy)

	labRequest := args[6]

	var lab []LabReq
	json.Unmarshal([]byte(labRequest), &lab)

	radiologyRequest := args[7]

	var radiology []RadiologyReq
	json.Unmarshal([]byte(radiologyRequest), &radiology)

	nurseService := args[8]

	var nurse []NurseServ
	json.Unmarshal([]byte(nurseService), &nurse)

	dischargeNotes := args[9]
	otherNotes := args[10]
	dDate := args[11]
	dStatus := args[12]
	dRequester := args[13]
	medicine := args[14]

	var med []Med
	json.Unmarshal([]byte(medicine), &med)

	drStatus := args[15]
	inProcessDate := args[16]
	completionDate := args[17]
	status := args[18]
	triageAssessment := args[19]

	var triage []Triage
	json.Unmarshal([]byte(triageAssessment), &triage)

	followUp := args[20]

	var follow []Follow
	json.Unmarshal([]byte(followUp), &follow)

	requestType := args[21]
	functionalUnit := args[22]
	verified := args[23]
	insurerId := args[24]
	paymentMethod := args[25]
	claimed := args[26]
	createdAt := args[27]
	updatedAt := args[28]

	// ======Check if PatientINFO Already exists

	InsuranceAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted IPR ID already Exists: " + requestNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "IPR"
	objectType1 := "DischargeRequest"
	IPRSchema := &IPRSchema{objectType, requestNo, patientId, generatedBy, append(IPRSchema{}.ConsultationNote, consult...),
		append(IPRSchema{}.ResidentNotesIPR, resident...), append(IPRSchema{}.PharmacyRequest, pharmacy...),
		append(IPRSchema{}.LabRequest, lab...), append(IPRSchema{}.RadiologyRequest, radiology...), append(IPRSchema{}.NurseService, nurse...),
		DischargeRequest{objectType1, DischargeSummary{dischargeNotes, otherNotes}, DischargeMedication{dDate, dStatus, dRequester,
			append(IPRSchema{}.DischargeRequest.DischargeMedication.Medicine, med...)}, drStatus, inProcessDate, completionDate}, status,
		append(IPRSchema{}.TriageAssessment, triage...), append(IPRSchema{}.FollowUp, follow...), requestType, functionalUnit, verified, insurerId, paymentMethod, claimed, createdAt, updatedAt}
	InsuranceJSONasBytes, err := json.Marshal(IPRSchema)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PatientINFO to State

	err = stub.PutState(requestNo, InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("IPR", InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Insurance Info")
	return shim.Success(nil)
}

func (t *SmartContract) queryIPRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requestNo := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"IPR\",\"requestNo\":\"%s\"}}", requestNo)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) updateEDRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 26 {
		return shim.Error("Incorrect number of arguments. Expecting 63")
	}

	requestNo := args[0]
	patientId := args[1]
	generatedBy := args[2]
	consultationNote := args[3]

	var consult []Consultation
	json.Unmarshal([]byte(consultationNote), &consult)

	residentNotes := args[4]

	var resident []Resident
	json.Unmarshal([]byte(residentNotes), &resident)

	pharmacyRequest := args[5]

	var pharmacy []PharmacyReq
	json.Unmarshal([]byte(pharmacyRequest), &pharmacy)

	labRequest := args[6]

	var lab []LabReq
	json.Unmarshal([]byte(labRequest), &lab)

	radiologyRequest := args[7]

	var radiology []RadiologyReq
	json.Unmarshal([]byte(radiologyRequest), &radiology)

	dischargeNotes := args[8]
	otherNotes := args[9]
	dDate := args[10]
	dStatus := args[11]
	dRequester := args[12]

	medicine := args[13]

	var med []Med
	json.Unmarshal([]byte(medicine), &med)

	drStatus := args[14]
	inProcessDate := args[15]
	completionDate := args[16]
	status := args[17]
	triageAssessment := args[18]

	var triage []Triage
	json.Unmarshal([]byte(triageAssessment), &triage)

	requestType := args[19]
	verified := args[20]
	insurerId := args[21]
	paymentMethod := args[22]
	claimed := args[23]
	createdAt := args[24]
	updatedAt := args[25]
	fmt.Println("- start  ", requestNo, patientId, generatedBy, consultationNote, residentNotes, pharmacyRequest, labRequest, radiologyRequest, dischargeNotes, otherNotes, dDate, dStatus, dRequester, medicine, drStatus, inProcessDate,
		completionDate, status, triageAssessment, requestType, verified, insurerId,
		paymentMethod, claimed, createdAt, updatedAt)

	PatientAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PatientAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	EDRToUpdate := EDRSchema{}
	err = json.Unmarshal(PatientAsBytes, &EDRToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	EDRToUpdate.PatientId = patientId
	EDRToUpdate.GeneratedBy = generatedBy
	EDRToUpdate.ConsultationNote = append(EDRToUpdate.ConsultationNote, consult...)
	EDRToUpdate.ResidentNotes = append(EDRToUpdate.ResidentNotes, resident...)
	EDRToUpdate.PharmacyRequest = append(EDRToUpdate.PharmacyRequest, pharmacy...)
	EDRToUpdate.LabRequest = append(EDRToUpdate.LabRequest, lab...)
	EDRToUpdate.RadiologyRequest = append(EDRToUpdate.RadiologyRequest, radiology...)
	EDRToUpdate.DischargeRequest.DischargeSummary.DischargeNotes = dischargeNotes
	EDRToUpdate.DischargeRequest.DischargeSummary.OtherNotes = otherNotes
	EDRToUpdate.DischargeRequest.DischargeMedication.Date = dDate
	EDRToUpdate.DischargeRequest.DischargeMedication.Status = dStatus
	EDRToUpdate.DischargeRequest.DischargeMedication.Requester = dRequester
	EDRToUpdate.DischargeRequest.DischargeMedication.Medicine = append(EDRToUpdate.DischargeRequest.DischargeMedication.Medicine, med...)
	EDRToUpdate.DischargeRequest.Status = drStatus
	EDRToUpdate.DischargeRequest.InProcessDate = inProcessDate
	EDRToUpdate.DischargeRequest.CompletionDate = completionDate
	EDRToUpdate.Status = status
	EDRToUpdate.TriageAssessment = append(EDRToUpdate.TriageAssessment, triage...)
	EDRToUpdate.RequestType = requestType
	EDRToUpdate.Verified = verified
	EDRToUpdate.InsurerId = insurerId
	EDRToUpdate.PaymentMethod = paymentMethod
	EDRToUpdate.Claimed = claimed
	EDRToUpdate.CreatedAt = createdAt
	EDRToUpdate.UpdatedAt = updatedAt

	PatientJSONasBytes, _ := json.Marshal(EDRToUpdate)
	err = stub.PutState(requestNo, PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("EDR", PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateIPRSchema(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 29 {
		return shim.Error("Incorrect number of arguments. Expecting 63")
	}

	requestNo := args[0]
	patientId := args[1]
	generatedBy := args[2]
	consultationNote := args[3]

	var consult []Consultation
	json.Unmarshal([]byte(consultationNote), &consult)

	residentNotes := args[4]

	var resident []ResidentIPR
	json.Unmarshal([]byte(residentNotes), &resident)

	pharmacyRequest := args[5]

	var pharmacy []PharmacyReq
	json.Unmarshal([]byte(pharmacyRequest), &pharmacy)

	labRequest := args[6]

	var lab []LabReq
	json.Unmarshal([]byte(labRequest), &lab)

	radiologyRequest := args[7]

	var radiology []RadiologyReq
	json.Unmarshal([]byte(radiologyRequest), &radiology)

	nurseService := args[8]

	var nurse []NurseServ
	json.Unmarshal([]byte(nurseService), &nurse)

	dischargeNotes := args[9]
	otherNotes := args[10]
	dDate := args[11]
	dStatus := args[12]
	dRequester := args[13]
	medicine := args[14]

	var med []Med
	json.Unmarshal([]byte(medicine), &med)

	drStatus := args[15]
	inProcessDate := args[16]
	completionDate := args[17]
	status := args[18]
	triageAssessment := args[19]

	var triage []Triage
	json.Unmarshal([]byte(triageAssessment), &triage)

	followUp := args[20]

	var follow []Follow
	json.Unmarshal([]byte(followUp), &follow)

	requestType := args[21]
	functionalUnit := args[22]
	verified := args[23]
	insurerId := args[24]
	paymentMethod := args[25]
	claimed := args[26]
	createdAt := args[27]
	updatedAt := args[28]

	fmt.Println("- start  ", requestNo, patientId, generatedBy, consultationNote, residentNotes, pharmacyRequest, labRequest, radiologyRequest, nurseService, dischargeNotes, otherNotes, dDate, dStatus, dRequester, medicine, drStatus, inProcessDate,
		completionDate, status, triageAssessment, followUp, requestType, functionalUnit, verified, insurerId,
		paymentMethod, claimed, createdAt, updatedAt)

	PatientAsBytes, err := stub.GetState(requestNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PatientAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	IPRToUpdate := IPRSchema{}
	err = json.Unmarshal(PatientAsBytes, &IPRToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	IPRToUpdate.PatientId = patientId
	IPRToUpdate.GeneratedBy = generatedBy
	IPRToUpdate.ConsultationNote = append(IPRToUpdate.ConsultationNote, consult...)
	IPRToUpdate.ResidentNotesIPR = append(IPRToUpdate.ResidentNotesIPR, resident...)
	IPRToUpdate.PharmacyRequest = append(IPRToUpdate.PharmacyRequest, pharmacy...)
	IPRToUpdate.LabRequest = append(IPRToUpdate.LabRequest, lab...)
	IPRToUpdate.RadiologyRequest = append(IPRToUpdate.RadiologyRequest, radiology...)
	IPRToUpdate.NurseService = append(IPRToUpdate.NurseService, nurse...)
	IPRToUpdate.DischargeRequest.DischargeSummary.DischargeNotes = dischargeNotes
	IPRToUpdate.DischargeRequest.DischargeSummary.OtherNotes = otherNotes
	IPRToUpdate.DischargeRequest.DischargeMedication.Date = dDate
	IPRToUpdate.DischargeRequest.DischargeMedication.Status = dStatus
	IPRToUpdate.DischargeRequest.DischargeMedication.Requester = dRequester
	IPRToUpdate.DischargeRequest.DischargeMedication.Medicine = append(IPRToUpdate.DischargeRequest.DischargeMedication.Medicine, med...)
	IPRToUpdate.DischargeRequest.Status = drStatus
	IPRToUpdate.DischargeRequest.InProcessDate = inProcessDate
	IPRToUpdate.DischargeRequest.CompletionDate = completionDate
	IPRToUpdate.Status = status
	IPRToUpdate.TriageAssessment = append(IPRToUpdate.TriageAssessment, triage...)
	IPRToUpdate.FollowUp = append(IPRToUpdate.FollowUp, follow...)
	IPRToUpdate.RequestType = requestType
	IPRToUpdate.FunctionalUnit = functionalUnit
	IPRToUpdate.Verified = verified
	IPRToUpdate.InsurerId = insurerId
	IPRToUpdate.PaymentMethod = paymentMethod
	IPRToUpdate.Claimed = claimed
	IPRToUpdate.CreatedAt = createdAt
	IPRToUpdate.UpdatedAt = updatedAt

	PatientJSONasBytes, _ := json.Marshal(IPRToUpdate)
	err = stub.PutState(requestNo, PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("IPR", PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updatePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 42 {
		return shim.Error("Incorrect number of arguments. Expecting 63")
	}

	profileNo := args[0]
	SIN := args[1]
	title := args[2]
	firstName := args[3]
	lastName := args[4]
	fullName := args[5]
	gender := args[6]
	nationality := args[7]
	age := args[8]
	height := args[9]
	weight := args[10]
	bloodGroup := args[11]
	dob := args[12]
	drugAllergy := strings.Split(args[13], ",")
	//drugAllergy := args[13]
	phoneNumber := args[14]
	mobileNumber := args[15]
	email := args[16]
	country := args[17]
	city := args[18]
	address := args[19]
	otherDetails := args[20]
	paymentMethod := args[21]
	depositAmount := args[22]
	amountReceived := args[23]
	bankName := args[24]
	depositorName := args[25]
	depositSlip := args[26]
	insuranceNo := args[27]
	insuranceVendor := args[28]
	coverageDetails := args[29]
	coverageTerms := args[30]
	payment := args[31]
	registeredIn := args[32]
	receivedBy := args[33]
	emergencyName := args[34]
	emergencyContactNo := args[35]
	emergencyRelation := args[36]
	coveredFamilyMembers := args[37]
	otherCoverageDetails := args[38]
	otherCity := args[39]
	paymentDate := args[40]
	createdAt := args[41]
	updatedAt := args[42]

	fmt.Println("- start  ", profileNo, profileNo, SIN, title, firstName, lastName, fullName, gender, nationality,
		age, height, weight, bloodGroup, dob, drugAllergy, phoneNumber, mobileNumber, email, country, city, address, otherDetails, paymentMethod,
		depositAmount, amountReceived, bankName, depositorName, depositSlip, insuranceNo, insuranceVendor, coverageDetails, coverageTerms,
		payment, registeredIn, receivedBy, emergencyName, emergencyContactNo, emergencyRelation, coveredFamilyMembers, otherCoverageDetails,
		otherCity, paymentDate, createdAt, updatedAt)

	PatientAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if PatientAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	PatientToUpdate := Patient{}
	err = json.Unmarshal(PatientAsBytes, &PatientToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	PatientToUpdate.SIN = SIN
	PatientToUpdate.Title = title
	PatientToUpdate.FirstName = firstName
	PatientToUpdate.LastName = lastName
	PatientToUpdate.FullName = fullName
	PatientToUpdate.Gender = gender
	PatientToUpdate.Nationality = nationality
	PatientToUpdate.Age = age
	PatientToUpdate.Height = height
	PatientToUpdate.Weight = weight
	PatientToUpdate.BloodGroup = bloodGroup
	PatientToUpdate.Dob = dob
	PatientToUpdate.DrugAllergy = append(PatientToUpdate.DrugAllergy, drugAllergy...)
	PatientToUpdate.PhoneNumber = phoneNumber
	PatientToUpdate.MobileNumber = mobileNumber
	PatientToUpdate.Email = email
	PatientToUpdate.Country = country
	PatientToUpdate.City = city
	PatientToUpdate.Address = address
	PatientToUpdate.OtherDetails = otherDetails
	PatientToUpdate.PaymentMethod = paymentMethod
	PatientToUpdate.DepositAmount = depositAmount
	PatientToUpdate.AmountReceived = amountReceived
	PatientToUpdate.BankName = bankName
	PatientToUpdate.DepositorName = depositorName
	PatientToUpdate.DepositSlip = depositSlip
	PatientToUpdate.InsuranceNo = insuranceNo
	PatientToUpdate.InsuranceVendor = insuranceVendor
	PatientToUpdate.CoverageDetails = coverageDetails
	PatientToUpdate.CoverageTerms = coverageTerms
	PatientToUpdate.Payment = payment
	PatientToUpdate.RegisteredIn = registeredIn
	PatientToUpdate.ReceivedBy = receivedBy
	PatientToUpdate.EmergencyName = emergencyName
	PatientToUpdate.EmergencyContactNo = emergencyContactNo
	PatientToUpdate.EmergencyRelation = emergencyRelation
	PatientToUpdate.CoveredFamilyMembers = coveredFamilyMembers
	PatientToUpdate.OtherCoverageDetails = otherCoverageDetails
	PatientToUpdate.OtherCity = otherCity
	PatientToUpdate.PaymentDate = paymentDate
	PatientToUpdate.CreatedAt = createdAt
	PatientToUpdate.UpdatedAt = updatedAt

	PatientJSONasBytes, _ := json.Marshal(PatientToUpdate)
	err = stub.PutState(profileNo, PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("PatientInfo", PatientJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateDrugAllergy(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	profileNo := args[0]
	newAllergy := args[1]
	fmt.Println("- start  ", profileNo, newAllergy)

	AllergyAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if AllergyAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	AllergyToUpdate := Patient{}
	err = json.Unmarshal(AllergyAsBytes, &AllergyToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	AllergyToUpdate.DrugAllergy = append(AllergyToUpdate.DrugAllergy, newAllergy) //change the status

	AllergyJSONasBytes, _ := json.Marshal(AllergyToUpdate)
	err = stub.PutState(profileNo, AllergyJSONasBytes) //rewrite the marble
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
