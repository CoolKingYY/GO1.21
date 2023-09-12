package green

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

// AddVideoDna invokes the green.AddVideoDna API synchronously
func (client *Client) AddVideoDna(request *AddVideoDnaRequest) (response *AddVideoDnaResponse, err error) {
	response = CreateAddVideoDnaResponse()
	err = client.DoAction(request, response)
	return
}

// AddVideoDnaWithChan invokes the green.AddVideoDna API asynchronously
func (client *Client) AddVideoDnaWithChan(request *AddVideoDnaRequest) (<-chan *AddVideoDnaResponse, <-chan error) {
	responseChan := make(chan *AddVideoDnaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVideoDna(request)
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

// AddVideoDnaWithCallback invokes the green.AddVideoDna API asynchronously
func (client *Client) AddVideoDnaWithCallback(request *AddVideoDnaRequest, callback func(response *AddVideoDnaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVideoDnaResponse
		var err error
		defer close(result)
		response, err = client.AddVideoDna(request)
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

// AddVideoDnaRequest is the request struct for api AddVideoDna
type AddVideoDnaRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// AddVideoDnaResponse is the response struct for api AddVideoDna
type AddVideoDnaResponse struct {
	*responses.BaseResponse
}

// CreateAddVideoDnaRequest creates a request to invoke AddVideoDna API
func CreateAddVideoDnaRequest() (request *AddVideoDnaRequest) {
	request = &AddVideoDnaRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "AddVideoDna", "/green/video/dna/add", "", "")
	request.Method = requests.POST
	return
}

// CreateAddVideoDnaResponse creates a response to parse from AddVideoDna response
func CreateAddVideoDnaResponse() (response *AddVideoDnaResponse) {
	response = &AddVideoDnaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
