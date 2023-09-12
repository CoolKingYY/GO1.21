package airec

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

// QuerySyncReportAggregation invokes the airec.QuerySyncReportAggregation API synchronously
func (client *Client) QuerySyncReportAggregation(request *QuerySyncReportAggregationRequest) (response *QuerySyncReportAggregationResponse, err error) {
	response = CreateQuerySyncReportAggregationResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySyncReportAggregationWithChan invokes the airec.QuerySyncReportAggregation API asynchronously
func (client *Client) QuerySyncReportAggregationWithChan(request *QuerySyncReportAggregationRequest) (<-chan *QuerySyncReportAggregationResponse, <-chan error) {
	responseChan := make(chan *QuerySyncReportAggregationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySyncReportAggregation(request)
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

// QuerySyncReportAggregationWithCallback invokes the airec.QuerySyncReportAggregation API asynchronously
func (client *Client) QuerySyncReportAggregationWithCallback(request *QuerySyncReportAggregationRequest, callback func(response *QuerySyncReportAggregationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySyncReportAggregationResponse
		var err error
		defer close(result)
		response, err = client.QuerySyncReportAggregation(request)
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

// QuerySyncReportAggregationRequest is the request struct for api QuerySyncReportAggregation
type QuerySyncReportAggregationRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"instanceId"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
}

// QuerySyncReportAggregationResponse is the response struct for api QuerySyncReportAggregation
type QuerySyncReportAggregationResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"code" xml:"code"`
	Message   string                 `json:"message" xml:"message"`
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateQuerySyncReportAggregationRequest creates a request to invoke QuerySyncReportAggregation API
func CreateQuerySyncReportAggregationRequest() (request *QuerySyncReportAggregationRequest) {
	request = &QuerySyncReportAggregationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "QuerySyncReportAggregation", "/v2/openapi/instances/[instanceId]/sync-reports/aggregation", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQuerySyncReportAggregationResponse creates a response to parse from QuerySyncReportAggregation response
func CreateQuerySyncReportAggregationResponse() (response *QuerySyncReportAggregationResponse) {
	response = &QuerySyncReportAggregationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
