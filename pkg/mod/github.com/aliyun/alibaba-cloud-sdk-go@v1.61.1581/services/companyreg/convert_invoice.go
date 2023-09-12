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

// ConvertInvoice invokes the companyreg.ConvertInvoice API synchronously
func (client *Client) ConvertInvoice(request *ConvertInvoiceRequest) (response *ConvertInvoiceResponse, err error) {
	response = CreateConvertInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// ConvertInvoiceWithChan invokes the companyreg.ConvertInvoice API asynchronously
func (client *Client) ConvertInvoiceWithChan(request *ConvertInvoiceRequest) (<-chan *ConvertInvoiceResponse, <-chan error) {
	responseChan := make(chan *ConvertInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConvertInvoice(request)
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

// ConvertInvoiceWithCallback invokes the companyreg.ConvertInvoice API asynchronously
func (client *Client) ConvertInvoiceWithCallback(request *ConvertInvoiceRequest, callback func(response *ConvertInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConvertInvoiceResponse
		var err error
		defer close(result)
		response, err = client.ConvertInvoice(request)
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

// ConvertInvoiceRequest is the request struct for api ConvertInvoice
type ConvertInvoiceRequest struct {
	*requests.RpcRequest
	IsFee          requests.Boolean `position:"Query" name:"IsFee"`
	ShellMethod    string           `position:"Query" name:"ShellMethod"`
	Kind           string           `position:"Query" name:"Kind"`
	Use            string           `position:"Query" name:"Use"`
	ThirdKey       string           `position:"Query" name:"ThirdKey"`
	Payer          string           `position:"Query" name:"Payer"`
	SecondKey      string           `position:"Query" name:"SecondKey"`
	PayMethod      string           `position:"Query" name:"PayMethod"`
	BuyMethod      string           `position:"Query" name:"BuyMethod"`
	FixedAssetType string           `position:"Query" name:"FixedAssetType"`
	FirstKey       string           `position:"Query" name:"FirstKey"`
	BizId          string           `position:"Query" name:"BizId"`
	Id             requests.Integer `position:"Query" name:"Id"`
	BuyTarget      string           `position:"Query" name:"BuyTarget"`
}

// ConvertInvoiceResponse is the response struct for api ConvertInvoice
type ConvertInvoiceResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConvertInvoiceRequest creates a request to invoke ConvertInvoice API
func CreateConvertInvoiceRequest() (request *ConvertInvoiceRequest) {
	request = &ConvertInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "ConvertInvoice", "", "")
	request.Method = requests.POST
	return
}

// CreateConvertInvoiceResponse creates a response to parse from ConvertInvoice response
func CreateConvertInvoiceResponse() (response *ConvertInvoiceResponse) {
	response = &ConvertInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}