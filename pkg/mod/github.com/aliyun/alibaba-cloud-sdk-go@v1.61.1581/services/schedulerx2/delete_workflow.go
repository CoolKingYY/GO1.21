package schedulerx2

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

// DeleteWorkflow invokes the schedulerx2.DeleteWorkflow API synchronously
func (client *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
	response = CreateDeleteWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWorkflowWithChan invokes the schedulerx2.DeleteWorkflow API asynchronously
func (client *Client) DeleteWorkflowWithChan(request *DeleteWorkflowRequest) (<-chan *DeleteWorkflowResponse, <-chan error) {
	responseChan := make(chan *DeleteWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWorkflow(request)
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

// DeleteWorkflowWithCallback invokes the schedulerx2.DeleteWorkflow API asynchronously
func (client *Client) DeleteWorkflowWithCallback(request *DeleteWorkflowRequest, callback func(response *DeleteWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWorkflowResponse
		var err error
		defer close(result)
		response, err = client.DeleteWorkflow(request)
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

// DeleteWorkflowRequest is the request struct for api DeleteWorkflow
type DeleteWorkflowRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	Namespace       string           `position:"Query" name:"Namespace"`
	WorkflowId      requests.Integer `position:"Query" name:"WorkflowId"`
}

// DeleteWorkflowResponse is the response struct for api DeleteWorkflow
type DeleteWorkflowResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteWorkflowRequest creates a request to invoke DeleteWorkflow API
func CreateDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
	request = &DeleteWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "DeleteWorkflow", "", "")
	request.Method = requests.GET
	return
}

// CreateDeleteWorkflowResponse creates a response to parse from DeleteWorkflow response
func CreateDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
	response = &DeleteWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
