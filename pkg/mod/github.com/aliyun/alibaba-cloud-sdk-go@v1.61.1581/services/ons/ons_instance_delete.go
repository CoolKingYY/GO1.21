package ons

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

// OnsInstanceDelete invokes the ons.OnsInstanceDelete API synchronously
func (client *Client) OnsInstanceDelete(request *OnsInstanceDeleteRequest) (response *OnsInstanceDeleteResponse, err error) {
	response = CreateOnsInstanceDeleteResponse()
	err = client.DoAction(request, response)
	return
}

// OnsInstanceDeleteWithChan invokes the ons.OnsInstanceDelete API asynchronously
func (client *Client) OnsInstanceDeleteWithChan(request *OnsInstanceDeleteRequest) (<-chan *OnsInstanceDeleteResponse, <-chan error) {
	responseChan := make(chan *OnsInstanceDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsInstanceDelete(request)
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

// OnsInstanceDeleteWithCallback invokes the ons.OnsInstanceDelete API asynchronously
func (client *Client) OnsInstanceDeleteWithCallback(request *OnsInstanceDeleteRequest, callback func(response *OnsInstanceDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsInstanceDeleteResponse
		var err error
		defer close(result)
		response, err = client.OnsInstanceDelete(request)
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

// OnsInstanceDeleteRequest is the request struct for api OnsInstanceDelete
type OnsInstanceDeleteRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// OnsInstanceDeleteResponse is the response struct for api OnsInstanceDelete
type OnsInstanceDeleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsInstanceDeleteRequest creates a request to invoke OnsInstanceDelete API
func CreateOnsInstanceDeleteRequest() (request *OnsInstanceDeleteRequest) {
	request = &OnsInstanceDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsInstanceDelete", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsInstanceDeleteResponse creates a response to parse from OnsInstanceDelete response
func CreateOnsInstanceDeleteResponse() (response *OnsInstanceDeleteResponse) {
	response = &OnsInstanceDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}