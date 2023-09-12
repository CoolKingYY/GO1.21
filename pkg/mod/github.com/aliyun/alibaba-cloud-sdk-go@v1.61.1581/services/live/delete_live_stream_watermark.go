package live

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

// DeleteLiveStreamWatermark invokes the live.DeleteLiveStreamWatermark API synchronously
func (client *Client) DeleteLiveStreamWatermark(request *DeleteLiveStreamWatermarkRequest) (response *DeleteLiveStreamWatermarkResponse, err error) {
	response = CreateDeleteLiveStreamWatermarkResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveStreamWatermarkWithChan invokes the live.DeleteLiveStreamWatermark API asynchronously
func (client *Client) DeleteLiveStreamWatermarkWithChan(request *DeleteLiveStreamWatermarkRequest) (<-chan *DeleteLiveStreamWatermarkResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveStreamWatermarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveStreamWatermark(request)
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

// DeleteLiveStreamWatermarkWithCallback invokes the live.DeleteLiveStreamWatermark API asynchronously
func (client *Client) DeleteLiveStreamWatermarkWithCallback(request *DeleteLiveStreamWatermarkRequest, callback func(response *DeleteLiveStreamWatermarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveStreamWatermarkResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveStreamWatermark(request)
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

// DeleteLiveStreamWatermarkRequest is the request struct for api DeleteLiveStreamWatermark
type DeleteLiveStreamWatermarkRequest struct {
	*requests.RpcRequest
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId string           `position:"Query" name:"TemplateId"`
}

// DeleteLiveStreamWatermarkResponse is the response struct for api DeleteLiveStreamWatermark
type DeleteLiveStreamWatermarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveStreamWatermarkRequest creates a request to invoke DeleteLiveStreamWatermark API
func CreateDeleteLiveStreamWatermarkRequest() (request *DeleteLiveStreamWatermarkRequest) {
	request = &DeleteLiveStreamWatermarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamWatermark", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveStreamWatermarkResponse creates a response to parse from DeleteLiveStreamWatermark response
func CreateDeleteLiveStreamWatermarkResponse() (response *DeleteLiveStreamWatermarkResponse) {
	response = &DeleteLiveStreamWatermarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
