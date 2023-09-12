package scsp

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

// StartHotlineService invokes the scsp.StartHotlineService API synchronously
func (client *Client) StartHotlineService(request *StartHotlineServiceRequest) (response *StartHotlineServiceResponse, err error) {
	response = CreateStartHotlineServiceResponse()
	err = client.DoAction(request, response)
	return
}

// StartHotlineServiceWithChan invokes the scsp.StartHotlineService API asynchronously
func (client *Client) StartHotlineServiceWithChan(request *StartHotlineServiceRequest) (<-chan *StartHotlineServiceResponse, <-chan error) {
	responseChan := make(chan *StartHotlineServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartHotlineService(request)
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

// StartHotlineServiceWithCallback invokes the scsp.StartHotlineService API asynchronously
func (client *Client) StartHotlineServiceWithCallback(request *StartHotlineServiceRequest, callback func(response *StartHotlineServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartHotlineServiceResponse
		var err error
		defer close(result)
		response, err = client.StartHotlineService(request)
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

// StartHotlineServiceRequest is the request struct for api StartHotlineService
type StartHotlineServiceRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Body"`
	InstanceId  string `position:"Body"`
	AccountName string `position:"Body"`
}

// StartHotlineServiceResponse is the response struct for api StartHotlineService
type StartHotlineServiceResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           string `json:"Data" xml:"Data"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateStartHotlineServiceRequest creates a request to invoke StartHotlineService API
func CreateStartHotlineServiceRequest() (request *StartHotlineServiceRequest) {
	request = &StartHotlineServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "StartHotlineService", "", "")
	request.Method = requests.POST
	return
}

// CreateStartHotlineServiceResponse creates a response to parse from StartHotlineService response
func CreateStartHotlineServiceResponse() (response *StartHotlineServiceResponse) {
	response = &StartHotlineServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
