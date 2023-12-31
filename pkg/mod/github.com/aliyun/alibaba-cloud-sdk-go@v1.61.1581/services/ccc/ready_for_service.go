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

// ReadyForService invokes the ccc.ReadyForService API synchronously
func (client *Client) ReadyForService(request *ReadyForServiceRequest) (response *ReadyForServiceResponse, err error) {
	response = CreateReadyForServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ReadyForServiceWithChan invokes the ccc.ReadyForService API asynchronously
func (client *Client) ReadyForServiceWithChan(request *ReadyForServiceRequest) (<-chan *ReadyForServiceResponse, <-chan error) {
	responseChan := make(chan *ReadyForServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReadyForService(request)
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

// ReadyForServiceWithCallback invokes the ccc.ReadyForService API asynchronously
func (client *Client) ReadyForServiceWithCallback(request *ReadyForServiceRequest, callback func(response *ReadyForServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReadyForServiceResponse
		var err error
		defer close(result)
		response, err = client.ReadyForService(request)
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

// ReadyForServiceRequest is the request struct for api ReadyForService
type ReadyForServiceRequest struct {
	*requests.RpcRequest
	OutboundScenario requests.Boolean `position:"Query" name:"OutboundScenario"`
	UserId           string           `position:"Query" name:"UserId"`
	DeviceId         string           `position:"Query" name:"DeviceId"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
}

// ReadyForServiceResponse is the response struct for api ReadyForService
type ReadyForServiceResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateReadyForServiceRequest creates a request to invoke ReadyForService API
func CreateReadyForServiceRequest() (request *ReadyForServiceRequest) {
	request = &ReadyForServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ReadyForService", "", "")
	request.Method = requests.POST
	return
}

// CreateReadyForServiceResponse creates a response to parse from ReadyForService response
func CreateReadyForServiceResponse() (response *ReadyForServiceResponse) {
	response = &ReadyForServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
