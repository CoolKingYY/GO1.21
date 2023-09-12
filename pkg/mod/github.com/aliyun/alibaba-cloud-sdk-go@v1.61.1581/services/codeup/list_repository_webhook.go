package codeup

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

// ListRepositoryWebhook invokes the codeup.ListRepositoryWebhook API synchronously
func (client *Client) ListRepositoryWebhook(request *ListRepositoryWebhookRequest) (response *ListRepositoryWebhookResponse, err error) {
	response = CreateListRepositoryWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepositoryWebhookWithChan invokes the codeup.ListRepositoryWebhook API asynchronously
func (client *Client) ListRepositoryWebhookWithChan(request *ListRepositoryWebhookRequest) (<-chan *ListRepositoryWebhookResponse, <-chan error) {
	responseChan := make(chan *ListRepositoryWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepositoryWebhook(request)
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

// ListRepositoryWebhookWithCallback invokes the codeup.ListRepositoryWebhook API asynchronously
func (client *Client) ListRepositoryWebhookWithCallback(request *ListRepositoryWebhookRequest, callback func(response *ListRepositoryWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepositoryWebhookResponse
		var err error
		defer close(result)
		response, err = client.ListRepositoryWebhook(request)
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

// ListRepositoryWebhookRequest is the request struct for api ListRepositoryWebhook
type ListRepositoryWebhookRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	Page           requests.Integer `position:"Query" name:"Page"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// ListRepositoryWebhookResponse is the response struct for api ListRepositoryWebhook
type ListRepositoryWebhookResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool         `json:"Success" xml:"Success"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int64        `json:"Total" xml:"Total"`
	Result       []ResultItem `json:"Result" xml:"Result"`
}

// CreateListRepositoryWebhookRequest creates a request to invoke ListRepositoryWebhook API
func CreateListRepositoryWebhookRequest() (request *ListRepositoryWebhookRequest) {
	request = &ListRepositoryWebhookRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListRepositoryWebhook", "/api/v3/projects/[ProjectId]/hooks", "", "")
	request.Method = requests.GET
	return
}

// CreateListRepositoryWebhookResponse creates a response to parse from ListRepositoryWebhook response
func CreateListRepositoryWebhookResponse() (response *ListRepositoryWebhookResponse) {
	response = &ListRepositoryWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
