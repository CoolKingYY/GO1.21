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

// QuerySchedulePeriodList invokes the iot.QuerySchedulePeriodList API synchronously
func (client *Client) QuerySchedulePeriodList(request *QuerySchedulePeriodListRequest) (response *QuerySchedulePeriodListResponse, err error) {
	response = CreateQuerySchedulePeriodListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySchedulePeriodListWithChan invokes the iot.QuerySchedulePeriodList API asynchronously
func (client *Client) QuerySchedulePeriodListWithChan(request *QuerySchedulePeriodListRequest) (<-chan *QuerySchedulePeriodListResponse, <-chan error) {
	responseChan := make(chan *QuerySchedulePeriodListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySchedulePeriodList(request)
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

// QuerySchedulePeriodListWithCallback invokes the iot.QuerySchedulePeriodList API asynchronously
func (client *Client) QuerySchedulePeriodListWithCallback(request *QuerySchedulePeriodListRequest, callback func(response *QuerySchedulePeriodListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySchedulePeriodListResponse
		var err error
		defer close(result)
		response, err = client.QuerySchedulePeriodList(request)
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

// QuerySchedulePeriodListRequest is the request struct for api QuerySchedulePeriodList
type QuerySchedulePeriodListRequest struct {
	*requests.RpcRequest
	ScheduleCode  string           `position:"Body" name:"ScheduleCode"`
	PageId        requests.Integer `position:"Body" name:"PageId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QuerySchedulePeriodListResponse is the response struct for api QuerySchedulePeriodList
type QuerySchedulePeriodListResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	Success      bool                          `json:"Success" xml:"Success"`
	Code         string                        `json:"Code" xml:"Code"`
	ErrorMessage string                        `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySchedulePeriodList `json:"Data" xml:"Data"`
}

// CreateQuerySchedulePeriodListRequest creates a request to invoke QuerySchedulePeriodList API
func CreateQuerySchedulePeriodListRequest() (request *QuerySchedulePeriodListRequest) {
	request = &QuerySchedulePeriodListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySchedulePeriodList", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySchedulePeriodListResponse creates a response to parse from QuerySchedulePeriodList response
func CreateQuerySchedulePeriodListResponse() (response *QuerySchedulePeriodListResponse) {
	response = &QuerySchedulePeriodListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
