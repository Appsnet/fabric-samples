#!/bin/bash

source scriptUtils.sh

function createOrg1() {

  infoln "Enroll the CA admin"
  mkdir -p organizations/peerOrganizations/org1.troops.io/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.troops.io/
  #  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
  #  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' >${PWD}/organizations/peerOrganizations/org1.troops.io/msp/config.yaml

  infoln "Register peer0"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Register user"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Register the org admin"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  mkdir -p organizations/peerOrganizations/org1.troops.io/peers
  mkdir -p organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io

  infoln "Generate the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/msp --csr.hosts peer0.org1.troops.io --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/msp/config.yaml

  infoln "Generate the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls --enrollment.profile tls --csr.hosts peer0.org1.troops.io --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/keystore/* ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.troops.io/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.troops.io/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.troops.io/tlsca
  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.troops.io/tlsca/tlsca.org1.troops.io-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.troops.io/ca
  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/peers/peer0.org1.troops.io/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org1.troops.io/ca/ca.org1.troops.io-cert.pem

  mkdir -p organizations/peerOrganizations/org1.troops.io/users
  mkdir -p organizations/peerOrganizations/org1.troops.io/users/User1@org1.troops.io

  infoln "Generate the user msp"
  set -x
  fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.troops.io/users/User1@org1.troops.io/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.troops.io/users/User1@org1.troops.io/msp/config.yaml

  mkdir -p organizations/peerOrganizations/org1.troops.io/users/Admin@org1.troops.io

  infoln "Generate the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.troops.io/users/Admin@org1.troops.io/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org1.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.troops.io/users/Admin@org1.troops.io/msp/config.yaml

}

function createOrg2() {

  infoln "Enroll the CA admin"
  mkdir -p organizations/peerOrganizations/org2.troops.io/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org2.troops.io/
  #  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
  #  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-org2 --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: orderer' >${PWD}/organizations/peerOrganizations/org2.troops.io/msp/config.yaml

  infoln "Register peer0"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Register user"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Register the org admin"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name org2admin --id.secret org2adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  mkdir -p organizations/peerOrganizations/org2.troops.io/peers
  mkdir -p organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io

  infoln "Generate the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/msp --csr.hosts peer0.org2.troops.io --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/msp/config.yaml

  infoln "Generate the peer0-tls certificates"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls --enrollment.profile tls --csr.hosts peer0.org2.troops.io --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/keystore/* ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.troops.io/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.troops.io/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.troops.io/tlsca
  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org2.troops.io/tlsca/tlsca.org2.troops.io-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org2.troops.io/ca
  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/peers/peer0.org2.troops.io/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org2.troops.io/ca/ca.org2.troops.io-cert.pem

  mkdir -p organizations/peerOrganizations/org2.troops.io/users
  mkdir -p organizations/peerOrganizations/org2.troops.io/users/User1@org2.troops.io

  infoln "Generate the user msp"
  set -x
  fabric-ca-client enroll -u https://user1:user1pw@localhost:8054 --caname ca-org2 -M ${PWD}/organizations/peerOrganizations/org2.troops.io/users/User1@org2.troops.io/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.troops.io/users/User1@org2.troops.io/msp/config.yaml

  mkdir -p organizations/peerOrganizations/org2.troops.io/users/Admin@org2.troops.io

  infoln "Generate the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org2admin:org2adminpw@localhost:8054 --caname ca-org2 -M ${PWD}/organizations/peerOrganizations/org2.troops.io/users/Admin@org2.troops.io/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org2/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/peerOrganizations/org2.troops.io/msp/config.yaml ${PWD}/organizations/peerOrganizations/org2.troops.io/users/Admin@org2.troops.io/msp/config.yaml

}

function createOrderer() {

  infoln "Enroll the CA admin"
  mkdir -p organizations/ordererOrganizations/troops.io

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/troops.io
  #  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
  #  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' >${PWD}/organizations/ordererOrganizations/troops.io/msp/config.yaml

  infoln "Register orderer"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  infoln "Register the orderer admin"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  mkdir -p organizations/ordererOrganizations/troops.io/orderers
  mkdir -p organizations/ordererOrganizations/troops.io/orderers/troops.io

  mkdir -p organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io

  infoln "Generate the orderer msp"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/msp --csr.hosts orderer.troops.io --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/troops.io/msp/config.yaml ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/msp/config.yaml

  infoln "Generate the orderer-tls certificates"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls --enrollment.profile tls --csr.hosts orderer.troops.io --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/keystore/* ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/msp/tlscacerts/tlsca.troops.io-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/troops.io/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/troops.io/orderers/orderer.troops.io/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/troops.io/msp/tlscacerts/tlsca.troops.io-cert.pem

  mkdir -p organizations/ordererOrganizations/troops.io/users
  mkdir -p organizations/ordererOrganizations/troops.io/users/Admin@troops.io

  infoln "Generate the admin msp"
  set -x
  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/troops.io/users/Admin@troops.io/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  { set +x; } 2>/dev/null

  cp ${PWD}/organizations/ordererOrganizations/troops.io/msp/config.yaml ${PWD}/organizations/ordererOrganizations/troops.io/users/Admin@troops.io/msp/config.yaml

}
