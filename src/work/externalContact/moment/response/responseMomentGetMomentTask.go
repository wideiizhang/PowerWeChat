package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMomentGetMomentTask struct {
	*response.ResponseWork

	TaskList   []*power.HashMap `json:"task_list"`
	NextCursor string           `json:"next_cursor"`
}
