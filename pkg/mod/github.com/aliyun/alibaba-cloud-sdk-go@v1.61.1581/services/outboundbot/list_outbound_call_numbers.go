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

// ListOutboundCallNumbers invokes the outboundbot.ListOutboundCallNumbers API synchronously
func (client *Client) ListOutboundCallNumbers(request *ListOutboundCallNumbersRequest) (response *ListOutboundCallNumbersResponse, err error) {
	response = CreateListOutboundCallNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// ListOutboundCallNumbersWithChan invokes the outboundbot.ListOutboundCallNumbers API asynchronously
func (client *Client) ListOutboundCallNumbersWithChan(request *ListOutboundCallNumbersRequest) (<-chan *ListOutboundCallNumbersResponse, <-chan error) {
	responseChan := make(chan *ListOutboundCallNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOutboundCallNumbers(request)
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

// ListOutboundCallNumbersWithCallback invokes the outboundbot.ListOutboundCallNumbers API asynchronously
func (client *Client) ListOutboundCallNumbersWithCallback(request *ListOutboundCallNumbersRequest, callback func(response *ListOutboundCallNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOutboundCallNumbersResponse
		var err error
		defer close(result)
		response, err = client.ListOutboundCallNumbers(request)
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

// ListOutboundCallNumbersRequest is the request struct for api ListOutboundCallNumbers
type ListOutboundCallNumbersRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListOutboundCallNumbersResponse is the response struct for api ListOutboundCallNumbers
type ListOutboundCallNumbersResponse struct {
	*responses.BaseResponse
	HttpStatusCode      int                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                string              `json:"Code" xml:"Code"`
	Message             string              `json:"Message" xml:"Message"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	Success             bool                `json:"Success" xml:"Success"`
	OutboundCallNumbers OutboundCallNumbers `json:"OutboundCallNumbers" xml:"OutboundCallNumbers"`
}

// CreateListOutboundCallNumbersRequest creates a request to invoke ListOutboundCallNumbers API
func CreateListOutboundCallNumbersRequest() (request *ListOutboundCallNumbersRequest) {
	request = &ListOutboundCallNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListOutboundCallNumbers", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListOutboundCallNumbersResponse creates a response to parse from ListOutboundCallNumbers response
func CreateListOutboundCallNumbersResponse() (response *ListOutboundCallNumbersResponse) {
	response = &ListOutboundCallNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
