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

// DeleteServiceMirror invokes the eas.DeleteServiceMirror API synchronously
func (client *Client) DeleteServiceMirror(request *DeleteServiceMirrorRequest) (response *DeleteServiceMirrorResponse, err error) {
	response = CreateDeleteServiceMirrorResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteServiceMirrorWithChan invokes the eas.DeleteServiceMirror API asynchronously
func (client *Client) DeleteServiceMirrorWithChan(request *DeleteServiceMirrorRequest) (<-chan *DeleteServiceMirrorResponse, <-chan error) {
	responseChan := make(chan *DeleteServiceMirrorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteServiceMirror(request)
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

// DeleteServiceMirrorWithCallback invokes the eas.DeleteServiceMirror API asynchronously
func (client *Client) DeleteServiceMirrorWithCallback(request *DeleteServiceMirrorRequest, callback func(response *DeleteServiceMirrorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteServiceMirrorResponse
		var err error
		defer close(result)
		response, err = client.DeleteServiceMirror(request)
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

// DeleteServiceMirrorRequest is the request struct for api DeleteServiceMirror
type DeleteServiceMirrorRequest struct {
	*requests.RoaRequest
	ServiceName string `position:"Path" name:"ServiceName"`
	ClusterId   string `position:"Path" name:"ClusterId"`
}

// DeleteServiceMirrorResponse is the response struct for api DeleteServiceMirror
type DeleteServiceMirrorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteServiceMirrorRequest creates a request to invoke DeleteServiceMirror API
func CreateDeleteServiceMirrorRequest() (request *DeleteServiceMirrorRequest) {
	request = &DeleteServiceMirrorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "DeleteServiceMirror", "/api/v2/services/[ClusterId]/[ServiceName]/mirror", "eas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteServiceMirrorResponse creates a response to parse from DeleteServiceMirror response
func CreateDeleteServiceMirrorResponse() (response *DeleteServiceMirrorResponse) {
	response = &DeleteServiceMirrorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
