package ehpc

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

// ListCommands invokes the ehpc.ListCommands API synchronously
func (client *Client) ListCommands(request *ListCommandsRequest) (response *ListCommandsResponse, err error) {
	response = CreateListCommandsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCommandsWithChan invokes the ehpc.ListCommands API asynchronously
func (client *Client) ListCommandsWithChan(request *ListCommandsRequest) (<-chan *ListCommandsResponse, <-chan error) {
	responseChan := make(chan *ListCommandsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCommands(request)
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

// ListCommandsWithCallback invokes the ehpc.ListCommands API asynchronously
func (client *Client) ListCommandsWithCallback(request *ListCommandsRequest, callback func(response *ListCommandsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCommandsResponse
		var err error
		defer close(result)
		response, err = client.ListCommands(request)
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

// ListCommandsRequest is the request struct for api ListCommands
type ListCommandsRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	CommandId  string           `position:"Query" name:"CommandId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListCommandsResponse is the response struct for api ListCommands
type ListCommandsResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	Commands   Commands `json:"Commands" xml:"Commands"`
}

// CreateListCommandsRequest creates a request to invoke ListCommands API
func CreateListCommandsRequest() (request *ListCommandsRequest) {
	request = &ListCommandsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListCommands", "", "")
	request.Method = requests.GET
	return
}

// CreateListCommandsResponse creates a response to parse from ListCommands response
func CreateListCommandsResponse() (response *ListCommandsResponse) {
	response = &ListCommandsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
