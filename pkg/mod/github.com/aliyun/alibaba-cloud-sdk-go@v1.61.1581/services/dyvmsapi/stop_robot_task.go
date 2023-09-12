package dyvmsapi

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

// StopRobotTask invokes the dyvmsapi.StopRobotTask API synchronously
func (client *Client) StopRobotTask(request *StopRobotTaskRequest) (response *StopRobotTaskResponse, err error) {
	response = CreateStopRobotTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopRobotTaskWithChan invokes the dyvmsapi.StopRobotTask API asynchronously
func (client *Client) StopRobotTaskWithChan(request *StopRobotTaskRequest) (<-chan *StopRobotTaskResponse, <-chan error) {
	responseChan := make(chan *StopRobotTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopRobotTask(request)
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

// StopRobotTaskWithCallback invokes the dyvmsapi.StopRobotTask API asynchronously
func (client *Client) StopRobotTaskWithCallback(request *StopRobotTaskRequest, callback func(response *StopRobotTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopRobotTaskResponse
		var err error
		defer close(result)
		response, err = client.StopRobotTask(request)
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

// StopRobotTaskRequest is the request struct for api StopRobotTask
type StopRobotTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TaskId               requests.Integer `position:"Query" name:"TaskId"`
}

// StopRobotTaskResponse is the response struct for api StopRobotTask
type StopRobotTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateStopRobotTaskRequest creates a request to invoke StopRobotTask API
func CreateStopRobotTaskRequest() (request *StopRobotTaskRequest) {
	request = &StopRobotTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "StopRobotTask", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopRobotTaskResponse creates a response to parse from StopRobotTask response
func CreateStopRobotTaskResponse() (response *StopRobotTaskResponse) {
	response = &StopRobotTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}