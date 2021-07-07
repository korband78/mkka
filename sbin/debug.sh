#!/bin/bash

BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`
VAR_DIR=${BASE_DIR}/var
PACKAGE_DIR=${VAR_DIR}/package
SRC_DIR=${BASE_DIR}/src

mkdir -p ${PACKAGE_DIR} ${SRC_DIR}

PORT=8100
DEBUG=" -debug"
while [ "$#" -gt "0" ]
do
  case $1 in
  		-p|--port)
  			shift
  			PORT=$1
  		;;
      -q|--quiet)
        shift
        DEBUG=
      ;;
      *)
      ;;
  esac
  shift
done

if [ x${PORT} == x ]; then
  echo
  echo "Usage: "
  echo "$0 --port <PORT>"
  echo
  exit 1
else
  cd ${SRC_DIR}
  ${PACKAGE_DIR}/go/bin/go run ${SRC_DIR}/*.go -port ${PORT} -rdir ${BASE_DIR}${DEBUG}
fi
