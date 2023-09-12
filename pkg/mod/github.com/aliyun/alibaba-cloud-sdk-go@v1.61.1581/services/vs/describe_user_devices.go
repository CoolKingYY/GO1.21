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

// DescribeUserDevices invokes the vs.DescribeUserDevices API synchronously
func (client *Client) DescribeUserDevices(request *DescribeUserDevicesRequest) (response *DescribeUserDevicesResponse, err error) {
	response = CreateDescribeUserDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserDevicesWithChan invokes the vs.DescribeUserDevices API asynchronously
func (client *Client) DescribeUserDevicesWithChan(request *DescribeUserDevicesRequest) (<-chan *DescribeUserDevicesResponse, <-chan error) {
	responseChan := make(chan *DescribeUserDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserDevices(request)
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

// DescribeUserDevicesWithCallback invokes the vs.DescribeUserDevices API asynchronously
func (client *Client) DescribeUserDevicesWithCallback(request *DescribeUserDevicesRequest, callback func(response *DescribeUserDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserDevicesResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserDevices(request)
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

// DescribeUserDevicesRequest is the request struct for api DescribeUserDevices
type DescribeUserDevicesRequest struct {
	*requests.RpcRequest
	EnsInstanceIds string           `position:"Query" name:"EnsInstanceIds"`
	ShowLog        string           `position:"Query" name:"ShowLog"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	ServerName     string           `position:"Query" name:"ServerName"`
}

// DescribeUserDevicesResponse is the response struct for api DescribeUserDevices
type DescribeUserDevicesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	List      []UserDevice `json:"List" xml:"List"`
}

// CreateDescribeUserDevicesRequest creates a request to invoke DescribeUserDevices API
func CreateDescribeUserDevicesRequest() (request *DescribeUserDevicesRequest) {
	request = &DescribeUserDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeUserDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUserDevicesResponse creates a response to parse from DescribeUserDevices response
func CreateDescribeUserDevicesResponse() (response *DescribeUserDevicesResponse) {
	response = &DescribeUserDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
