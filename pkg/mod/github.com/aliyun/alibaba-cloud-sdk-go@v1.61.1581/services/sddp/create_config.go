package sddp

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

// CreateConfig invokes the sddp.CreateConfig API synchronously
func (client *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
	response = CreateCreateConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateConfigWithChan invokes the sddp.CreateConfig API asynchronously
func (client *Client) CreateConfigWithChan(request *CreateConfigRequest) (<-chan *CreateConfigResponse, <-chan error) {
	responseChan := make(chan *CreateConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateConfig(request)
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

// CreateConfigWithCallback invokes the sddp.CreateConfig API asynchronously
func (client *Client) CreateConfigWithCallback(request *CreateConfigRequest, callback func(response *CreateConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateConfig(request)
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

// CreateConfigRequest is the request struct for api CreateConfig
type CreateConfigRequest struct {
	*requests.RpcRequest
	Code        string           `position:"Query" name:"Code"`
	Description string           `position:"Query" name:"Description"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Lang        string           `position:"Query" name:"Lang"`
	Value       string           `position:"Query" name:"Value"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	ConfigList  string           `position:"Query" name:"ConfigList"`
}

// CreateConfigResponse is the response struct for api CreateConfig
type CreateConfigResponse struct {
	*responses.BaseResponse
	Id        int64  `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateConfigRequest creates a request to invoke CreateConfig API
func CreateCreateConfigRequest() (request *CreateConfigRequest) {
	request = &CreateConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "CreateConfig", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateConfigResponse creates a response to parse from CreateConfig response
func CreateCreateConfigResponse() (response *CreateConfigResponse) {
	response = &CreateConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
