package cloudcallcenter

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

// ListAvailableTestPhoneNumbers invokes the cloudcallcenter.ListAvailableTestPhoneNumbers API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listavailabletestphonenumbers.html
func (client *Client) ListAvailableTestPhoneNumbers(request *ListAvailableTestPhoneNumbersRequest) (response *ListAvailableTestPhoneNumbersResponse, err error) {
	response = CreateListAvailableTestPhoneNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// ListAvailableTestPhoneNumbersWithChan invokes the cloudcallcenter.ListAvailableTestPhoneNumbers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listavailabletestphonenumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAvailableTestPhoneNumbersWithChan(request *ListAvailableTestPhoneNumbersRequest) (<-chan *ListAvailableTestPhoneNumbersResponse, <-chan error) {
	responseChan := make(chan *ListAvailableTestPhoneNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAvailableTestPhoneNumbers(request)
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

// ListAvailableTestPhoneNumbersWithCallback invokes the cloudcallcenter.ListAvailableTestPhoneNumbers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listavailabletestphonenumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAvailableTestPhoneNumbersWithCallback(request *ListAvailableTestPhoneNumbersRequest, callback func(response *ListAvailableTestPhoneNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAvailableTestPhoneNumbersResponse
		var err error
		defer close(result)
		response, err = client.ListAvailableTestPhoneNumbers(request)
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

// ListAvailableTestPhoneNumbersRequest is the request struct for api ListAvailableTestPhoneNumbers
type ListAvailableTestPhoneNumbersRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListAvailableTestPhoneNumbersResponse is the response struct for api ListAvailableTestPhoneNumbers
type ListAvailableTestPhoneNumbersResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbers `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateListAvailableTestPhoneNumbersRequest creates a request to invoke ListAvailableTestPhoneNumbers API
func CreateListAvailableTestPhoneNumbersRequest() (request *ListAvailableTestPhoneNumbersRequest) {
	request = &ListAvailableTestPhoneNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListAvailableTestPhoneNumbers", "", "")
	request.Method = requests.POST
	return
}

// CreateListAvailableTestPhoneNumbersResponse creates a response to parse from ListAvailableTestPhoneNumbers response
func CreateListAvailableTestPhoneNumbersResponse() (response *ListAvailableTestPhoneNumbersResponse) {
	response = &ListAvailableTestPhoneNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
