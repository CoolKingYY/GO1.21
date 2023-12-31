package iot

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

// GenerateFileUploadURL invokes the iot.GenerateFileUploadURL API synchronously
func (client *Client) GenerateFileUploadURL(request *GenerateFileUploadURLRequest) (response *GenerateFileUploadURLResponse, err error) {
	response = CreateGenerateFileUploadURLResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateFileUploadURLWithChan invokes the iot.GenerateFileUploadURL API asynchronously
func (client *Client) GenerateFileUploadURLWithChan(request *GenerateFileUploadURLRequest) (<-chan *GenerateFileUploadURLResponse, <-chan error) {
	responseChan := make(chan *GenerateFileUploadURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateFileUploadURL(request)
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

// GenerateFileUploadURLWithCallback invokes the iot.GenerateFileUploadURL API asynchronously
func (client *Client) GenerateFileUploadURLWithCallback(request *GenerateFileUploadURLRequest, callback func(response *GenerateFileUploadURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateFileUploadURLResponse
		var err error
		defer close(result)
		response, err = client.GenerateFileUploadURL(request)
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

// GenerateFileUploadURLRequest is the request struct for api GenerateFileUploadURL
type GenerateFileUploadURLRequest struct {
	*requests.RpcRequest
	FileSuffix    string `position:"Query" name:"FileSuffix"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	FileName      string `position:"Query" name:"FileName"`
	BizCode       string `position:"Query" name:"BizCode"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GenerateFileUploadURLResponse is the response struct for api GenerateFileUploadURL
type GenerateFileUploadURLResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGenerateFileUploadURLRequest creates a request to invoke GenerateFileUploadURL API
func CreateGenerateFileUploadURLRequest() (request *GenerateFileUploadURLRequest) {
	request = &GenerateFileUploadURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GenerateFileUploadURL", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateFileUploadURLResponse creates a response to parse from GenerateFileUploadURL response
func CreateGenerateFileUploadURLResponse() (response *GenerateFileUploadURLResponse) {
	response = &GenerateFileUploadURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
