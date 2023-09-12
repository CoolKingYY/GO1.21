package saf

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

// ExecuteRequestML invokes the saf.ExecuteRequestML API synchronously
func (client *Client) ExecuteRequestML(request *ExecuteRequestMLRequest) (response *ExecuteRequestMLResponse, err error) {
	response = CreateExecuteRequestMLResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteRequestMLWithChan invokes the saf.ExecuteRequestML API asynchronously
func (client *Client) ExecuteRequestMLWithChan(request *ExecuteRequestMLRequest) (<-chan *ExecuteRequestMLResponse, <-chan error) {
	responseChan := make(chan *ExecuteRequestMLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteRequestML(request)
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

// ExecuteRequestMLWithCallback invokes the saf.ExecuteRequestML API asynchronously
func (client *Client) ExecuteRequestMLWithCallback(request *ExecuteRequestMLRequest, callback func(response *ExecuteRequestMLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteRequestMLResponse
		var err error
		defer close(result)
		response, err = client.ExecuteRequestML(request)
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

// ExecuteRequestMLRequest is the request struct for api ExecuteRequestML
type ExecuteRequestMLRequest struct {
	*requests.RpcRequest
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
	Lang              string `position:"Query" name:"Lang"`
}

// ExecuteRequestMLResponse is the response struct for api ExecuteRequestML
type ExecuteRequestMLResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateExecuteRequestMLRequest creates a request to invoke ExecuteRequestML API
func CreateExecuteRequestMLRequest() (request *ExecuteRequestMLRequest) {
	request = &ExecuteRequestMLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("saf", "2019-05-21", "ExecuteRequestML", "", "")
	request.Method = requests.POST
	return
}

// CreateExecuteRequestMLResponse creates a response to parse from ExecuteRequestML response
func CreateExecuteRequestMLResponse() (response *ExecuteRequestMLResponse) {
	response = &ExecuteRequestMLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
