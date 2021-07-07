#!/bin/bash

ulimit -n 60000
cat > /etc/security/limits.conf << EOF
*              soft      nofile          10240
*              hard     nofile          10240
EOF

iptables -I INPUT 5 -i ens3 -p tcp --dport 80 -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -I INPUT 5 -i ens3 -p tcp --dport 443 -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -I INPUT 5 -i ens3 -p tcp --dport 8100 -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -I INPUT 5 -i ens3 -p tcp --dport 8101 -m state --state NEW,ESTABLISHED -j ACCEPT

############################################################
# OS 체크
############################################################
OSNAME=
case "$OSTYPE" in
  darwin*)
    OSNAME="darwin"
  ;;
  linux*)
    OSNAME="linux"

    yum install -y wget gcc
    systemctl stop iptables
    systemctl stop ip6tables
    systemctl disable iptables
    systemctl disable ip6tables
  ;;
  *)
    echo "This operating system is not supported.($OSTYPE)"
    exit 1
  ;;
esac

############################################################
# 경로 설정
############################################################
BASE_DIR=$(dirname $0)
cd ${BASE_DIR}/..
BASE_DIR=`pwd`
VAR_DIR=${BASE_DIR}/var
CACHE_DIR=${VAR_DIR}/.cache
LOG_DIR=${VAR_DIR}/log
DOWNLOAD_DIR=${VAR_DIR}/download
PACKAGE_DIR=${VAR_DIR}/package
SBIN_DIR=${BASE_DIR}/sbin
SRC_DIR=${BASE_DIR}/src

PKG_DIR=${BASE_DIR}/pkg

mkdir -p ${CACHE_DIR} ${LOG_DIR} ${DOWNLOAD_DIR} ${PACKAGE_DIR} ${SBIN_DIR} ${SRC_DIR} ${PKG_DIR}

rm -rf ${DOWNLOAD_DIR} && mkdir -p ${DOWNLOAD_DIR} && cd ${DOWNLOAD_DIR}

# golang 패키지 다운로드
VERSION=go1.15.${OSNAME}-amd64
wget https://dl.google.com/go/${VERSION}.tar.gz .

# golang 설치
tar xvf ${VERSION}.tar.gz
mv go ${PACKAGE_DIR}/${VERSION}
rm -rf ${PACKAGE_DIR}/go && ln -fs ${PACKAGE_DIR}/${VERSION} ${PACKAGE_DIR}/go
rm -rf ${DOWNLOAD_DIR}
