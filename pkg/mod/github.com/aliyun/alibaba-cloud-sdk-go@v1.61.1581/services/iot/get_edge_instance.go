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

// GetEdgeInstance invokes the iot.GetEdgeInstance API synchronously
func (client *Client) GetEdgeInstance(request *GetEdgeInstanceRequest) (response *GetEdgeInstanceResponse, err error) {
	response = CreateGetEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetEdgeInstanceWithChan invokes the iot.GetEdgeInstance API asynchronously
func (client *Client) GetEdgeInstanceWithChan(request *GetEdgeInstanceRequest) (<-chan *GetEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *GetEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEdgeInstance(request)
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

// GetEdgeInstanceWithCallback invokes the iot.GetEdgeInstance API asynchronously
func (client *Client) GetEdgeInstanceWithCallback(request *GetEdgeInstanceRequest, callback func(response *GetEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetEdgeInstance(request)
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

// GetEdgeInstanceRequest is the request struct for api GetEdgeInstance
type GetEdgeInstanceRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetEdgeInstanceResponse is the response struct for api GetEdgeInstance
type GetEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetEdgeInstanceRequest creates a request to invoke GetEdgeInstance API
func CreateGetEdgeInstanceRequest() (request *GetEdgeInstanceRequest) {
	request = &GetEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetEdgeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateGetEdgeInstanceResponse creates a response to parse from GetEdgeInstance response
func CreateGetEdgeInstanceResponse() (response *GetEdgeInstanceResponse) {
	response = &GetEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
