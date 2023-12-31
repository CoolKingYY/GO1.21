package outboundbot

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

// CreateDialogueFlow invokes the outboundbot.CreateDialogueFlow API synchronously
func (client *Client) CreateDialogueFlow(request *CreateDialogueFlowRequest) (response *CreateDialogueFlowResponse, err error) {
	response = CreateCreateDialogueFlowResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDialogueFlowWithChan invokes the outboundbot.CreateDialogueFlow API asynchronously
func (client *Client) CreateDialogueFlowWithChan(request *CreateDialogueFlowRequest) (<-chan *CreateDialogueFlowResponse, <-chan error) {
	responseChan := make(chan *CreateDialogueFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDialogueFlow(request)
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

// CreateDialogueFlowWithCallback invokes the outboundbot.CreateDialogueFlow API asynchronously
func (client *Client) CreateDialogueFlowWithCallback(request *CreateDialogueFlowRequest, callback func(response *CreateDialogueFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDialogueFlowResponse
		var err error
		defer close(result)
		response, err = client.CreateDialogueFlow(request)
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

// CreateDialogueFlowRequest is the request struct for api CreateDialogueFlow
type CreateDialogueFlowRequest struct {
	*requests.RpcRequest
	DialogueFlowType string `position:"Query" name:"DialogueFlowType"`
	DialogueName     string `position:"Query" name:"DialogueName"`
	ScriptId         string `position:"Query" name:"ScriptId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
}

// CreateDialogueFlowResponse is the response struct for api CreateDialogueFlow
type CreateDialogueFlowResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DialogueFlowId string `json:"DialogueFlowId" xml:"DialogueFlowId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateCreateDialogueFlowRequest creates a request to invoke CreateDialogueFlow API
func CreateCreateDialogueFlowRequest() (request *CreateDialogueFlowRequest) {
	request = &CreateDialogueFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateDialogueFlow", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDialogueFlowResponse creates a response to parse from CreateDialogueFlow response
func CreateCreateDialogueFlowResponse() (response *CreateDialogueFlowResponse) {
	response = &CreateDialogueFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
