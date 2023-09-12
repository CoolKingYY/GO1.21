package sae

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

// DescribeAppServiceDetail invokes the sae.DescribeAppServiceDetail API synchronously
func (client *Client) DescribeAppServiceDetail(request *DescribeAppServiceDetailRequest) (response *DescribeAppServiceDetailResponse, err error) {
	response = CreateDescribeAppServiceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppServiceDetailWithChan invokes the sae.DescribeAppServiceDetail API asynchronously
func (client *Client) DescribeAppServiceDetailWithChan(request *DescribeAppServiceDetailRequest) (<-chan *DescribeAppServiceDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeAppServiceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppServiceDetail(request)
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

// DescribeAppServiceDetailWithCallback invokes the sae.DescribeAppServiceDetail API asynchronously
func (client *Client) DescribeAppServiceDetailWithCallback(request *DescribeAppServiceDetailRequest, callback func(response *DescribeAppServiceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppServiceDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppServiceDetail(request)
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

// DescribeAppServiceDetailRequest is the request struct for api DescribeAppServiceDetail
type DescribeAppServiceDetailRequest struct {
	*requests.RoaRequest
	ServiceType    string `position:"Query" name:"ServiceType"`
	AppId          string `position:"Query" name:"AppId"`
	ServiceVersion string `position:"Query" name:"ServiceVersion"`
	ServiceName    string `position:"Query" name:"ServiceName"`
	ServiceGroup   string `position:"Query" name:"ServiceGroup"`
}

// DescribeAppServiceDetailResponse is the response struct for api DescribeAppServiceDetail
type DescribeAppServiceDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAppServiceDetailRequest creates a request to invoke DescribeAppServiceDetail API
func CreateDescribeAppServiceDetailRequest() (request *DescribeAppServiceDetailRequest) {
	request = &DescribeAppServiceDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeAppServiceDetail", "/pop/v1/sam/service/describeAppServiceDetail", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeAppServiceDetailResponse creates a response to parse from DescribeAppServiceDetail response
func CreateDescribeAppServiceDetailResponse() (response *DescribeAppServiceDetailResponse) {
	response = &DescribeAppServiceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
