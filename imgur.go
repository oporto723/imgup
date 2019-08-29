package main

import (
	"errors"
	"fmt"
)

// Sender sends a stream of bytes to a remote system.
type Sender interface {
	Send([]byte) error
}

// ImgurClient is an implementation of Sender that can send a stream of bytes to the imgur.com service.
type ImgurClient struct{}

// Send method implements Sender.
func (c ImgurClient) Send([]byte) error {
	fmt.Println("[ImgurClient] Sending file...")
	return nil
}

// GistClient is an implementation of Sender that can send a stream of bytes to the Gist service (part of GitHub).
type GistClient struct{}

// Send method implements Sender.
func (c GistClient) Send([]byte) error {
	fmt.Println("[GistClient] Sending file...")
	return nil
}

// client returns the corresponding Sender according to the service requested.
func client(service string) (Sender, error) {
	if service == "gist" {
		return GistClient{}, nil
	} else if service == "imgur" {
		return ImgurClient{}, nil
	}
	return nil, errors.New("unknown service")
}

// upload sends a stream of bytes to a Sender.
func upload(client Sender, stream []byte) error {
	return client.Send(stream)
}
