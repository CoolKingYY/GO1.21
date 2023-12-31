package outboundbot

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

// QueryJobs invokes the outboundbot.QueryJobs API synchronously
func (client *Client) QueryJobs(request *QueryJobsRequest) (response *QueryJobsResponse, err error) {
	response = CreateQueryJobsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryJobsWithChan invokes the outboundbot.QueryJobs API asynchronously
func (client *Client) QueryJobsWithChan(request *QueryJobsRequest) (<-chan *QueryJobsResponse, <-chan error) {
	responseChan := make(chan *QueryJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryJobs(request)
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

// QueryJobsWithCallback invokes the outboundbot.QueryJobs API asynchronously
func (client *Client) QueryJobsWithCallback(request *QueryJobsRequest, callback func(response *QueryJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryJobsResponse
		var err error
		defer close(result)
		response, err = client.QueryJobs(request)
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

// QueryJobsRequest is the request struct for api QueryJobs
type QueryJobsRequest struct {
	*requests.RpcRequest
	TimeAlignment string           `position:"Query" name:"TimeAlignment"`
	PhoneNumber   string           `position:"Query" name:"PhoneNumber"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	ContactName   string           `position:"Query" name:"ContactName"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	JobGroupId    string           `position:"Query" name:"JobGroupId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ScenarioId    string           `position:"Query" name:"ScenarioId"`
}

// QueryJobsResponse is the response struct for api QueryJobs
type QueryJobsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Jobs           Jobs   `json:"Jobs" xml:"Jobs"`
}

// CreateQueryJobsRequest creates a request to invoke QueryJobs API
func CreateQueryJobsRequest() (request *QueryJobsRequest) {
	request = &QueryJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "QueryJobs", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryJobsResponse creates a response to parse from QueryJobs response
func CreateQueryJobsResponse() (response *QueryJobsResponse) {
	response = &QueryJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
