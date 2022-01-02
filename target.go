package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

type Target struct {
	Name        string `yaml:"name"`
	Url         string `yaml:"url"`
	StatusCode  int    `yaml:"status_code"`
	InsecureReq bool   `yaml:"insecure_req"`
	ComponentId string `yaml:"component_id"`
}

func (target Target) request_target() (TargetStatus, error) {
	var client *http.Client

	log.Println("Debug:", target.Name, "Initiating request")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: target.InsecureReq},
	}
	client = &http.Client{Transport: tr, Timeout: 5 * time.Second}

	start := time.Now()
	req, err := http.NewRequest("GET", target.Url, nil)
	if err != nil {
		log.Println("Error:", target.Name, err)
		return MajorOutage, err
	}
	log.Println("Debug:", target.Name, "First byte:", time.Since(start))
	log.Println("Debug:", target.Name, "Requesting:", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error:", target.Name, target.Url, err)
		return MajorOutage, err
	}
	defer resp.Body.Close()
	log.Println("Debug:", target.Name, "Response Status:", resp.Status)
	log.Println("Debug:", target.Name, "Everything:", time.Since(start))
	return Operational, nil
}
