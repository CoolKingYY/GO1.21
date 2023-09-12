package ecd

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

// ModifyOfficeSiteAttribute invokes the ecd.ModifyOfficeSiteAttribute API synchronously
func (client *Client) ModifyOfficeSiteAttribute(request *ModifyOfficeSiteAttributeRequest) (response *ModifyOfficeSiteAttributeResponse, err error) {
	response = CreateModifyOfficeSiteAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOfficeSiteAttributeWithChan invokes the ecd.ModifyOfficeSiteAttribute API asynchronously
func (client *Client) ModifyOfficeSiteAttributeWithChan(request *ModifyOfficeSiteAttributeRequest) (<-chan *ModifyOfficeSiteAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyOfficeSiteAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOfficeSiteAttribute(request)
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

// ModifyOfficeSiteAttributeWithCallback invokes the ecd.ModifyOfficeSiteAttribute API asynchronously
func (client *Client) ModifyOfficeSiteAttributeWithCallback(request *ModifyOfficeSiteAttributeRequest, callback func(response *ModifyOfficeSiteAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOfficeSiteAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyOfficeSiteAttribute(request)
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

// ModifyOfficeSiteAttributeRequest is the request struct for api ModifyOfficeSiteAttribute
type ModifyOfficeSiteAttributeRequest struct {
	*requests.RpcRequest
	OfficeSiteId         string           `position:"Query" name:"OfficeSiteId"`
	DesktopAccessType    string           `position:"Query" name:"DesktopAccessType"`
	OfficeSiteName       string           `position:"Query" name:"OfficeSiteName"`
	NeedVerifyLoginRisk  requests.Boolean `position:"Query" name:"NeedVerifyLoginRisk"`
	NeedVerifyZeroDevice requests.Boolean `position:"Query" name:"NeedVerifyZeroDevice"`
}

// ModifyOfficeSiteAttributeResponse is the response struct for api ModifyOfficeSiteAttribute
type ModifyOfficeSiteAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyOfficeSiteAttributeRequest creates a request to invoke ModifyOfficeSiteAttribute API
func CreateModifyOfficeSiteAttributeRequest() (request *ModifyOfficeSiteAttributeRequest) {
	request = &ModifyOfficeSiteAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyOfficeSiteAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyOfficeSiteAttributeResponse creates a response to parse from ModifyOfficeSiteAttribute response
func CreateModifyOfficeSiteAttributeResponse() (response *ModifyOfficeSiteAttributeResponse) {
	response = &ModifyOfficeSiteAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}