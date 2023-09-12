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

// ImportCustomCallTagging invokes the ccc.ImportCustomCallTagging API synchronously
func (client *Client) ImportCustomCallTagging(request *ImportCustomCallTaggingRequest) (response *ImportCustomCallTaggingResponse, err error) {
	response = CreateImportCustomCallTaggingResponse()
	err = client.DoAction(request, response)
	return
}

// ImportCustomCallTaggingWithChan invokes the ccc.ImportCustomCallTagging API asynchronously
func (client *Client) ImportCustomCallTaggingWithChan(request *ImportCustomCallTaggingRequest) (<-chan *ImportCustomCallTaggingResponse, <-chan error) {
	responseChan := make(chan *ImportCustomCallTaggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportCustomCallTagging(request)
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

// ImportCustomCallTaggingWithCallback invokes the ccc.ImportCustomCallTagging API asynchronously
func (client *Client) ImportCustomCallTaggingWithCallback(request *ImportCustomCallTaggingRequest, callback func(response *ImportCustomCallTaggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportCustomCallTaggingResponse
		var err error
		defer close(result)
		response, err = client.ImportCustomCallTagging(request)
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

// ImportCustomCallTaggingRequest is the request struct for api ImportCustomCallTagging
type ImportCustomCallTaggingRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	FilePath   string `position:"Query" name:"FilePath"`
}

// ImportCustomCallTaggingResponse is the response struct for api ImportCustomCallTagging
type ImportCustomCallTaggingResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	Data           []FailureItem `json:"Data" xml:"Data"`
}

// CreateImportCustomCallTaggingRequest creates a request to invoke ImportCustomCallTagging API
func CreateImportCustomCallTaggingRequest() (request *ImportCustomCallTaggingRequest) {
	request = &ImportCustomCallTaggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ImportCustomCallTagging", "", "")
	request.Method = requests.POST
	return
}

// CreateImportCustomCallTaggingResponse creates a response to parse from ImportCustomCallTagging response
func CreateImportCustomCallTaggingResponse() (response *ImportCustomCallTaggingResponse) {
	response = &ImportCustomCallTaggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
