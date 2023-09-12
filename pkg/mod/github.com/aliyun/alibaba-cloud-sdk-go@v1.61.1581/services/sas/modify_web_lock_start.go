package sas

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

// ModifyWebLockStart invokes the sas.ModifyWebLockStart API synchronously
func (client *Client) ModifyWebLockStart(request *ModifyWebLockStartRequest) (response *ModifyWebLockStartResponse, err error) {
	response = CreateModifyWebLockStartResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebLockStartWithChan invokes the sas.ModifyWebLockStart API asynchronously
func (client *Client) ModifyWebLockStartWithChan(request *ModifyWebLockStartRequest) (<-chan *ModifyWebLockStartResponse, <-chan error) {
	responseChan := make(chan *ModifyWebLockStartResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebLockStart(request)
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

// ModifyWebLockStartWithCallback invokes the sas.ModifyWebLockStart API asynchronously
func (client *Client) ModifyWebLockStartWithCallback(request *ModifyWebLockStartRequest, callback func(response *ModifyWebLockStartResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebLockStartResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebLockStart(request)
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

// ModifyWebLockStartRequest is the request struct for api ModifyWebLockStart
type ModifyWebLockStartRequest struct {
	*requests.RpcRequest
	LocalBackupDir    string `position:"Query" name:"LocalBackupDir"`
	ExclusiveFile     string `position:"Query" name:"ExclusiveFile"`
	ExclusiveFileType string `position:"Query" name:"ExclusiveFileType"`
	Dir               string `position:"Query" name:"Dir"`
	Uuid              string `position:"Query" name:"Uuid"`
	Mode              string `position:"Query" name:"Mode"`
	SourceIp          string `position:"Query" name:"SourceIp"`
	ExclusiveDir      string `position:"Query" name:"ExclusiveDir"`
	InclusiveFileType string `position:"Query" name:"InclusiveFileType"`
}

// ModifyWebLockStartResponse is the response struct for api ModifyWebLockStart
type ModifyWebLockStartResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebLockStartRequest creates a request to invoke ModifyWebLockStart API
func CreateModifyWebLockStartRequest() (request *ModifyWebLockStartRequest) {
	request = &ModifyWebLockStartRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyWebLockStart", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyWebLockStartResponse creates a response to parse from ModifyWebLockStart response
func CreateModifyWebLockStartResponse() (response *ModifyWebLockStartResponse) {
	response = &ModifyWebLockStartResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
