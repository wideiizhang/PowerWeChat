package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseJournalGetRecordList struct {
	*response.ResponseWork

	JournalUUIDList []string `json:"journaluuid_list"`
	NextCursor      int      `json:"next_cursor"`
	EndFlag         bool     `json:"endflag"`
}
