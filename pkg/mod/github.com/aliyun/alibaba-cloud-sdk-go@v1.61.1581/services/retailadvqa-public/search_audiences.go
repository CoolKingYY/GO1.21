package retailadvqa_public

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

// SearchAudiences invokes the retailadvqa_public.SearchAudiences API synchronously
func (client *Client) SearchAudiences(request *SearchAudiencesRequest) (response *SearchAudiencesResponse, err error) {
	response = CreateSearchAudiencesResponse()
	err = client.DoAction(request, response)
	return
}

// SearchAudiencesWithChan invokes the retailadvqa_public.SearchAudiences API asynchronously
func (client *Client) SearchAudiencesWithChan(request *SearchAudiencesRequest) (<-chan *SearchAudiencesResponse, <-chan error) {
	responseChan := make(chan *SearchAudiencesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchAudiences(request)
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

// SearchAudiencesWithCallback invokes the retailadvqa_public.SearchAudiences API asynchronously
func (client *Client) SearchAudiencesWithCallback(request *SearchAudiencesRequest, callback func(response *SearchAudiencesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchAudiencesResponse
		var err error
		defer close(result)
		response, err = client.SearchAudiences(request)
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

// SearchAudiencesRequest is the request struct for api SearchAudiences
type SearchAudiencesRequest struct {
	*requests.RpcRequest
	AccessId    string           `position:"Query" name:"AccessId"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	ParentId    string           `position:"Query" name:"ParentId"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// SearchAudiencesResponse is the response struct for api SearchAudiences
type SearchAudiencesResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"RequestId" xml:"RequestId"`
	ErrorDesc string                `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string                `json:"TraceId" xml:"TraceId"`
	ErrorCode string                `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool                  `json:"Success" xml:"Success"`
	Data      DataInSearchAudiences `json:"Data" xml:"Data"`
}

// CreateSearchAudiencesRequest creates a request to invoke SearchAudiences API
func CreateSearchAudiencesRequest() (request *SearchAudiencesRequest) {
	request = &SearchAudiencesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "SearchAudiences", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchAudiencesResponse creates a response to parse from SearchAudiences response
func CreateSearchAudiencesResponse() (response *SearchAudiencesResponse) {
	response = &SearchAudiencesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
