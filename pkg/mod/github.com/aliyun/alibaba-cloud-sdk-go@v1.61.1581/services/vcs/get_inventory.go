package vcs

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

// GetInventory invokes the vcs.GetInventory API synchronously
func (client *Client) GetInventory(request *GetInventoryRequest) (response *GetInventoryResponse, err error) {
	response = CreateGetInventoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetInventoryWithChan invokes the vcs.GetInventory API asynchronously
func (client *Client) GetInventoryWithChan(request *GetInventoryRequest) (<-chan *GetInventoryResponse, <-chan error) {
	responseChan := make(chan *GetInventoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInventory(request)
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

// GetInventoryWithCallback invokes the vcs.GetInventory API asynchronously
func (client *Client) GetInventoryWithCallback(request *GetInventoryRequest, callback func(response *GetInventoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInventoryResponse
		var err error
		defer close(result)
		response, err = client.GetInventory(request)
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

// GetInventoryRequest is the request struct for api GetInventory
type GetInventoryRequest struct {
	*requests.RpcRequest
	CommodityCode string `position:"Body" name:"CommodityCode"`
}

// GetInventoryResponse is the response struct for api GetInventory
type GetInventoryResponse struct {
	*responses.BaseResponse
	Success bool `json:"Success" xml:"Success"`
	Data    Data `json:"Data" xml:"Data"`
}

// CreateGetInventoryRequest creates a request to invoke GetInventory API
func CreateGetInventoryRequest() (request *GetInventoryRequest) {
	request = &GetInventoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetInventory", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInventoryResponse creates a response to parse from GetInventory response
func CreateGetInventoryResponse() (response *GetInventoryResponse) {
	response = &GetInventoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
