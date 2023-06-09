package request

import "encoding/xml"

type RequestSendMiniProgramNormal struct {
	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ActName     string   `xml:"act_name"`
	MchBillNO   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	NonceStr    string   `xml:"nonce_str"`
	NotifyWay   string   `xml:"notify_way"`
	ReOpenID    string   `xml:"re_openid"`
	Remark      string   `xml:"remark"`
	SendName    string   `xml:"send_name"`
	TotalAmount string   `xml:"total_amount"`
	TotalNum    int      `xml:"total_num"`
	Wishing     int      `xml:"wishing"`
	WXAppID     string   `xml:"wxappid"`
	Sign        string   `xml:"sign"`
}
