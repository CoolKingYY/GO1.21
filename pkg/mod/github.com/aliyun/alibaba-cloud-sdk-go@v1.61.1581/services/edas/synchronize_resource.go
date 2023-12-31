package edas

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

// SynchronizeResource invokes the edas.SynchronizeResource API synchronously
func (client *Client) SynchronizeResource(request *SynchronizeResourceRequest) (response *SynchronizeResourceResponse, err error) {
	response = CreateSynchronizeResourceResponse()
	err = client.DoAction(request, response)
	return
}

// SynchronizeResourceWithChan invokes the edas.SynchronizeResource API asynchronously
func (client *Client) SynchronizeResourceWithChan(request *SynchronizeResourceRequest) (<-chan *SynchronizeResourceResponse, <-chan error) {
	responseChan := make(chan *SynchronizeResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SynchronizeResource(request)
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

// SynchronizeResourceWithCallback invokes the edas.SynchronizeResource API asynchronously
func (client *Client) SynchronizeResourceWithCallback(request *SynchronizeResourceRequest, callback func(response *SynchronizeResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SynchronizeResourceResponse
		var err error
		defer close(result)
		response, err = client.SynchronizeResource(request)
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

// SynchronizeResourceRequest is the request struct for api SynchronizeResource
type SynchronizeResourceRequest struct {
	*requests.RoaRequest
	Type        string `position:"Query" name:"Type"`
	ResourceIds string `position:"Query" name:"ResourceIds"`
}

// SynchronizeResourceResponse is the response struct for api SynchronizeResource
type SynchronizeResourceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSynchronizeResourceRequest creates a request to invoke SynchronizeResource API
func CreateSynchronizeResourceRequest() (request *SynchronizeResourceRequest) {
	request = &SynchronizeResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "SynchronizeResource", "/pop/v5/resource/pop_sync_resource", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateSynchronizeResourceResponse creates a response to parse from SynchronizeResource response
func CreateSynchronizeResourceResponse() (response *SynchronizeResourceResponse) {
	response = &SynchronizeResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
