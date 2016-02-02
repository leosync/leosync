package lingualeo

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
	"strings"
)

type linguaLeoClient struct {
	loginResp  loginResp
	connection connection
}

func (c linguaLeoClient) getRequest() *gorequest.SuperAgent {
	return c.connection.getClient()
}

func (c linguaLeoClient) AddWord(word, translation string) {
	resp, _, _ := c.getRequest().Get(linguaAddWordUrl).Query("word=" + word).Query("tword=" + translation).End()
	c.connection.validate(resp)
}

func (c linguaLeoClient) GetTranslations(word string) Word {
	resp, body, errs := c.getRequest().Get(linguaTranslateUrl).Query("word=" + word).End()
	c.connection.validate(resp)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	var translations Word
	translations.Value = word
	decodedError := json.NewDecoder(strings.NewReader(body)).Decode(&translations)
	if decodedError != nil {
		log.Fatalf("%v \n", errs)
	}
	return translations
}
