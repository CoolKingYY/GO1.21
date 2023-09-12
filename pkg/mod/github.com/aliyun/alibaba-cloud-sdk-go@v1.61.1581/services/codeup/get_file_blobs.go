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

// GetFileBlobs invokes the codeup.GetFileBlobs API synchronously
func (client *Client) GetFileBlobs(request *GetFileBlobsRequest) (response *GetFileBlobsResponse, err error) {
	response = CreateGetFileBlobsResponse()
	err = client.DoAction(request, response)
	return
}

// GetFileBlobsWithChan invokes the codeup.GetFileBlobs API asynchronously
func (client *Client) GetFileBlobsWithChan(request *GetFileBlobsRequest) (<-chan *GetFileBlobsResponse, <-chan error) {
	responseChan := make(chan *GetFileBlobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFileBlobs(request)
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

// GetFileBlobsWithCallback invokes the codeup.GetFileBlobs API asynchronously
func (client *Client) GetFileBlobsWithCallback(request *GetFileBlobsRequest, callback func(response *GetFileBlobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFileBlobsResponse
		var err error
		defer close(result)
		response, err = client.GetFileBlobs(request)
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

// GetFileBlobsRequest is the request struct for api GetFileBlobs
type GetFileBlobsRequest struct {
	*requests.RoaRequest
	AccessToken    string           `position:"Query" name:"AccessToken"`
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	Ref            string           `position:"Query" name:"Ref"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	FilePath       string           `position:"Query" name:"FilePath"`
	From           requests.Integer `position:"Query" name:"From"`
	To             requests.Integer `position:"Query" name:"To"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// GetFileBlobsResponse is the response struct for api GetFileBlobs
type GetFileBlobsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateGetFileBlobsRequest creates a request to invoke GetFileBlobs API
func CreateGetFileBlobsRequest() (request *GetFileBlobsRequest) {
	request = &GetFileBlobsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "GetFileBlobs", "/api/v4/projects/[ProjectId]/repository/blobs", "", "")
	request.Method = requests.GET
	return
}

// CreateGetFileBlobsResponse creates a response to parse from GetFileBlobs response
func CreateGetFileBlobsResponse() (response *GetFileBlobsResponse) {
	response = &GetFileBlobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
