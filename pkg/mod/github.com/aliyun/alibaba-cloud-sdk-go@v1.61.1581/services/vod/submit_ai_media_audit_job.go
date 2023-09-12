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

// SubmitAIMediaAuditJob invokes the vod.SubmitAIMediaAuditJob API synchronously
func (client *Client) SubmitAIMediaAuditJob(request *SubmitAIMediaAuditJobRequest) (response *SubmitAIMediaAuditJobResponse, err error) {
	response = CreateSubmitAIMediaAuditJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAIMediaAuditJobWithChan invokes the vod.SubmitAIMediaAuditJob API asynchronously
func (client *Client) SubmitAIMediaAuditJobWithChan(request *SubmitAIMediaAuditJobRequest) (<-chan *SubmitAIMediaAuditJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAIMediaAuditJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAIMediaAuditJob(request)
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

// SubmitAIMediaAuditJobWithCallback invokes the vod.SubmitAIMediaAuditJob API asynchronously
func (client *Client) SubmitAIMediaAuditJobWithCallback(request *SubmitAIMediaAuditJobRequest, callback func(response *SubmitAIMediaAuditJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAIMediaAuditJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAIMediaAuditJob(request)
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

// SubmitAIMediaAuditJobRequest is the request struct for api SubmitAIMediaAuditJob
type SubmitAIMediaAuditJobRequest struct {
	*requests.RpcRequest
	UserData                string `position:"Query" name:"UserData"`
	MediaId                 string `position:"Query" name:"MediaId"`
	TemplateId              string `position:"Query" name:"TemplateId"`
	MediaAuditConfiguration string `position:"Query" name:"MediaAuditConfiguration"`
	MediaType               string `position:"Query" name:"MediaType"`
}

// SubmitAIMediaAuditJobResponse is the response struct for api SubmitAIMediaAuditJob
type SubmitAIMediaAuditJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MediaId   string `json:"MediaId" xml:"MediaId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitAIMediaAuditJobRequest creates a request to invoke SubmitAIMediaAuditJob API
func CreateSubmitAIMediaAuditJobRequest() (request *SubmitAIMediaAuditJobRequest) {
	request = &SubmitAIMediaAuditJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitAIMediaAuditJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitAIMediaAuditJobResponse creates a response to parse from SubmitAIMediaAuditJob response
func CreateSubmitAIMediaAuditJobResponse() (response *SubmitAIMediaAuditJobResponse) {
	response = &SubmitAIMediaAuditJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}