package scsp

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

// HangupCall invokes the scsp.HangupCall API synchronously
func (client *Client) HangupCall(request *HangupCallRequest) (response *HangupCallResponse, err error) {
	response = CreateHangupCallResponse()
	err = client.DoAction(request, response)
	return
}

// HangupCallWithChan invokes the scsp.HangupCall API asynchronously
func (client *Client) HangupCallWithChan(request *HangupCallRequest) (<-chan *HangupCallResponse, <-chan error) {
	responseChan := make(chan *HangupCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HangupCall(request)
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

// HangupCallWithCallback invokes the scsp.HangupCall API asynchronously
func (client *Client) HangupCallWithCallback(request *HangupCallRequest, callback func(response *HangupCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HangupCallResponse
		var err error
		defer close(result)
		response, err = client.HangupCall(request)
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

// HangupCallRequest is the request struct for api HangupCall
type HangupCallRequest struct {
	*requests.RpcRequest
	ClientToken  string `position:"Body"`
	InstanceId   string `position:"Body"`
	AccountName  string `position:"Body"`
	CallId       string `position:"Body"`
	JobId        string `position:"Body"`
	ConnectionId string `position:"Body"`
}

// HangupCallResponse is the response struct for api HangupCall
type HangupCallResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateHangupCallRequest creates a request to invoke HangupCall API
func CreateHangupCallRequest() (request *HangupCallRequest) {
	request = &HangupCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "HangupCall", "", "")
	request.Method = requests.POST
	return
}

// CreateHangupCallResponse creates a response to parse from HangupCall response
func CreateHangupCallResponse() (response *HangupCallResponse) {
	response = &HangupCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}