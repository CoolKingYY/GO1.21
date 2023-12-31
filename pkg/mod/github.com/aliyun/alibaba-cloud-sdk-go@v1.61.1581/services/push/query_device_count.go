package push

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

// QueryDeviceCount invokes the push.QueryDeviceCount API synchronously
func (client *Client) QueryDeviceCount(request *QueryDeviceCountRequest) (response *QueryDeviceCountResponse, err error) {
	response = CreateQueryDeviceCountResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceCountWithChan invokes the push.QueryDeviceCount API asynchronously
func (client *Client) QueryDeviceCountWithChan(request *QueryDeviceCountRequest) (<-chan *QueryDeviceCountResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceCount(request)
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

// QueryDeviceCountWithCallback invokes the push.QueryDeviceCount API asynchronously
func (client *Client) QueryDeviceCountWithCallback(request *QueryDeviceCountRequest, callback func(response *QueryDeviceCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceCountResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceCount(request)
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

// QueryDeviceCountRequest is the request struct for api QueryDeviceCount
type QueryDeviceCountRequest struct {
	*requests.RpcRequest
	Target      string           `position:"Query" name:"Target"`
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	TargetValue string           `position:"Query" name:"TargetValue"`
}

// QueryDeviceCountResponse is the response struct for api QueryDeviceCount
type QueryDeviceCountResponse struct {
	*responses.BaseResponse
	DeviceCount int64  `json:"DeviceCount" xml:"DeviceCount"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateQueryDeviceCountRequest creates a request to invoke QueryDeviceCount API
func CreateQueryDeviceCountRequest() (request *QueryDeviceCountRequest) {
	request = &QueryDeviceCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryDeviceCount", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceCountResponse creates a response to parse from QueryDeviceCount response
func CreateQueryDeviceCountResponse() (response *QueryDeviceCountResponse) {
	response = &QueryDeviceCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
