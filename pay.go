package wechatpay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//统一下单
func (this *WechatPay) Pay(param UnitOrder) (*UnifyOrderResult, error) {
	param.Appid = this.AppId
	param.Mchid = this.MchId

	bytes_req, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	str_req := string(bytes_req)
	req, err := http.NewRequest("POST", UNIT_ORDER_URL, bytes.NewReader([]byte(str_req)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	w_req := http.Client{}
	resp, err := w_req.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var pay_result UnifyOrderResult
	err = json.Unmarshal(body, &pay_result)
	if err != nil {
		return nil, err
	}
	return &pay_result, nil
}
