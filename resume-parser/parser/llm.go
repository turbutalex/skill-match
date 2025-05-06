package parser

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"resume-parser/models"
	"resume-parser/utils"
)

type OpenRouterRequestData struct {
	Model    string
	Messages []MessageStructure
}

type MessageStructure struct {
	Role    string
	Content string
}

var prompt = "I am going to send you a raw text containing some information processed from a CV.\nI want you process this data and from it to output the email, parsedName, phone and skills\nin JSON format like so:\n{\nemail: String\nparsedName: String\nphone: String\nskills: String[]\n}\nI do not want any other boilerplate text, so do not respond anything else before or after the JSON response.\nRaw Text:\n"

func SendRequest(message string) (models.Resume, error) {
	promptRequest := prompt + message

	messageStructure := MessageStructure{
		Role:    "user",
		Content: promptRequest,
	}
	requestData := OpenRouterRequestData{
		Model:    "",
		Messages: []MessageStructure{messageStructure},
	}

	requestDataJSON, err := json.Marshal(requestData)
	if err != nil {
		return models.Resume{}, err
	}

	req, err := http.NewRequest("POST", utils.GetEnv(utils.LLM_APIURL, "DefaultUrl"), bytes.NewBuffer(requestDataJSON))
	if err != nil {
		return models.Resume{}, err
	}

	req.Header.Add("Authorization", "Bearer "+utils.GetEnv(utils.LLM_APIKEY, "SomeKey"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.Resume{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Resume{}, err
	}

	var response any
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.Resume{}, err
	}

	resume := models.Resume{
		ParsedName: "ABC",
	}

	return models.Resume{}, errors.New("Not Yet functional")

	print(string(body))
	return resume, nil
}
