package lingualeo

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/parnurzeal/gorequest"
)

func createConnection(config *config) connection {
	conn := connection{
		config: config,
		client: gorequest.New(),
	}
	return conn
}

type connection struct {
	config    *config
	loginResp *loginResp
	client    *gorequest.SuperAgent
}

func (c connection) getClient() *gorequest.SuperAgent {
	if c.client == nil {
		c.client = gorequest.New()
	}
	if c.loginResp == nil {
		c.authLeo()
	}

	return c.client
}

type loginResp struct {
	ErrorMsg string `json:"error_msg"`
	User     User   `json:"user"`
}

func (c connection) validate(resp gorequest.Response) {
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Something went wrong: %v", resp.Status)
	}
}

func (c connection) authLeo() {
	connClient := c.client //.Timeout(httpTimeout*time.Second)
	config := c.config
	resp, body, errs := connClient.Get(linguaLoginUrl).Query("email=" + config.Email).Query("password=" + config.Password).End()
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed login %v", resp.Status)
	}
	var loginResp loginResp
	err := json.NewDecoder(strings.NewReader(body)).Decode(&loginResp)
	if err != nil {
		log.Fatalf("Failed decode %v", body)
	}
	c.loginResp = &loginResp
}
