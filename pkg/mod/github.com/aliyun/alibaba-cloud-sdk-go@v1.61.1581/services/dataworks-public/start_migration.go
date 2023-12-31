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

// StartMigration invokes the dataworks_public.StartMigration API synchronously
func (client *Client) StartMigration(request *StartMigrationRequest) (response *StartMigrationResponse, err error) {
	response = CreateStartMigrationResponse()
	err = client.DoAction(request, response)
	return
}

// StartMigrationWithChan invokes the dataworks_public.StartMigration API asynchronously
func (client *Client) StartMigrationWithChan(request *StartMigrationRequest) (<-chan *StartMigrationResponse, <-chan error) {
	responseChan := make(chan *StartMigrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartMigration(request)
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

// StartMigrationWithCallback invokes the dataworks_public.StartMigration API asynchronously
func (client *Client) StartMigrationWithCallback(request *StartMigrationRequest, callback func(response *StartMigrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartMigrationResponse
		var err error
		defer close(result)
		response, err = client.StartMigration(request)
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

// StartMigrationRequest is the request struct for api StartMigration
type StartMigrationRequest struct {
	*requests.RpcRequest
	MigrationId requests.Integer `position:"Body" name:"MigrationId"`
	ProjectId   requests.Integer `position:"Body" name:"ProjectId"`
}

// StartMigrationResponse is the response struct for api StartMigration
type StartMigrationResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateStartMigrationRequest creates a request to invoke StartMigration API
func CreateStartMigrationRequest() (request *StartMigrationRequest) {
	request = &StartMigrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "StartMigration", "", "")
	request.Method = requests.POST
	return
}

// CreateStartMigrationResponse creates a response to parse from StartMigration response
func CreateStartMigrationResponse() (response *StartMigrationResponse) {
	response = &StartMigrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
