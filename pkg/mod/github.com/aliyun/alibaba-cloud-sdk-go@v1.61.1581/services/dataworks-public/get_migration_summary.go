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

// GetMigrationSummary invokes the dataworks_public.GetMigrationSummary API synchronously
func (client *Client) GetMigrationSummary(request *GetMigrationSummaryRequest) (response *GetMigrationSummaryResponse, err error) {
	response = CreateGetMigrationSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetMigrationSummaryWithChan invokes the dataworks_public.GetMigrationSummary API asynchronously
func (client *Client) GetMigrationSummaryWithChan(request *GetMigrationSummaryRequest) (<-chan *GetMigrationSummaryResponse, <-chan error) {
	responseChan := make(chan *GetMigrationSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMigrationSummary(request)
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

// GetMigrationSummaryWithCallback invokes the dataworks_public.GetMigrationSummary API asynchronously
func (client *Client) GetMigrationSummaryWithCallback(request *GetMigrationSummaryRequest, callback func(response *GetMigrationSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMigrationSummaryResponse
		var err error
		defer close(result)
		response, err = client.GetMigrationSummary(request)
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

// GetMigrationSummaryRequest is the request struct for api GetMigrationSummary
type GetMigrationSummaryRequest struct {
	*requests.RpcRequest
	MigrationId requests.Integer `position:"Body" name:"MigrationId"`
	ProjectId   requests.Integer `position:"Body" name:"ProjectId"`
}

// GetMigrationSummaryResponse is the response struct for api GetMigrationSummary
type GetMigrationSummaryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetMigrationSummaryRequest creates a request to invoke GetMigrationSummary API
func CreateGetMigrationSummaryRequest() (request *GetMigrationSummaryRequest) {
	request = &GetMigrationSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMigrationSummary", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMigrationSummaryResponse creates a response to parse from GetMigrationSummary response
func CreateGetMigrationSummaryResponse() (response *GetMigrationSummaryResponse) {
	response = &GetMigrationSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
