 #!/bin/sh

 . $(dirname "$0")/config.sh
 
 go build  -ldflags "-w -s" -o main && \
 rm deploy.zip && \
 zip deploy.zip main && \
 aws lambda update-function-code \
    --function-name $FUNCNAME \
    --zip-file fileb://deploy.zip  \
    --region $REGNAME