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

// QueryTicketActions invokes the scsp.QueryTicketActions API synchronously
func (client *Client) QueryTicketActions(request *QueryTicketActionsRequest) (response *QueryTicketActionsResponse, err error) {
	response = CreateQueryTicketActionsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTicketActionsWithChan invokes the scsp.QueryTicketActions API asynchronously
func (client *Client) QueryTicketActionsWithChan(request *QueryTicketActionsRequest) (<-chan *QueryTicketActionsResponse, <-chan error) {
	responseChan := make(chan *QueryTicketActionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTicketActions(request)
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

// QueryTicketActionsWithCallback invokes the scsp.QueryTicketActions API asynchronously
func (client *Client) QueryTicketActionsWithCallback(request *QueryTicketActionsRequest, callback func(response *QueryTicketActionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTicketActionsResponse
		var err error
		defer close(result)
		response, err = client.QueryTicketActions(request)
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

// QueryTicketActionsRequest is the request struct for api QueryTicketActions
type QueryTicketActionsRequest struct {
	*requests.RpcRequest
	InstanceId     string    `position:"Body"`
	TicketId       string    `position:"Body"`
	ActionCodeList *[]string `position:"Body" name:"ActionCodeList"  type:"Repeated"`
}

// QueryTicketActionsResponse is the response struct for api QueryTicketActions
type QueryTicketActionsResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryTicketActionsRequest creates a request to invoke QueryTicketActions API
func CreateQueryTicketActionsRequest() (request *QueryTicketActionsRequest) {
	request = &QueryTicketActionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "QueryTicketActions", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTicketActionsResponse creates a response to parse from QueryTicketActions response
func CreateQueryTicketActionsResponse() (response *QueryTicketActionsResponse) {
	response = &QueryTicketActionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}