package contract

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken(refresh bool) (resToken *response2.ResponseGetToken, err error)
		Refresh() AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
	}
)
