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

// QueryDeviceDetail invokes the iot.QueryDeviceDetail API synchronously
func (client *Client) QueryDeviceDetail(request *QueryDeviceDetailRequest) (response *QueryDeviceDetailResponse, err error) {
	response = CreateQueryDeviceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceDetailWithChan invokes the iot.QueryDeviceDetail API asynchronously
func (client *Client) QueryDeviceDetailWithChan(request *QueryDeviceDetailRequest) (<-chan *QueryDeviceDetailResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceDetail(request)
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

// QueryDeviceDetailWithCallback invokes the iot.QueryDeviceDetail API asynchronously
func (client *Client) QueryDeviceDetailWithCallback(request *QueryDeviceDetailRequest, callback func(response *QueryDeviceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceDetail(request)
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

// QueryDeviceDetailRequest is the request struct for api QueryDeviceDetail
type QueryDeviceDetailRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotId             string `position:"Query" name:"IotId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	DeviceName        string `position:"Query" name:"DeviceName"`
}

// QueryDeviceDetailResponse is the response struct for api QueryDeviceDetail
type QueryDeviceDetailResponse struct {
	*responses.BaseResponse
	RequestId    string                  `json:"RequestId" xml:"RequestId"`
	Success      bool                    `json:"Success" xml:"Success"`
	Code         string                  `json:"Code" xml:"Code"`
	ErrorMessage string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceDetail `json:"Data" xml:"Data"`
}

// CreateQueryDeviceDetailRequest creates a request to invoke QueryDeviceDetail API
func CreateQueryDeviceDetailRequest() (request *QueryDeviceDetailRequest) {
	request = &QueryDeviceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceDetailResponse creates a response to parse from QueryDeviceDetail response
func CreateQueryDeviceDetailResponse() (response *QueryDeviceDetailResponse) {
	response = &QueryDeviceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}