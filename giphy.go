package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Giphyclient is an implementation of Sender that can send a stream of bytes
// to the Giphy service.
type GiphyClient struct{}

const (
	// Environment variable where the user places the API key.
	GiphyAPIKeyEnv = "GIPHY_API_KEY"

	// Giphy API base URL.
	GiphyUploadURL = "http://upload.giphy.com/v1/gifs"
)

// Send method implements Sender.
//
// Giphy expects a Multipart message:
// https://en.wikipedia.org/wiki/MIME#Multipart_messages.
//
func (c GiphyClient) Send(stream []byte) error {
	// Look up the API key from the environment variable.
	apiKey, ok := os.LookupEnv(GiphyAPIKeyEnv)
	if !ok {
		return fmt.Errorf("Environment variable %s not defined", GiphyAPIKeyEnv)
	} else if apiKey == "" {
		return fmt.Errorf("Environment variable %s is empty", GiphyAPIKeyEnv)
	}

	body := &bytes.Buffer{}

	// Our multipart message writer.
	writer := multipart.NewWriter(body)

	// Include "api_key" field.
	label, _ := writer.CreateFormField("api_key")
	label.Write([]byte(apiKey))

	// Include "file" field.
	part, err := writer.CreateFormFile("file", "image.gif")
	if err != nil {
		return err
	}
	_, err = io.Copy(part, bytes.NewReader(stream))
	if err != nil {
		return err
	}

	// Close the writer.
	if err := writer.Close(); err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", GiphyUploadURL, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("[GiphyClient] Status code: %d\n", resp.StatusCode)
	fmt.Printf("[GiphyClient] Response from the server: %s\n", string(respBody))

	return nil
}
