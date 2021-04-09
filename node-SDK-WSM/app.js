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

app.get('/api/queryByTrasanction', async function (req, res) {

  const request = req.query.txId;
  let response = await queryTransaction.invokeQuery(request)
  if (response) {
    if(response.status == 200){

     let data = JSON.parse(response.message);
     // console.log("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaa"+data.transactionEnvelope);

    res.status(response.status).send({ message: data});
    }
    else
    res.status(response.status).send({ message: response.message });
  }
});