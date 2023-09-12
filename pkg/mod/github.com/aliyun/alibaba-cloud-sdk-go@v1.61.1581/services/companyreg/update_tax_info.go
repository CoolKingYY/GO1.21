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

// UpdateTaxInfo invokes the companyreg.UpdateTaxInfo API synchronously
func (client *Client) UpdateTaxInfo(request *UpdateTaxInfoRequest) (response *UpdateTaxInfoResponse, err error) {
	response = CreateUpdateTaxInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaxInfoWithChan invokes the companyreg.UpdateTaxInfo API asynchronously
func (client *Client) UpdateTaxInfoWithChan(request *UpdateTaxInfoRequest) (<-chan *UpdateTaxInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateTaxInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaxInfo(request)
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

// UpdateTaxInfoWithCallback invokes the companyreg.UpdateTaxInfo API asynchronously
func (client *Client) UpdateTaxInfoWithCallback(request *UpdateTaxInfoRequest, callback func(response *UpdateTaxInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaxInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaxInfo(request)
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

// UpdateTaxInfoRequest is the request struct for api UpdateTaxInfo
type UpdateTaxInfoRequest struct {
	*requests.RpcRequest
	UpdatedTaxInfo string `position:"Query" name:"UpdatedTaxInfo"`
	BizId          string `position:"Query" name:"BizId"`
}

// UpdateTaxInfoResponse is the response struct for api UpdateTaxInfo
type UpdateTaxInfoResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateTaxInfoRequest creates a request to invoke UpdateTaxInfo API
func CreateUpdateTaxInfoRequest() (request *UpdateTaxInfoRequest) {
	request = &UpdateTaxInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "UpdateTaxInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTaxInfoResponse creates a response to parse from UpdateTaxInfo response
func CreateUpdateTaxInfoResponse() (response *UpdateTaxInfoResponse) {
	response = &UpdateTaxInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}