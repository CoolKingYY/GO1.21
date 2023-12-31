package live

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

// ModifyCasterComponent invokes the live.ModifyCasterComponent API synchronously
func (client *Client) ModifyCasterComponent(request *ModifyCasterComponentRequest) (response *ModifyCasterComponentResponse, err error) {
	response = CreateModifyCasterComponentResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCasterComponentWithChan invokes the live.ModifyCasterComponent API asynchronously
func (client *Client) ModifyCasterComponentWithChan(request *ModifyCasterComponentRequest) (<-chan *ModifyCasterComponentResponse, <-chan error) {
	responseChan := make(chan *ModifyCasterComponentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCasterComponent(request)
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

// ModifyCasterComponentWithCallback invokes the live.ModifyCasterComponent API asynchronously
func (client *Client) ModifyCasterComponentWithCallback(request *ModifyCasterComponentRequest, callback func(response *ModifyCasterComponentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCasterComponentResponse
		var err error
		defer close(result)
		response, err = client.ModifyCasterComponent(request)
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

// ModifyCasterComponentRequest is the request struct for api ModifyCasterComponent
type ModifyCasterComponentRequest struct {
	*requests.RpcRequest
	ImageLayerContent   string           `position:"Query" name:"ImageLayerContent"`
	ComponentName       string           `position:"Query" name:"ComponentName"`
	ComponentId         string           `position:"Query" name:"ComponentId"`
	CasterId            string           `position:"Query" name:"CasterId"`
	ComponentLayer      string           `position:"Query" name:"ComponentLayer"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	ComponentType       string           `position:"Query" name:"ComponentType"`
	Effect              string           `position:"Query" name:"Effect"`
	CaptionLayerContent string           `position:"Query" name:"CaptionLayerContent"`
	TextLayerContent    string           `position:"Query" name:"TextLayerContent"`
}

// ModifyCasterComponentResponse is the response struct for api ModifyCasterComponent
type ModifyCasterComponentResponse struct {
	*responses.BaseResponse
	ComponentId string `json:"ComponentId" xml:"ComponentId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCasterComponentRequest creates a request to invoke ModifyCasterComponent API
func CreateModifyCasterComponentRequest() (request *ModifyCasterComponentRequest) {
	request = &ModifyCasterComponentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyCasterComponent", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCasterComponentResponse creates a response to parse from ModifyCasterComponent response
func CreateModifyCasterComponentResponse() (response *ModifyCasterComponentResponse) {
	response = &ModifyCasterComponentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
