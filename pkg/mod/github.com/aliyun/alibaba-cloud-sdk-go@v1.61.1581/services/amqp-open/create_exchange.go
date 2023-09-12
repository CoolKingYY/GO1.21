package amqp_open

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

// CreateExchange invokes the amqp_open.CreateExchange API synchronously
// api document: https://help.aliyun.com/api/amqp-open/createexchange.html
func (client *Client) CreateExchange(request *CreateExchangeRequest) (response *CreateExchangeResponse, err error) {
	response = CreateCreateExchangeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExchangeWithChan invokes the amqp_open.CreateExchange API asynchronously
// api document: https://help.aliyun.com/api/amqp-open/createexchange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateExchangeWithChan(request *CreateExchangeRequest) (<-chan *CreateExchangeResponse, <-chan error) {
	responseChan := make(chan *CreateExchangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExchange(request)
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

// CreateExchangeWithCallback invokes the amqp_open.CreateExchange API asynchronously
// api document: https://help.aliyun.com/api/amqp-open/createexchange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateExchangeWithCallback(request *CreateExchangeRequest, callback func(response *CreateExchangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExchangeResponse
		var err error
		defer close(result)
		response, err = client.CreateExchange(request)
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

// CreateExchangeRequest is the request struct for api CreateExchange
type CreateExchangeRequest struct {
	*requests.RpcRequest
	Internal          requests.Boolean `position:"Body" name:"Internal"`
	ExchangeName      string           `position:"Body" name:"ExchangeName"`
	InstanceId        string           `position:"Body" name:"InstanceId"`
	AlternateExchange string           `position:"Body" name:"AlternateExchange"`
	AutoDeleteState   requests.Boolean `position:"Body" name:"AutoDeleteState"`
	ExchangeType      string           `position:"Body" name:"ExchangeType"`
	VirtualHost       string           `position:"Body" name:"VirtualHost"`
}

// CreateExchangeResponse is the response struct for api CreateExchange
type CreateExchangeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateExchangeRequest creates a request to invoke CreateExchange API
func CreateCreateExchangeRequest() (request *CreateExchangeRequest) {
	request = &CreateExchangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("amqp-open", "2019-12-12", "CreateExchange", "onsproxy", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateExchangeResponse creates a response to parse from CreateExchange response
func CreateCreateExchangeResponse() (response *CreateExchangeResponse) {
	response = &CreateExchangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
