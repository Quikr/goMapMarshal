package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

/*
   Scenario: when ClientMetadata of a json payload on the server is a key value scenario.
    metadata is returned as "client_metadata" : {}{}    // Unknown struct

   This example is about how to read and write out a map[string]string in a streamlined manner.

*/

var jsoncontent = `
{
	"tenant": "workforce",
	"client_metadata": {
		"ClientIds": "1,2125,5",
		"tenantDB": "ABC-001",
		"tenantId": "9"
	}
}
`

type Client struct {
	Tenant string `json:"tenant"`

	// ClientMetadata struct {}{}  this what client struct is but it's a key map
	CM *CMeta `json:"client_metadata"`
}

// Wrap the map as we need a instance and some supporting methods
type CMeta struct {
	CMetadata map[string]string
}

func (cm CMeta) UnmarshalJSON(b []byte) (err error) {
	// Unmarshal it into the map - if you don't know what it is then consider using map[string]json.RawMessage
	// but place it into a temp map then transpose it.
	if cm.CMetadata != nil {
		if err = json.Unmarshal(b, &cm.CMetadata); err != nil {
			return err
		}
	} else {
		return errors.New("instance of CMeta is nil, Parent has not been instantiated")
	}

	return nil
}

func (cm CMeta) MarshalJSON() ([]byte, error) {
	// write out in single manner  NOTE: if you remove this func then you will get a child json
	out := make(map[string]string)

	for i, v := range cm.CMetadata {
		out[i] = v
	}

	return json.Marshal(out)
}

// need an instance that wraps the map
func NewCMeta() *CMeta {
	return &CMeta{
		CMetadata: make(map[string]string),
	}
}

func main() {
	fmt.Println(fmt.Sprintf("Starting with -  %s", jsoncontent))
	Client := &Client{
		Tenant: "",
		CM:     NewCMeta(),
	}

	json.Unmarshal([]byte(jsoncontent), Client)
	fmt.Println(fmt.Sprintf("%+v", Client))

	// Write the content out does it match?
	Data, _ := json.Marshal(Client)

	fmt.Println(fmt.Sprintf("Ends up with - %s", Data))
}
