package cms

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

// DeleteMonitorGroup invokes the cms.DeleteMonitorGroup API synchronously
func (client *Client) DeleteMonitorGroup(request *DeleteMonitorGroupRequest) (response *DeleteMonitorGroupResponse, err error) {
	response = CreateDeleteMonitorGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMonitorGroupWithChan invokes the cms.DeleteMonitorGroup API asynchronously
func (client *Client) DeleteMonitorGroupWithChan(request *DeleteMonitorGroupRequest) (<-chan *DeleteMonitorGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteMonitorGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMonitorGroup(request)
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

// DeleteMonitorGroupWithCallback invokes the cms.DeleteMonitorGroup API asynchronously
func (client *Client) DeleteMonitorGroupWithCallback(request *DeleteMonitorGroupRequest, callback func(response *DeleteMonitorGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMonitorGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteMonitorGroup(request)
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

// DeleteMonitorGroupRequest is the request struct for api DeleteMonitorGroup
type DeleteMonitorGroupRequest struct {
	*requests.RpcRequest
	GroupId requests.Integer `position:"Query" name:"GroupId"`
}

// DeleteMonitorGroupResponse is the response struct for api DeleteMonitorGroup
type DeleteMonitorGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Group     Group  `json:"Group" xml:"Group"`
}

// CreateDeleteMonitorGroupRequest creates a request to invoke DeleteMonitorGroup API
func CreateDeleteMonitorGroupRequest() (request *DeleteMonitorGroupRequest) {
	request = &DeleteMonitorGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMonitorGroup", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMonitorGroupResponse creates a response to parse from DeleteMonitorGroup response
func CreateDeleteMonitorGroupResponse() (response *DeleteMonitorGroupResponse) {
	response = &DeleteMonitorGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
