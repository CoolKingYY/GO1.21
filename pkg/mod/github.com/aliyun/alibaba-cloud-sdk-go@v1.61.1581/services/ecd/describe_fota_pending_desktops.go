package ecd

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

// DescribeFotaPendingDesktops invokes the ecd.DescribeFotaPendingDesktops API synchronously
func (client *Client) DescribeFotaPendingDesktops(request *DescribeFotaPendingDesktopsRequest) (response *DescribeFotaPendingDesktopsResponse, err error) {
	response = CreateDescribeFotaPendingDesktopsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFotaPendingDesktopsWithChan invokes the ecd.DescribeFotaPendingDesktops API asynchronously
func (client *Client) DescribeFotaPendingDesktopsWithChan(request *DescribeFotaPendingDesktopsRequest) (<-chan *DescribeFotaPendingDesktopsResponse, <-chan error) {
	responseChan := make(chan *DescribeFotaPendingDesktopsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFotaPendingDesktops(request)
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

// DescribeFotaPendingDesktopsWithCallback invokes the ecd.DescribeFotaPendingDesktops API asynchronously
func (client *Client) DescribeFotaPendingDesktopsWithCallback(request *DescribeFotaPendingDesktopsRequest, callback func(response *DescribeFotaPendingDesktopsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFotaPendingDesktopsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFotaPendingDesktops(request)
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

// DescribeFotaPendingDesktopsRequest is the request struct for api DescribeFotaPendingDesktops
type DescribeFotaPendingDesktopsRequest struct {
	*requests.RpcRequest
	TaskUid    string           `position:"Query" name:"TaskUid"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeFotaPendingDesktopsResponse is the response struct for api DescribeFotaPendingDesktops
type DescribeFotaPendingDesktopsResponse struct {
	*responses.BaseResponse
	NextToken           string               `json:"NextToken" xml:"NextToken"`
	RequestId           string               `json:"RequestId" xml:"RequestId"`
	FotaPendingDesktops []FotaPendingDesktop `json:"FotaPendingDesktops" xml:"FotaPendingDesktops"`
}

// CreateDescribeFotaPendingDesktopsRequest creates a request to invoke DescribeFotaPendingDesktops API
func CreateDescribeFotaPendingDesktopsRequest() (request *DescribeFotaPendingDesktopsRequest) {
	request = &DescribeFotaPendingDesktopsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeFotaPendingDesktops", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeFotaPendingDesktopsResponse creates a response to parse from DescribeFotaPendingDesktops response
func CreateDescribeFotaPendingDesktopsResponse() (response *DescribeFotaPendingDesktopsResponse) {
	response = &DescribeFotaPendingDesktopsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
