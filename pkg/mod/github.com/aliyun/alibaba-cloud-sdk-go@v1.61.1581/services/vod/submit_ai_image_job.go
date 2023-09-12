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

// SubmitAIImageJob invokes the vod.SubmitAIImageJob API synchronously
func (client *Client) SubmitAIImageJob(request *SubmitAIImageJobRequest) (response *SubmitAIImageJobResponse, err error) {
	response = CreateSubmitAIImageJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAIImageJobWithChan invokes the vod.SubmitAIImageJob API asynchronously
func (client *Client) SubmitAIImageJobWithChan(request *SubmitAIImageJobRequest) (<-chan *SubmitAIImageJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAIImageJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAIImageJob(request)
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

// SubmitAIImageJobWithCallback invokes the vod.SubmitAIImageJob API asynchronously
func (client *Client) SubmitAIImageJobWithCallback(request *SubmitAIImageJobRequest, callback func(response *SubmitAIImageJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAIImageJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAIImageJob(request)
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

// SubmitAIImageJobRequest is the request struct for api SubmitAIImageJob
type SubmitAIImageJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	AIPipelineId         string `position:"Query" name:"AIPipelineId"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	AITemplateId         string `position:"Query" name:"AITemplateId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

// SubmitAIImageJobResponse is the response struct for api SubmitAIImageJob
type SubmitAIImageJobResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitAIImageJobRequest creates a request to invoke SubmitAIImageJob API
func CreateSubmitAIImageJobRequest() (request *SubmitAIImageJobRequest) {
	request = &SubmitAIImageJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitAIImageJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitAIImageJobResponse creates a response to parse from SubmitAIImageJob response
func CreateSubmitAIImageJobResponse() (response *SubmitAIImageJobResponse) {
	response = &SubmitAIImageJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
