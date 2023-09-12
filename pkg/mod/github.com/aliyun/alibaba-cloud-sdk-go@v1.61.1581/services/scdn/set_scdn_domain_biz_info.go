package scdn

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

// SetScdnDomainBizInfo invokes the scdn.SetScdnDomainBizInfo API synchronously
func (client *Client) SetScdnDomainBizInfo(request *SetScdnDomainBizInfoRequest) (response *SetScdnDomainBizInfoResponse, err error) {
	response = CreateSetScdnDomainBizInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SetScdnDomainBizInfoWithChan invokes the scdn.SetScdnDomainBizInfo API asynchronously
func (client *Client) SetScdnDomainBizInfoWithChan(request *SetScdnDomainBizInfoRequest) (<-chan *SetScdnDomainBizInfoResponse, <-chan error) {
	responseChan := make(chan *SetScdnDomainBizInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetScdnDomainBizInfo(request)
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

// SetScdnDomainBizInfoWithCallback invokes the scdn.SetScdnDomainBizInfo API asynchronously
func (client *Client) SetScdnDomainBizInfoWithCallback(request *SetScdnDomainBizInfoRequest, callback func(response *SetScdnDomainBizInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetScdnDomainBizInfoResponse
		var err error
		defer close(result)
		response, err = client.SetScdnDomainBizInfo(request)
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

// SetScdnDomainBizInfoRequest is the request struct for api SetScdnDomainBizInfo
type SetScdnDomainBizInfoRequest struct {
	*requests.RpcRequest
	BizName    string           `position:"Query" name:"BizName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// SetScdnDomainBizInfoResponse is the response struct for api SetScdnDomainBizInfo
type SetScdnDomainBizInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetScdnDomainBizInfoRequest creates a request to invoke SetScdnDomainBizInfo API
func CreateSetScdnDomainBizInfoRequest() (request *SetScdnDomainBizInfoRequest) {
	request = &SetScdnDomainBizInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "SetScdnDomainBizInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateSetScdnDomainBizInfoResponse creates a response to parse from SetScdnDomainBizInfo response
func CreateSetScdnDomainBizInfoResponse() (response *SetScdnDomainBizInfoResponse) {
	response = &SetScdnDomainBizInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}