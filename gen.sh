#!/bin/bash

basepath=$(cd `dirname $0`; pwd)
echo $basepath
mkdir -p $basepath/s2s
protoc -I $basepath --go_out=plugins=grpc:./s2s $basepath/protofiles/s2s.proto
mkdir -p $basepath/ss
protoc -I $basepath --go_out=plugins=grpc:./ss $basepath/protofiles/ss.proto
mkdir -p $basepath/s2s
protoc -I $basepath --go_out=plugins=grpc:./s2s $basepath/app/user/s2suser.proto