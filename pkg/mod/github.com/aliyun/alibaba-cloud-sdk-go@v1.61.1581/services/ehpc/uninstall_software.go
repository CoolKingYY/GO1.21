package ehpc

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

// UninstallSoftware invokes the ehpc.UninstallSoftware API synchronously
func (client *Client) UninstallSoftware(request *UninstallSoftwareRequest) (response *UninstallSoftwareResponse, err error) {
	response = CreateUninstallSoftwareResponse()
	err = client.DoAction(request, response)
	return
}

// UninstallSoftwareWithChan invokes the ehpc.UninstallSoftware API asynchronously
func (client *Client) UninstallSoftwareWithChan(request *UninstallSoftwareRequest) (<-chan *UninstallSoftwareResponse, <-chan error) {
	responseChan := make(chan *UninstallSoftwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UninstallSoftware(request)
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

// UninstallSoftwareWithCallback invokes the ehpc.UninstallSoftware API asynchronously
func (client *Client) UninstallSoftwareWithCallback(request *UninstallSoftwareRequest, callback func(response *UninstallSoftwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UninstallSoftwareResponse
		var err error
		defer close(result)
		response, err = client.UninstallSoftware(request)
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

// UninstallSoftwareRequest is the request struct for api UninstallSoftware
type UninstallSoftwareRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	Application string `position:"Query" name:"Application"`
}

// UninstallSoftwareResponse is the response struct for api UninstallSoftware
type UninstallSoftwareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUninstallSoftwareRequest creates a request to invoke UninstallSoftware API
func CreateUninstallSoftwareRequest() (request *UninstallSoftwareRequest) {
	request = &UninstallSoftwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "UninstallSoftware", "", "")
	request.Method = requests.GET
	return
}

// CreateUninstallSoftwareResponse creates a response to parse from UninstallSoftware response
func CreateUninstallSoftwareResponse() (response *UninstallSoftwareResponse) {
	response = &UninstallSoftwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}