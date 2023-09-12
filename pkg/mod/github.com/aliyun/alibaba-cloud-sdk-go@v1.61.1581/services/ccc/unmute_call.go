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

// UnmuteCall invokes the ccc.UnmuteCall API synchronously
func (client *Client) UnmuteCall(request *UnmuteCallRequest) (response *UnmuteCallResponse, err error) {
	response = CreateUnmuteCallResponse()
	err = client.DoAction(request, response)
	return
}

// UnmuteCallWithChan invokes the ccc.UnmuteCall API asynchronously
func (client *Client) UnmuteCallWithChan(request *UnmuteCallRequest) (<-chan *UnmuteCallResponse, <-chan error) {
	responseChan := make(chan *UnmuteCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnmuteCall(request)
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

// UnmuteCallWithCallback invokes the ccc.UnmuteCall API asynchronously
func (client *Client) UnmuteCallWithCallback(request *UnmuteCallRequest, callback func(response *UnmuteCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnmuteCallResponse
		var err error
		defer close(result)
		response, err = client.UnmuteCall(request)
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

// UnmuteCallRequest is the request struct for api UnmuteCall
type UnmuteCallRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	DeviceId   string `position:"Query" name:"DeviceId"`
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	ChannelId  string `position:"Query" name:"ChannelId"`
}

// UnmuteCallResponse is the response struct for api UnmuteCall
type UnmuteCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateUnmuteCallRequest creates a request to invoke UnmuteCall API
func CreateUnmuteCallRequest() (request *UnmuteCallRequest) {
	request = &UnmuteCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "UnmuteCall", "", "")
	request.Method = requests.POST
	return
}

// CreateUnmuteCallResponse creates a response to parse from UnmuteCall response
func CreateUnmuteCallResponse() (response *UnmuteCallResponse) {
	response = &UnmuteCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
