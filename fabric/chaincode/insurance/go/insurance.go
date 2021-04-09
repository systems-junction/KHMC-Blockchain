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
	ObjectType            string   `json:"Type"`
	ProfileNo             string   `json:"profileNo"` //MRN
	SIN                   string   `json:"SIN"`
	Title                 string   `json:"title"`
	FirstName             string   `json:"firstName"` //name
	LastName              string   `json:"lastName"`  //name
	FullName              string   `json:"fullName"`
	Gender                string   `json:"gender"`
	Nationality           string   `json:"nationality"`
	Age                   string   `json:"age"`
	Height                string   `json:"height"`
	Weight                string   `json:"weight"`
	BloodGroup            string   `json:"bloodGroup"`
	Dob                   string   `json:"dob"`
	DrugAllergy           []string `json:"drugAllergy"`
	PhoneNumber           string   `json:"phoneNumber"`
	MobileNumber          string   `json:"mobileNumber"` //phone
	Email                 string   `json:"email"`        //email
	Country               string   `json:"country"`
	City                  string   `json:"city"`
	Address               string   `json:"address"` //address
	OtherDetails          string   `json:"otherDetails"`
	PaymentMethod         string   `json:"paymentMethod"`
	DepositAmount         string   `json:"depositAmount"` //Amount
	AmountReceived        string   `json:"amountReceived"`
	BankName              string   `json:"bankName"`
	DepositorName         string   `json:"depositorName"` //Depositor
	DepositSlip           string   `json:"depositSlip"`
	InsuranceNo           string   `json:"insuranceNo"`
	InsuranceVendor       string   `json:"insuranceVendor"`
	CoverageDetails       string   `json:"coverageDetails"`
	CoverageTerms         string   `json:"coverageTerms"`
	Payment               string   `json:"payment"`
	RegisteredIn          string   `json:"registeredIn"`
	ReceivedBy            string   `json:"receivedBy"`
	EmergencyName         string   `json:"emergencyName"`
	EmergencyContactNo    string   `json:"emergencyContactNo"`
	EmergencyRelation     string   `json:"emergencyRelation"`
	CoveredFamilyMembers  string   `json:"coveredFamilyMembers"`
	OtherCoverageDetails  string   `json:"otherCoverageDetails"`
	OtherCity             string   `json:"otherCity"`
	QR                    string   `json:"QR"`
	CreatedAt             string   `json:"createdAt"` //Registered Date
	UpdatedAt             string   `json:"updatedAt"`
	UserProfile           UserProfile
	PatientMedicalProfile PatientMedicalProfile
}

type UserProfile struct {
	ObjectType            string   `json:"Type"`
	Email                 string   `json:"email"`
	Contact               string   `json:"contact"`
	FirstName             string   `json:"firstName"`
	LastName              string   `json:"lastName"`
	UserName              string   `json:"userName"`
	Gender                string   `json:"gender"`
	Dob                   string   `json:"dob"`
	IsActive              string   `json:"isActive"`
	MaritalStatus         string   `json:"maritalStatus"`
	Address               string   `json:"address"`
	CommunicationLanguage string   `json:"communicationLanguage"`
	ProfilePicture        string   `json:"profilePicture"`
	GeneticDisease        []string `json:"geneticDisease"`
}

