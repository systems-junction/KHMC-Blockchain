'use strict';

//get libraries
const express = require('express');
var queue = require('express-queue');
const bodyParser = require('body-parser');
const request = require('request');
const path = require('path');

//create express web-app
const app = express();
const invoke = require('./invokeNetwork');
const query = require('./queryNetwork');
const queryHistory = require('./queryHistory');

var _time = "T00:00:00Z";

//declare port
var port = process.env.PORT || 8001;
if (process.env.VCAP_APPLICATION) {
  port = process.env.PORT;
}

app.use(bodyParser.json());

app.use(bodyParser.urlencoded({
  extended: true
 }));

app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  next();
  });

//Using queue middleware
app.use(queue({ activeLimit: 30, queuedLimit: -1 }));

//run app on port
app.listen(port, function () {
  console.log('app running on port: %d', port);
});

//-------------------------------------------------------------
//----------------------  POST API'S    -----------------------
//-------------------------------------------------------------

app.post('/api/addInsuranceInfo', async function (req, res) {

  var request = {
    chaincodeId: 'insurance',
    fcn: 'addInsuranceInfo',
    args: [

      req.body.patientsID,
      req.body.insuranceIDNo,
      req.body.patientName,
      req.body.insuranceStatus,
      req.body.claimedBy,
      req.body.totalFee,
      req.body.coveredAmount,
      req.body.details,
      req.body.prescriberSign


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The InsuranceInfo with ID: "+req.body.patientsID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addPatient', async function (req, res) {

  var request = {
    chaincodeId: 'insurance',
    fcn: 'addPatient',
    args: [

      req.body.profileNo,
      req.body.SIN,
      req.body.title,
      req.body.firstName,
      req.body.lastName ,
      req.body.fullName ,
      req.body.gender,
      req.body.nationality,
      req.body.age,
      req.body.height ,
      req.body.weight,
      req.body.bloodGroup, 
      req.body.dob ,
      req.body.drugAllergy.toString(),
      req.body.phoneNumber ,
      req.body.mobileNumber ,
      req.body.email ,
      req.body.country, 
      req.body.city ,
      req.body.address, 
      req.body.otherDetails, 
      req.body.paymentMethod, 
      req.body.depositAmount, 
      req.body.amountReceived, 
      req.body.bankName ,
      req.body.depositorName, 
      req.body.depositSlip, 
      req.body.insuranceNo, 
      req.body.insuranceVendor, 
      req.body.coverageDetails, 
      req.body.coverageTerms, 
      req.body.payment, 
      req.body.registeredIn, 
      req.body.receivedBy ,
      req.body.emergencyName, 
      req.body.emergencyContactNo, 
      req.body.emergencyRelation, 
      req.body.coveredFamilyMembers, 
      req.body.otherCoverageDetails, 
      req.body.otherCity, 
      req.body.QR, 
      req.body.createdAt, 
      req.body.updatedAt, 
      req.body.userProfile.email, 
      req.body.userProfile.contact, 
      req.body.userProfile.firstName, 
      req.body.userProfile.lastName, 
      req.body.userProfile.userName, 
      req.body.userProfile.gender, 
      req.body.userProfile.dob, 
      req.body.userProfile.isActive, 
      req.body.userProfile.maritalStatus, 
      req.body.userProfile.address, 
      req.body.userProfile.communicationLanguage, 
      req.body.userProfile.profilePicture, 
      req.body.userProfile.geneticDisease.toString(), 
      req.body.patientMedicalProfile.knownAllergies.toString(), 
      req.body.patientMedicalProfile.currentMedication.toString(), 
      req.body.patientMedicalProfile.surgeries.toString(), 
      req.body.patientMedicalProfile.chronicIllness.toString(), 
      req.body.patientMedicalProfile.bloodGroup, 
      req.body.patientMedicalProfile.pregnancy

    ]
  };
console.log(req.body.drugAllergy.toString());
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Patient Info with ID: "+req.body.profileNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePatient', async function (req, res) {

  var request = {
    chaincodeId: 'insurance',
    fcn: 'updatePatient',
    args: [

      req.body.profileNo,
      req.body.SIN,
      req.body.title,
      req.body.firstName,
      req.body.lastName ,
      req.body.fullName ,
      req.body.gender,
      req.body.nationality,
      req.body.age,
      req.body.height ,
      req.body.weight,
      req.body.bloodGroup, 
      req.body.dob ,
      req.body.drugAllergy.toString(),
      req.body.phoneNumber ,
      req.body.mobileNumber ,
      req.body.email ,
      req.body.country, 
      req.body.city ,
      req.body.address, 
      req.body.otherDetails, 
      req.body.paymentMethod, 
      req.body.depositAmount, 
      req.body.amountReceived, 
      req.body.bankName ,
      req.body.depositorName, 
      req.body.depositSlip, 
      req.body.insuranceNo, 
      req.body.insuranceVendor, 
      req.body.coverageDetails, 
      req.body.coverageTerms, 
      req.body.payment, 
      req.body.registeredIn, 
      req.body.receivedBy ,
      req.body.emergencyName, 
      req.body.emergencyContactNo, 
      req.body.emergencyRelation, 
      req.body.coveredFamilyMembers, 
      req.body.otherCoverageDetails, 
      req.body.otherCity, 
      req.body.QR, 
      req.body.createdAt, 
      req.body.updatedAt, 
      req.body.userProfile.email, 
      req.body.userProfile.contact, 
      req.body.userProfile.firstName, 
      req.body.userProfile.lastName, 
      req.body.userProfile.userName, 
      req.body.userProfile.gender, 
      req.body.userProfile.dob, 
      req.body.userProfile.isActive, 
      req.body.userProfile.maritalStatus, 
      req.body.userProfile.address, 
      req.body.userProfile.communicationLanguage, 
      req.body.userProfile.profilePicture, 
      req.body.userProfile.geneticDisease.toString(), 
      req.body.patientMedicalProfile.knownAllergies.toString(), 
      req.body.patientMedicalProfile.currentMedication.toString(), 
      req.body.patientMedicalProfile.surgeries.toString(), 
      req.body.patientMedicalProfile.chronicIllness.toString(), 
      req.body.patientMedicalProfile.bloodGroup, 
      req.body.patientMedicalProfile.pregnancy

    ]
  };
console.log( req.body.drugAllergy.toString());
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Patient Info with ID: "+req.body.profileNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});


//-------------------------------------------------------------
//----------------------  GET API'S  --------------------------
//-------------------------------------------------------------

app.get('/api/queryInsuranceInfo', async function (req, res) {

  const request = {
    chaincodeId: 'insurance',
    fcn: 'queryInsuranceInfo',
    args: [
      req.query.patientsID
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/queryPatient', async function (req, res) {

  const request = {
    chaincodeId: 'insurance',
    fcn: 'queryPatient',
    args: [
      req.query.profileNo
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200)
      res.status(response.status).send(JSON.parse(response.message));
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistory', async function (req, res) {
  var resp = {};
  var resparr = [];
  const request = {
    chaincodeId: 'insurance',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    let data = JSON.parse(response.message);
    for(var i = 0 ; i<data.length ; i++){
      console.log(i,"BBBBBBBBBBBBBBBBB")
    // resp={Txid:data[i].TxId,profileNo:data[i].Value.profileNo,
    //   firstName:data[i].Value.firstName,lastName:data[i].Value.lastName,mobileNumber:data[i].Value.mobileNumber,
    //   email:data[i].Value.email,address:data[i].Value.address,depositAmount:data[i].Value.depositAmount,
    //   depositorName:data[i].Value.depositorName,createdAt:data[i].Value.createdAt}

      resp['Txid']=data[i].TxId
      resp['profileNo']=data[i].Value.profileNo
      resp['firstName']=data[i].Value.firstName
      resp['lastName']=data[i].Value.lastName
      resp['mobileNumber']=data[i].Value.mobileNumber
      resp['email']=data[i].Value.email
      resp['address']=data[i].Value.address
      resp['depositAmount']=data[i].Value.depositAmount
      resp['depositorName']=data[i].Value.depositorName
      resp['createdAt']=data[i].Value.createdAt

      resp.Txid=resp['Txid']
      resp.profileNo=resp['profileNo']
      resp.firstName=resp['firstName']
      resp.lastName=resp['lastName']
      resp.mobileNumber=resp['mobileNumber']
      resp.email=resp['email']
      resp.address=resp['address']
      resp.depositAmount=resp['depositAmount']
      resp.depositorName=resp['depositorName']
      resp.createdAt=resp['createdAt']

    console.log(resp,"AAAAAAAAAAAAAAAAAA");
    resparr.push(resp)
    }
    res.status(response.status).send((resparr));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});
