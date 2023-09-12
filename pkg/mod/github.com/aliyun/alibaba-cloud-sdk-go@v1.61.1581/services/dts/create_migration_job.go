package dts

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

// CreateMigrationJob invokes the dts.CreateMigrationJob API synchronously
func (client *Client) CreateMigrationJob(request *CreateMigrationJobRequest) (response *CreateMigrationJobResponse, err error) {
	response = CreateCreateMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMigrationJobWithChan invokes the dts.CreateMigrationJob API asynchronously
func (client *Client) CreateMigrationJobWithChan(request *CreateMigrationJobRequest) (<-chan *CreateMigrationJobResponse, <-chan error) {
	responseChan := make(chan *CreateMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMigrationJob(request)
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

// CreateMigrationJobWithCallback invokes the dts.CreateMigrationJob API asynchronously
func (client *Client) CreateMigrationJobWithCallback(request *CreateMigrationJobRequest, callback func(response *CreateMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.CreateMigrationJob(request)
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

// CreateMigrationJobRequest is the request struct for api CreateMigrationJob
type CreateMigrationJobRequest struct {
	*requests.RpcRequest
	ClientToken       string `position:"Query" name:"ClientToken"`
	OwnerId           string `position:"Query" name:"OwnerId"`
	AccountId         string `position:"Query" name:"AccountId"`
	Region            string `position:"Query" name:"Region"`
	MigrationJobClass string `position:"Query" name:"MigrationJobClass"`
}

// CreateMigrationJobResponse is the response struct for api CreateMigrationJob
type CreateMigrationJobResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	MigrationJobId string `json:"MigrationJobId" xml:"MigrationJobId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        string `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateCreateMigrationJobRequest creates a request to invoke CreateMigrationJob API
func CreateCreateMigrationJobRequest() (request *CreateMigrationJobRequest) {
	request = &CreateMigrationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "CreateMigrationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMigrationJobResponse creates a response to parse from CreateMigrationJob response
func CreateCreateMigrationJobResponse() (response *CreateMigrationJobResponse) {
	response = &CreateMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
