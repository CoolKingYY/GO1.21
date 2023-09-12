package domain

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

// QueryTransferOutInfo invokes the domain.QueryTransferOutInfo API synchronously
func (client *Client) QueryTransferOutInfo(request *QueryTransferOutInfoRequest) (response *QueryTransferOutInfoResponse, err error) {
	response = CreateQueryTransferOutInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTransferOutInfoWithChan invokes the domain.QueryTransferOutInfo API asynchronously
func (client *Client) QueryTransferOutInfoWithChan(request *QueryTransferOutInfoRequest) (<-chan *QueryTransferOutInfoResponse, <-chan error) {
	responseChan := make(chan *QueryTransferOutInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTransferOutInfo(request)
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

// QueryTransferOutInfoWithCallback invokes the domain.QueryTransferOutInfo API asynchronously
func (client *Client) QueryTransferOutInfoWithCallback(request *QueryTransferOutInfoRequest, callback func(response *QueryTransferOutInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTransferOutInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryTransferOutInfo(request)
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

// QueryTransferOutInfoRequest is the request struct for api QueryTransferOutInfo
type QueryTransferOutInfoRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryTransferOutInfoResponse is the response struct for api QueryTransferOutInfo
type QueryTransferOutInfoResponse struct {
	*responses.BaseResponse
	RequestId                         string `json:"RequestId" xml:"RequestId"`
	Status                            int    `json:"Status" xml:"Status"`
	Email                             string `json:"Email" xml:"Email"`
	TransferAuthorizationCodeSendDate string `json:"TransferAuthorizationCodeSendDate" xml:"TransferAuthorizationCodeSendDate"`
	ExpirationDate                    string `json:"ExpirationDate" xml:"ExpirationDate"`
	PendingRequestDate                string `json:"PendingRequestDate" xml:"PendingRequestDate"`
	ResultCode                        string `json:"ResultCode" xml:"ResultCode"`
	ResultMsg                         string `json:"ResultMsg" xml:"ResultMsg"`
}

// CreateQueryTransferOutInfoRequest creates a request to invoke QueryTransferOutInfo API
func CreateQueryTransferOutInfoRequest() (request *QueryTransferOutInfoRequest) {
	request = &QueryTransferOutInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryTransferOutInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTransferOutInfoResponse creates a response to parse from QueryTransferOutInfo response
func CreateQueryTransferOutInfoResponse() (response *QueryTransferOutInfoResponse) {
	response = &QueryTransferOutInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
