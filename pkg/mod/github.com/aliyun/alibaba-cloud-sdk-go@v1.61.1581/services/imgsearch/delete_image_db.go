package imgsearch

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

// DeleteImageDb invokes the imgsearch.DeleteImageDb API synchronously
func (client *Client) DeleteImageDb(request *DeleteImageDbRequest) (response *DeleteImageDbResponse, err error) {
	response = CreateDeleteImageDbResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageDbWithChan invokes the imgsearch.DeleteImageDb API asynchronously
func (client *Client) DeleteImageDbWithChan(request *DeleteImageDbRequest) (<-chan *DeleteImageDbResponse, <-chan error) {
	responseChan := make(chan *DeleteImageDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImageDb(request)
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

// DeleteImageDbWithCallback invokes the imgsearch.DeleteImageDb API asynchronously
func (client *Client) DeleteImageDbWithCallback(request *DeleteImageDbRequest, callback func(response *DeleteImageDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageDbResponse
		var err error
		defer close(result)
		response, err = client.DeleteImageDb(request)
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

// DeleteImageDbRequest is the request struct for api DeleteImageDb
type DeleteImageDbRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
}

// DeleteImageDbResponse is the response struct for api DeleteImageDb
type DeleteImageDbResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteImageDbRequest creates a request to invoke DeleteImageDb API
func CreateDeleteImageDbRequest() (request *DeleteImageDbRequest) {
	request = &DeleteImageDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imgsearch", "2020-03-20", "DeleteImageDb", "imgsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteImageDbResponse creates a response to parse from DeleteImageDb response
func CreateDeleteImageDbResponse() (response *DeleteImageDbResponse) {
	response = &DeleteImageDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
