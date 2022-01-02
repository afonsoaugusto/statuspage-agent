package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	ApiKey  string `yaml:"apiKey"`
	PageId  string `yaml:"pageId"`
	ApiBase string `yaml:"apiBase"`
}
type ConfigurationFile struct {
	Configuration Configuration `yaml:"configuration"`
	Targets       []Target      `yaml:"targets,flow"`
}

func request_targets(configuration ConfigurationFile) {
	for {
		for _, target := range configuration.Targets {
			// sleep 1 second between requests
			// https://developer.statuspage.io/#section/Rate-Limiting
			time.Sleep(time.Second)
			targetStatus, err := target.request_target()
			if err != nil {
				log.Println("Error:", target.Name, err)
			}
			log.Println("Info:", target.Name, "Target Status:", targetStatus)
			err = send_status_component(configuration.Configuration, target, targetStatus)
			if err != nil {
				log.Println("Error:", target.Name, err)
			}
		}
		if !true {
			break
		}
		// sleep 1 minute between interations of the loop
		time.Sleep(time.Second * 5)
	}
}
func send_status_component(configuration Configuration, target Target, targetStatus TargetStatus) error {
	log.Println("Info:", "send_status_component", "Target:", target.Name)
	params := url.Values{}
	params.Add("component[status]", targetStatus.String())
	body := strings.NewReader(params.Encode())

	endpoint := fmt.Sprintf("%s/pages/%s/components/%s", configuration.ApiBase, configuration.PageId, target.ComponentId)
	req, err := http.NewRequest("PATCH", endpoint, body)
	if err != nil {
		log.Println("Error:", endpoint, err)
	}
	req.Header.Set("Authorization", "OAuth "+configuration.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error:", target.Name, target.Url, err)
	}
	defer resp.Body.Close()

	return nil
}

func read_configration(filename string) (ConfigurationFile, error) {
	configuration := ConfigurationFile{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Info:", "Reading configuration file:", filename)
	err = yaml.Unmarshal(content, &configuration)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return configuration, nil
}

func main() {
	configuration, err := read_configration("targets.yaml")
	log.Println(configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	request_targets(configuration)
}
