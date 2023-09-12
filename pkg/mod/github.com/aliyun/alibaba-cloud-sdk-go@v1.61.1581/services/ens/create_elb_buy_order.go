package ens

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

// CreateElbBuyOrder invokes the ens.CreateElbBuyOrder API synchronously
func (client *Client) CreateElbBuyOrder(request *CreateElbBuyOrderRequest) (response *CreateElbBuyOrderResponse, err error) {
	response = CreateCreateElbBuyOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateElbBuyOrderWithChan invokes the ens.CreateElbBuyOrder API asynchronously
func (client *Client) CreateElbBuyOrderWithChan(request *CreateElbBuyOrderRequest) (<-chan *CreateElbBuyOrderResponse, <-chan error) {
	responseChan := make(chan *CreateElbBuyOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateElbBuyOrder(request)
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

// CreateElbBuyOrderWithCallback invokes the ens.CreateElbBuyOrder API asynchronously
func (client *Client) CreateElbBuyOrderWithCallback(request *CreateElbBuyOrderRequest, callback func(response *CreateElbBuyOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateElbBuyOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateElbBuyOrder(request)
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

// CreateElbBuyOrderRequest is the request struct for api CreateElbBuyOrder
type CreateElbBuyOrderRequest struct {
	*requests.RpcRequest
	OrderDetails string `position:"Query" name:"OrderDetails"`
}

// CreateElbBuyOrderResponse is the response struct for api CreateElbBuyOrder
type CreateElbBuyOrderResponse struct {
	*responses.BaseResponse
	RequestId       string   `json:"RequestId" xml:"RequestId"`
	LoadBalancerIds []string `json:"LoadBalancerIds" xml:"LoadBalancerIds"`
}

// CreateCreateElbBuyOrderRequest creates a request to invoke CreateElbBuyOrder API
func CreateCreateElbBuyOrderRequest() (request *CreateElbBuyOrderRequest) {
	request = &CreateElbBuyOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateElbBuyOrder", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateElbBuyOrderResponse creates a response to parse from CreateElbBuyOrder response
func CreateCreateElbBuyOrderResponse() (response *CreateElbBuyOrderResponse) {
	response = &CreateElbBuyOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
