package scsp

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

// CloseTicket invokes the scsp.CloseTicket API synchronously
func (client *Client) CloseTicket(request *CloseTicketRequest) (response *CloseTicketResponse, err error) {
	response = CreateCloseTicketResponse()
	err = client.DoAction(request, response)
	return
}

// CloseTicketWithChan invokes the scsp.CloseTicket API asynchronously
func (client *Client) CloseTicketWithChan(request *CloseTicketRequest) (<-chan *CloseTicketResponse, <-chan error) {
	responseChan := make(chan *CloseTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloseTicket(request)
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

// CloseTicketWithCallback invokes the scsp.CloseTicket API asynchronously
func (client *Client) CloseTicketWithCallback(request *CloseTicketRequest, callback func(response *CloseTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloseTicketResponse
		var err error
		defer close(result)
		response, err = client.CloseTicket(request)
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

// CloseTicketRequest is the request struct for api CloseTicket
type CloseTicketRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Body"`
	InstanceId  string           `position:"Body"`
	TicketId    requests.Integer `position:"Body"`
	ActionItems string           `position:"Body"`
	OperatorId  requests.Integer `position:"Body"`
}

// CloseTicketResponse is the response struct for api CloseTicket
type CloseTicketResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCloseTicketRequest creates a request to invoke CloseTicket API
func CreateCloseTicketRequest() (request *CloseTicketRequest) {
	request = &CloseTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "CloseTicket", "", "")
	request.Method = requests.POST
	return
}

// CreateCloseTicketResponse creates a response to parse from CloseTicket response
func CreateCloseTicketResponse() (response *CloseTicketResponse) {
	response = &CloseTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
