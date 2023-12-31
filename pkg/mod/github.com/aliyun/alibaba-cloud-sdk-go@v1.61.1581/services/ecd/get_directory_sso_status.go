package ecd

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

// GetDirectorySsoStatus invokes the ecd.GetDirectorySsoStatus API synchronously
func (client *Client) GetDirectorySsoStatus(request *GetDirectorySsoStatusRequest) (response *GetDirectorySsoStatusResponse, err error) {
	response = CreateGetDirectorySsoStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetDirectorySsoStatusWithChan invokes the ecd.GetDirectorySsoStatus API asynchronously
func (client *Client) GetDirectorySsoStatusWithChan(request *GetDirectorySsoStatusRequest) (<-chan *GetDirectorySsoStatusResponse, <-chan error) {
	responseChan := make(chan *GetDirectorySsoStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDirectorySsoStatus(request)
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

// GetDirectorySsoStatusWithCallback invokes the ecd.GetDirectorySsoStatus API asynchronously
func (client *Client) GetDirectorySsoStatusWithCallback(request *GetDirectorySsoStatusRequest, callback func(response *GetDirectorySsoStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDirectorySsoStatusResponse
		var err error
		defer close(result)
		response, err = client.GetDirectorySsoStatus(request)
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

// GetDirectorySsoStatusRequest is the request struct for api GetDirectorySsoStatus
type GetDirectorySsoStatusRequest struct {
	*requests.RpcRequest
	DirectoryId string `position:"Query" name:"DirectoryId"`
}

// GetDirectorySsoStatusResponse is the response struct for api GetDirectorySsoStatus
type GetDirectorySsoStatusResponse struct {
	*responses.BaseResponse
	SsoStatus bool   `json:"SsoStatus" xml:"SsoStatus"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetDirectorySsoStatusRequest creates a request to invoke GetDirectorySsoStatus API
func CreateGetDirectorySsoStatusRequest() (request *GetDirectorySsoStatusRequest) {
	request = &GetDirectorySsoStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "GetDirectorySsoStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDirectorySsoStatusResponse creates a response to parse from GetDirectorySsoStatus response
func CreateGetDirectorySsoStatusResponse() (response *GetDirectorySsoStatusResponse) {
	response = &GetDirectorySsoStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
