package bitbank

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/bitbank-go-client/v1/models"
)

func (c *Client) GetAssets(ctx context.Context) (*models.Assets, error) {
	spath := fmt.Sprintf("/user/assets")
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var assets models.Assets
	if err := decodeBody(res, &assets); err != nil {
		return nil, err
	}

	return &assets, nil
}
