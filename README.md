# linechatbot-on-lambda
This is a line chatbot demo code running on AWS lambda which calling AWS translate service in background.  You might be charged by AWS and please check AWS free-tier servides. You might be also charged by Line business account and please refer to Line Offical website for details.

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

# Test your line chatbot in your line messager
## Supported Commands
Translating English to Traditional Chinese by AWS translate service
> '2tw hello, I like to say hello!'

Translating Tranditional Chinese to English by AWS translate service
> '2en 你是在哈囉嗎？'

Command helper
> '@@' 

# Demo Video Link
https://www.youtube.com/embed/nTyU3pXkN_Q
