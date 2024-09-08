package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://api-free.deepl.com/v2/translate"

type DeeplClient struct {
}

func (client *DeeplClient) Translate(sourceLang, targetLang, text string) (string, error) {
	payload := strings.NewReader(
		fmt.Sprintf(
			`{
			  "text": [
				%q
			  ],
			  "source_lang": %q,
			  "target_lang": %q
			}`,
			text,
			sourceLang,
			targetLang,
		),
	)

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "DeepL-Auth-Key d568f7f5-af8f-4f68-917f-3e50e7b0f1ac:fx")
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resultData map[string][]map[string]string

	err = json.Unmarshal(body, &resultData)
	if err != nil {
		return "", err
	}

	return resultData["translations"][0]["text"], nil
}
