package retailadvqa_public

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

// CheckSmsReceiptExportStatus invokes the retailadvqa_public.CheckSmsReceiptExportStatus API synchronously
func (client *Client) CheckSmsReceiptExportStatus(request *CheckSmsReceiptExportStatusRequest) (response *CheckSmsReceiptExportStatusResponse, err error) {
	response = CreateCheckSmsReceiptExportStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSmsReceiptExportStatusWithChan invokes the retailadvqa_public.CheckSmsReceiptExportStatus API asynchronously
func (client *Client) CheckSmsReceiptExportStatusWithChan(request *CheckSmsReceiptExportStatusRequest) (<-chan *CheckSmsReceiptExportStatusResponse, <-chan error) {
	responseChan := make(chan *CheckSmsReceiptExportStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSmsReceiptExportStatus(request)
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

// CheckSmsReceiptExportStatusWithCallback invokes the retailadvqa_public.CheckSmsReceiptExportStatus API asynchronously
func (client *Client) CheckSmsReceiptExportStatusWithCallback(request *CheckSmsReceiptExportStatusRequest, callback func(response *CheckSmsReceiptExportStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSmsReceiptExportStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckSmsReceiptExportStatus(request)
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

// CheckSmsReceiptExportStatusRequest is the request struct for api CheckSmsReceiptExportStatus
type CheckSmsReceiptExportStatusRequest struct {
	*requests.RpcRequest
	AccessId     string `position:"Query" name:"AccessId"`
	MarketTaskId string `position:"Query" name:"MarketTaskId"`
	WorkspaceId  string `position:"Query" name:"WorkspaceId"`
}

// CheckSmsReceiptExportStatusResponse is the response struct for api CheckSmsReceiptExportStatus
type CheckSmsReceiptExportStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCheckSmsReceiptExportStatusRequest creates a request to invoke CheckSmsReceiptExportStatus API
func CreateCheckSmsReceiptExportStatusRequest() (request *CheckSmsReceiptExportStatusRequest) {
	request = &CheckSmsReceiptExportStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CheckSmsReceiptExportStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckSmsReceiptExportStatusResponse creates a response to parse from CheckSmsReceiptExportStatus response
func CreateCheckSmsReceiptExportStatusResponse() (response *CheckSmsReceiptExportStatusResponse) {
	response = &CheckSmsReceiptExportStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
