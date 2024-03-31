package req

import (
	errors "canvas/error"
	"fmt"
	"net/http"
	"os"

	yml "gopkg.in/yaml.v3"
)

type Config struct {
	URL   string `yml:"url"`
	TOKEN string `yml:"token"`
}

func Auth() *http.Response {
	config, err := readConfig()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", config.URL, nil)

	if err != nil {
		errors.ErrorLog(err)
	}

	token := "Bearer" + config.TOKEN
	req.Header = http.Header{
		"Authorization": {token},
	}

	resp, err := client.Do(req)

	if err != nil {
		errors.ErrorLog(err)
	}

	return resp

}

func readConfig() (*Config, error) {
	config := &Config{}

	ymlConfig, err := os.ReadFile("/Users/nikhilgudur/go/src/github.com/nikhilgudur/canvas-cli/config.yml")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	err = yml.Unmarshal(ymlConfig, config)

	return config, err
}
