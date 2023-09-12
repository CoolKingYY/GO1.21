package ecd

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

// DeleteImages invokes the ecd.DeleteImages API synchronously
func (client *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
	response = CreateDeleteImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImagesWithChan invokes the ecd.DeleteImages API asynchronously
func (client *Client) DeleteImagesWithChan(request *DeleteImagesRequest) (<-chan *DeleteImagesResponse, <-chan error) {
	responseChan := make(chan *DeleteImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImages(request)
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

// DeleteImagesWithCallback invokes the ecd.DeleteImages API asynchronously
func (client *Client) DeleteImagesWithCallback(request *DeleteImagesRequest, callback func(response *DeleteImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImagesResponse
		var err error
		defer close(result)
		response, err = client.DeleteImages(request)
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

// DeleteImagesRequest is the request struct for api DeleteImages
type DeleteImagesRequest struct {
	*requests.RpcRequest
	ImageId *[]string `position:"Query" name:"ImageId"  type:"Repeated"`
}

// DeleteImagesResponse is the response struct for api DeleteImages
type DeleteImagesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteImagesRequest creates a request to invoke DeleteImages API
func CreateDeleteImagesRequest() (request *DeleteImagesRequest) {
	request = &DeleteImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeleteImages", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteImagesResponse creates a response to parse from DeleteImages response
func CreateDeleteImagesResponse() (response *DeleteImagesResponse) {
	response = &DeleteImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}