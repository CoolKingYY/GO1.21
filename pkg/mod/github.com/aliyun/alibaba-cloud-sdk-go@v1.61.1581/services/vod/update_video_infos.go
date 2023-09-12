package vod

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

// UpdateVideoInfos invokes the vod.UpdateVideoInfos API synchronously
func (client *Client) UpdateVideoInfos(request *UpdateVideoInfosRequest) (response *UpdateVideoInfosResponse, err error) {
	response = CreateUpdateVideoInfosResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVideoInfosWithChan invokes the vod.UpdateVideoInfos API asynchronously
func (client *Client) UpdateVideoInfosWithChan(request *UpdateVideoInfosRequest) (<-chan *UpdateVideoInfosResponse, <-chan error) {
	responseChan := make(chan *UpdateVideoInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVideoInfos(request)
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

// UpdateVideoInfosWithCallback invokes the vod.UpdateVideoInfos API asynchronously
func (client *Client) UpdateVideoInfosWithCallback(request *UpdateVideoInfosRequest, callback func(response *UpdateVideoInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVideoInfosResponse
		var err error
		defer close(result)
		response, err = client.UpdateVideoInfos(request)
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

// UpdateVideoInfosRequest is the request struct for api UpdateVideoInfos
type UpdateVideoInfosRequest struct {
	*requests.RpcRequest
	UpdateContent string `position:"Query" name:"UpdateContent"`
}

// UpdateVideoInfosResponse is the response struct for api UpdateVideoInfos
type UpdateVideoInfosResponse struct {
	*responses.BaseResponse
	RequestId         string   `json:"RequestId" xml:"RequestId"`
	ForbiddenVideoIds []string `json:"ForbiddenVideoIds" xml:"ForbiddenVideoIds"`
	NonExistVideoIds  []string `json:"NonExistVideoIds" xml:"NonExistVideoIds"`
}

// CreateUpdateVideoInfosRequest creates a request to invoke UpdateVideoInfos API
func CreateUpdateVideoInfosRequest() (request *UpdateVideoInfosRequest) {
	request = &UpdateVideoInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateVideoInfos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateVideoInfosResponse creates a response to parse from UpdateVideoInfos response
func CreateUpdateVideoInfosResponse() (response *UpdateVideoInfosResponse) {
	response = &UpdateVideoInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
