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

// ListCampaignTrendingReport invokes the ccc.ListCampaignTrendingReport API synchronously
func (client *Client) ListCampaignTrendingReport(request *ListCampaignTrendingReportRequest) (response *ListCampaignTrendingReportResponse, err error) {
	response = CreateListCampaignTrendingReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListCampaignTrendingReportWithChan invokes the ccc.ListCampaignTrendingReport API asynchronously
func (client *Client) ListCampaignTrendingReportWithChan(request *ListCampaignTrendingReportRequest) (<-chan *ListCampaignTrendingReportResponse, <-chan error) {
	responseChan := make(chan *ListCampaignTrendingReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCampaignTrendingReport(request)
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

// ListCampaignTrendingReportWithCallback invokes the ccc.ListCampaignTrendingReport API asynchronously
func (client *Client) ListCampaignTrendingReportWithCallback(request *ListCampaignTrendingReportRequest, callback func(response *ListCampaignTrendingReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCampaignTrendingReportResponse
		var err error
		defer close(result)
		response, err = client.ListCampaignTrendingReport(request)
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

// ListCampaignTrendingReportRequest is the request struct for api ListCampaignTrendingReport
type ListCampaignTrendingReportRequest struct {
	*requests.RpcRequest
	CampaignId string           `position:"Query" name:"CampaignId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// ListCampaignTrendingReportResponse is the response struct for api ListCampaignTrendingReport
type ListCampaignTrendingReportResponse struct {
	*responses.BaseResponse
	Code           string         `json:"Code" xml:"Code"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Data           []TrendingList `json:"Data" xml:"Data"`
}

// CreateListCampaignTrendingReportRequest creates a request to invoke ListCampaignTrendingReport API
func CreateListCampaignTrendingReportRequest() (request *ListCampaignTrendingReportRequest) {
	request = &ListCampaignTrendingReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListCampaignTrendingReport", "", "")
	request.Method = requests.GET
	return
}

// CreateListCampaignTrendingReportResponse creates a response to parse from ListCampaignTrendingReport response
func CreateListCampaignTrendingReportResponse() (response *ListCampaignTrendingReportResponse) {
	response = &ListCampaignTrendingReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
