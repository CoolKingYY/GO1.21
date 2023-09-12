package emas_appmonitor

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

// QueryApiErrorGroupTrend invokes the emas_appmonitor.QueryApiErrorGroupTrend API synchronously
func (client *Client) QueryApiErrorGroupTrend(request *QueryApiErrorGroupTrendRequest) (response *QueryApiErrorGroupTrendResponse, err error) {
	response = CreateQueryApiErrorGroupTrendResponse()
	err = client.DoAction(request, response)
	return
}

// QueryApiErrorGroupTrendWithChan invokes the emas_appmonitor.QueryApiErrorGroupTrend API asynchronously
func (client *Client) QueryApiErrorGroupTrendWithChan(request *QueryApiErrorGroupTrendRequest) (<-chan *QueryApiErrorGroupTrendResponse, <-chan error) {
	responseChan := make(chan *QueryApiErrorGroupTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryApiErrorGroupTrend(request)
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

// QueryApiErrorGroupTrendWithCallback invokes the emas_appmonitor.QueryApiErrorGroupTrend API asynchronously
func (client *Client) QueryApiErrorGroupTrendWithCallback(request *QueryApiErrorGroupTrendRequest, callback func(response *QueryApiErrorGroupTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryApiErrorGroupTrendResponse
		var err error
		defer close(result)
		response, err = client.QueryApiErrorGroupTrend(request)
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

// QueryApiErrorGroupTrendRequest is the request struct for api QueryApiErrorGroupTrend
type QueryApiErrorGroupTrendRequest struct {
	*requests.RpcRequest
	AppVersionStrategy string           `position:"Body" name:"AppVersionStrategy"`
	StartTime          requests.Integer `position:"Body" name:"StartTime"`
	Group              string           `position:"Body" name:"Group"`
	IntervalMinutes    requests.Integer `position:"Body" name:"IntervalMinutes"`
	UniqueAppId        string           `position:"Body" name:"UniqueAppId"`
	EndTime            requests.Integer `position:"Body" name:"EndTime"`
	AppVersion         *[]string        `position:"Body" name:"AppVersion"  type:"Repeated"`
}

// QueryApiErrorGroupTrendResponse is the response struct for api QueryApiErrorGroupTrend
type QueryApiErrorGroupTrendResponse struct {
	*responses.BaseResponse
	RequestId        string             `json:"RequestId" xml:"RequestId"`
	MetricResultList []MetricResultItem `json:"MetricResultList" xml:"MetricResultList"`
}

// CreateQueryApiErrorGroupTrendRequest creates a request to invoke QueryApiErrorGroupTrend API
func CreateQueryApiErrorGroupTrendRequest() (request *QueryApiErrorGroupTrendRequest) {
	request = &QueryApiErrorGroupTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "QueryApiErrorGroupTrend", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryApiErrorGroupTrendResponse creates a response to parse from QueryApiErrorGroupTrend response
func CreateQueryApiErrorGroupTrendResponse() (response *QueryApiErrorGroupTrendResponse) {
	response = &QueryApiErrorGroupTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
