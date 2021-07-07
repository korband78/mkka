#!/bin/bash

BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`
VAR_DIR=${BASE_DIR}/var
PACKAGE_DIR=${VAR_DIR}/package
SRC_DIR=${BASE_DIR}/src
SBIN_DIR=${BASE_DIR}/sbin

mkdir -p ${PACKAGE_DIR} ${SRC_DIR} ${SBIN_DIR}

cd ${SRC_DIR}
time ${PACKAGE_DIR}/go/bin/go build -o ${SBIN_DIR}/run ${SRC_DIR}/*.go && ${PACKAGE_DIR}/go/bin/go mod tidy -v
