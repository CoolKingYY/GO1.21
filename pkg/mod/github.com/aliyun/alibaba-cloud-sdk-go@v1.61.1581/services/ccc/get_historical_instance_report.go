package ccc

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

// GetHistoricalInstanceReport invokes the ccc.GetHistoricalInstanceReport API synchronously
func (client *Client) GetHistoricalInstanceReport(request *GetHistoricalInstanceReportRequest) (response *GetHistoricalInstanceReportResponse, err error) {
	response = CreateGetHistoricalInstanceReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetHistoricalInstanceReportWithChan invokes the ccc.GetHistoricalInstanceReport API asynchronously
func (client *Client) GetHistoricalInstanceReportWithChan(request *GetHistoricalInstanceReportRequest) (<-chan *GetHistoricalInstanceReportResponse, <-chan error) {
	responseChan := make(chan *GetHistoricalInstanceReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHistoricalInstanceReport(request)
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

// GetHistoricalInstanceReportWithCallback invokes the ccc.GetHistoricalInstanceReport API asynchronously
func (client *Client) GetHistoricalInstanceReportWithCallback(request *GetHistoricalInstanceReportRequest, callback func(response *GetHistoricalInstanceReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHistoricalInstanceReportResponse
		var err error
		defer close(result)
		response, err = client.GetHistoricalInstanceReport(request)
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

// GetHistoricalInstanceReportRequest is the request struct for api GetHistoricalInstanceReport
type GetHistoricalInstanceReportRequest struct {
	*requests.RpcRequest
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// GetHistoricalInstanceReportResponse is the response struct for api GetHistoricalInstanceReport
type GetHistoricalInstanceReportResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetHistoricalInstanceReportRequest creates a request to invoke GetHistoricalInstanceReport API
func CreateGetHistoricalInstanceReportRequest() (request *GetHistoricalInstanceReportRequest) {
	request = &GetHistoricalInstanceReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetHistoricalInstanceReport", "", "")
	request.Method = requests.POST
	return
}

// CreateGetHistoricalInstanceReportResponse creates a response to parse from GetHistoricalInstanceReport response
func CreateGetHistoricalInstanceReportResponse() (response *GetHistoricalInstanceReportResponse) {
	response = &GetHistoricalInstanceReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
