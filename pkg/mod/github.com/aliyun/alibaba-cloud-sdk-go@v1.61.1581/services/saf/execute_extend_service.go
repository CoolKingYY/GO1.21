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

// ExecuteExtendService invokes the saf.ExecuteExtendService API synchronously
func (client *Client) ExecuteExtendService(request *ExecuteExtendServiceRequest) (response *ExecuteExtendServiceResponse, err error) {
	response = CreateExecuteExtendServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteExtendServiceWithChan invokes the saf.ExecuteExtendService API asynchronously
func (client *Client) ExecuteExtendServiceWithChan(request *ExecuteExtendServiceRequest) (<-chan *ExecuteExtendServiceResponse, <-chan error) {
	responseChan := make(chan *ExecuteExtendServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteExtendService(request)
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

// ExecuteExtendServiceWithCallback invokes the saf.ExecuteExtendService API asynchronously
func (client *Client) ExecuteExtendServiceWithCallback(request *ExecuteExtendServiceRequest, callback func(response *ExecuteExtendServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteExtendServiceResponse
		var err error
		defer close(result)
		response, err = client.ExecuteExtendService(request)
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

// ExecuteExtendServiceRequest is the request struct for api ExecuteExtendService
type ExecuteExtendServiceRequest struct {
	*requests.RpcRequest
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
	Region            string `position:"Query" name:"Region"`
}

// ExecuteExtendServiceResponse is the response struct for api ExecuteExtendService
type ExecuteExtendServiceResponse struct {
	*responses.BaseResponse
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateExecuteExtendServiceRequest creates a request to invoke ExecuteExtendService API
func CreateExecuteExtendServiceRequest() (request *ExecuteExtendServiceRequest) {
	request = &ExecuteExtendServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("saf", "2019-05-21", "ExecuteExtendService", "", "")
	request.Method = requests.POST
	return
}

// CreateExecuteExtendServiceResponse creates a response to parse from ExecuteExtendService response
func CreateExecuteExtendServiceResponse() (response *ExecuteExtendServiceResponse) {
	response = &ExecuteExtendServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}