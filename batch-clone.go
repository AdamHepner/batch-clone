package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

// RepoEntry is the format used by GH to report all basic information about a repo
// I only need one position here, so I'm sticking to basics
type RepoEntry struct {
	GitURL string `json:"ssh_url"`
}

func main() {
	app := cli.App{}

	var username string
	var token string

	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:        "username",
		Destination: &username,
	}, &cli.StringFlag{
		// @todo Add access-token flag as environment variable
		Name:        "access-token",
		Destination: &token,
	})

	app.Action = func(c *cli.Context) error {
		req := http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "https",
				Host:   "api.github.com",
				Path:   "/user/repos",
			},
			Header: make(http.Header, 0),
		}

		headers := req.Header
		headers.Add("accept", "application/vnd.github.v3+json")
		if len(token) > 0 {
			headers.Add("authorization", "token "+token)
		}

		client := http.Client{}

		resp, err := client.Do(&req)
		if err != nil {
			return err
		}

		log.WithFields(log.Fields{"Status": resp.Status}).Info("Got response")

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		data := []RepoEntry{}

		err = json.Unmarshal(buf.Bytes(), &data)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Fatal("Cannot unmarshal GH's respnse")
		}

		for _, e := range data {
			// @todo add functionality to actually clone the repositories, instead of only listing them
			fmt.Println(e.GitURL)
		}

		return nil
	}

	app.Run(os.Args)
}
