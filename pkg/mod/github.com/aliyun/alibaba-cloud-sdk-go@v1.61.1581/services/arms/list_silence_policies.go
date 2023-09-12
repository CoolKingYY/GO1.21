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

// ListSilencePolicies invokes the arms.ListSilencePolicies API synchronously
func (client *Client) ListSilencePolicies(request *ListSilencePoliciesRequest) (response *ListSilencePoliciesResponse, err error) {
	response = CreateListSilencePoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSilencePoliciesWithChan invokes the arms.ListSilencePolicies API asynchronously
func (client *Client) ListSilencePoliciesWithChan(request *ListSilencePoliciesRequest) (<-chan *ListSilencePoliciesResponse, <-chan error) {
	responseChan := make(chan *ListSilencePoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSilencePolicies(request)
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

// ListSilencePoliciesWithCallback invokes the arms.ListSilencePolicies API asynchronously
func (client *Client) ListSilencePoliciesWithCallback(request *ListSilencePoliciesRequest, callback func(response *ListSilencePoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSilencePoliciesResponse
		var err error
		defer close(result)
		response, err = client.ListSilencePolicies(request)
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

// ListSilencePoliciesRequest is the request struct for api ListSilencePolicies
type ListSilencePoliciesRequest struct {
	*requests.RpcRequest
	Size     requests.Integer `position:"Query" name:"Size"`
	Name     string           `position:"Query" name:"Name"`
	IsDetail requests.Boolean `position:"Query" name:"IsDetail"`
	Page     requests.Integer `position:"Query" name:"Page"`
}

// ListSilencePoliciesResponse is the response struct for api ListSilencePolicies
type ListSilencePoliciesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateListSilencePoliciesRequest creates a request to invoke ListSilencePolicies API
func CreateListSilencePoliciesRequest() (request *ListSilencePoliciesRequest) {
	request = &ListSilencePoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListSilencePolicies", "", "")
	request.Method = requests.POST
	return
}

// CreateListSilencePoliciesResponse creates a response to parse from ListSilencePolicies response
func CreateListSilencePoliciesResponse() (response *ListSilencePoliciesResponse) {
	response = &ListSilencePoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
