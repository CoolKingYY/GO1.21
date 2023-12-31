package live

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

// DescribeLiveEdgeTransfer invokes the live.DescribeLiveEdgeTransfer API synchronously
func (client *Client) DescribeLiveEdgeTransfer(request *DescribeLiveEdgeTransferRequest) (response *DescribeLiveEdgeTransferResponse, err error) {
	response = CreateDescribeLiveEdgeTransferResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveEdgeTransferWithChan invokes the live.DescribeLiveEdgeTransfer API asynchronously
func (client *Client) DescribeLiveEdgeTransferWithChan(request *DescribeLiveEdgeTransferRequest) (<-chan *DescribeLiveEdgeTransferResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveEdgeTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveEdgeTransfer(request)
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

// DescribeLiveEdgeTransferWithCallback invokes the live.DescribeLiveEdgeTransfer API asynchronously
func (client *Client) DescribeLiveEdgeTransferWithCallback(request *DescribeLiveEdgeTransferRequest, callback func(response *DescribeLiveEdgeTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveEdgeTransferResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveEdgeTransfer(request)
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

// DescribeLiveEdgeTransferRequest is the request struct for api DescribeLiveEdgeTransfer
type DescribeLiveEdgeTransferRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveEdgeTransferResponse is the response struct for api DescribeLiveEdgeTransfer
type DescribeLiveEdgeTransferResponse struct {
	*responses.BaseResponse
	HttpDns          string `json:"HttpDns" xml:"HttpDns"`
	AppName          string `json:"AppName" xml:"AppName"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TransferArgs     string `json:"TransferArgs" xml:"TransferArgs"`
	StreamName       string `json:"StreamName" xml:"StreamName"`
	TargetDomainList string `json:"TargetDomainList" xml:"TargetDomainList"`
	DomainName       string `json:"DomainName" xml:"DomainName"`
}

// CreateDescribeLiveEdgeTransferRequest creates a request to invoke DescribeLiveEdgeTransfer API
func CreateDescribeLiveEdgeTransferRequest() (request *DescribeLiveEdgeTransferRequest) {
	request = &DescribeLiveEdgeTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveEdgeTransfer", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveEdgeTransferResponse creates a response to parse from DescribeLiveEdgeTransfer response
func CreateDescribeLiveEdgeTransferResponse() (response *DescribeLiveEdgeTransferResponse) {
	response = &DescribeLiveEdgeTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
