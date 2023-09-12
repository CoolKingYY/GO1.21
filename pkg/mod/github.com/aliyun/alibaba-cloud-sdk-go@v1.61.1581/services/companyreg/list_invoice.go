package companyreg

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

// ListInvoice invokes the companyreg.ListInvoice API synchronously
func (client *Client) ListInvoice(request *ListInvoiceRequest) (response *ListInvoiceResponse, err error) {
	response = CreateListInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// ListInvoiceWithChan invokes the companyreg.ListInvoice API asynchronously
func (client *Client) ListInvoiceWithChan(request *ListInvoiceRequest) (<-chan *ListInvoiceResponse, <-chan error) {
	responseChan := make(chan *ListInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInvoice(request)
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

// ListInvoiceWithCallback invokes the companyreg.ListInvoice API asynchronously
func (client *Client) ListInvoiceWithCallback(request *ListInvoiceRequest, callback func(response *ListInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInvoiceResponse
		var err error
		defer close(result)
		response, err = client.ListInvoice(request)
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

// ListInvoiceRequest is the request struct for api ListInvoice
type ListInvoiceRequest struct {
	*requests.RpcRequest
	BizId    string           `position:"Query" name:"BizId"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Periods  string           `position:"Query" name:"Periods"`
	Page     requests.Integer `position:"Query" name:"Page"`
	Key      string           `position:"Query" name:"Key"`
}

// ListInvoiceResponse is the response struct for api ListInvoice
type ListInvoiceResponse struct {
	*responses.BaseResponse
	Amount    string     `json:"Amount" xml:"Amount"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Count     int        `json:"Count" xml:"Count"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListInvoiceRequest creates a request to invoke ListInvoice API
func CreateListInvoiceRequest() (request *ListInvoiceRequest) {
	request = &ListInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "ListInvoice", "", "")
	request.Method = requests.GET
	return
}

// CreateListInvoiceResponse creates a response to parse from ListInvoice response
func CreateListInvoiceResponse() (response *ListInvoiceResponse) {
	response = &ListInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}