#!/bin/bash

############################################################
# 경로 설정
############################################################
BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`

OSNAME=
case "$OSTYPE" in
  darwin*)
    OSNAME="darwin"
  ;;
  *)
    echo "This operating system is not supported.($OSTYPE)"
    exit 1
  ;;
esac

cd ${BASE_DIR}/app

rm -rf ${BASE_DIR}/app/git/docs
mkdir -p ${BASE_DIR}/app/git/docs/static

rm -rf ${BASE_DIR}/www
mkdir -p ${BASE_DIR}/www

ng run app:build:production --base-href=/ --deploy-url=/v/ --output-path=${BASE_DIR}/app/git/docs/v

mv ${BASE_DIR}/app/git/docs/v/index.html ${BASE_DIR}/app/git/docs/index.html
mv ${BASE_DIR}/app/git/docs/v/assets ${BASE_DIR}/app/git/docs/assets
mv ${BASE_DIR}/app/git/docs/v/svg ${BASE_DIR}/app/git/docs/svg
cp ${BASE_DIR}/app/git/docs/index.html ${BASE_DIR}/app/git/docs/404.html
cp -r ${BASE_DIR}/app/src/index/* ${BASE_DIR}/app/git/docs
cp -r ${BASE_DIR}/app/git/docs/* ${BASE_DIR}/www
