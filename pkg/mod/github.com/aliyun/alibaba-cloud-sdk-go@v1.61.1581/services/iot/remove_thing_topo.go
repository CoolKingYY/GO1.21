package iot

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

// RemoveThingTopo invokes the iot.RemoveThingTopo API synchronously
func (client *Client) RemoveThingTopo(request *RemoveThingTopoRequest) (response *RemoveThingTopoResponse, err error) {
	response = CreateRemoveThingTopoResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveThingTopoWithChan invokes the iot.RemoveThingTopo API asynchronously
func (client *Client) RemoveThingTopoWithChan(request *RemoveThingTopoRequest) (<-chan *RemoveThingTopoResponse, <-chan error) {
	responseChan := make(chan *RemoveThingTopoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveThingTopo(request)
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

// RemoveThingTopoWithCallback invokes the iot.RemoveThingTopo API asynchronously
func (client *Client) RemoveThingTopoWithCallback(request *RemoveThingTopoRequest, callback func(response *RemoveThingTopoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveThingTopoResponse
		var err error
		defer close(result)
		response, err = client.RemoveThingTopo(request)
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

// RemoveThingTopoRequest is the request struct for api RemoveThingTopo
type RemoveThingTopoRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// RemoveThingTopoResponse is the response struct for api RemoveThingTopo
type RemoveThingTopoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         bool   `json:"Data" xml:"Data"`
}

// CreateRemoveThingTopoRequest creates a request to invoke RemoveThingTopo API
func CreateRemoveThingTopoRequest() (request *RemoveThingTopoRequest) {
	request = &RemoveThingTopoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "RemoveThingTopo", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveThingTopoResponse creates a response to parse from RemoveThingTopo response
func CreateRemoveThingTopoResponse() (response *RemoveThingTopoResponse) {
	response = &RemoveThingTopoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}