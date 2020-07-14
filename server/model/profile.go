package model

import (
	"github.com/satori/go.uuid"
	"github.com/team142/angrychess/io/ws"
)

//CreateProfile for easy access to a new client
func createProfile(client *ws.Client) *Profile {
	return &Profile{Client: client, Secret: uuid.NewV4().String(), ID: uuid.NewV4().String()}
}

//Profile describes a client
type Profile struct {
	Client *ws.Client `json:"-"`
	Nick   string     `json:"nick"`
	ID     string     `json:"id"`
	Secret string     `json:"-"`
}
