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

// GetTaskByUuid invokes the outboundbot.GetTaskByUuid API synchronously
func (client *Client) GetTaskByUuid(request *GetTaskByUuidRequest) (response *GetTaskByUuidResponse, err error) {
	response = CreateGetTaskByUuidResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskByUuidWithChan invokes the outboundbot.GetTaskByUuid API asynchronously
func (client *Client) GetTaskByUuidWithChan(request *GetTaskByUuidRequest) (<-chan *GetTaskByUuidResponse, <-chan error) {
	responseChan := make(chan *GetTaskByUuidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskByUuid(request)
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

// GetTaskByUuidWithCallback invokes the outboundbot.GetTaskByUuid API asynchronously
func (client *Client) GetTaskByUuidWithCallback(request *GetTaskByUuidRequest, callback func(response *GetTaskByUuidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskByUuidResponse
		var err error
		defer close(result)
		response, err = client.GetTaskByUuid(request)
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

// GetTaskByUuidRequest is the request struct for api GetTaskByUuid
type GetTaskByUuidRequest struct {
	*requests.RpcRequest
	WithConversations requests.Boolean `position:"Query" name:"WithConversations"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	TaskId            string           `position:"Query" name:"TaskId"`
}

// GetTaskByUuidResponse is the response struct for api GetTaskByUuid
type GetTaskByUuidResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Task      TaskInGetTaskByUuid `json:"Task" xml:"Task"`
}

// CreateGetTaskByUuidRequest creates a request to invoke GetTaskByUuid API
func CreateGetTaskByUuidRequest() (request *GetTaskByUuidRequest) {
	request = &GetTaskByUuidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetTaskByUuid", "outboundbot", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetTaskByUuidResponse creates a response to parse from GetTaskByUuid response
func CreateGetTaskByUuidResponse() (response *GetTaskByUuidResponse) {
	response = &GetTaskByUuidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
