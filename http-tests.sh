#!/bin/bash

ENDPOINT="http://localhost:8080"
PROJECT_NAME="http-service"
GIT_COMMIT=$(git rev-list -1 HEAD)
PASSED=true



#####################
# hello stranger test
#####################

TESR_NR=0

WANT_BODY="Hello Stranger"
WANT_CODE=200

RESPONSE=$(curl -i "$ENDPOINT"/helloworld 2> /dev/null)

RES_CODE=$(echo "$RESPONSE" | awk '{print $2}' | head -n 1)
RES_BODY=$(echo "$RESPONSE" | tail -n 1)

if [ "$WANT_BODY" != "$RES_BODY" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_BODY, got $RES_BODY"
fi

if [ "$WANT_CODE" != "$RES_CODE" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_CODE, got $RES_CODE"
fi

#####################
# versionz test
#####################

TESR_NR=1

WANT_BODY="{\"git_commit\":\"$GIT_COMMIT\",\"project_name\":\"$PROJECT_NAME\"}"
WANT_CODE=200

RESPONSE=$(curl -i "$ENDPOINT"/versionz 2> /dev/null)

RES_CODE=$(echo "$RESPONSE" | awk '{print $2}' | head -n 1)
RES_BODY=$(echo "$RESPONSE" | tail -n 1)

if [ "$WANT_BODY" != "$RES_BODY" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_BODY, got $RES_BODY"
fi

if [ "$WANT_CODE" != "$RES_CODE" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_CODE, got $RES_CODE"
fi

#######################
# helloworld query test
#######################

TESR_NR=2

WANT_BODY="Hello Edoardo Rizzardi"
WANT_CODE=200

RESPONSE=$(curl -i "$ENDPOINT/helloworld?name=EdoardoRizzardi" 2> /dev/null)

RES_CODE=$(echo "$RESPONSE" | awk '{print $2}' | head -n 1)
RES_BODY=$(echo "$RESPONSE" | tail -n 1)

if [ "$WANT_BODY" != "$RES_BODY" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_BODY, got $RES_BODY"
fi

if [ "$WANT_CODE" != "$RES_CODE" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_CODE, got $RES_CODE"
fi

###########################################
# helloworld query test with no query param
###########################################

TESR_NR=3

WANT_BODY="Invalid Request!"
WANT_CODE=400

RESPONSE=$(curl -i "$ENDPOINT/helloworld?pippo=franco" 2> /dev/null)

RES_CODE=$(echo "$RESPONSE" | awk '{print $2}' | head -n 1)
RES_BODY=$(echo "$RESPONSE" | tail -n 1)

if [ "$WANT_BODY" != "$RES_BODY" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_BODY, got $RES_BODY"
fi

if [ "$WANT_CODE" != "$RES_CODE" ]; then
    PASSED=false
    echo "Error in $TESR_NR"
    echo "Expected $WANT_CODE, got $RES_CODE"
fi

#####################
# ending
#####################

if [ "$PASSED" != true ]; then
    exit 1
fi

exit 0