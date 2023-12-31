package vod

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

// SearchEditingProject invokes the vod.SearchEditingProject API synchronously
func (client *Client) SearchEditingProject(request *SearchEditingProjectRequest) (response *SearchEditingProjectResponse, err error) {
	response = CreateSearchEditingProjectResponse()
	err = client.DoAction(request, response)
	return
}

// SearchEditingProjectWithChan invokes the vod.SearchEditingProject API asynchronously
func (client *Client) SearchEditingProjectWithChan(request *SearchEditingProjectRequest) (<-chan *SearchEditingProjectResponse, <-chan error) {
	responseChan := make(chan *SearchEditingProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchEditingProject(request)
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

// SearchEditingProjectWithCallback invokes the vod.SearchEditingProject API asynchronously
func (client *Client) SearchEditingProjectWithCallback(request *SearchEditingProjectRequest, callback func(response *SearchEditingProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchEditingProjectResponse
		var err error
		defer close(result)
		response, err = client.SearchEditingProject(request)
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

// SearchEditingProjectRequest is the request struct for api SearchEditingProject
type SearchEditingProjectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Title                string           `position:"Query" name:"Title"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	SortBy               string           `position:"Query" name:"SortBy"`
	Status               string           `position:"Query" name:"Status"`
}

// SearchEditingProjectResponse is the response struct for api SearchEditingProject
type SearchEditingProjectResponse struct {
	*responses.BaseResponse
	Total       int         `json:"Total" xml:"Total"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ProjectList ProjectList `json:"ProjectList" xml:"ProjectList"`
}

// CreateSearchEditingProjectRequest creates a request to invoke SearchEditingProject API
func CreateSearchEditingProjectRequest() (request *SearchEditingProjectRequest) {
	request = &SearchEditingProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SearchEditingProject", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchEditingProjectResponse creates a response to parse from SearchEditingProject response
func CreateSearchEditingProjectResponse() (response *SearchEditingProjectResponse) {
	response = &SearchEditingProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
