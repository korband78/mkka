#!/bin/bash

BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`

cd ${BASE_DIR}/app/git
git add .
git commit -m "auto commit"
git push origin main
