package agent

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/work/agent/workbench"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *workbench.Client) {

	client := NewClient(app)

	Workbench := workbench.NewClient(app)

	return client, Workbench

}
