package main

import (
	"app-generator/helper"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	apiKey := os.Getenv("API_KEY")
	var input string
	fmt.Println("Enter a topic")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)
	fmt.Println("IPIS", input)

	qnsAndCode := helper.QnsAndCode + input
	qnsAndTopic := helper.QnsAndTopic + input
	answer, err := getAnswerFromOpenAI(qnsAndCode, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	folderStructure, err := getAnswerFromOpenAI(qnsAndTopic, apiKey)
	if err != nil {
		fmt.Println(err)
	}

	if err = os.WriteFile("./code/code.go", []byte(answer), 0644); err != nil {
		fmt.Println("Here is the error", err)
	}
	if err = os.WriteFile("./code/output.txt", []byte(folderStructure), 0644); err != nil {
		fmt.Println("Here is the error", err)
	}

	baseDir := "./project" // Set the base directory
	os.MkdirAll(baseDir, os.ModePerm)
	helper.CreateFolderStructure(folderStructure, baseDir)

}

func getAnswerFromOpenAI(question string, apiKey string) (string, error) {
	client := resty.New()
	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]interface{}{{"role": "user", "content": question}},
		}).
		Post(helper.ApiEndpoint)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return "", err
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, nil

}
