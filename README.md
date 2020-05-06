# linechatbot-on-lambda
This is a line chatbot demo code running on AWS lambda which calling AWS translate service in background.  You might be charged by AWS and please check AWS free-tier servides. You might be also charged by Line business account and please refer to Line Offical website for details.

# Cloud Architecture

![alt text](https://github.com/webbhlin/linechatbot-on-lambda/blob/master/img/AWSServerlesslinechatbot.png" Serverless Line ChatBot")



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

# Serverless framework command reference and code changes
## Create a serverless go project
> sls serverless create -t aws-go-dep -p go-linechatbotlambda

## Copy the main.go to your serverless project.
if you use serverless command to create your project, your main.go should be under /hello.  This is also what I use for this demo, and you are feel to change but require to update your serverless.yml file.

In your terminal:
1. copy the main.go file to your project by typeing "cp main.go to_your_serverless_project/hello/"
2. compile the go binary by typeing "make" - ensure your check your Makefile. Please check my Makefile sample.
3. fix the errors if any.  do you go get all the library you need? any typo? 
4. if there is no error, congrats~ you are free to go next and try "sls deploy"
![alt text](https://github.com/webbhlin/linechatbot-on-lambda/blob/master/img/makefile_sample.png "make file")

## deploy your serverless project to your lambda function
> sls deploy

fix the errors if any. 
![alt text](https://github.com/webbhlin/linechatbot-on-lambda/blob/master/img/sls_deploy.png "sls deploy")


## add your enpoint url to your line developer management console
It might take 3-5 minutes to be ready after you implement your lambda function.  In the line console, you have to update your webhook url by clicking "edit" button and test the webhook by clicking on "verify"

![alt text](https://github.com/webbhlin/linechatbot-on-lambda/blob/master/img/webhook_line_console.png "line console")


You will see a success message in a pop window of line console.  If you don't see the success message, go to your cloudwatch log to see what happened.
![alt text](https://github.com/webbhlin/linechatbot-on-lambda/blob/master/img/cloudwatch_log.png "cloudwatch log")


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


# Reference links

serverless.yml guideline - https://www.serverless.com/framework/docs/providers/aws/guide/functions/

linechatbot sample code - https://github.com/kkdai/LineBotTemplate/blob/master/main.go

serverless golang examples - https://www.serverless.com/blog/framework-example-golang-lambda-support/

gin gonic github - https://github.com/gin-gonic

gin-lambda github - https://github.com/appleboy/gin-lambda

serverless examples [very useful]  - https://github.com/serverless/examples

### special thanks for @pahuddev :-) 
