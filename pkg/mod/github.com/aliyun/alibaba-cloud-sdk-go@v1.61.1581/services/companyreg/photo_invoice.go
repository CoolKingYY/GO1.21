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

// PhotoInvoice invokes the companyreg.PhotoInvoice API synchronously
func (client *Client) PhotoInvoice(request *PhotoInvoiceRequest) (response *PhotoInvoiceResponse, err error) {
	response = CreatePhotoInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// PhotoInvoiceWithChan invokes the companyreg.PhotoInvoice API asynchronously
func (client *Client) PhotoInvoiceWithChan(request *PhotoInvoiceRequest) (<-chan *PhotoInvoiceResponse, <-chan error) {
	responseChan := make(chan *PhotoInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PhotoInvoice(request)
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

// PhotoInvoiceWithCallback invokes the companyreg.PhotoInvoice API asynchronously
func (client *Client) PhotoInvoiceWithCallback(request *PhotoInvoiceRequest, callback func(response *PhotoInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PhotoInvoiceResponse
		var err error
		defer close(result)
		response, err = client.PhotoInvoice(request)
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

// PhotoInvoiceRequest is the request struct for api PhotoInvoice
type PhotoInvoiceRequest struct {
	*requests.RpcRequest
	UploadedNum   requests.Integer `position:"Query" name:"UploadedNum"`
	FileUrlList   string           `position:"Query" name:"FileUrlList"`
	UploadedStamp requests.Integer `position:"Query" name:"UploadedStamp"`
	BizId         string           `position:"Query" name:"BizId"`
	FileUrl       string           `position:"Query" name:"FileUrl"`
	IsMobile      requests.Boolean `position:"Query" name:"IsMobile"`
}

// PhotoInvoiceResponse is the response struct for api PhotoInvoice
type PhotoInvoiceResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Item      []ItemItem `json:"Item" xml:"Item"`
}

// CreatePhotoInvoiceRequest creates a request to invoke PhotoInvoice API
func CreatePhotoInvoiceRequest() (request *PhotoInvoiceRequest) {
	request = &PhotoInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "PhotoInvoice", "", "")
	request.Method = requests.GET
	return
}

// CreatePhotoInvoiceResponse creates a response to parse from PhotoInvoice response
func CreatePhotoInvoiceResponse() (response *PhotoInvoiceResponse) {
	response = &PhotoInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
