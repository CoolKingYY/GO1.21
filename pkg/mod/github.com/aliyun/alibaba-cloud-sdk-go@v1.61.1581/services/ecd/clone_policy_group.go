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

// ClonePolicyGroup invokes the ecd.ClonePolicyGroup API synchronously
func (client *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (response *ClonePolicyGroupResponse, err error) {
	response = CreateClonePolicyGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ClonePolicyGroupWithChan invokes the ecd.ClonePolicyGroup API asynchronously
func (client *Client) ClonePolicyGroupWithChan(request *ClonePolicyGroupRequest) (<-chan *ClonePolicyGroupResponse, <-chan error) {
	responseChan := make(chan *ClonePolicyGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClonePolicyGroup(request)
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

// ClonePolicyGroupWithCallback invokes the ecd.ClonePolicyGroup API asynchronously
func (client *Client) ClonePolicyGroupWithCallback(request *ClonePolicyGroupRequest, callback func(response *ClonePolicyGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClonePolicyGroupResponse
		var err error
		defer close(result)
		response, err = client.ClonePolicyGroup(request)
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

// ClonePolicyGroupRequest is the request struct for api ClonePolicyGroup
type ClonePolicyGroupRequest struct {
	*requests.RpcRequest
	Name          string `position:"Query" name:"Name"`
	PolicyGroupId string `position:"Query" name:"PolicyGroupId"`
}

// ClonePolicyGroupResponse is the response struct for api ClonePolicyGroup
type ClonePolicyGroupResponse struct {
	*responses.BaseResponse
	PolicyGroupId string `json:"PolicyGroupId" xml:"PolicyGroupId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateClonePolicyGroupRequest creates a request to invoke ClonePolicyGroup API
func CreateClonePolicyGroupRequest() (request *ClonePolicyGroupRequest) {
	request = &ClonePolicyGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ClonePolicyGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateClonePolicyGroupResponse creates a response to parse from ClonePolicyGroup response
func CreateClonePolicyGroupResponse() (response *ClonePolicyGroupResponse) {
	response = &ClonePolicyGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
