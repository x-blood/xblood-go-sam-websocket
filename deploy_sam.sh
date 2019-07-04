#!/bin/sh

# Build on_connect
GOOS=linux go build ./handlers/on_connect/on_connect.go
mv ./on_connect ./handlers/on_connect

# Build on_disconnect
GOOS=linux go build ./handlers/on_disconnect/on_disconnect.go
mv ./on_disconnect ./handlers/on_disconnect

# Build send_message
GOOS=linux go build ./handlers/send_message/send_message.go
mv ./send_message ./handlers/send_message

# Create Package
sam package \
  --template-file ./template.yml \
  --output-template-file ./template-output.yml \
  --s3-bucket aws-sam-nested-application-packages \
  --profile temp_profile

# Deploy
sam deploy \
  --template-file ./template-output.yml \
  --stack-name xblood-go-sam-websocket \
  --capabilities CAPABILITY_IAM \
  --profile temp_profile
