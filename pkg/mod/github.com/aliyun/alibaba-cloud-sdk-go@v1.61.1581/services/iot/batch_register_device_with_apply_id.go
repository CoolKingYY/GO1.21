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

// BatchRegisterDeviceWithApplyId invokes the iot.BatchRegisterDeviceWithApplyId API synchronously
func (client *Client) BatchRegisterDeviceWithApplyId(request *BatchRegisterDeviceWithApplyIdRequest) (response *BatchRegisterDeviceWithApplyIdResponse, err error) {
	response = CreateBatchRegisterDeviceWithApplyIdResponse()
	err = client.DoAction(request, response)
	return
}

// BatchRegisterDeviceWithApplyIdWithChan invokes the iot.BatchRegisterDeviceWithApplyId API asynchronously
func (client *Client) BatchRegisterDeviceWithApplyIdWithChan(request *BatchRegisterDeviceWithApplyIdRequest) (<-chan *BatchRegisterDeviceWithApplyIdResponse, <-chan error) {
	responseChan := make(chan *BatchRegisterDeviceWithApplyIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchRegisterDeviceWithApplyId(request)
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

// BatchRegisterDeviceWithApplyIdWithCallback invokes the iot.BatchRegisterDeviceWithApplyId API asynchronously
func (client *Client) BatchRegisterDeviceWithApplyIdWithCallback(request *BatchRegisterDeviceWithApplyIdRequest, callback func(response *BatchRegisterDeviceWithApplyIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchRegisterDeviceWithApplyIdResponse
		var err error
		defer close(result)
		response, err = client.BatchRegisterDeviceWithApplyId(request)
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

// BatchRegisterDeviceWithApplyIdRequest is the request struct for api BatchRegisterDeviceWithApplyId
type BatchRegisterDeviceWithApplyIdRequest struct {
	*requests.RpcRequest
	RealTenantId      string           `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string           `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	ApplyId           requests.Integer `position:"Query" name:"ApplyId"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
}

// BatchRegisterDeviceWithApplyIdResponse is the response struct for api BatchRegisterDeviceWithApplyId
type BatchRegisterDeviceWithApplyIdResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateBatchRegisterDeviceWithApplyIdRequest creates a request to invoke BatchRegisterDeviceWithApplyId API
func CreateBatchRegisterDeviceWithApplyIdRequest() (request *BatchRegisterDeviceWithApplyIdRequest) {
	request = &BatchRegisterDeviceWithApplyIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchRegisterDeviceWithApplyId", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchRegisterDeviceWithApplyIdResponse creates a response to parse from BatchRegisterDeviceWithApplyId response
func CreateBatchRegisterDeviceWithApplyIdResponse() (response *BatchRegisterDeviceWithApplyIdResponse) {
	response = &BatchRegisterDeviceWithApplyIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
