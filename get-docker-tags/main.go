package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/get", handle)
	err := r.Run("0.0.0.0:9000")

	if err != nil {
		panic(err)
	}
}

func handle(c *gin.Context) {
	fullname := c.Query("name")

	if fullname == "" {
		log.Printf("name is empty")
		return
	}

	domain, name := parseName(fullname)

	versions, err := getVersion(domain, name)

	if err != nil {
		log.Printf("get version error: %s", err.Error())
		return
	}

	c.String(http.StatusOK, strings.Join(versions, ","))
	return
}

func getVersion(domain, name string) ([]string, error) {
	urlStr := fmt.Sprintf("https://%s/v2/%s/tags/list", domain, name)

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)

	if err != nil {
		return nil, err
	}

	if "ghcr.io" != domain {
		token, err := getToken(domain, name)

		if err != nil {
			return nil, err
		}

		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	type Data struct {
		Tags []string `json:"tags"`
	}

	var data Data

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return data.Tags, nil
}

func getToken(domain, name string) (string, error) {
	type Data struct {
		Token string `json:"token"`
	}

	var data Data

	urlStr := fmt.Sprintf("https://%s/v2", domain)

	req, err := http.Get(urlStr)

	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	if req.Header.Get("Docker-Distribution-Api-Version") != "registry/2.0" {
		return "", errors.New("不支持 v2")
	}

	bearer := req.Header.Get("Www-Authenticate")

	if bearer == "" {
		return "", errors.New("没有 Www-Authenticate")
	}

	bearers := strings.Split(bearer, "\"")

	urlStr = fmt.Sprintf("%s?service=%s&scope=repository:%s:pull", bearers[1], bearers[3], name)

	req, err = http.Get(urlStr)

	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", err
	}

	return data.Token, nil
}

func parseName(name string) (string, string) {
	if strings.Count(name, "/") <= 1 {
		if strings.Count(name, "/") == 0 {
			name = "library/" + name
		}

		return "registry-1.docker.io", name
	}

	p, err := url.Parse("https://" + name)

	if err != nil {
		return "", ""
	}

	return p.Host, p.Path[1:]
}
