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

// ResetUserStatus invokes the cloudcallcenter.ResetUserStatus API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resetuserstatus.html
func (client *Client) ResetUserStatus(request *ResetUserStatusRequest) (response *ResetUserStatusResponse, err error) {
	response = CreateResetUserStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ResetUserStatusWithChan invokes the cloudcallcenter.ResetUserStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resetuserstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetUserStatusWithChan(request *ResetUserStatusRequest) (<-chan *ResetUserStatusResponse, <-chan error) {
	responseChan := make(chan *ResetUserStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetUserStatus(request)
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

// ResetUserStatusWithCallback invokes the cloudcallcenter.ResetUserStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resetuserstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetUserStatusWithCallback(request *ResetUserStatusRequest, callback func(response *ResetUserStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetUserStatusResponse
		var err error
		defer close(result)
		response, err = client.ResetUserStatus(request)
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

// ResetUserStatusRequest is the request struct for api ResetUserStatus
type ResetUserStatusRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	RamIdList  *[]string `position:"Query" name:"RamIdList"  type:"Repeated"`
}

// ResetUserStatusResponse is the response struct for api ResetUserStatus
type ResetUserStatusResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateResetUserStatusRequest creates a request to invoke ResetUserStatus API
func CreateResetUserStatusRequest() (request *ResetUserStatusRequest) {
	request = &ResetUserStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ResetUserStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateResetUserStatusResponse creates a response to parse from ResetUserStatus response
func CreateResetUserStatusResponse() (response *ResetUserStatusResponse) {
	response = &ResetUserStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
