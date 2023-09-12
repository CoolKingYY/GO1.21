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

// EnablePhoneNumber invokes the cloudcallcenter.EnablePhoneNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablephonenumber.html
func (client *Client) EnablePhoneNumber(request *EnablePhoneNumberRequest) (response *EnablePhoneNumberResponse, err error) {
	response = CreateEnablePhoneNumberResponse()
	err = client.DoAction(request, response)
	return
}

// EnablePhoneNumberWithChan invokes the cloudcallcenter.EnablePhoneNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablephonenumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnablePhoneNumberWithChan(request *EnablePhoneNumberRequest) (<-chan *EnablePhoneNumberResponse, <-chan error) {
	responseChan := make(chan *EnablePhoneNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnablePhoneNumber(request)
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

// EnablePhoneNumberWithCallback invokes the cloudcallcenter.EnablePhoneNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablephonenumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnablePhoneNumberWithCallback(request *EnablePhoneNumberRequest, callback func(response *EnablePhoneNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnablePhoneNumberResponse
		var err error
		defer close(result)
		response, err = client.EnablePhoneNumber(request)
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

// EnablePhoneNumberRequest is the request struct for api EnablePhoneNumber
type EnablePhoneNumberRequest struct {
	*requests.RpcRequest
	CommodityInstanceId string `position:"Query" name:"CommodityInstanceId"`
}

// EnablePhoneNumberResponse is the response struct for api EnablePhoneNumber
type EnablePhoneNumberResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateEnablePhoneNumberRequest creates a request to invoke EnablePhoneNumber API
func CreateEnablePhoneNumberRequest() (request *EnablePhoneNumberRequest) {
	request = &EnablePhoneNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "EnablePhoneNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateEnablePhoneNumberResponse creates a response to parse from EnablePhoneNumber response
func CreateEnablePhoneNumberResponse() (response *EnablePhoneNumberResponse) {
	response = &EnablePhoneNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
