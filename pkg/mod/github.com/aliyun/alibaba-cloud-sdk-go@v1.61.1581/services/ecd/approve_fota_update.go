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

// ApproveFotaUpdate invokes the ecd.ApproveFotaUpdate API synchronously
func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (response *ApproveFotaUpdateResponse, err error) {
	response = CreateApproveFotaUpdateResponse()
	err = client.DoAction(request, response)
	return
}

// ApproveFotaUpdateWithChan invokes the ecd.ApproveFotaUpdate API asynchronously
func (client *Client) ApproveFotaUpdateWithChan(request *ApproveFotaUpdateRequest) (<-chan *ApproveFotaUpdateResponse, <-chan error) {
	responseChan := make(chan *ApproveFotaUpdateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApproveFotaUpdate(request)
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

// ApproveFotaUpdateWithCallback invokes the ecd.ApproveFotaUpdate API asynchronously
func (client *Client) ApproveFotaUpdateWithCallback(request *ApproveFotaUpdateRequest, callback func(response *ApproveFotaUpdateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApproveFotaUpdateResponse
		var err error
		defer close(result)
		response, err = client.ApproveFotaUpdate(request)
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

// ApproveFotaUpdateRequest is the request struct for api ApproveFotaUpdate
type ApproveFotaUpdateRequest struct {
	*requests.RpcRequest
	AppVersion string `position:"Query" name:"AppVersion"`
	DesktopId  string `position:"Query" name:"DesktopId"`
}

// ApproveFotaUpdateResponse is the response struct for api ApproveFotaUpdate
type ApproveFotaUpdateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateApproveFotaUpdateRequest creates a request to invoke ApproveFotaUpdate API
func CreateApproveFotaUpdateRequest() (request *ApproveFotaUpdateRequest) {
	request = &ApproveFotaUpdateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ApproveFotaUpdate", "", "")
	request.Method = requests.POST
	return
}

// CreateApproveFotaUpdateResponse creates a response to parse from ApproveFotaUpdate response
func CreateApproveFotaUpdateResponse() (response *ApproveFotaUpdateResponse) {
	response = &ApproveFotaUpdateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
