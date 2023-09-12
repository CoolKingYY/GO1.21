package rds

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

// DescribeLocalAvailableRecoveryTime invokes the rds.DescribeLocalAvailableRecoveryTime API synchronously
func (client *Client) DescribeLocalAvailableRecoveryTime(request *DescribeLocalAvailableRecoveryTimeRequest) (response *DescribeLocalAvailableRecoveryTimeResponse, err error) {
	response = CreateDescribeLocalAvailableRecoveryTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLocalAvailableRecoveryTimeWithChan invokes the rds.DescribeLocalAvailableRecoveryTime API asynchronously
func (client *Client) DescribeLocalAvailableRecoveryTimeWithChan(request *DescribeLocalAvailableRecoveryTimeRequest) (<-chan *DescribeLocalAvailableRecoveryTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeLocalAvailableRecoveryTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLocalAvailableRecoveryTime(request)
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

// DescribeLocalAvailableRecoveryTimeWithCallback invokes the rds.DescribeLocalAvailableRecoveryTime API asynchronously
func (client *Client) DescribeLocalAvailableRecoveryTimeWithCallback(request *DescribeLocalAvailableRecoveryTimeRequest, callback func(response *DescribeLocalAvailableRecoveryTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLocalAvailableRecoveryTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLocalAvailableRecoveryTime(request)
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

// DescribeLocalAvailableRecoveryTimeRequest is the request struct for api DescribeLocalAvailableRecoveryTime
type DescribeLocalAvailableRecoveryTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Region               string           `position:"Query" name:"Region"`
}

// DescribeLocalAvailableRecoveryTimeResponse is the response struct for api DescribeLocalAvailableRecoveryTime
type DescribeLocalAvailableRecoveryTimeResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string `json:"DBInstanceId" xml:"DBInstanceId"`
	RecoveryBeginTime string `json:"RecoveryBeginTime" xml:"RecoveryBeginTime"`
	RecoveryEndTime   string `json:"RecoveryEndTime" xml:"RecoveryEndTime"`
}

// CreateDescribeLocalAvailableRecoveryTimeRequest creates a request to invoke DescribeLocalAvailableRecoveryTime API
func CreateDescribeLocalAvailableRecoveryTimeRequest() (request *DescribeLocalAvailableRecoveryTimeRequest) {
	request = &DescribeLocalAvailableRecoveryTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeLocalAvailableRecoveryTime", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLocalAvailableRecoveryTimeResponse creates a response to parse from DescribeLocalAvailableRecoveryTime response
func CreateDescribeLocalAvailableRecoveryTimeResponse() (response *DescribeLocalAvailableRecoveryTimeResponse) {
	response = &DescribeLocalAvailableRecoveryTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
