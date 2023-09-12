package mts

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

// QueryPipelineList invokes the mts.QueryPipelineList API synchronously
func (client *Client) QueryPipelineList(request *QueryPipelineListRequest) (response *QueryPipelineListResponse, err error) {
	response = CreateQueryPipelineListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPipelineListWithChan invokes the mts.QueryPipelineList API asynchronously
func (client *Client) QueryPipelineListWithChan(request *QueryPipelineListRequest) (<-chan *QueryPipelineListResponse, <-chan error) {
	responseChan := make(chan *QueryPipelineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPipelineList(request)
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

// QueryPipelineListWithCallback invokes the mts.QueryPipelineList API asynchronously
func (client *Client) QueryPipelineListWithCallback(request *QueryPipelineListRequest, callback func(response *QueryPipelineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPipelineListResponse
		var err error
		defer close(result)
		response, err = client.QueryPipelineList(request)
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

// QueryPipelineListRequest is the request struct for api QueryPipelineList
type QueryPipelineListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineIds          string           `position:"Query" name:"PipelineIds"`
}

// QueryPipelineListResponse is the response struct for api QueryPipelineList
type QueryPipelineListResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	NonExistPids NonExistPids                    `json:"NonExistPids" xml:"NonExistPids"`
	PipelineList PipelineListInQueryPipelineList `json:"PipelineList" xml:"PipelineList"`
}

// CreateQueryPipelineListRequest creates a request to invoke QueryPipelineList API
func CreateQueryPipelineListRequest() (request *QueryPipelineListRequest) {
	request = &QueryPipelineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryPipelineList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryPipelineListResponse creates a response to parse from QueryPipelineList response
func CreateQueryPipelineListResponse() (response *QueryPipelineListResponse) {
	response = &QueryPipelineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
