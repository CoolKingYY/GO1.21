package dataworks_public

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

// ListReminds invokes the dataworks_public.ListReminds API synchronously
func (client *Client) ListReminds(request *ListRemindsRequest) (response *ListRemindsResponse, err error) {
	response = CreateListRemindsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRemindsWithChan invokes the dataworks_public.ListReminds API asynchronously
func (client *Client) ListRemindsWithChan(request *ListRemindsRequest) (<-chan *ListRemindsResponse, <-chan error) {
	responseChan := make(chan *ListRemindsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListReminds(request)
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

// ListRemindsWithCallback invokes the dataworks_public.ListReminds API asynchronously
func (client *Client) ListRemindsWithCallback(request *ListRemindsRequest, callback func(response *ListRemindsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRemindsResponse
		var err error
		defer close(result)
		response, err = client.ListReminds(request)
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

// ListRemindsRequest is the request struct for api ListReminds
type ListRemindsRequest struct {
	*requests.RpcRequest
	SearchText  string           `position:"Body" name:"SearchText"`
	Founder     string           `position:"Body" name:"Founder"`
	RemindTypes string           `position:"Body" name:"RemindTypes"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	AlertTarget string           `position:"Body" name:"AlertTarget"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	NodeId      requests.Integer `position:"Body" name:"NodeId"`
}

// ListRemindsResponse is the response struct for api ListReminds
type ListRemindsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListRemindsRequest creates a request to invoke ListReminds API
func CreateListRemindsRequest() (request *ListRemindsRequest) {
	request = &ListRemindsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListReminds", "", "")
	request.Method = requests.POST
	return
}

// CreateListRemindsResponse creates a response to parse from ListReminds response
func CreateListRemindsResponse() (response *ListRemindsResponse) {
	response = &ListRemindsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
