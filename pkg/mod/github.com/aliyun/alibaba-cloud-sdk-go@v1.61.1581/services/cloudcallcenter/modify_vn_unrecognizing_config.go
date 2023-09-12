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

// ModifyVnUnrecognizingConfig invokes the cloudcallcenter.ModifyVnUnrecognizingConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnunrecognizingconfig.html
func (client *Client) ModifyVnUnrecognizingConfig(request *ModifyVnUnrecognizingConfigRequest) (response *ModifyVnUnrecognizingConfigResponse, err error) {
	response = CreateModifyVnUnrecognizingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnUnrecognizingConfigWithChan invokes the cloudcallcenter.ModifyVnUnrecognizingConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnunrecognizingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnUnrecognizingConfigWithChan(request *ModifyVnUnrecognizingConfigRequest) (<-chan *ModifyVnUnrecognizingConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyVnUnrecognizingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnUnrecognizingConfig(request)
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

// ModifyVnUnrecognizingConfigWithCallback invokes the cloudcallcenter.ModifyVnUnrecognizingConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnunrecognizingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnUnrecognizingConfigWithCallback(request *ModifyVnUnrecognizingConfigRequest, callback func(response *ModifyVnUnrecognizingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnUnrecognizingConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnUnrecognizingConfig(request)
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

// ModifyVnUnrecognizingConfigRequest is the request struct for api ModifyVnUnrecognizingConfig
type ModifyVnUnrecognizingConfigRequest struct {
	*requests.RpcRequest
	FinalAction       string           `position:"Query" name:"FinalAction"`
	FinalPrompt       string           `position:"Query" name:"FinalPrompt"`
	Threshold         requests.Integer `position:"Query" name:"Threshold"`
	InstanceId        string           `position:"Query" name:"InstanceId"`
	FinalActionParams string           `position:"Query" name:"FinalActionParams"`
	Prompt            string           `position:"Query" name:"Prompt"`
}

// ModifyVnUnrecognizingConfigResponse is the response struct for api ModifyVnUnrecognizingConfig
type ModifyVnUnrecognizingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnUnrecognizingConfigRequest creates a request to invoke ModifyVnUnrecognizingConfig API
func CreateModifyVnUnrecognizingConfigRequest() (request *ModifyVnUnrecognizingConfigRequest) {
	request = &ModifyVnUnrecognizingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnUnrecognizingConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnUnrecognizingConfigResponse creates a response to parse from ModifyVnUnrecognizingConfig response
func CreateModifyVnUnrecognizingConfigResponse() (response *ModifyVnUnrecognizingConfigResponse) {
	response = &ModifyVnUnrecognizingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
