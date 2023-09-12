package dms_enterprise

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

// ListInstanceLoginAuditLog invokes the dms_enterprise.ListInstanceLoginAuditLog API synchronously
func (client *Client) ListInstanceLoginAuditLog(request *ListInstanceLoginAuditLogRequest) (response *ListInstanceLoginAuditLogResponse, err error) {
	response = CreateListInstanceLoginAuditLogResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceLoginAuditLogWithChan invokes the dms_enterprise.ListInstanceLoginAuditLog API asynchronously
func (client *Client) ListInstanceLoginAuditLogWithChan(request *ListInstanceLoginAuditLogRequest) (<-chan *ListInstanceLoginAuditLogResponse, <-chan error) {
	responseChan := make(chan *ListInstanceLoginAuditLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceLoginAuditLog(request)
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

// ListInstanceLoginAuditLogWithCallback invokes the dms_enterprise.ListInstanceLoginAuditLog API asynchronously
func (client *Client) ListInstanceLoginAuditLogWithCallback(request *ListInstanceLoginAuditLogRequest, callback func(response *ListInstanceLoginAuditLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceLoginAuditLogResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceLoginAuditLog(request)
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

// ListInstanceLoginAuditLogRequest is the request struct for api ListInstanceLoginAuditLog
type ListInstanceLoginAuditLogRequest struct {
	*requests.RpcRequest
	SearchName string           `position:"Query" name:"SearchName"`
	OpUserName string           `position:"Query" name:"OpUserName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListInstanceLoginAuditLogResponse is the response struct for api ListInstanceLoginAuditLog
type ListInstanceLoginAuditLogResponse struct {
	*responses.BaseResponse
	TotalCount                int64                     `json:"TotalCount" xml:"TotalCount"`
	RequestId                 string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode                 string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage              string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success                   bool                      `json:"Success" xml:"Success"`
	InstanceLoginAuditLogList InstanceLoginAuditLogList `json:"InstanceLoginAuditLogList" xml:"InstanceLoginAuditLogList"`
}

// CreateListInstanceLoginAuditLogRequest creates a request to invoke ListInstanceLoginAuditLog API
func CreateListInstanceLoginAuditLogRequest() (request *ListInstanceLoginAuditLogRequest) {
	request = &ListInstanceLoginAuditLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListInstanceLoginAuditLog", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstanceLoginAuditLogResponse creates a response to parse from ListInstanceLoginAuditLog response
func CreateListInstanceLoginAuditLogResponse() (response *ListInstanceLoginAuditLogResponse) {
	response = &ListInstanceLoginAuditLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
