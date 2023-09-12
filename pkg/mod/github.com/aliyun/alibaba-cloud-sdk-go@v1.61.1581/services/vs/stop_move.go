package vs

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

// StopMove invokes the vs.StopMove API synchronously
func (client *Client) StopMove(request *StopMoveRequest) (response *StopMoveResponse, err error) {
	response = CreateStopMoveResponse()
	err = client.DoAction(request, response)
	return
}

// StopMoveWithChan invokes the vs.StopMove API asynchronously
func (client *Client) StopMoveWithChan(request *StopMoveRequest) (<-chan *StopMoveResponse, <-chan error) {
	responseChan := make(chan *StopMoveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopMove(request)
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

// StopMoveWithCallback invokes the vs.StopMove API asynchronously
func (client *Client) StopMoveWithCallback(request *StopMoveRequest, callback func(response *StopMoveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopMoveResponse
		var err error
		defer close(result)
		response, err = client.StopMove(request)
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

// StopMoveRequest is the request struct for api StopMove
type StopMoveRequest struct {
	*requests.RpcRequest
	Tilt        requests.Boolean `position:"Query" name:"Tilt"`
	SubProtocol string           `position:"Query" name:"SubProtocol"`
	Id          string           `position:"Query" name:"Id"`
	Pan         requests.Boolean `position:"Query" name:"Pan"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	Zoom        requests.Boolean `position:"Query" name:"Zoom"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// StopMoveResponse is the response struct for api StopMove
type StopMoveResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopMoveRequest creates a request to invoke StopMove API
func CreateStopMoveRequest() (request *StopMoveRequest) {
	request = &StopMoveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "StopMove", "", "")
	request.Method = requests.POST
	return
}

// CreateStopMoveResponse creates a response to parse from StopMove response
func CreateStopMoveResponse() (response *StopMoveResponse) {
	response = &StopMoveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}