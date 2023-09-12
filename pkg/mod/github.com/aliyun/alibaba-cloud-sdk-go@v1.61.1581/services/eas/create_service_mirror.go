package eas

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

// CreateServiceMirror invokes the eas.CreateServiceMirror API synchronously
func (client *Client) CreateServiceMirror(request *CreateServiceMirrorRequest) (response *CreateServiceMirrorResponse, err error) {
	response = CreateCreateServiceMirrorResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceMirrorWithChan invokes the eas.CreateServiceMirror API asynchronously
func (client *Client) CreateServiceMirrorWithChan(request *CreateServiceMirrorRequest) (<-chan *CreateServiceMirrorResponse, <-chan error) {
	responseChan := make(chan *CreateServiceMirrorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServiceMirror(request)
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

// CreateServiceMirrorWithCallback invokes the eas.CreateServiceMirror API asynchronously
func (client *Client) CreateServiceMirrorWithCallback(request *CreateServiceMirrorRequest, callback func(response *CreateServiceMirrorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceMirrorResponse
		var err error
		defer close(result)
		response, err = client.CreateServiceMirror(request)
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

// CreateServiceMirrorRequest is the request struct for api CreateServiceMirror
type CreateServiceMirrorRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
	Body        string `position:"Body" name:"body"`
}

// CreateServiceMirrorResponse is the response struct for api CreateServiceMirror
type CreateServiceMirrorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateServiceMirrorRequest creates a request to invoke CreateServiceMirror API
func CreateCreateServiceMirrorRequest() (request *CreateServiceMirrorRequest) {
	request = &CreateServiceMirrorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "CreateServiceMirror", "/api/v2/services/[ClusterId]/[ServiceName]/mirror", "eas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateServiceMirrorResponse creates a response to parse from CreateServiceMirror response
func CreateCreateServiceMirrorResponse() (response *CreateServiceMirrorResponse) {
	response = &CreateServiceMirrorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
