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

// DescribeSuspEventOverview invokes the ecd.DescribeSuspEventOverview API synchronously
func (client *Client) DescribeSuspEventOverview(request *DescribeSuspEventOverviewRequest) (response *DescribeSuspEventOverviewResponse, err error) {
	response = CreateDescribeSuspEventOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspEventOverviewWithChan invokes the ecd.DescribeSuspEventOverview API asynchronously
func (client *Client) DescribeSuspEventOverviewWithChan(request *DescribeSuspEventOverviewRequest) (<-chan *DescribeSuspEventOverviewResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspEventOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspEventOverview(request)
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

// DescribeSuspEventOverviewWithCallback invokes the ecd.DescribeSuspEventOverview API asynchronously
func (client *Client) DescribeSuspEventOverviewWithCallback(request *DescribeSuspEventOverviewRequest, callback func(response *DescribeSuspEventOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspEventOverviewResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspEventOverview(request)
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

// DescribeSuspEventOverviewRequest is the request struct for api DescribeSuspEventOverview
type DescribeSuspEventOverviewRequest struct {
	*requests.RpcRequest
}

// DescribeSuspEventOverviewResponse is the response struct for api DescribeSuspEventOverview
type DescribeSuspEventOverviewResponse struct {
	*responses.BaseResponse
	SuspiciousCount int    `json:"SuspiciousCount" xml:"SuspiciousCount"`
	RemindCount     int    `json:"RemindCount" xml:"RemindCount"`
	SeriousCount    int    `json:"SeriousCount" xml:"SeriousCount"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeSuspEventOverviewRequest creates a request to invoke DescribeSuspEventOverview API
func CreateDescribeSuspEventOverviewRequest() (request *DescribeSuspEventOverviewRequest) {
	request = &DescribeSuspEventOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeSuspEventOverview", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSuspEventOverviewResponse creates a response to parse from DescribeSuspEventOverview response
func CreateDescribeSuspEventOverviewResponse() (response *DescribeSuspEventOverviewResponse) {
	response = &DescribeSuspEventOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
