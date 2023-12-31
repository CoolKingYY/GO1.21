package dataworks_public

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

// GetTopicInfluence invokes the dataworks_public.GetTopicInfluence API synchronously
func (client *Client) GetTopicInfluence(request *GetTopicInfluenceRequest) (response *GetTopicInfluenceResponse, err error) {
	response = CreateGetTopicInfluenceResponse()
	err = client.DoAction(request, response)
	return
}

// GetTopicInfluenceWithChan invokes the dataworks_public.GetTopicInfluence API asynchronously
func (client *Client) GetTopicInfluenceWithChan(request *GetTopicInfluenceRequest) (<-chan *GetTopicInfluenceResponse, <-chan error) {
	responseChan := make(chan *GetTopicInfluenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTopicInfluence(request)
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

// GetTopicInfluenceWithCallback invokes the dataworks_public.GetTopicInfluence API asynchronously
func (client *Client) GetTopicInfluenceWithCallback(request *GetTopicInfluenceRequest, callback func(response *GetTopicInfluenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTopicInfluenceResponse
		var err error
		defer close(result)
		response, err = client.GetTopicInfluence(request)
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

// GetTopicInfluenceRequest is the request struct for api GetTopicInfluence
type GetTopicInfluenceRequest struct {
	*requests.RpcRequest
	TopicId requests.Integer `position:"Body" name:"TopicId"`
}

// GetTopicInfluenceResponse is the response struct for api GetTopicInfluence
type GetTopicInfluenceResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetTopicInfluenceRequest creates a request to invoke GetTopicInfluence API
func CreateGetTopicInfluenceRequest() (request *GetTopicInfluenceRequest) {
	request = &GetTopicInfluenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetTopicInfluence", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTopicInfluenceResponse creates a response to parse from GetTopicInfluence response
func CreateGetTopicInfluenceResponse() (response *GetTopicInfluenceResponse) {
	response = &GetTopicInfluenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
