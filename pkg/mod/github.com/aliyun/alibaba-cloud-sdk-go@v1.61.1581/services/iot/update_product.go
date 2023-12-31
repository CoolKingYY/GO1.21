package iot

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

// UpdateProduct invokes the iot.UpdateProduct API synchronously
func (client *Client) UpdateProduct(request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
	response = CreateUpdateProductResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProductWithChan invokes the iot.UpdateProduct API asynchronously
func (client *Client) UpdateProductWithChan(request *UpdateProductRequest) (<-chan *UpdateProductResponse, <-chan error) {
	responseChan := make(chan *UpdateProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProduct(request)
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

// UpdateProductWithCallback invokes the iot.UpdateProduct API asynchronously
func (client *Client) UpdateProductWithCallback(request *UpdateProductRequest, callback func(response *UpdateProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProductResponse
		var err error
		defer close(result)
		response, err = client.UpdateProduct(request)
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

// UpdateProductRequest is the request struct for api UpdateProduct
type UpdateProductRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	Description       string `position:"Query" name:"Description"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductName       string `position:"Query" name:"ProductName"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// UpdateProductResponse is the response struct for api UpdateProduct
type UpdateProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateProductRequest creates a request to invoke UpdateProduct API
func CreateUpdateProductRequest() (request *UpdateProductRequest) {
	request = &UpdateProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateProduct", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateProductResponse creates a response to parse from UpdateProduct response
func CreateUpdateProductResponse() (response *UpdateProductResponse) {
	response = &UpdateProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
