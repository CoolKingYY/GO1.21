package arms

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

// ListDispatchRule invokes the arms.ListDispatchRule API synchronously
func (client *Client) ListDispatchRule(request *ListDispatchRuleRequest) (response *ListDispatchRuleResponse, err error) {
	response = CreateListDispatchRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ListDispatchRuleWithChan invokes the arms.ListDispatchRule API asynchronously
func (client *Client) ListDispatchRuleWithChan(request *ListDispatchRuleRequest) (<-chan *ListDispatchRuleResponse, <-chan error) {
	responseChan := make(chan *ListDispatchRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDispatchRule(request)
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

// ListDispatchRuleWithCallback invokes the arms.ListDispatchRule API asynchronously
func (client *Client) ListDispatchRuleWithCallback(request *ListDispatchRuleRequest, callback func(response *ListDispatchRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDispatchRuleResponse
		var err error
		defer close(result)
		response, err = client.ListDispatchRule(request)
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

// ListDispatchRuleRequest is the request struct for api ListDispatchRule
type ListDispatchRuleRequest struct {
	*requests.RpcRequest
	System      requests.Boolean `position:"Query" name:"System"`
	Name        string           `position:"Query" name:"Name"`
	ProxyUserId string           `position:"Query" name:"ProxyUserId"`
}

// ListDispatchRuleResponse is the response struct for api ListDispatchRule
type ListDispatchRuleResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	DispatchRules []DispatchRule `json:"DispatchRules" xml:"DispatchRules"`
}

// CreateListDispatchRuleRequest creates a request to invoke ListDispatchRule API
func CreateListDispatchRuleRequest() (request *ListDispatchRuleRequest) {
	request = &ListDispatchRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListDispatchRule", "", "")
	request.Method = requests.POST
	return
}

// CreateListDispatchRuleResponse creates a response to parse from ListDispatchRule response
func CreateListDispatchRuleResponse() (response *ListDispatchRuleResponse) {
	response = &ListDispatchRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
