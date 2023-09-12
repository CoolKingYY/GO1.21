package opensearch

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

// RemoveSecondRank invokes the opensearch.RemoveSecondRank API synchronously
func (client *Client) RemoveSecondRank(request *RemoveSecondRankRequest) (response *RemoveSecondRankResponse, err error) {
	response = CreateRemoveSecondRankResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveSecondRankWithChan invokes the opensearch.RemoveSecondRank API asynchronously
func (client *Client) RemoveSecondRankWithChan(request *RemoveSecondRankRequest) (<-chan *RemoveSecondRankResponse, <-chan error) {
	responseChan := make(chan *RemoveSecondRankResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveSecondRank(request)
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

// RemoveSecondRankWithCallback invokes the opensearch.RemoveSecondRank API asynchronously
func (client *Client) RemoveSecondRankWithCallback(request *RemoveSecondRankRequest, callback func(response *RemoveSecondRankResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveSecondRankResponse
		var err error
		defer close(result)
		response, err = client.RemoveSecondRank(request)
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

// RemoveSecondRankRequest is the request struct for api RemoveSecondRank
type RemoveSecondRankRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	Name             string           `position:"Path" name:"name"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// RemoveSecondRankResponse is the response struct for api RemoveSecondRank
type RemoveSecondRankResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateRemoveSecondRankRequest creates a request to invoke RemoveSecondRank API
func CreateRemoveSecondRankRequest() (request *RemoveSecondRankRequest) {
	request = &RemoveSecondRankRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RemoveSecondRank", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/second-ranks/[name]", "opensearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateRemoveSecondRankResponse creates a response to parse from RemoveSecondRank response
func CreateRemoveSecondRankResponse() (response *RemoveSecondRankResponse) {
	response = &RemoveSecondRankResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
