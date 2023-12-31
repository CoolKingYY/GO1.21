package ecd

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

// CreateUserTags invokes the ecd.CreateUserTags API synchronously
func (client *Client) CreateUserTags(request *CreateUserTagsRequest) (response *CreateUserTagsResponse, err error) {
	response = CreateCreateUserTagsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserTagsWithChan invokes the ecd.CreateUserTags API asynchronously
func (client *Client) CreateUserTagsWithChan(request *CreateUserTagsRequest) (<-chan *CreateUserTagsResponse, <-chan error) {
	responseChan := make(chan *CreateUserTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserTags(request)
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

// CreateUserTagsWithCallback invokes the ecd.CreateUserTags API asynchronously
func (client *Client) CreateUserTagsWithCallback(request *CreateUserTagsRequest, callback func(response *CreateUserTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserTagsResponse
		var err error
		defer close(result)
		response, err = client.CreateUserTags(request)
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

// CreateUserTagsRequest is the request struct for api CreateUserTags
type CreateUserTagsRequest struct {
	*requests.RpcRequest
	Tags *[]string `position:"Query" name:"Tags"  type:"Repeated"`
}

// CreateUserTagsResponse is the response struct for api CreateUserTags
type CreateUserTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateUserTagsRequest creates a request to invoke CreateUserTags API
func CreateCreateUserTagsRequest() (request *CreateUserTagsRequest) {
	request = &CreateUserTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateUserTags", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateUserTagsResponse creates a response to parse from CreateUserTags response
func CreateCreateUserTagsResponse() (response *CreateUserTagsResponse) {
	response = &CreateUserTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
