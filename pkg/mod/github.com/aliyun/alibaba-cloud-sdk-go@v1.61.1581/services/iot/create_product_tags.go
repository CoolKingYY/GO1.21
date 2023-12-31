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

// CreateProductTags invokes the iot.CreateProductTags API synchronously
func (client *Client) CreateProductTags(request *CreateProductTagsRequest) (response *CreateProductTagsResponse, err error) {
	response = CreateCreateProductTagsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProductTagsWithChan invokes the iot.CreateProductTags API asynchronously
func (client *Client) CreateProductTagsWithChan(request *CreateProductTagsRequest) (<-chan *CreateProductTagsResponse, <-chan error) {
	responseChan := make(chan *CreateProductTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProductTags(request)
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

// CreateProductTagsWithCallback invokes the iot.CreateProductTags API asynchronously
func (client *Client) CreateProductTagsWithCallback(request *CreateProductTagsRequest, callback func(response *CreateProductTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProductTagsResponse
		var err error
		defer close(result)
		response, err = client.CreateProductTags(request)
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

// CreateProductTagsRequest is the request struct for api CreateProductTags
type CreateProductTagsRequest struct {
	*requests.RpcRequest
	RealTenantId      string                         `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string                         `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string                         `position:"Query" name:"IotInstanceId"`
	ProductKey        string                         `position:"Query" name:"ProductKey"`
	ProductTag        *[]CreateProductTagsProductTag `position:"Query" name:"ProductTag"  type:"Repeated"`
	ApiProduct        string                         `position:"Body" name:"ApiProduct"`
	ApiRevision       string                         `position:"Body" name:"ApiRevision"`
}

// CreateProductTagsProductTag is a repeated param struct in CreateProductTagsRequest
type CreateProductTagsProductTag struct {
	TagValue string `name:"TagValue"`
	TagKey   string `name:"TagKey"`
}

// CreateProductTagsResponse is the response struct for api CreateProductTags
type CreateProductTagsResponse struct {
	*responses.BaseResponse
	RequestId          string                                `json:"RequestId" xml:"RequestId"`
	Success            bool                                  `json:"Success" xml:"Success"`
	ErrorMessage       string                                `json:"ErrorMessage" xml:"ErrorMessage"`
	Code               string                                `json:"Code" xml:"Code"`
	InvalidProductTags InvalidProductTagsInCreateProductTags `json:"InvalidProductTags" xml:"InvalidProductTags"`
}

// CreateCreateProductTagsRequest creates a request to invoke CreateProductTags API
func CreateCreateProductTagsRequest() (request *CreateProductTagsRequest) {
	request = &CreateProductTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateProductTags", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateProductTagsResponse creates a response to parse from CreateProductTags response
func CreateCreateProductTagsResponse() (response *CreateProductTagsResponse) {
	response = &CreateProductTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
