package vod

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

// DescribeVodDomainConfigs invokes the vod.DescribeVodDomainConfigs API synchronously
func (client *Client) DescribeVodDomainConfigs(request *DescribeVodDomainConfigsRequest) (response *DescribeVodDomainConfigsResponse, err error) {
	response = CreateDescribeVodDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainConfigsWithChan invokes the vod.DescribeVodDomainConfigs API asynchronously
func (client *Client) DescribeVodDomainConfigsWithChan(request *DescribeVodDomainConfigsRequest) (<-chan *DescribeVodDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainConfigs(request)
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

// DescribeVodDomainConfigsWithCallback invokes the vod.DescribeVodDomainConfigs API asynchronously
func (client *Client) DescribeVodDomainConfigsWithCallback(request *DescribeVodDomainConfigsRequest, callback func(response *DescribeVodDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainConfigs(request)
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

// DescribeVodDomainConfigsRequest is the request struct for api DescribeVodDomainConfigs
type DescribeVodDomainConfigsRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeVodDomainConfigsResponse is the response struct for api DescribeVodDomainConfigs
type DescribeVodDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	DomainConfigs DomainConfigs `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeVodDomainConfigsRequest creates a request to invoke DescribeVodDomainConfigs API
func CreateDescribeVodDomainConfigsRequest() (request *DescribeVodDomainConfigsRequest) {
	request = &DescribeVodDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainConfigs", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodDomainConfigsResponse creates a response to parse from DescribeVodDomainConfigs response
func CreateDescribeVodDomainConfigsResponse() (response *DescribeVodDomainConfigsResponse) {
	response = &DescribeVodDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
