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

// GetTaxationCalendar invokes the companyreg.GetTaxationCalendar API synchronously
func (client *Client) GetTaxationCalendar(request *GetTaxationCalendarRequest) (response *GetTaxationCalendarResponse, err error) {
	response = CreateGetTaxationCalendarResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaxationCalendarWithChan invokes the companyreg.GetTaxationCalendar API asynchronously
func (client *Client) GetTaxationCalendarWithChan(request *GetTaxationCalendarRequest) (<-chan *GetTaxationCalendarResponse, <-chan error) {
	responseChan := make(chan *GetTaxationCalendarResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaxationCalendar(request)
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

// GetTaxationCalendarWithCallback invokes the companyreg.GetTaxationCalendar API asynchronously
func (client *Client) GetTaxationCalendarWithCallback(request *GetTaxationCalendarRequest, callback func(response *GetTaxationCalendarResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaxationCalendarResponse
		var err error
		defer close(result)
		response, err = client.GetTaxationCalendar(request)
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

// GetTaxationCalendarRequest is the request struct for api GetTaxationCalendar
type GetTaxationCalendarRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// GetTaxationCalendarResponse is the response struct for api GetTaxationCalendar
type GetTaxationCalendarResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	CalendarList []CalendarListItem `json:"CalendarList" xml:"CalendarList"`
}

// CreateGetTaxationCalendarRequest creates a request to invoke GetTaxationCalendar API
func CreateGetTaxationCalendarRequest() (request *GetTaxationCalendarRequest) {
	request = &GetTaxationCalendarRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "GetTaxationCalendar", "", "")
	request.Method = requests.GET
	return
}

// CreateGetTaxationCalendarResponse creates a response to parse from GetTaxationCalendar response
func CreateGetTaxationCalendarResponse() (response *GetTaxationCalendarResponse) {
	response = &GetTaxationCalendarResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
