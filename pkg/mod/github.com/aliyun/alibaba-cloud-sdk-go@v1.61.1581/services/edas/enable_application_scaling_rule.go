package edas

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

// EnableApplicationScalingRule invokes the edas.EnableApplicationScalingRule API synchronously
func (client *Client) EnableApplicationScalingRule(request *EnableApplicationScalingRuleRequest) (response *EnableApplicationScalingRuleResponse, err error) {
	response = CreateEnableApplicationScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableApplicationScalingRuleWithChan invokes the edas.EnableApplicationScalingRule API asynchronously
func (client *Client) EnableApplicationScalingRuleWithChan(request *EnableApplicationScalingRuleRequest) (<-chan *EnableApplicationScalingRuleResponse, <-chan error) {
	responseChan := make(chan *EnableApplicationScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableApplicationScalingRule(request)
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

// EnableApplicationScalingRuleWithCallback invokes the edas.EnableApplicationScalingRule API asynchronously
func (client *Client) EnableApplicationScalingRuleWithCallback(request *EnableApplicationScalingRuleRequest, callback func(response *EnableApplicationScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableApplicationScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableApplicationScalingRule(request)
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

// EnableApplicationScalingRuleRequest is the request struct for api EnableApplicationScalingRule
type EnableApplicationScalingRuleRequest struct {
	*requests.RoaRequest
	ScalingRuleName string `position:"Query" name:"ScalingRuleName"`
	AppId           string `position:"Query" name:"AppId"`
}

// EnableApplicationScalingRuleResponse is the response struct for api EnableApplicationScalingRule
type EnableApplicationScalingRuleResponse struct {
	*responses.BaseResponse
	Code           int            `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	AppScalingRule AppScalingRule `json:"AppScalingRule" xml:"AppScalingRule"`
}

// CreateEnableApplicationScalingRuleRequest creates a request to invoke EnableApplicationScalingRule API
func CreateEnableApplicationScalingRuleRequest() (request *EnableApplicationScalingRuleRequest) {
	request = &EnableApplicationScalingRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "EnableApplicationScalingRule", "/pop/v1/eam/scale/enable_application_scaling_rule", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateEnableApplicationScalingRuleResponse creates a response to parse from EnableApplicationScalingRule response
func CreateEnableApplicationScalingRuleResponse() (response *EnableApplicationScalingRuleResponse) {
	response = &EnableApplicationScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
