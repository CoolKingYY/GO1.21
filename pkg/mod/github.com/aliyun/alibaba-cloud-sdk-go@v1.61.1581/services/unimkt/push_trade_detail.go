package unimkt

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// PushTradeDetail invokes the unimkt.PushTradeDetail API synchronously
func (client *Client) PushTradeDetail(request *PushTradeDetailRequest) (response *PushTradeDetailResponse, err error) {
	response = CreatePushTradeDetailResponse()
	err = client.DoAction(request, response)
	return
}

// PushTradeDetailWithChan invokes the unimkt.PushTradeDetail API asynchronously
func (client *Client) PushTradeDetailWithChan(request *PushTradeDetailRequest) (<-chan *PushTradeDetailResponse, <-chan error) {
	responseChan := make(chan *PushTradeDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushTradeDetail(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// PushTradeDetailWithCallback invokes the unimkt.PushTradeDetail API asynchronously
func (client *Client) PushTradeDetailWithCallback(request *PushTradeDetailRequest, callback func(response *PushTradeDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushTradeDetailResponse
		var err error
		defer close(result)
		response, err = client.PushTradeDetail(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// PushTradeDetailRequest is the request struct for api PushTradeDetail
type PushTradeDetailRequest struct {
	*requests.RpcRequest
	SalePrice          requests.Float   `position:"Body" name:"SalePrice"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	TradeStatus        requests.Integer `position:"Body" name:"TradeStatus"`
	CommodityId        string           `position:"Body" name:"CommodityId"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	TradeOrderId       string           `position:"Body" name:"TradeOrderId"`
	DeviceSn           string           `position:"Body" name:"DeviceSn"`
	CommodityName      string           `position:"Body" name:"CommodityName"`
	VerificationStatus requests.Integer `position:"Body" name:"VerificationStatus"`
	AlipayOrderId      string           `position:"Body" name:"AlipayOrderId"`
	ChannelId          string           `position:"Body" name:"ChannelId"`
	OuterTradeId       string           `position:"Body" name:"OuterTradeId"`
	TradeTime          requests.Integer `position:"Body" name:"TradeTime"`
	TradePrice         requests.Float   `position:"Body" name:"TradePrice"`
}

// PushTradeDetailResponse is the response struct for api PushTradeDetail
type PushTradeDetailResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePushTradeDetailRequest creates a request to invoke PushTradeDetail API
func CreatePushTradeDetailRequest() (request *PushTradeDetailRequest) {
	request = &PushTradeDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "PushTradeDetail", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushTradeDetailResponse creates a response to parse from PushTradeDetail response
func CreatePushTradeDetailResponse() (response *PushTradeDetailResponse) {
	response = &PushTradeDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
