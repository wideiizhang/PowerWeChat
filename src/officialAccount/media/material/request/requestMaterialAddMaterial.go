package request

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"

type RequestMaterialAddMaterial struct {
	Type         string         `json:"type"`
	Media        *power.HashMap `json:"media"`
	Title        string         `json:"title"`
	Introduction string         `json:"introduction"`

}
