#!/bin/bash

export PATH=${PWD}/../bin:$PATH

export FABRIC_CFG_PATH=$PWD/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="KhmcMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode package khmc.tar.gz --path ../chaincode/khmc/go/ --lang golang --label khmc_1
peer lifecycle chaincode install khmc.tar.gz

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:8051
peer channel join -b ${PWD}/channel-artifacts/khmcchannel.block
peer lifecycle chaincode install khmc.tar.gz

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer2.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:9051
peer channel join -b ${PWD}/channel-artifacts/khmcchannel.block
peer lifecycle chaincode install khmc.tar.gz

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer3.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:10051
peer channel join -b ${PWD}/channel-artifacts/khmcchannel.block
peer lifecycle chaincode install khmc.tar.gz

peer lifecycle chaincode queryinstalled

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export CORE_PEER_ADDRESS=localhost:7051



