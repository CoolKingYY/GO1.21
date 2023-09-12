package codeup

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

// CreateMergeRequestComment invokes the codeup.CreateMergeRequestComment API synchronously
func (client *Client) CreateMergeRequestComment(request *CreateMergeRequestCommentRequest) (response *CreateMergeRequestCommentResponse, err error) {
	response = CreateCreateMergeRequestCommentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMergeRequestCommentWithChan invokes the codeup.CreateMergeRequestComment API asynchronously
func (client *Client) CreateMergeRequestCommentWithChan(request *CreateMergeRequestCommentRequest) (<-chan *CreateMergeRequestCommentResponse, <-chan error) {
	responseChan := make(chan *CreateMergeRequestCommentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMergeRequestComment(request)
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

// CreateMergeRequestCommentWithCallback invokes the codeup.CreateMergeRequestComment API asynchronously
func (client *Client) CreateMergeRequestCommentWithCallback(request *CreateMergeRequestCommentRequest, callback func(response *CreateMergeRequestCommentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMergeRequestCommentResponse
		var err error
		defer close(result)
		response, err = client.CreateMergeRequestComment(request)
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

// CreateMergeRequestCommentRequest is the request struct for api CreateMergeRequestComment
type CreateMergeRequestCommentRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	MergeRequestId requests.Integer `position:"Path" name:"MergeRequestId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// CreateMergeRequestCommentResponse is the response struct for api CreateMergeRequestComment
type CreateMergeRequestCommentResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateCreateMergeRequestCommentRequest creates a request to invoke CreateMergeRequestComment API
func CreateCreateMergeRequestCommentRequest() (request *CreateMergeRequestCommentRequest) {
	request = &CreateMergeRequestCommentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "CreateMergeRequestComment", "/api/v4/projects/[ProjectId]/merge_request/[MergeRequestId]/comments", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMergeRequestCommentResponse creates a response to parse from CreateMergeRequestComment response
func CreateCreateMergeRequestCommentResponse() (response *CreateMergeRequestCommentResponse) {
	response = &CreateMergeRequestCommentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
