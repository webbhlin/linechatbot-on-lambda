# linechatbot-on-lambda
This is a line chatbot running on AWS lambda

# Prerequesites
1. need to enable your line developer account
2. need to add your webhook url - this is your api gateway URL
3. need to install serverless framework 
4. need to install aws cli and export your 

# Required Packages
1. gin gonic web framework
2. line-bot-sdk-go
3. aws-go-sdk

# Environment Variables
## export your AWS environment variables in your bashrc or bashprofile 
> export AWS_ACCESS_KEY_ID=<your access key>
> export AWS_SECRET_ACCESS_KEY=<your secret key>

# serverless command reference
## create a serverless go project
> sls serverless create -t aws-go-dep -p go-linechatbotlambda
deploy your serverless project
> sls deploy

## add your enpoint url to your line developer management console
> endpoints:
>   POST - https://xxxxx.execute-api.us-east-2.amazonaws.com/dev/hello

