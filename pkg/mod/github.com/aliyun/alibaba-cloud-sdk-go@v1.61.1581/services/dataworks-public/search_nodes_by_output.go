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

// SearchNodesByOutput invokes the dataworks_public.SearchNodesByOutput API synchronously
func (client *Client) SearchNodesByOutput(request *SearchNodesByOutputRequest) (response *SearchNodesByOutputResponse, err error) {
	response = CreateSearchNodesByOutputResponse()
	err = client.DoAction(request, response)
	return
}

// SearchNodesByOutputWithChan invokes the dataworks_public.SearchNodesByOutput API asynchronously
func (client *Client) SearchNodesByOutputWithChan(request *SearchNodesByOutputRequest) (<-chan *SearchNodesByOutputResponse, <-chan error) {
	responseChan := make(chan *SearchNodesByOutputResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchNodesByOutput(request)
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

// SearchNodesByOutputWithCallback invokes the dataworks_public.SearchNodesByOutput API asynchronously
func (client *Client) SearchNodesByOutputWithCallback(request *SearchNodesByOutputRequest, callback func(response *SearchNodesByOutputResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchNodesByOutputResponse
		var err error
		defer close(result)
		response, err = client.SearchNodesByOutput(request)
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

// SearchNodesByOutputRequest is the request struct for api SearchNodesByOutput
type SearchNodesByOutputRequest struct {
	*requests.RpcRequest
	ProjectEnv string `position:"Body" name:"ProjectEnv"`
	Outputs    string `position:"Body" name:"Outputs"`
}

// SearchNodesByOutputResponse is the response struct for api SearchNodesByOutput
type SearchNodesByOutputResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           map[string]interface{} `json:"Data" xml:"Data"`
	ErrorMessage   string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Success        bool                   `json:"Success" xml:"Success"`
	ErrorCode      string                 `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateSearchNodesByOutputRequest creates a request to invoke SearchNodesByOutput API
func CreateSearchNodesByOutputRequest() (request *SearchNodesByOutputRequest) {
	request = &SearchNodesByOutputRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SearchNodesByOutput", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchNodesByOutputResponse creates a response to parse from SearchNodesByOutput response
func CreateSearchNodesByOutputResponse() (response *SearchNodesByOutputResponse) {
	response = &SearchNodesByOutputResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
