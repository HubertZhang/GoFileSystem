#!/bin/sh

export GOPATH="`pwd`"

cd $GOPATH/src/GoFileSystem/primary
go install

cd  $GOPATH/src/GoFileSystem/backup
go install
