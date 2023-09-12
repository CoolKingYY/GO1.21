package teambition_aliyun

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

// GetUserByUid invokes the teambition_aliyun.GetUserByUid API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getuserbyuid.html
func (client *Client) GetUserByUid(request *GetUserByUidRequest) (response *GetUserByUidResponse, err error) {
	response = CreateGetUserByUidResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserByUidWithChan invokes the teambition_aliyun.GetUserByUid API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getuserbyuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByUidWithChan(request *GetUserByUidRequest) (<-chan *GetUserByUidResponse, <-chan error) {
	responseChan := make(chan *GetUserByUidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserByUid(request)
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

// GetUserByUidWithCallback invokes the teambition_aliyun.GetUserByUid API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getuserbyuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByUidWithCallback(request *GetUserByUidRequest, callback func(response *GetUserByUidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserByUidResponse
		var err error
		defer close(result)
		response, err = client.GetUserByUid(request)
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

// GetUserByUidRequest is the request struct for api GetUserByUid
type GetUserByUidRequest struct {
	*requests.RpcRequest
	UserPk string `position:"Body" name:"UserPk"`
	OrgId  string `position:"Body" name:"OrgId"`
}

// GetUserByUidResponse is the response struct for api GetUserByUid
type GetUserByUidResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateGetUserByUidRequest creates a request to invoke GetUserByUid API
func CreateGetUserByUidRequest() (request *GetUserByUidRequest) {
	request = &GetUserByUidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "GetUserByUid", "", "")
	request.Method = requests.POST
	return
}

// CreateGetUserByUidResponse creates a response to parse from GetUserByUid response
func CreateGetUserByUidResponse() (response *GetUserByUidResponse) {
	response = &GetUserByUidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
