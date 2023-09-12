package outboundbot

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

// SuspendCallWithFile invokes the outboundbot.SuspendCallWithFile API synchronously
func (client *Client) SuspendCallWithFile(request *SuspendCallWithFileRequest) (response *SuspendCallWithFileResponse, err error) {
	response = CreateSuspendCallWithFileResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendCallWithFileWithChan invokes the outboundbot.SuspendCallWithFile API asynchronously
func (client *Client) SuspendCallWithFileWithChan(request *SuspendCallWithFileRequest) (<-chan *SuspendCallWithFileResponse, <-chan error) {
	responseChan := make(chan *SuspendCallWithFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendCallWithFile(request)
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

// SuspendCallWithFileWithCallback invokes the outboundbot.SuspendCallWithFile API asynchronously
func (client *Client) SuspendCallWithFileWithCallback(request *SuspendCallWithFileRequest, callback func(response *SuspendCallWithFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendCallWithFileResponse
		var err error
		defer close(result)
		response, err = client.SuspendCallWithFile(request)
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

// SuspendCallWithFileRequest is the request struct for api SuspendCallWithFile
type SuspendCallWithFileRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	FilePath   string `position:"Query" name:"FilePath"`
}

// SuspendCallWithFileResponse is the response struct for api SuspendCallWithFile
type SuspendCallWithFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSuspendCallWithFileRequest creates a request to invoke SuspendCallWithFile API
func CreateSuspendCallWithFileRequest() (request *SuspendCallWithFileRequest) {
	request = &SuspendCallWithFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SuspendCallWithFile", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendCallWithFileResponse creates a response to parse from SuspendCallWithFile response
func CreateSuspendCallWithFileResponse() (response *SuspendCallWithFileResponse) {
	response = &SuspendCallWithFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
