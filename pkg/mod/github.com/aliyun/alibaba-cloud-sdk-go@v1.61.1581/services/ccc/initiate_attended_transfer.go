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

// InitiateAttendedTransfer invokes the ccc.InitiateAttendedTransfer API synchronously
func (client *Client) InitiateAttendedTransfer(request *InitiateAttendedTransferRequest) (response *InitiateAttendedTransferResponse, err error) {
	response = CreateInitiateAttendedTransferResponse()
	err = client.DoAction(request, response)
	return
}

// InitiateAttendedTransferWithChan invokes the ccc.InitiateAttendedTransfer API asynchronously
func (client *Client) InitiateAttendedTransferWithChan(request *InitiateAttendedTransferRequest) (<-chan *InitiateAttendedTransferResponse, <-chan error) {
	responseChan := make(chan *InitiateAttendedTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitiateAttendedTransfer(request)
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

// InitiateAttendedTransferWithCallback invokes the ccc.InitiateAttendedTransfer API asynchronously
func (client *Client) InitiateAttendedTransferWithCallback(request *InitiateAttendedTransferRequest, callback func(response *InitiateAttendedTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitiateAttendedTransferResponse
		var err error
		defer close(result)
		response, err = client.InitiateAttendedTransfer(request)
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

// InitiateAttendedTransferRequest is the request struct for api InitiateAttendedTransfer
type InitiateAttendedTransferRequest struct {
	*requests.RpcRequest
	Transferee     string           `position:"Query" name:"Transferee"`
	Transferor     string           `position:"Query" name:"Transferor"`
	UserId         string           `position:"Query" name:"UserId"`
	DeviceId       string           `position:"Query" name:"DeviceId"`
	TimeoutSeconds requests.Integer `position:"Query" name:"TimeoutSeconds"`
	JobId          string           `position:"Query" name:"JobId"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
}

// InitiateAttendedTransferResponse is the response struct for api InitiateAttendedTransfer
type InitiateAttendedTransferResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateInitiateAttendedTransferRequest creates a request to invoke InitiateAttendedTransfer API
func CreateInitiateAttendedTransferRequest() (request *InitiateAttendedTransferRequest) {
	request = &InitiateAttendedTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "InitiateAttendedTransfer", "", "")
	request.Method = requests.POST
	return
}

// CreateInitiateAttendedTransferResponse creates a response to parse from InitiateAttendedTransfer response
func CreateInitiateAttendedTransferResponse() (response *InitiateAttendedTransferResponse) {
	response = &InitiateAttendedTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
