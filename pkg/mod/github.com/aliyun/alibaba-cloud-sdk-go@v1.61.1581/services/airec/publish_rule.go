package airec

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

// PublishRule invokes the airec.PublishRule API synchronously
func (client *Client) PublishRule(request *PublishRuleRequest) (response *PublishRuleResponse, err error) {
	response = CreatePublishRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PublishRuleWithChan invokes the airec.PublishRule API asynchronously
func (client *Client) PublishRuleWithChan(request *PublishRuleRequest) (<-chan *PublishRuleResponse, <-chan error) {
	responseChan := make(chan *PublishRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishRule(request)
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

// PublishRuleWithCallback invokes the airec.PublishRule API asynchronously
func (client *Client) PublishRuleWithCallback(request *PublishRuleRequest, callback func(response *PublishRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishRuleResponse
		var err error
		defer close(result)
		response, err = client.PublishRule(request)
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

// PublishRuleRequest is the request struct for api PublishRule
type PublishRuleRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	RuleType   string `position:"Query" name:"ruleType"`
	SceneId    string `position:"Query" name:"sceneId"`
	RuleId     string `position:"Path" name:"ruleId"`
}

// PublishRuleResponse is the response struct for api PublishRule
type PublishRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	Result    Result `json:"result" xml:"result"`
}

// CreatePublishRuleRequest creates a request to invoke PublishRule API
func CreatePublishRuleRequest() (request *PublishRuleRequest) {
	request = &PublishRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "PublishRule", "/v2/openapi/instances/[instanceId]/rules/[ruleId]/actions/publish", "airec", "openAPI")
	request.Method = requests.PUT
	return
}

// CreatePublishRuleResponse creates a response to parse from PublishRule response
func CreatePublishRuleResponse() (response *PublishRuleResponse) {
	response = &PublishRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
