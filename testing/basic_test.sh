#!/bin/bash

accountId=123
age=25

echo "--- Start Test ---"
resp=$(curl http://localhost:8081/start\?accountId\=$accountId)
echo $resp
workflowId=$(echo $resp | jq -r .workflowId)
sleep 3

echo "--- Continue Test ---"
curl http://localhost:8081/continue\?workflowId\=$workflowId\&age\=$age