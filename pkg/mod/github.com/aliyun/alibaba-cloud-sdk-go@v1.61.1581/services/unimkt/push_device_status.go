package unimkt

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

// PushDeviceStatus invokes the unimkt.PushDeviceStatus API synchronously
func (client *Client) PushDeviceStatus(request *PushDeviceStatusRequest) (response *PushDeviceStatusResponse, err error) {
	response = CreatePushDeviceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// PushDeviceStatusWithChan invokes the unimkt.PushDeviceStatus API asynchronously
func (client *Client) PushDeviceStatusWithChan(request *PushDeviceStatusRequest) (<-chan *PushDeviceStatusResponse, <-chan error) {
	responseChan := make(chan *PushDeviceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushDeviceStatus(request)
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

// PushDeviceStatusWithCallback invokes the unimkt.PushDeviceStatus API asynchronously
func (client *Client) PushDeviceStatusWithCallback(request *PushDeviceStatusRequest, callback func(response *PushDeviceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushDeviceStatusResponse
		var err error
		defer close(result)
		response, err = client.PushDeviceStatus(request)
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

// PushDeviceStatusRequest is the request struct for api PushDeviceStatus
type PushDeviceStatusRequest struct {
	*requests.RpcRequest
	DeviceSn  string           `position:"Body" name:"DeviceSn"`
	ChannelId string           `position:"Body" name:"ChannelId"`
	Status    requests.Integer `position:"Body" name:"Status"`
}

// PushDeviceStatusResponse is the response struct for api PushDeviceStatus
type PushDeviceStatusResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePushDeviceStatusRequest creates a request to invoke PushDeviceStatus API
func CreatePushDeviceStatusRequest() (request *PushDeviceStatusRequest) {
	request = &PushDeviceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "PushDeviceStatus", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushDeviceStatusResponse creates a response to parse from PushDeviceStatus response
func CreatePushDeviceStatusResponse() (response *PushDeviceStatusResponse) {
	response = &PushDeviceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
