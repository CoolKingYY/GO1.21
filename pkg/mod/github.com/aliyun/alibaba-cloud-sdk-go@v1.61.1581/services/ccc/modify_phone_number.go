package ccc

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

// ModifyPhoneNumber invokes the ccc.ModifyPhoneNumber API synchronously
func (client *Client) ModifyPhoneNumber(request *ModifyPhoneNumberRequest) (response *ModifyPhoneNumberResponse, err error) {
	response = CreateModifyPhoneNumberResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPhoneNumberWithChan invokes the ccc.ModifyPhoneNumber API asynchronously
func (client *Client) ModifyPhoneNumberWithChan(request *ModifyPhoneNumberRequest) (<-chan *ModifyPhoneNumberResponse, <-chan error) {
	responseChan := make(chan *ModifyPhoneNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPhoneNumber(request)
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

// ModifyPhoneNumberWithCallback invokes the ccc.ModifyPhoneNumber API asynchronously
func (client *Client) ModifyPhoneNumberWithCallback(request *ModifyPhoneNumberRequest, callback func(response *ModifyPhoneNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPhoneNumberResponse
		var err error
		defer close(result)
		response, err = client.ModifyPhoneNumber(request)
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

// ModifyPhoneNumberRequest is the request struct for api ModifyPhoneNumber
type ModifyPhoneNumberRequest struct {
	*requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	Usage         string `position:"Query" name:"Usage"`
	Number        string `position:"Query" name:"Number"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// ModifyPhoneNumberResponse is the response struct for api ModifyPhoneNumber
type ModifyPhoneNumberResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPhoneNumberRequest creates a request to invoke ModifyPhoneNumber API
func CreateModifyPhoneNumberRequest() (request *ModifyPhoneNumberRequest) {
	request = &ModifyPhoneNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ModifyPhoneNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPhoneNumberResponse creates a response to parse from ModifyPhoneNumber response
func CreateModifyPhoneNumberResponse() (response *ModifyPhoneNumberResponse) {
	response = &ModifyPhoneNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
