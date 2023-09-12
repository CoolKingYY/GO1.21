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

// UpdateFile invokes the codeup.UpdateFile API synchronously
func (client *Client) UpdateFile(request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
	response = CreateUpdateFileResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFileWithChan invokes the codeup.UpdateFile API asynchronously
func (client *Client) UpdateFileWithChan(request *UpdateFileRequest) (<-chan *UpdateFileResponse, <-chan error) {
	responseChan := make(chan *UpdateFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFile(request)
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

// UpdateFileWithCallback invokes the codeup.UpdateFile API asynchronously
func (client *Client) UpdateFileWithCallback(request *UpdateFileRequest, callback func(response *UpdateFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFileResponse
		var err error
		defer close(result)
		response, err = client.UpdateFile(request)
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

// UpdateFileRequest is the request struct for api UpdateFile
type UpdateFileRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// UpdateFileResponse is the response struct for api UpdateFile
type UpdateFileResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateUpdateFileRequest creates a request to invoke UpdateFile API
func CreateUpdateFileRequest() (request *UpdateFileRequest) {
	request = &UpdateFileRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "UpdateFile", "/api/v4/projects/[ProjectId]/repository/files", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateFileResponse creates a response to parse from UpdateFile response
func CreateUpdateFileResponse() (response *UpdateFileResponse) {
	response = &UpdateFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
