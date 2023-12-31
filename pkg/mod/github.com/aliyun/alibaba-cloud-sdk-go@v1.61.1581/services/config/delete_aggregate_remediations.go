package config

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

// DeleteAggregateRemediations invokes the config.DeleteAggregateRemediations API synchronously
func (client *Client) DeleteAggregateRemediations(request *DeleteAggregateRemediationsRequest) (response *DeleteAggregateRemediationsResponse, err error) {
	response = CreateDeleteAggregateRemediationsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAggregateRemediationsWithChan invokes the config.DeleteAggregateRemediations API asynchronously
func (client *Client) DeleteAggregateRemediationsWithChan(request *DeleteAggregateRemediationsRequest) (<-chan *DeleteAggregateRemediationsResponse, <-chan error) {
	responseChan := make(chan *DeleteAggregateRemediationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAggregateRemediations(request)
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

// DeleteAggregateRemediationsWithCallback invokes the config.DeleteAggregateRemediations API asynchronously
func (client *Client) DeleteAggregateRemediationsWithCallback(request *DeleteAggregateRemediationsRequest, callback func(response *DeleteAggregateRemediationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAggregateRemediationsResponse
		var err error
		defer close(result)
		response, err = client.DeleteAggregateRemediations(request)
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

// DeleteAggregateRemediationsRequest is the request struct for api DeleteAggregateRemediations
type DeleteAggregateRemediationsRequest struct {
	*requests.RpcRequest
	RemediationIds string `position:"Body" name:"RemediationIds"`
	AggregatorId   string `position:"Body" name:"AggregatorId"`
}

// DeleteAggregateRemediationsResponse is the response struct for api DeleteAggregateRemediations
type DeleteAggregateRemediationsResponse struct {
	*responses.BaseResponse
	RequestId                string                    `json:"RequestId" xml:"RequestId"`
	RemediationDeleteResults []RemediationDeleteResult `json:"RemediationDeleteResults" xml:"RemediationDeleteResults"`
}

// CreateDeleteAggregateRemediationsRequest creates a request to invoke DeleteAggregateRemediations API
func CreateDeleteAggregateRemediationsRequest() (request *DeleteAggregateRemediationsRequest) {
	request = &DeleteAggregateRemediationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "DeleteAggregateRemediations", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAggregateRemediationsResponse creates a response to parse from DeleteAggregateRemediations response
func CreateDeleteAggregateRemediationsResponse() (response *DeleteAggregateRemediationsResponse) {
	response = &DeleteAggregateRemediationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
