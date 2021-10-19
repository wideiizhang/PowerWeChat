package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	*response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
