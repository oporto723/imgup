package main

import (
	"fmt"
)

// GistClient is an implementation of Sender that can send a stream of bytes to the Gist service (part of GitHub).
type GistClient struct{}

// Send method implements Sender.
func (c GistClient) Send([]byte) error {
	fmt.Println("[GistClient] Sending file...")
	return nil
}