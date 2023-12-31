package sddp

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

// ModifyRuleStatus invokes the sddp.ModifyRuleStatus API synchronously
func (client *Client) ModifyRuleStatus(request *ModifyRuleStatusRequest) (response *ModifyRuleStatusResponse, err error) {
	response = CreateModifyRuleStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRuleStatusWithChan invokes the sddp.ModifyRuleStatus API asynchronously
func (client *Client) ModifyRuleStatusWithChan(request *ModifyRuleStatusRequest) (<-chan *ModifyRuleStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyRuleStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRuleStatus(request)
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

// ModifyRuleStatusWithCallback invokes the sddp.ModifyRuleStatus API asynchronously
func (client *Client) ModifyRuleStatusWithCallback(request *ModifyRuleStatusRequest, callback func(response *ModifyRuleStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRuleStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyRuleStatus(request)
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

// ModifyRuleStatusRequest is the request struct for api ModifyRuleStatus
type ModifyRuleStatusRequest struct {
	*requests.RpcRequest
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Ids         string           `position:"Query" name:"Ids"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
	Status      requests.Integer `position:"Query" name:"Status"`
}

// ModifyRuleStatusResponse is the response struct for api ModifyRuleStatus
type ModifyRuleStatusResponse struct {
	*responses.BaseResponse
	FailedIds string `json:"FailedIds" xml:"FailedIds"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRuleStatusRequest creates a request to invoke ModifyRuleStatus API
func CreateModifyRuleStatusRequest() (request *ModifyRuleStatusRequest) {
	request = &ModifyRuleStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "ModifyRuleStatus", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyRuleStatusResponse creates a response to parse from ModifyRuleStatus response
func CreateModifyRuleStatusResponse() (response *ModifyRuleStatusResponse) {
	response = &ModifyRuleStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
