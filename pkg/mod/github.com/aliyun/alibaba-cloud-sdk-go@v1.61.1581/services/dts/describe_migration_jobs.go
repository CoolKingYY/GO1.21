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

// DescribeMigrationJobs invokes the dts.DescribeMigrationJobs API synchronously
func (client *Client) DescribeMigrationJobs(request *DescribeMigrationJobsRequest) (response *DescribeMigrationJobsResponse, err error) {
	response = CreateDescribeMigrationJobsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMigrationJobsWithChan invokes the dts.DescribeMigrationJobs API asynchronously
func (client *Client) DescribeMigrationJobsWithChan(request *DescribeMigrationJobsRequest) (<-chan *DescribeMigrationJobsResponse, <-chan error) {
	responseChan := make(chan *DescribeMigrationJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMigrationJobs(request)
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

// DescribeMigrationJobsWithCallback invokes the dts.DescribeMigrationJobs API asynchronously
func (client *Client) DescribeMigrationJobsWithCallback(request *DescribeMigrationJobsRequest, callback func(response *DescribeMigrationJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMigrationJobsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMigrationJobs(request)
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

// DescribeMigrationJobsRequest is the request struct for api DescribeMigrationJobs
type DescribeMigrationJobsRequest struct {
	*requests.RpcRequest
	InstFilterRegion string                      `position:"Query" name:"InstFilterRegion"`
	PageNum          requests.Integer            `position:"Query" name:"PageNum"`
	OwnerId          string                      `position:"Query" name:"OwnerId"`
	AccountId        string                      `position:"Query" name:"AccountId"`
	PageSize         requests.Integer            `position:"Query" name:"PageSize"`
	MigrationJobName string                      `position:"Query" name:"MigrationJobName"`
	Tag              *[]DescribeMigrationJobsTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// DescribeMigrationJobsTag is a repeated param struct in DescribeMigrationJobsRequest
type DescribeMigrationJobsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeMigrationJobsResponse is the response struct for api DescribeMigrationJobs
type DescribeMigrationJobsResponse struct {
	*responses.BaseResponse
	RequestId        string        `json:"RequestId" xml:"RequestId"`
	ErrCode          string        `json:"ErrCode" xml:"ErrCode"`
	PageRecordCount  int           `json:"PageRecordCount" xml:"PageRecordCount"`
	Success          string        `json:"Success" xml:"Success"`
	TotalRecordCount int64         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	ErrMessage       string        `json:"ErrMessage" xml:"ErrMessage"`
	PageNumber       int           `json:"PageNumber" xml:"PageNumber"`
	MigrationJobs    MigrationJobs `json:"MigrationJobs" xml:"MigrationJobs"`
}

// CreateDescribeMigrationJobsRequest creates a request to invoke DescribeMigrationJobs API
func CreateDescribeMigrationJobsRequest() (request *DescribeMigrationJobsRequest) {
	request = &DescribeMigrationJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeMigrationJobs", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMigrationJobsResponse creates a response to parse from DescribeMigrationJobs response
func CreateDescribeMigrationJobsResponse() (response *DescribeMigrationJobsResponse) {
	response = &DescribeMigrationJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