type PatientMedicalProfile struct {
	ObjectType        string   `json:"Type"`
	KnownAllergies    []string `json:"knownAllergies"`
	CurrentMedication []string `json:"currentMedication"`
	Surgeries         []string `json:"surgeries"`
	ChronicIllness    []string `json:"chronicIllness"`
	BloodGroup        string   `json:"bloodGroup"`
	Pregnancy         string   `json:"pregnancy"`
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
	if function == "updateGeneticDisease" {
		return t.updateGeneticDisease(stub, args)
	}
	if function == "updateKnownAllergies" {
		return t.updateKnownAllergies(stub, args)
	}
	if function == "updateCurrentMedication" {
		return t.updateCurrentMedication(stub, args)
	}
	if function == "updateSurgeries" {
		return t.updateSurgeries(stub, args)
	}
	if function == "updateChronicIllness" {
		return t.updateChronicIllness(stub, args)
	}
	if function == "updatePatient" {
		return t.updatePatient(stub, args)
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

	if len(args) != 62 {
		return shim.Error("Incorrect Number of Aruments. Expecting 62")
	}

	fmt.Println("Adding new Patient Info")

	for i := 0; i < 61; i++ {
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
	QR := args[40]
	createdAt := args[41]
	updatedAt := args[42]
	uemail := args[43]
	contact := args[44]
	ufirstName := args[45]
	ulastName := args[46]
	userName := args[47]
	ugender := args[48]
	udob := args[49]
	isActive := args[50]
	maritalStatus := args[51]
	uaddress := args[52]
	communicationLanguage := args[53]
	profilePicture := args[54]
	geneticDisease := strings.Split(args[55], ",")
	//geneticDisease := args[55]
	// var geneticDisease []string
	// geneticDisease = append(geneticDisease, geneticDiseaseTest)
	knownAllergies := strings.Split(args[56], ",")
	//knownAllergies := args[56]
	// var knownAllergies []string
	// knownAllergies = append(knownAllergies, knownAllergiesTest)
	currentMedication := strings.Split(args[57], ",")
	//currentMedication := args[57]
	// var currentMedication []string
	// currentMedication = append(currentMedication, currentMedicationTest)
	surgeries := strings.Split(args[58], ",")
	//surgeries := args[58]
	// var surgeries []string
	// surgeries = append(surgeries, surgeriesTest)
	chronicIllness := strings.Split(args[59], ",")
	//chronicIllness := args[59]
	// var chronicIllness []string
	// chronicIllness = append(chronicIllness, chronicIllnessTest)
	pbloodGroup := args[60]
	pregnancy := args[61]

	// ======Check if PatientINFO Already exists

	InsuranceAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted Patient ID already Exists: " + profileNo)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "Patient"
	objectType1 := "UserProfile"
	objectType2 := "PatientMedicalProfile"
	Patient := &Patient{objectType, profileNo, SIN, title, firstName, lastName, fullName, gender, nationality,
		age, height, weight, bloodGroup, dob, append(Patient{}.DrugAllergy, drugAllergy...), phoneNumber, mobileNumber, email, country, city, address, otherDetails, paymentMethod,
		depositAmount, amountReceived, bankName, depositorName, depositSlip, insuranceNo, insuranceVendor, coverageDetails, coverageTerms,
		payment, registeredIn, receivedBy, emergencyName, emergencyContactNo, emergencyRelation, coveredFamilyMembers, otherCoverageDetails,
		otherCity, QR, createdAt, updatedAt, UserProfile{objectType1, uemail, contact, ufirstName, ulastName, userName, ugender, udob, isActive,
			maritalStatus, uaddress, communicationLanguage, profilePicture, append(Patient{}.UserProfile.GeneticDisease, geneticDisease...)}, PatientMedicalProfile{objectType2,
			append(Patient{}.PatientMedicalProfile.KnownAllergies, knownAllergies...), append(Patient{}.PatientMedicalProfile.CurrentMedication, currentMedication...),
			append(Patient{}.PatientMedicalProfile.Surgeries, surgeries...), append(Patient{}.PatientMedicalProfile.ChronicIllness, chronicIllness...), pbloodGroup, pregnancy}}
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

func (t *SmartContract) updatePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 62 {
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
	QR := args[40]
	createdAt := args[41]
	updatedAt := args[42]
	uemail := args[43]
	contact := args[44]
	ufirstName := args[45]
	ulastName := args[46]
	userName := args[47]
	ugender := args[48]
	udob := args[49]
	isActive := args[50]
	maritalStatus := args[51]
	uaddress := args[52]
	communicationLanguage := args[53]
	profilePicture := args[54]
	geneticDisease := strings.Split(args[55], ",")
	//geneticDisease := args[55]
	knownAllergies := strings.Split(args[56], ",")
	currentMedication := strings.Split(args[57], ",")
	surgeries := strings.Split(args[58], ",")
	chronicIllness := strings.Split(args[59], ",")
	pbloodGroup := args[60]
	pregnancy := args[61]
	fmt.Println("- start  ", profileNo, profileNo, SIN, title, firstName, lastName, fullName, gender, nationality,
		age, height, weight, bloodGroup, dob, drugAllergy, phoneNumber, mobileNumber, email, country, city, address, otherDetails, paymentMethod,
		depositAmount, amountReceived, bankName, depositorName, depositSlip, insuranceNo, insuranceVendor, coverageDetails, coverageTerms,
		payment, registeredIn, receivedBy, emergencyName, emergencyContactNo, emergencyRelation, coveredFamilyMembers, otherCoverageDetails,
		otherCity, QR, createdAt, updatedAt, uemail, contact, ufirstName, ulastName, userName, ugender, udob, isActive,
		maritalStatus, uaddress, communicationLanguage, profilePicture, geneticDisease, knownAllergies, currentMedication, surgeries, chronicIllness, pbloodGroup, pregnancy)

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
	PatientToUpdate.QR = QR
	PatientToUpdate.CreatedAt = createdAt
	PatientToUpdate.UpdatedAt = updatedAt
	PatientToUpdate.UserProfile.Email = uemail
	PatientToUpdate.UserProfile.Contact = contact
	PatientToUpdate.UserProfile.FirstName = ufirstName
	PatientToUpdate.UserProfile.LastName = ulastName
	PatientToUpdate.UserProfile.UserName = userName
	PatientToUpdate.UserProfile.Gender = ugender
	PatientToUpdate.UserProfile.Dob = udob
	PatientToUpdate.UserProfile.IsActive = isActive
	PatientToUpdate.UserProfile.MaritalStatus = maritalStatus
	PatientToUpdate.UserProfile.Address = uaddress
	PatientToUpdate.UserProfile.CommunicationLanguage = communicationLanguage
	PatientToUpdate.UserProfile.ProfilePicture = profilePicture
	PatientToUpdate.UserProfile.GeneticDisease = append(PatientToUpdate.UserProfile.GeneticDisease, geneticDisease...)
	PatientToUpdate.PatientMedicalProfile.KnownAllergies = append(PatientToUpdate.PatientMedicalProfile.KnownAllergies, knownAllergies...)
	PatientToUpdate.PatientMedicalProfile.CurrentMedication = nil
	PatientToUpdate.PatientMedicalProfile.CurrentMedication = append(PatientToUpdate.PatientMedicalProfile.CurrentMedication, currentMedication...)
	PatientToUpdate.PatientMedicalProfile.Surgeries = append(PatientToUpdate.PatientMedicalProfile.Surgeries, surgeries...)
	PatientToUpdate.PatientMedicalProfile.ChronicIllness = append(PatientToUpdate.PatientMedicalProfile.ChronicIllness, chronicIllness...)
	PatientToUpdate.PatientMedicalProfile.BloodGroup = pbloodGroup
	PatientToUpdate.PatientMedicalProfile.Pregnancy = pregnancy

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

func (t *SmartContract) updateGeneticDisease(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	profileNo := args[0]
	newDisease := args[1]
	fmt.Println("- start  ", profileNo, newDisease)

	DiseaseAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if DiseaseAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	DiseaseToUpdate := Patient{}
	err = json.Unmarshal(DiseaseAsBytes, &DiseaseToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	DiseaseToUpdate.UserProfile.GeneticDisease = append(DiseaseToUpdate.UserProfile.GeneticDisease, newDisease) //change the status

	DiseaseJSONasBytes, _ := json.Marshal(DiseaseToUpdate)
	err = stub.PutState(profileNo, DiseaseJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateKnownAllergies(stub shim.ChaincodeStubInterface, args []string) peer.Response {

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
	AllergyToUpdate.PatientMedicalProfile.KnownAllergies = append(AllergyToUpdate.PatientMedicalProfile.KnownAllergies, newAllergy) //change the status

	AllergyJSONasBytes, _ := json.Marshal(AllergyToUpdate)
	err = stub.PutState(profileNo, AllergyJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateCurrentMedication(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	profileNo := args[0]
	newCurrMed := args[1]
	fmt.Println("- start  ", profileNo, newCurrMed)

	CurrMedAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if CurrMedAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	CurrMedToUpdate := Patient{}
	err = json.Unmarshal(CurrMedAsBytes, &CurrMedToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	CurrMedToUpdate.PatientMedicalProfile.CurrentMedication = append(CurrMedToUpdate.PatientMedicalProfile.CurrentMedication, newCurrMed) //change the status

	CurrMedJSONasBytes, _ := json.Marshal(CurrMedToUpdate)
	err = stub.PutState(profileNo, CurrMedJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateSurgeries(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	profileNo := args[0]
	newSurgeri := args[1]
	fmt.Println("- start  ", profileNo, newSurgeri)

	SurgeriAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if SurgeriAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	SurgeriToUpdate := Patient{}
	err = json.Unmarshal(SurgeriAsBytes, &SurgeriToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	SurgeriToUpdate.PatientMedicalProfile.Surgeries = append(SurgeriToUpdate.PatientMedicalProfile.Surgeries, newSurgeri) //change the status

	SurgeriJSONasBytes, _ := json.Marshal(SurgeriToUpdate)
	err = stub.PutState(profileNo, SurgeriJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateChronicIllness(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	profileNo := args[0]
	newChronicIllness := args[1]
	fmt.Println("- start  ", profileNo, newChronicIllness)

	ChronicIllnessAsBytes, err := stub.GetState(profileNo)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if ChronicIllnessAsBytes == nil {
		return shim.Error("Patient Info does not exist")
	}

	ChronicIllnessToUpdate := Patient{}
	err = json.Unmarshal(ChronicIllnessAsBytes, &ChronicIllnessToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	ChronicIllnessToUpdate.PatientMedicalProfile.ChronicIllness = append(ChronicIllnessToUpdate.PatientMedicalProfile.ChronicIllness, newChronicIllness) //change the status

	ChronicIllnessJSONasBytes, _ := json.Marshal(ChronicIllnessToUpdate)
	err = stub.PutState(profileNo, ChronicIllnessJSONasBytes) //rewrite the marble
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
