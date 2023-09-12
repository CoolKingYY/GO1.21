package imageenhan

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

// IntelligentComposition invokes the imageenhan.IntelligentComposition API synchronously
func (client *Client) IntelligentComposition(request *IntelligentCompositionRequest) (response *IntelligentCompositionResponse, err error) {
	response = CreateIntelligentCompositionResponse()
	err = client.DoAction(request, response)
	return
}

// IntelligentCompositionWithChan invokes the imageenhan.IntelligentComposition API asynchronously
func (client *Client) IntelligentCompositionWithChan(request *IntelligentCompositionRequest) (<-chan *IntelligentCompositionResponse, <-chan error) {
	responseChan := make(chan *IntelligentCompositionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.IntelligentComposition(request)
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

// IntelligentCompositionWithCallback invokes the imageenhan.IntelligentComposition API asynchronously
func (client *Client) IntelligentCompositionWithCallback(request *IntelligentCompositionRequest, callback func(response *IntelligentCompositionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *IntelligentCompositionResponse
		var err error
		defer close(result)
		response, err = client.IntelligentComposition(request)
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

// IntelligentCompositionRequest is the request struct for api IntelligentComposition
type IntelligentCompositionRequest struct {
	*requests.RpcRequest
	NumBoxes requests.Integer `position:"Body" name:"NumBoxes"`
	ImageURL string           `position:"Body" name:"ImageURL"`
}

// IntelligentCompositionResponse is the response struct for api IntelligentComposition
type IntelligentCompositionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateIntelligentCompositionRequest creates a request to invoke IntelligentComposition API
func CreateIntelligentCompositionRequest() (request *IntelligentCompositionRequest) {
	request = &IntelligentCompositionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "IntelligentComposition", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateIntelligentCompositionResponse creates a response to parse from IntelligentComposition response
func CreateIntelligentCompositionResponse() (response *IntelligentCompositionResponse) {
	response = &IntelligentCompositionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
