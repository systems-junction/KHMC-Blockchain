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
const queryTransaction = require('./query');
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

app.post('/api/addPurchaseOrder', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addPurchaseOrder',
    args: [

      req.body.purchaseOrderNo,
      req.body.purchaseRequestId,
      req.body.date,
      req.body.generated,
      req.body.generatedBy,
      req.body.commentNotes,
      req.body.approvedBy,
      req.body.vendorId,
      req.body.status,
      req.body.committeeStatus,
      req.body.inProgressTime,
      req.body.createdAt,
      req.body.sentAt,
      req.body.updatedAt


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseOrder with ID: "+req.body.purchaseOrderNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addPurchaseRequest', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addPurchaseRequest',
    args: [

      req.body.requestNo,
      req.body.generatedBy,
      req.body.status,
      req.body.committeeStatus,
      req.body.availability,
      req.body.reason,
      req.body.vendorId,
      req.body.rr,
      req.body.itemId,
      req.body.currQty,
      req.body.reqQty,
      req.body.comments,
      req.body.name,
      req.body.description,
      req.body.itemCode,
      req.body.istatus,
      req.body.secondStatus,
      req.body.requesterName,
      req.body.rejectionReason,
      req.body.department,
      req.body.commentNotes,
      req.body.orderType,
      req.body.generated,
      req.body.approvedBy,
      req.body.inProgressTime,
      req.body.createdAt,
      req.body.updatedAt




    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addPatient', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addPatient',
    args: [

      req.body.patientID,
      req.body.name,
      req.body.age,
      req.body.gender



    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Patient with ID: "+req.body.patientID+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addFunctionalUnit', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addFunctionalUnit',
    args: [
    
      req.body.uuid,
      req.body.fuName,
      req.body.description,
      req.body.fuHead,
      req.body.status,
      req.body.buId,
      req.body.fuLogId,
      req.body.createdAt,
      req.body.updatedAt

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Functional Unit with ID: "+req.body.uuid+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addReplenishmentRequest', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addReplenishmentRequest',
    args: [

      req.body.requestNo,
      req.body.generated,
      req.body.generatedBy,
      req.body.dateGenerated,
      req.body.reason,
      req.body.fuId,
      req.body.to,
      req.body.from,
      req.body.comments,
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.recieptUnit,
      req.body.issueUnit,
      req.body.fuItemCost,
      req.body.description,
      req.body.rstatus,
      req.body.rsecondStatus,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity,
      req.body.status,
      req.body.secondStatus,
      req.body.rrB,
      req.body.approvedBy,
      req.body.requesterName,
      req.body.orderType,
      req.body.department,
      req.body.commentNote,
      req.body.inProgressTime,
      req.body.completedTime,
      req.body.createdAt,
      req.body.updatedAt

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addFuInventory', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addFuInventory',
    args: [

      req.body.fuId,
      req.body.itemId,
      req.body.qty,
      req.body.maximumLevel,
      req.body.reorderLevel,
      req.body.minimumLevel,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The fuInventory with ID: "+req.body.fuId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addWarehouseInventory', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addWarehouseInventory',
    args: [

      req.body.itemId,
      req.body.qty,
      req.body.maximumLevel,
      req.body.minimumLevel,
      req.body.reorderLevel,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.price,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity,
      req.body.tempprice

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The WarehouseInventory with ID: "+req.body.itemId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addReceiveItem', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addReceiveItem',
    args: [
      
      req.body.itemId,
      req.body.prId,
      req.body.status,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount,
      req.body.discountAmount,
      req.body.tax,
      req.body.taxAmount,
      req.body.finalUnitPrice,
      req.body.subTotal,
      req.body.discountAmount2,
      req.body.totalPrice,
      req.body.invoice,
      req.body.dateInvoice,
      req.body.dateReceived,
      req.body.notes,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.returnedQty,
      req.body.batchNumberArr,
      req.body.expiryDateArr,
      req.body.quantity,
      req.body.price,
      req.body.qrCode,
      req.body.unitPrice

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItem with ID: "+req.body.itemId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addReceiveItemFU', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addReceiveItemFUSchema',
    args: [
      
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount ,
      req.body.discountAmount ,
      req.body.tax ,
      req.body.taxAmount ,
      req.body.finalUnitPrice ,
      req.body.subTotal ,
      req.body.discountAmount2 ,
      req.body.totalPrice ,
      req.body.invoice ,
      req.body.dateInvoice ,
      req.body.dateReceived ,
      req.body.notes ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.replenishmentRequestId ,
      req.body.batchNumberArr ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price 

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItemFU with ID: "+req.body.itemId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addReceiveItemBU', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addReceiveItemBUSchema',
    args: [
      
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount ,
      req.body.discountAmount ,
      req.body.tax ,
      req.body.taxAmount ,
      req.body.finalUnitPrice ,
      req.body.subTotal ,
      req.body.discountAmount2 ,
      req.body.totalPrice ,
      req.body.invoice ,
      req.body.dateInvoice ,
      req.body.dateReceived ,
      req.body.notes ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.replenishmentRequestId ,
      req.body.replenishmentRequestItemId,
      req.body.qualityRate,
      req.body.batchNumberArr ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price 

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItemFU with ID: "+req.body.itemId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addStaff', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addStaff',
    args: [

      req.body.staffId,
      req.body.staffTypeId,
      req.body.firstName,
      req.body.lastName,
      req.body.designation,
      req.body.contactNumber,
      req.body.identificationNumber,
      req.body.email,
      req.body.password,
      req.body.gender,
      req.body.dob,
      req.body.address ,
      req.body.functionalUnit ,
      req.body.systemAdminId ,
      req.body.status ,
      req.body.routes.toString() 


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.staffId+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addItem', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addItem',
    args: [

      req.body.name,
      req.body.description,
      req.body.itemCode,
      req.body.form,
      req.body.drugAllergy.toString(),
      req.body.receiptUnit,
      req.body.issueUnit,
      req.body.vendorId,
      req.body.purchasePrice,
      req.body.minimumLevel,
      req.body.maximumLevel ,
      req.body.reorderLevel ,
      req.body.cls ,
      req.body.medClass ,
      req.body.subClass ,
      req.body.grandSubClass ,
      req.body.comments ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.receiptUnitCost ,
      req.body.issueUnitCost ,
      req.body.scientificName ,
      req.body.tradeName ,
      req.body.temprature ,
      req.body.humidity ,
      req.body.lightSensitive ,
      req.body.resuableItem ,
      req.body.storageCondition ,
      req.body.expiration ,
      req.body.tax  


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.itemCode+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addInternalReturnRequestSchema', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addInternalReturnRequestSchema',
    args: [

      req.body.returnRequestNo,
      req.body.generatedBy,
      req.body.dateGenerated,
      req.body.expiryDate,
      req.body.to,
      req.body.from,
      req.body.currentQty,
      req.body.returnedQty,
      req.body.itemId,
      req.body.description,
      req.body.fuId ,
      req.body.reason ,
      req.body.reasonDetail ,
      req.body.buId ,
      req.body.causedBy ,
      req.body.totalDamageCost ,
      req.body.date ,
      req.body.itemCostPerUnit ,
      req.body.status ,
      req.body.replenishmentRequestBU ,
      req.body.replenishmentRequestFU ,
      req.body.approvedBy ,
      req.body.commentNote ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.batchNo ,
      req.body.batchNumber ,
      req.body.expiryDatePerBatch ,
      req.body.receivedQtyPerBatch ,
      req.body.returnedQtyPerBatch ,
      req.body.price


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.returnRequestNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addExternalReturnRequestSchema', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addExternalReturnRequestSchema',
    args: [

      req.body.returnRequestNo,
      req.body.generatedBy,
      req.body.generated,
      req.body.dateGenerated,
      req.body.expiryDate,
      req.body.returnedQty,
      req.body.itemId,
      req.body.prId,
      req.body.description,
      req.body.reason,
      req.body.reasonDetail ,
      req.body.causedBy ,
      req.body.totalDamageCost ,
      req.body.date ,
      req.body.itemCostPerUnit ,
      req.body.status ,
      req.body.approvedBy ,
      req.body.commentNote ,
      req.body.inProgressTime ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.batchNumber ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.returnRequestNo+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/addVendor', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'addVendor',
    args: [

      req.body.uuid,
req.body.vendorNo,
req.body.englishName,
req.body.arabicName,
req.body.telephone1,
req.body.telephone2,
req.body.contactEmail,
req.body.address,
req.body.country,
req.body.city,
req.body.zipcode ,
req.body.faxno ,
req.body.taxno ,
req.body.contactPersonName ,
req.body.contactPersonTelephone ,
req.body.contactPersonEmail ,
req.body.paymentTerms ,
req.body.shippingTerms ,
req.body.rating ,
req.body.status ,
req.body.cls ,
req.body.subClass.toString() ,
req.body.compliance ,
req.body.createdAt ,
req.body.updatedAt 


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Vendor with ID: "+req.body.uuid+ " is stored in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReceiveItemFU', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReceiveItemFU',
    args: [
      
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount ,
      req.body.discountAmount ,
      req.body.tax ,
      req.body.taxAmount ,
      req.body.finalUnitPrice ,
      req.body.subTotal ,
      req.body.discountAmount2 ,
      req.body.totalPrice ,
      req.body.invoice ,
      req.body.dateInvoice ,
      req.body.dateReceived ,
      req.body.notes ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.replenishmentRequestId ,
      req.body.batchNumberArr ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price 

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItemFU with ID: "+req.body.itemId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReceiveItemBU', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReceiveItemBU',
    args: [
      
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount ,
      req.body.discountAmount ,
      req.body.tax ,
      req.body.taxAmount ,
      req.body.finalUnitPrice ,
      req.body.subTotal ,
      req.body.discountAmount2 ,
      req.body.totalPrice ,
      req.body.invoice ,
      req.body.dateInvoice ,
      req.body.dateReceived ,
      req.body.notes ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.replenishmentRequestId ,
      req.body.replenishmentRequestItemId,
      req.body.qualityRate,
      req.body.batchNumberArr ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price 

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItemFU with ID: "+req.body.itemId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateItem', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateItem',
    args: [

      req.body.name,
      req.body.description,
      req.body.itemCode,
      req.body.form,
      req.body.drugAllergy.toString(),
      req.body.receiptUnit,
      req.body.issueUnit,
      req.body.vendorId,
      req.body.purchasePrice,
      req.body.minimumLevel,
      req.body.maximumLevel ,
      req.body.reorderLevel ,
      req.body.cls ,
      req.body.medClass ,
      req.body.subClass ,
      req.body.grandSubClass ,
      req.body.comments ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.receiptUnitCost ,
      req.body.issueUnitCost ,
      req.body.scientificName ,
      req.body.tradeName ,
      req.body.temprature ,
      req.body.humidity ,
      req.body.lightSensitive ,
      req.body.resuableItem ,
      req.body.storageCondition ,
      req.body.expiration ,
      req.body.tax  


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.itemCode+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateExternalReturnRequestSchema', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateExternalReturnRequestSchema',
    args: [

      req.body.returnRequestNo,
      req.body.generatedBy,
      req.body.generated,
      req.body.dateGenerated,
      req.body.expiryDate,
      req.body.returnedQty,
      req.body.itemId,
      req.body.prId,
      req.body.description,
      req.body.reason,
      req.body.reasonDetail ,
      req.body.causedBy ,
      req.body.totalDamageCost ,
      req.body.date ,
      req.body.itemCostPerUnit ,
      req.body.status ,
      req.body.approvedBy ,
      req.body.commentNote ,
      req.body.inProgressTime ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.batchNumber ,
      req.body.expiryDateArr ,
      req.body.quantity ,
      req.body.price


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.returnRequestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateInternalReturnRequestSchema', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateInternalReturnRequestSchema',
    args: [

      req.body.returnRequestNo,
      req.body.generatedBy,
      req.body.dateGenerated,
      req.body.expiryDate,
      req.body.to,
      req.body.from,
      req.body.currentQty,
      req.body.returnedQty,
      req.body.itemId,
      req.body.description,
      req.body.fuId ,
      req.body.reason ,
      req.body.reasonDetail ,
      req.body.buId ,
      req.body.causedBy ,
      req.body.totalDamageCost ,
      req.body.date ,
      req.body.itemCostPerUnit ,
      req.body.status ,
      req.body.replenishmentRequestBU ,
      req.body.replenishmentRequestFU ,
      req.body.approvedBy ,
      req.body.commentNote ,
      req.body.createdAt ,
      req.body.updatedAt ,
      req.body.batchNo ,
      req.body.batchNumber ,
      req.body.expiryDatePerBatch ,
      req.body.receivedQtyPerBatch ,
      req.body.returnedQtyPerBatch ,
      req.body.price


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.returnRequestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseOrder', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseOrder',
    args: [

      req.body.purchaseOrderNo,
      req.body.purchaseRequestId,
      req.body.date,
      req.body.generated,
      req.body.generatedBy,
      req.body.commentNotes,
      req.body.approvedBy,
      req.body.vendorId,
      req.body.status,
      req.body.committeeStatus,
      req.body.inProgressTime,
      req.body.createdAt,
      req.body.sentAt,
      req.body.updatedAt


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseOrder with ID: "+req.body.purchaseOrderNo+ " is Updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseRequest', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseRequest',
    args: [

      req.body.requestNo,
      req.body.generatedBy,
      req.body.status,
      req.body.committeeStatus,
      req.body.availability,
      req.body.reason,
      req.body.vendorId,
      req.body.rr,
      req.body.itemId,
      req.body.currQty,
      req.body.reqQty,
      req.body.comments,
      req.body.name,
      req.body.description,
      req.body.itemCode,
      req.body.istatus,
      req.body.secondStatus,
      req.body.requesterName,
      req.body.rejectionReason,
      req.body.department,
      req.body.commentNotes,
      req.body.orderType,
      req.body.generated,
      req.body.approvedBy,
      req.body.inProgressTime,
      req.body.createdAt,
      req.body.updatedAt




    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateFunctionalUnit', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateFunctionalUnit',
    args: [
    
      req.body.uuid,
      req.body.fuName,
      req.body.description,
      req.body.fuHead,
      req.body.status,
      req.body.buId,
      req.body.fuLogId,
      req.body.createdAt,
      req.body.updatedAt

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Functional Unit with ID: "+req.body.uuid+ " is Updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReplenishmentRequest', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReplenishmentRequest',
    args: [

      req.body.requestNo,
      req.body.generated,
      req.body.generatedBy,
      req.body.dateGenerated,
      req.body.reason,
      req.body.fuId,
      req.body.to,
      req.body.from,
      req.body.comments,
      req.body.itemId,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.recieptUnit,
      req.body.issueUnit,
      req.body.fuItemCost,
      req.body.description,
      req.body.rstatus,
      req.body.rsecondStatus,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity,
      req.body.status,
      req.body.secondStatus,
      req.body.rrB,
      req.body.approvedBy,
      req.body.requesterName,
      req.body.orderType,
      req.body.department,
      req.body.commentNote,
      req.body.inProgressTime,
      req.body.completedTime,
      req.body.createdAt,
      req.body.updatedAt

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateFuInventory', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateFuInventory',
    args: [

      req.body.fuId,
      req.body.itemId,
      req.body.qty,
      req.body.maximumLevel,
      req.body.reorderLevel,
      req.body.minimumLevel,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The fuInventory with ID: "+req.body.fuId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateWarehouseInventory', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateWarehouseInventory',
    args: [

      req.body.itemId,
      req.body.qty,
      req.body.maximumLevel,
      req.body.minimumLevel,
      req.body.reorderLevel,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.batchNumber,
      req.body.expiryDate,
      req.body.quantity,
      req.body.price,
      req.body.tempbatchNumber,
      req.body.tempexpiryDate,
      req.body.tempquantity,
      req.body.tempprice

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The WarehouseInventory with ID: "+req.body.itemId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReceiveItem', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReceiveItem',
    args: [
      
      req.body.itemId,
      req.body.prId,
      req.body.status,
      req.body.currentQty,
      req.body.requestedQty,
      req.body.receivedQty,
      req.body.bonusQty,
      req.body.batchNumber,
      req.body.lotNumber,
      req.body.expiryDate,
      req.body.unit,
      req.body.discount,
      req.body.unitDiscount,
      req.body.discountAmount,
      req.body.tax,
      req.body.taxAmount,
      req.body.finalUnitPrice,
      req.body.subTotal,
      req.body.discountAmount2,
      req.body.totalPrice,
      req.body.invoice,
      req.body.dateInvoice,
      req.body.dateReceived,
      req.body.notes,
      req.body.createdAt,
      req.body.updatedAt,
      req.body.returnedQty,
      req.body.batchNumberArr,
      req.body.expiryDateArr,
      req.body.quantity,
      req.body.price,
      req.body.qrCode,
      req.body.unitPrice

    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItem with ID: "+req.body.itemId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateStaff', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateStaff',
    args: [

      req.body.staffId,
      req.body.staffTypeId,
      req.body.firstName,
      req.body.lastName,
      req.body.designation,
      req.body.contactNumber,
      req.body.identificationNumber,
      req.body.email,
      req.body.password,
      req.body.gender,
      req.body.dob,
      req.body.address ,
      req.body.functionalUnit ,
      req.body.systemAdminId ,
      req.body.status ,
      req.body.routes.toString() 


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Staff with ID: "+req.body.staffId+ " is Updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateVendor', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateVendor',
    args: [

      req.body.uuid,
req.body.vendorNo,
req.body.englishName,
req.body.arabicName,
req.body.telephone1,
req.body.telephone2,
req.body.contactEmail,
req.body.address,
req.body.country,
req.body.city,
req.body.zipcode ,
req.body.faxno ,
req.body.taxno ,
req.body.contactPersonName ,
req.body.contactPersonTelephone ,
req.body.contactPersonEmail ,
req.body.paymentTerms ,
req.body.shippingTerms ,
req.body.rating ,
req.body.status ,
req.body.cls ,
req.body.subClass.toString() ,
req.body.compliance ,
req.body.createdAt ,
req.body.updatedAt 


    ]
  };
console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The Vendor with ID: "+req.body.uuid+ " is Updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseOrderStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseOrderStatus',
    args: [

      req.body.purchaseOrderNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The purchaseOrderNo with ID: "+req.body.purchaseOrderNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseOrderCommitteeStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseOrderCommitteeStatus',
    args: [

      req.body.purchaseOrderNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The purchaseOrderNo with ID: "+req.body.purchaseOrderNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseRequestStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseRequestStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseRequestCommitteeStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseRequestCommitteeStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseRequestItemStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseRequestItemStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updatePurchaseRequestItemSecondStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updatePurchaseRequestItemSecondStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The PurchaseRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReplenishmentRequestStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReplenishmentRequestStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReplenishmentRequestSecondStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReplenishmentRequestSecondStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReplenishmentRequestItemStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReplenishmentRequestItemStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReplenishmentRequestItemSecondStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReplenishmentRequestItemSecondStatus',
    args: [

      req.body.requestNo,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReplenishmentRequest with ID: "+req.body.requestNo+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateFunctionalUnitStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateFunctionalUnitStatus',
    args: [

      req.body.uuid,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The FunctionalUnit with ID: "+req.body.uuid+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

app.post('/api/updateReceiveItemStatus', async function (req, res) {

  var request = {
    chaincodeId: 'khmc',
    fcn: 'updateReceiveItemStatus',
    args: [

      req.body.itemId,
      req.body.newStatus
    ]
  };
  console.log(req.body);
  let response = await invoke.invokeCreate(request);
  if (response) {
    if(response.status == 200)
    res.status(response.status).send({ message: "The ReceiveItem with ID: "+req.body.itemId+ " is updated in the blockchain with " +response.message  });
    else
    res.status(response.status).send({ message: response.message});
  }
});

//-------------------------------------------------------------
//----------------------  GET API'S  --------------------------
//-------------------------------------------------------------

app.get('/api/queryPurchaseOrder', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryPurchaseOrder',
    args: [
      req.query.purchaseOrderNo
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

app.get('/api/queryPurchaseRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryPurchaseRequest',
    args: [
      req.query.requestNo
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
    chaincodeId: 'khmc',
    fcn: 'queryPatient',
    args: [
      req.query.patientID
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

app.get('/api/queryPatientByName', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryPatientByName',
    args: [
      req.query.name
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

app.get('/api/queryFunctionalUnit', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryFunctionalUnit',
    args: [
      req.query.uuid
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

app.get('/api/queryReplenishmentRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryReplenishmentRequest',
    args: [
      req.query.requestNo
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

app.get('/api/queryFuInventory', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryFuInventory',
    args: [
      req.query.fuId
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

app.get('/api/queryWarehouseInventory', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryWarehouseInventory',
    args: [
      req.query.itemId
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

app.get('/api/queryReceiveItem', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryReceiveItem',
    args: [
      req.query.itemId
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

app.get('/api/queryItem', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryItem',
    args: [
      req.query.itemCode
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

app.get('/api/queryInternalReturnRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryInternalReturnRequest',
    args: [
      req.query.returnRequestNo
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

app.get('/api/queryExternalReturnRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryExternalReturnRequest',
    args: [
      req.query.returnRequestNo
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

app.get('/api/queryReceiveItemBU', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryReceiveItemBU',
    args: [
      req.query.itemId
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

app.get('/api/queryReceiveItemFU', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryReceiveItemFU',
    args: [
      req.query.itemId
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

app.get('/api/queryStaff', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryStaff',
    args: [
      req.query.staffId
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

app.get('/api/queryVendor', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'queryVendor',
    args: [
      req.query.uuid
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

app.get('/api/getHistoryPurchaseRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryPurchaseOrder', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryReplenishmentRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryFuInventory', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryFunctionalUnit', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryWarehouseInventory', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryReceiveItem', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryStaff', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryVendor', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryItem', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryReceiveItemFU', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryReceiveItemBU', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryInternalReturnRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

app.get('/api/getHistoryExternalReturnRequest', async function (req, res) {

  const request = {
    chaincodeId: 'khmc',
    fcn: 'getHistory',
    args: [
      req.query.info
    ]
  };
  console.log(req.query);
  let response = await query.invokeQuery(request)
  if (response) {
    if(response.status == 200){
    res.status(response.status).send((JSON.parse(response.message)));

  }
    else
      res.status(response.status).send({ message: response.message });
  }
});

