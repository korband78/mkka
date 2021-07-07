#!/bin/bash

BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`
VAR_DIR=${BASE_DIR}/var
CONF_DIR=${BASE_DIR}/conf
SBIN_DIR=${BASE_DIR}/sbin

mkdir -p ${VAR_DIR} ${CONF_DIR} ${SBIN_DIR}

PID=$(<${VAR_DIR}/pid)
ARGS=`xargs -0 < /proc/${PID}/cmdline | awk '{$1="";print $0;}'`
kill -9 ${PID}
echo "# execute cmd: "
echo "${SBIN_DIR}/run$ARGS"
nohup ${SBIN_DIR}/run${ARGS} > /dev/null 2>&1 &
