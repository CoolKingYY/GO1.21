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

// DownloadFiles invokes the sae.DownloadFiles API synchronously
func (client *Client) DownloadFiles(request *DownloadFilesRequest) (response *DownloadFilesResponse, err error) {
	response = CreateDownloadFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadFilesWithChan invokes the sae.DownloadFiles API asynchronously
func (client *Client) DownloadFilesWithChan(request *DownloadFilesRequest) (<-chan *DownloadFilesResponse, <-chan error) {
	responseChan := make(chan *DownloadFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadFiles(request)
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

// DownloadFilesWithCallback invokes the sae.DownloadFiles API asynchronously
func (client *Client) DownloadFilesWithCallback(request *DownloadFilesRequest, callback func(response *DownloadFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadFilesResponse
		var err error
		defer close(result)
		response, err = client.DownloadFiles(request)
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

// DownloadFilesRequest is the request struct for api DownloadFiles
type DownloadFilesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Localpath  string `position:"Query" name:"Localpath"`
	AppId      string `position:"Query" name:"AppId"`
}

// DownloadFilesResponse is the response struct for api DownloadFiles
type DownloadFilesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDownloadFilesRequest creates a request to invoke DownloadFiles API
func CreateDownloadFilesRequest() (request *DownloadFilesRequest) {
	request = &DownloadFilesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DownloadFiles", "/pop/v1/sam/app/downloadFiles", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDownloadFilesResponse creates a response to parse from DownloadFiles response
func CreateDownloadFilesResponse() (response *DownloadFilesResponse) {
	response = &DownloadFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
