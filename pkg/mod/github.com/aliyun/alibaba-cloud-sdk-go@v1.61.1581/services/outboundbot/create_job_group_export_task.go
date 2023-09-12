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

// CreateJobGroupExportTask invokes the outboundbot.CreateJobGroupExportTask API synchronously
func (client *Client) CreateJobGroupExportTask(request *CreateJobGroupExportTaskRequest) (response *CreateJobGroupExportTaskResponse, err error) {
	response = CreateCreateJobGroupExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateJobGroupExportTaskWithChan invokes the outboundbot.CreateJobGroupExportTask API asynchronously
func (client *Client) CreateJobGroupExportTaskWithChan(request *CreateJobGroupExportTaskRequest) (<-chan *CreateJobGroupExportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateJobGroupExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateJobGroupExportTask(request)
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

// CreateJobGroupExportTaskWithCallback invokes the outboundbot.CreateJobGroupExportTask API asynchronously
func (client *Client) CreateJobGroupExportTaskWithCallback(request *CreateJobGroupExportTaskRequest, callback func(response *CreateJobGroupExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateJobGroupExportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateJobGroupExportTask(request)
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

// CreateJobGroupExportTaskRequest is the request struct for api CreateJobGroupExportTask
type CreateJobGroupExportTaskRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	JobGroupId string    `position:"Query" name:"JobGroupId"`
	Option     *[]string `position:"Query" name:"Option"  type:"Repeated"`
}

// CreateJobGroupExportTaskResponse is the response struct for api CreateJobGroupExportTask
type CreateJobGroupExportTaskResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateJobGroupExportTaskRequest creates a request to invoke CreateJobGroupExportTask API
func CreateCreateJobGroupExportTaskRequest() (request *CreateJobGroupExportTaskRequest) {
	request = &CreateJobGroupExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateJobGroupExportTask", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateJobGroupExportTaskResponse creates a response to parse from CreateJobGroupExportTask response
func CreateCreateJobGroupExportTaskResponse() (response *CreateJobGroupExportTaskResponse) {
	response = &CreateJobGroupExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
