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

// ListCallTags invokes the ccc.ListCallTags API synchronously
func (client *Client) ListCallTags(request *ListCallTagsRequest) (response *ListCallTagsResponse, err error) {
	response = CreateListCallTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCallTagsWithChan invokes the ccc.ListCallTags API asynchronously
func (client *Client) ListCallTagsWithChan(request *ListCallTagsRequest) (<-chan *ListCallTagsResponse, <-chan error) {
	responseChan := make(chan *ListCallTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCallTags(request)
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

// ListCallTagsWithCallback invokes the ccc.ListCallTags API asynchronously
func (client *Client) ListCallTagsWithCallback(request *ListCallTagsRequest, callback func(response *ListCallTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCallTagsResponse
		var err error
		defer close(result)
		response, err = client.ListCallTags(request)
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

// ListCallTagsRequest is the request struct for api ListCallTags
type ListCallTagsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListCallTagsResponse is the response struct for api ListCallTags
type ListCallTagsResponse struct {
	*responses.BaseResponse
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string             `json:"Code" xml:"Code"`
	Message        string             `json:"Message" xml:"Message"`
	Data           DataInListCallTags `json:"Data" xml:"Data"`
}

// CreateListCallTagsRequest creates a request to invoke ListCallTags API
func CreateListCallTagsRequest() (request *ListCallTagsRequest) {
	request = &ListCallTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListCallTags", "", "")
	request.Method = requests.POST
	return
}

// CreateListCallTagsResponse creates a response to parse from ListCallTags response
func CreateListCallTagsResponse() (response *ListCallTagsResponse) {
	response = &ListCallTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
