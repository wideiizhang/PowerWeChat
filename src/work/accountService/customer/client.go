package customer

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/customer/request"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/customer/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 读取消息
// https://work.weixin.qq.com/api/doc/90000/90135/94670
func (comp *Client) BatchGet(externalUserIDList []string) (*response.ResponseCustomerBatchGet, error) {

	result := &response.ResponseCustomerBatchGet{}

	options := &object.HashMap{
		"external_userid_list": externalUserIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/customer/batchget", options, nil, nil, result)

	return result, err
}

// 获取配置的专员与客户群
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func (comp *Client) GetUpgradeServiceConfig() (*response.ResponseCustomerGetUpgradeServiceConfig, error) {

	result := &response.ResponseCustomerGetUpgradeServiceConfig{}

	_, err := comp.HttpGet("cgi-bin/kf/customer/get_upgrade_service_config", nil, nil, result)

	return result, err
}

// 为客户升级为专员或客户群服务
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func (comp *Client) UpgradeService(options *request.RequestUpgradeService) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/kf/customer/upgrade_service", options, nil, nil, result)

	return result, err
}

// 为客户取消推荐
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func (comp *Client) CancelUpgradeService(openKFID, externalUserID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}
	options := &power.StringMap{
		"open_kfid": openKFID,
		"external_userid": externalUserID,
	}

	_, err := comp.HttpPostJson("cgi-bin/kf/customer/cancel_upgrade_service", options, nil, nil, result)

	return result, err
}
