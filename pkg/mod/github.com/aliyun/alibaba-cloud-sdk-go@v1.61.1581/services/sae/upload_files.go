package sae

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

// UploadFiles invokes the sae.UploadFiles API synchronously
func (client *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
	response = CreateUploadFilesResponse()
	err = client.DoAction(request, response)
	return
}

// UploadFilesWithChan invokes the sae.UploadFiles API asynchronously
func (client *Client) UploadFilesWithChan(request *UploadFilesRequest) (<-chan *UploadFilesResponse, <-chan error) {
	responseChan := make(chan *UploadFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadFiles(request)
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

// UploadFilesWithCallback invokes the sae.UploadFiles API asynchronously
func (client *Client) UploadFilesWithCallback(request *UploadFilesRequest, callback func(response *UploadFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadFilesResponse
		var err error
		defer close(result)
		response, err = client.UploadFiles(request)
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

// UploadFilesRequest is the request struct for api UploadFiles
type UploadFilesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	CloudUrl   string `position:"Query" name:"CloudUrl"`
	Localpath  string `position:"Query" name:"Localpath"`
	AppId      string `position:"Query" name:"AppId"`
}

// UploadFilesResponse is the response struct for api UploadFiles
type UploadFilesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUploadFilesRequest creates a request to invoke UploadFiles API
func CreateUploadFilesRequest() (request *UploadFilesRequest) {
	request = &UploadFilesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "UploadFiles", "/pop/v1/sam/app/uploadFiles", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUploadFilesResponse creates a response to parse from UploadFiles response
func CreateUploadFilesResponse() (response *UploadFilesResponse) {
	response = &UploadFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}