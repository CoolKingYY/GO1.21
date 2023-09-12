package outboundbot

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

// GetContactWhiteList invokes the outboundbot.GetContactWhiteList API synchronously
func (client *Client) GetContactWhiteList(request *GetContactWhiteListRequest) (response *GetContactWhiteListResponse, err error) {
	response = CreateGetContactWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// GetContactWhiteListWithChan invokes the outboundbot.GetContactWhiteList API asynchronously
func (client *Client) GetContactWhiteListWithChan(request *GetContactWhiteListRequest) (<-chan *GetContactWhiteListResponse, <-chan error) {
	responseChan := make(chan *GetContactWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContactWhiteList(request)
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

// GetContactWhiteListWithCallback invokes the outboundbot.GetContactWhiteList API asynchronously
func (client *Client) GetContactWhiteListWithCallback(request *GetContactWhiteListRequest, callback func(response *GetContactWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContactWhiteListResponse
		var err error
		defer close(result)
		response, err = client.GetContactWhiteList(request)
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

// GetContactWhiteListRequest is the request struct for api GetContactWhiteList
type GetContactWhiteListRequest struct {
	*requests.RpcRequest
	CountTotalRow requests.Boolean `position:"Query" name:"CountTotalRow"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// GetContactWhiteListResponse is the response struct for api GetContactWhiteList
type GetContactWhiteListResponse struct {
	*responses.BaseResponse
	HttpStatusCode       int                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                 string               `json:"Code" xml:"Code"`
	Message              string               `json:"Message" xml:"Message"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Success              bool                 `json:"Success" xml:"Success"`
	ContactWhitelistList ContactWhitelistList `json:"ContactWhitelistList" xml:"ContactWhitelistList"`
}

// CreateGetContactWhiteListRequest creates a request to invoke GetContactWhiteList API
func CreateGetContactWhiteListRequest() (request *GetContactWhiteListRequest) {
	request = &GetContactWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetContactWhiteList", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetContactWhiteListResponse creates a response to parse from GetContactWhiteList response
func CreateGetContactWhiteListResponse() (response *GetContactWhiteListResponse) {
	response = &GetContactWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
