package quickbi_public

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

// UpdateTicketNum invokes the quickbi_public.UpdateTicketNum API synchronously
func (client *Client) UpdateTicketNum(request *UpdateTicketNumRequest) (response *UpdateTicketNumResponse, err error) {
	response = CreateUpdateTicketNumResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTicketNumWithChan invokes the quickbi_public.UpdateTicketNum API asynchronously
func (client *Client) UpdateTicketNumWithChan(request *UpdateTicketNumRequest) (<-chan *UpdateTicketNumResponse, <-chan error) {
	responseChan := make(chan *UpdateTicketNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTicketNum(request)
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

// UpdateTicketNumWithCallback invokes the quickbi_public.UpdateTicketNum API asynchronously
func (client *Client) UpdateTicketNumWithCallback(request *UpdateTicketNumRequest, callback func(response *UpdateTicketNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTicketNumResponse
		var err error
		defer close(result)
		response, err = client.UpdateTicketNum(request)
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

// UpdateTicketNumRequest is the request struct for api UpdateTicketNum
type UpdateTicketNumRequest struct {
	*requests.RpcRequest
	Ticket      string           `position:"Query" name:"Ticket"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	TicketNum   requests.Integer `position:"Query" name:"TicketNum"`
	SignType    string           `position:"Query" name:"SignType"`
}

// UpdateTicketNumResponse is the response struct for api UpdateTicketNum
type UpdateTicketNumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTicketNumRequest creates a request to invoke UpdateTicketNum API
func CreateUpdateTicketNumRequest() (request *UpdateTicketNumRequest) {
	request = &UpdateTicketNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateTicketNum", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTicketNumResponse creates a response to parse from UpdateTicketNum response
func CreateUpdateTicketNumResponse() (response *UpdateTicketNumResponse) {
	response = &UpdateTicketNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}