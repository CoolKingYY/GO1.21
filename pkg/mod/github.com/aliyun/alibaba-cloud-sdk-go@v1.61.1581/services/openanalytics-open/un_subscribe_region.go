package openanalytics_open

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

// UnSubscribeRegion invokes the openanalytics_open.UnSubscribeRegion API synchronously
func (client *Client) UnSubscribeRegion(request *UnSubscribeRegionRequest) (response *UnSubscribeRegionResponse, err error) {
	response = CreateUnSubscribeRegionResponse()
	err = client.DoAction(request, response)
	return
}

// UnSubscribeRegionWithChan invokes the openanalytics_open.UnSubscribeRegion API asynchronously
func (client *Client) UnSubscribeRegionWithChan(request *UnSubscribeRegionRequest) (<-chan *UnSubscribeRegionResponse, <-chan error) {
	responseChan := make(chan *UnSubscribeRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnSubscribeRegion(request)
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

// UnSubscribeRegionWithCallback invokes the openanalytics_open.UnSubscribeRegion API asynchronously
func (client *Client) UnSubscribeRegionWithCallback(request *UnSubscribeRegionRequest, callback func(response *UnSubscribeRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnSubscribeRegionResponse
		var err error
		defer close(result)
		response, err = client.UnSubscribeRegion(request)
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

// UnSubscribeRegionRequest is the request struct for api UnSubscribeRegion
type UnSubscribeRegionRequest struct {
	*requests.RpcRequest
}

// UnSubscribeRegionResponse is the response struct for api UnSubscribeRegion
type UnSubscribeRegionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
}

// CreateUnSubscribeRegionRequest creates a request to invoke UnSubscribeRegion API
func CreateUnSubscribeRegionRequest() (request *UnSubscribeRegionRequest) {
	request = &UnSubscribeRegionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "UnSubscribeRegion", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnSubscribeRegionResponse creates a response to parse from UnSubscribeRegion response
func CreateUnSubscribeRegionResponse() (response *UnSubscribeRegionResponse) {
	response = &UnSubscribeRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
