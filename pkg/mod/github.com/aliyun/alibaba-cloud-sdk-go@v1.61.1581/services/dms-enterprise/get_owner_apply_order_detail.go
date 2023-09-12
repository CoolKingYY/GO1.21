package dms_enterprise

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

// GetOwnerApplyOrderDetail invokes the dms_enterprise.GetOwnerApplyOrderDetail API synchronously
func (client *Client) GetOwnerApplyOrderDetail(request *GetOwnerApplyOrderDetailRequest) (response *GetOwnerApplyOrderDetailResponse, err error) {
	response = CreateGetOwnerApplyOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetOwnerApplyOrderDetailWithChan invokes the dms_enterprise.GetOwnerApplyOrderDetail API asynchronously
func (client *Client) GetOwnerApplyOrderDetailWithChan(request *GetOwnerApplyOrderDetailRequest) (<-chan *GetOwnerApplyOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetOwnerApplyOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOwnerApplyOrderDetail(request)
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

// GetOwnerApplyOrderDetailWithCallback invokes the dms_enterprise.GetOwnerApplyOrderDetail API asynchronously
func (client *Client) GetOwnerApplyOrderDetailWithCallback(request *GetOwnerApplyOrderDetailRequest, callback func(response *GetOwnerApplyOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOwnerApplyOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetOwnerApplyOrderDetail(request)
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

// GetOwnerApplyOrderDetailRequest is the request struct for api GetOwnerApplyOrderDetail
type GetOwnerApplyOrderDetailRequest struct {
	*requests.RpcRequest
	OrderId requests.Integer `position:"Query" name:"OrderId"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
}

// GetOwnerApplyOrderDetailResponse is the response struct for api GetOwnerApplyOrderDetail
type GetOwnerApplyOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               bool                  `json:"Success" xml:"Success"`
	ErrorMessage          string                `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode             string                `json:"ErrorCode" xml:"ErrorCode"`
	OwnerApplyOrderDetail OwnerApplyOrderDetail `json:"OwnerApplyOrderDetail" xml:"OwnerApplyOrderDetail"`
}

// CreateGetOwnerApplyOrderDetailRequest creates a request to invoke GetOwnerApplyOrderDetail API
func CreateGetOwnerApplyOrderDetailRequest() (request *GetOwnerApplyOrderDetailRequest) {
	request = &GetOwnerApplyOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetOwnerApplyOrderDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOwnerApplyOrderDetailResponse creates a response to parse from GetOwnerApplyOrderDetail response
func CreateGetOwnerApplyOrderDetailResponse() (response *GetOwnerApplyOrderDetailResponse) {
	response = &GetOwnerApplyOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
