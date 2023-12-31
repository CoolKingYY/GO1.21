package cms

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

// DeleteEventRuleTargets invokes the cms.DeleteEventRuleTargets API synchronously
func (client *Client) DeleteEventRuleTargets(request *DeleteEventRuleTargetsRequest) (response *DeleteEventRuleTargetsResponse, err error) {
	response = CreateDeleteEventRuleTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEventRuleTargetsWithChan invokes the cms.DeleteEventRuleTargets API asynchronously
func (client *Client) DeleteEventRuleTargetsWithChan(request *DeleteEventRuleTargetsRequest) (<-chan *DeleteEventRuleTargetsResponse, <-chan error) {
	responseChan := make(chan *DeleteEventRuleTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEventRuleTargets(request)
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

// DeleteEventRuleTargetsWithCallback invokes the cms.DeleteEventRuleTargets API asynchronously
func (client *Client) DeleteEventRuleTargetsWithCallback(request *DeleteEventRuleTargetsRequest, callback func(response *DeleteEventRuleTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEventRuleTargetsResponse
		var err error
		defer close(result)
		response, err = client.DeleteEventRuleTargets(request)
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

// DeleteEventRuleTargetsRequest is the request struct for api DeleteEventRuleTargets
type DeleteEventRuleTargetsRequest struct {
	*requests.RpcRequest
	RuleName string    `position:"Query" name:"RuleName"`
	Ids      *[]string `position:"Query" name:"Ids"  type:"Repeated"`
}

// DeleteEventRuleTargetsResponse is the response struct for api DeleteEventRuleTargets
type DeleteEventRuleTargetsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteEventRuleTargetsRequest creates a request to invoke DeleteEventRuleTargets API
func CreateDeleteEventRuleTargetsRequest() (request *DeleteEventRuleTargetsRequest) {
	request = &DeleteEventRuleTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteEventRuleTargets", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEventRuleTargetsResponse creates a response to parse from DeleteEventRuleTargets response
func CreateDeleteEventRuleTargetsResponse() (response *DeleteEventRuleTargetsResponse) {
	response = &DeleteEventRuleTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
