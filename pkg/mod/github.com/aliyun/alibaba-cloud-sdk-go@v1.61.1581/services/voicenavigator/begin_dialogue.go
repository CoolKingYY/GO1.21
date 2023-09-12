package voicenavigator

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

// BeginDialogue invokes the voicenavigator.BeginDialogue API synchronously
func (client *Client) BeginDialogue(request *BeginDialogueRequest) (response *BeginDialogueResponse, err error) {
	response = CreateBeginDialogueResponse()
	err = client.DoAction(request, response)
	return
}

// BeginDialogueWithChan invokes the voicenavigator.BeginDialogue API asynchronously
func (client *Client) BeginDialogueWithChan(request *BeginDialogueRequest) (<-chan *BeginDialogueResponse, <-chan error) {
	responseChan := make(chan *BeginDialogueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BeginDialogue(request)
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

// BeginDialogueWithCallback invokes the voicenavigator.BeginDialogue API asynchronously
func (client *Client) BeginDialogueWithCallback(request *BeginDialogueRequest, callback func(response *BeginDialogueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BeginDialogueResponse
		var err error
		defer close(result)
		response, err = client.BeginDialogue(request)
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

// BeginDialogueRequest is the request struct for api BeginDialogue
type BeginDialogueRequest struct {
	*requests.RpcRequest
	ConversationId  string           `position:"Query" name:"ConversationId"`
	InitialContext  string           `position:"Query" name:"InitialContext"`
	CallingNumber   string           `position:"Query" name:"CallingNumber"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	CalledNumber    string           `position:"Query" name:"CalledNumber"`
	InstanceOwnerId requests.Integer `position:"Query" name:"InstanceOwnerId"`
}

// BeginDialogueResponse is the response struct for api BeginDialogue
type BeginDialogueResponse struct {
	*responses.BaseResponse
	Action        string `json:"Action" xml:"Action"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
}

// CreateBeginDialogueRequest creates a request to invoke BeginDialogue API
func CreateBeginDialogueRequest() (request *BeginDialogueRequest) {
	request = &BeginDialogueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "BeginDialogue", "voicebot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBeginDialogueResponse creates a response to parse from BeginDialogue response
func CreateBeginDialogueResponse() (response *BeginDialogueResponse) {
	response = &BeginDialogueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
