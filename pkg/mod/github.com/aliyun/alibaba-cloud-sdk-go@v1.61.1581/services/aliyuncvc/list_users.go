package aliyuncvc

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

// ListUsers invokes the aliyuncvc.ListUsers API synchronously
func (client *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
	response = CreateListUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListUsersWithChan invokes the aliyuncvc.ListUsers API asynchronously
func (client *Client) ListUsersWithChan(request *ListUsersRequest) (<-chan *ListUsersResponse, <-chan error) {
	responseChan := make(chan *ListUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsers(request)
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

// ListUsersWithCallback invokes the aliyuncvc.ListUsers API asynchronously
func (client *Client) ListUsersWithCallback(request *ListUsersRequest, callback func(response *ListUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersResponse
		var err error
		defer close(result)
		response, err = client.ListUsers(request)
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

// ListUsersRequest is the request struct for api ListUsers
type ListUsersRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListUsersResponse is the response struct for api ListUsers
type ListUsersResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListUsersRequest creates a request to invoke ListUsers API
func CreateListUsersRequest() (request *ListUsersRequest) {
	request = &ListUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ListUsers", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUsersResponse creates a response to parse from ListUsers response
func CreateListUsersResponse() (response *ListUsersResponse) {
	response = &ListUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
