package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseJoinCode struct {
	*response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}
