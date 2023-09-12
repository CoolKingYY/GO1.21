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

// BatchUnbindDeviceFromEdgeInstance invokes the iot.BatchUnbindDeviceFromEdgeInstance API synchronously
func (client *Client) BatchUnbindDeviceFromEdgeInstance(request *BatchUnbindDeviceFromEdgeInstanceRequest) (response *BatchUnbindDeviceFromEdgeInstanceResponse, err error) {
	response = CreateBatchUnbindDeviceFromEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchUnbindDeviceFromEdgeInstanceWithChan invokes the iot.BatchUnbindDeviceFromEdgeInstance API asynchronously
func (client *Client) BatchUnbindDeviceFromEdgeInstanceWithChan(request *BatchUnbindDeviceFromEdgeInstanceRequest) (<-chan *BatchUnbindDeviceFromEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *BatchUnbindDeviceFromEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchUnbindDeviceFromEdgeInstance(request)
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

// BatchUnbindDeviceFromEdgeInstanceWithCallback invokes the iot.BatchUnbindDeviceFromEdgeInstance API asynchronously
func (client *Client) BatchUnbindDeviceFromEdgeInstanceWithCallback(request *BatchUnbindDeviceFromEdgeInstanceRequest, callback func(response *BatchUnbindDeviceFromEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchUnbindDeviceFromEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.BatchUnbindDeviceFromEdgeInstance(request)
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

// BatchUnbindDeviceFromEdgeInstanceRequest is the request struct for api BatchUnbindDeviceFromEdgeInstance
type BatchUnbindDeviceFromEdgeInstanceRequest struct {
	*requests.RpcRequest
	IotIds        *[]string `position:"Query" name:"IotIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// BatchUnbindDeviceFromEdgeInstanceResponse is the response struct for api BatchUnbindDeviceFromEdgeInstance
type BatchUnbindDeviceFromEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBatchUnbindDeviceFromEdgeInstanceRequest creates a request to invoke BatchUnbindDeviceFromEdgeInstance API
func CreateBatchUnbindDeviceFromEdgeInstanceRequest() (request *BatchUnbindDeviceFromEdgeInstanceRequest) {
	request = &BatchUnbindDeviceFromEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchUnbindDeviceFromEdgeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchUnbindDeviceFromEdgeInstanceResponse creates a response to parse from BatchUnbindDeviceFromEdgeInstance response
func CreateBatchUnbindDeviceFromEdgeInstanceResponse() (response *BatchUnbindDeviceFromEdgeInstanceResponse) {
	response = &BatchUnbindDeviceFromEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}