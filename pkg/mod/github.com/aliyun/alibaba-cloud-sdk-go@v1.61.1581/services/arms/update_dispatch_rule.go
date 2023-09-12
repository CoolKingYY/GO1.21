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

// UpdateDispatchRule invokes the arms.UpdateDispatchRule API synchronously
func (client *Client) UpdateDispatchRule(request *UpdateDispatchRuleRequest) (response *UpdateDispatchRuleResponse, err error) {
	response = CreateUpdateDispatchRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDispatchRuleWithChan invokes the arms.UpdateDispatchRule API asynchronously
func (client *Client) UpdateDispatchRuleWithChan(request *UpdateDispatchRuleRequest) (<-chan *UpdateDispatchRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateDispatchRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDispatchRule(request)
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

// UpdateDispatchRuleWithCallback invokes the arms.UpdateDispatchRule API asynchronously
func (client *Client) UpdateDispatchRuleWithCallback(request *UpdateDispatchRuleRequest, callback func(response *UpdateDispatchRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDispatchRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateDispatchRule(request)
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

// UpdateDispatchRuleRequest is the request struct for api UpdateDispatchRule
type UpdateDispatchRuleRequest struct {
	*requests.RpcRequest
	DispatchRule string `position:"Query" name:"DispatchRule"`
	ProxyUserId  string `position:"Query" name:"ProxyUserId"`
}

// UpdateDispatchRuleResponse is the response struct for api UpdateDispatchRule
type UpdateDispatchRuleResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDispatchRuleRequest creates a request to invoke UpdateDispatchRule API
func CreateUpdateDispatchRuleRequest() (request *UpdateDispatchRuleRequest) {
	request = &UpdateDispatchRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateDispatchRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDispatchRuleResponse creates a response to parse from UpdateDispatchRule response
func CreateUpdateDispatchRuleResponse() (response *UpdateDispatchRuleResponse) {
	response = &UpdateDispatchRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
