#!/bin/bash

BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`
VAR_DIR=${BASE_DIR}/var

mkdir -p ${VAR_DIR}

PID=$(<${VAR_DIR}/pid)
kill -9 ${PID}
