package dts

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

// ModifyDtsJobName invokes the dts.ModifyDtsJobName API synchronously
func (client *Client) ModifyDtsJobName(request *ModifyDtsJobNameRequest) (response *ModifyDtsJobNameResponse, err error) {
	response = CreateModifyDtsJobNameResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDtsJobNameWithChan invokes the dts.ModifyDtsJobName API asynchronously
func (client *Client) ModifyDtsJobNameWithChan(request *ModifyDtsJobNameRequest) (<-chan *ModifyDtsJobNameResponse, <-chan error) {
	responseChan := make(chan *ModifyDtsJobNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDtsJobName(request)
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

// ModifyDtsJobNameWithCallback invokes the dts.ModifyDtsJobName API asynchronously
func (client *Client) ModifyDtsJobNameWithCallback(request *ModifyDtsJobNameRequest, callback func(response *ModifyDtsJobNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDtsJobNameResponse
		var err error
		defer close(result)
		response, err = client.ModifyDtsJobName(request)
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

// ModifyDtsJobNameRequest is the request struct for api ModifyDtsJobName
type ModifyDtsJobNameRequest struct {
	*requests.RpcRequest
	DtsJobName string `position:"Query" name:"DtsJobName"`
	DtsJobId   string `position:"Query" name:"DtsJobId"`
}

// ModifyDtsJobNameResponse is the response struct for api ModifyDtsJobName
type ModifyDtsJobNameResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateModifyDtsJobNameRequest creates a request to invoke ModifyDtsJobName API
func CreateModifyDtsJobNameRequest() (request *ModifyDtsJobNameRequest) {
	request = &ModifyDtsJobNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ModifyDtsJobName", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDtsJobNameResponse creates a response to parse from ModifyDtsJobName response
func CreateModifyDtsJobNameResponse() (response *ModifyDtsJobNameResponse) {
	response = &ModifyDtsJobNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}