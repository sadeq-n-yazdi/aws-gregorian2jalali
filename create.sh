#!/bin/sh

. $(dirname "$0")/config.sh

go build -ldflags "-w -s" -o main &&\
rm deploy.zip &&\
zip deploy.zip main &&\
( 
  aws lambda delete-function --function-name $FUNCNAME \
    --region $REGNAME 
  aws lambda create-function --function-name $FUNCNAME \
    --zip-file fileb://deploy.zip \
    --runtime go1.x \
    --handler main \
    --role arn:aws:iam::785706276563:role/jalali  \
    --region $REGNAME
)
