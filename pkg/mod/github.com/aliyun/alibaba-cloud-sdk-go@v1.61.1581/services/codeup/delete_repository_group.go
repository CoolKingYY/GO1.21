package codeup

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

// DeleteRepositoryGroup invokes the codeup.DeleteRepositoryGroup API synchronously
func (client *Client) DeleteRepositoryGroup(request *DeleteRepositoryGroupRequest) (response *DeleteRepositoryGroupResponse, err error) {
	response = CreateDeleteRepositoryGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepositoryGroupWithChan invokes the codeup.DeleteRepositoryGroup API asynchronously
func (client *Client) DeleteRepositoryGroupWithChan(request *DeleteRepositoryGroupRequest) (<-chan *DeleteRepositoryGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteRepositoryGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepositoryGroup(request)
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

// DeleteRepositoryGroupWithCallback invokes the codeup.DeleteRepositoryGroup API asynchronously
func (client *Client) DeleteRepositoryGroupWithCallback(request *DeleteRepositoryGroupRequest, callback func(response *DeleteRepositoryGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepositoryGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepositoryGroup(request)
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

// DeleteRepositoryGroupRequest is the request struct for api DeleteRepositoryGroup
type DeleteRepositoryGroupRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	SubUserId      string           `position:"Query" name:"SubUserId"`
	GroupId        requests.Integer `position:"Path" name:"GroupId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
}

// DeleteRepositoryGroupResponse is the response struct for api DeleteRepositoryGroup
type DeleteRepositoryGroupResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateDeleteRepositoryGroupRequest creates a request to invoke DeleteRepositoryGroup API
func CreateDeleteRepositoryGroupRequest() (request *DeleteRepositoryGroupRequest) {
	request = &DeleteRepositoryGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "DeleteRepositoryGroup", "/api/v3/groups/[GroupId]/remove", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRepositoryGroupResponse creates a response to parse from DeleteRepositoryGroup response
func CreateDeleteRepositoryGroupResponse() (response *DeleteRepositoryGroupResponse) {
	response = &DeleteRepositoryGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}