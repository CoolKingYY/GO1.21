package arms

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

// QueryTraces invokes the arms.QueryTraces API synchronously
func (client *Client) QueryTraces(request *QueryTracesRequest) (response *QueryTracesResponse, err error) {
	response = CreateQueryTracesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTracesWithChan invokes the arms.QueryTraces API asynchronously
func (client *Client) QueryTracesWithChan(request *QueryTracesRequest) (<-chan *QueryTracesResponse, <-chan error) {
	responseChan := make(chan *QueryTracesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTraces(request)
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

// QueryTracesWithCallback invokes the arms.QueryTraces API asynchronously
func (client *Client) QueryTracesWithCallback(request *QueryTracesRequest, callback func(response *QueryTracesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTracesResponse
		var err error
		defer close(result)
		response, err = client.QueryTraces(request)
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

// QueryTracesRequest is the request struct for api QueryTraces
type QueryTracesRequest struct {
	*requests.RpcRequest
	TraceId   string           `position:"Query" name:"TraceId"`
	EndTime   requests.Integer `position:"Query" name:"EndTime"`
	StartTime requests.Integer `position:"Query" name:"StartTime"`
}

// QueryTracesResponse is the response struct for api QueryTraces
type QueryTracesResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	TraceLists []TraceLists `json:"TraceLists" xml:"TraceLists"`
}

// CreateQueryTracesRequest creates a request to invoke QueryTraces API
func CreateQueryTracesRequest() (request *QueryTracesRequest) {
	request = &QueryTracesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "QueryTraces", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTracesResponse creates a response to parse from QueryTraces response
func CreateQueryTracesResponse() (response *QueryTracesResponse) {
	response = &QueryTracesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
