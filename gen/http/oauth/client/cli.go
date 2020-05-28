// Code generated by goa v3.1.1, DO NOT EDIT.
//
// oauth HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sm43/goa-crud/design

package client

import (
	"encoding/json"
	"fmt"

	oauth "github.com/sm43/goa-crud/gen/oauth"
)

// BuildOauthPayload builds the payload for the oauth oauth endpoint from CLI
// flags.
func BuildOauthPayload(oauthOauthBody string) (*oauth.OauthPayload, error) {
	var err error
	var body OauthRequestBody
	{
		err = json.Unmarshal([]byte(oauthOauthBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"token\": \"Non eaque omnis.\"\n   }'")
		}
	}
	v := &oauth.OauthPayload{
		Token: body.Token,
	}

	return v, nil
}
