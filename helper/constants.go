package helper


var QnsAndCode = "Start fresh chat. You are a golang programmer, you need to create a golang project on the given topic.Give complete code without wrap.The topic is as follows:  \n" 
var QnsAndTopic = `Start fresh chat. You are a golang programmer, you need to create a golang project on the given topic.Give the complete folder structure of the project without wrap,the folder should be looks like ecommerce/
	├── main.go
	├── models/
	│   ├── user.go
	│   ├── product.go
	├── handlers/.
	Add or delete folders and files according your concept reagrding the topic.The topic is as follows:  \n` 

const ApiEndpoint = "https://api.openai.com/v1/chat/completions"