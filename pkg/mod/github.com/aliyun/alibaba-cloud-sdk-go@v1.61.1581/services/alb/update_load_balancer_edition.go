package alb

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

// UpdateLoadBalancerEdition invokes the alb.UpdateLoadBalancerEdition API synchronously
func (client *Client) UpdateLoadBalancerEdition(request *UpdateLoadBalancerEditionRequest) (response *UpdateLoadBalancerEditionResponse, err error) {
	response = CreateUpdateLoadBalancerEditionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLoadBalancerEditionWithChan invokes the alb.UpdateLoadBalancerEdition API asynchronously
func (client *Client) UpdateLoadBalancerEditionWithChan(request *UpdateLoadBalancerEditionRequest) (<-chan *UpdateLoadBalancerEditionResponse, <-chan error) {
	responseChan := make(chan *UpdateLoadBalancerEditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLoadBalancerEdition(request)
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

// UpdateLoadBalancerEditionWithCallback invokes the alb.UpdateLoadBalancerEdition API asynchronously
func (client *Client) UpdateLoadBalancerEditionWithCallback(request *UpdateLoadBalancerEditionRequest, callback func(response *UpdateLoadBalancerEditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLoadBalancerEditionResponse
		var err error
		defer close(result)
		response, err = client.UpdateLoadBalancerEdition(request)
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

// UpdateLoadBalancerEditionRequest is the request struct for api UpdateLoadBalancerEdition
type UpdateLoadBalancerEditionRequest struct {
	*requests.RpcRequest
	LoadBalancerEdition string           `position:"Query" name:"LoadBalancerEdition"`
	ClientToken         string           `position:"Query" name:"ClientToken"`
	DryRun              requests.Boolean `position:"Query" name:"DryRun"`
	LoadBalancerId      string           `position:"Query" name:"LoadBalancerId"`
}

// UpdateLoadBalancerEditionResponse is the response struct for api UpdateLoadBalancerEdition
type UpdateLoadBalancerEditionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLoadBalancerEditionRequest creates a request to invoke UpdateLoadBalancerEdition API
func CreateUpdateLoadBalancerEditionRequest() (request *UpdateLoadBalancerEditionRequest) {
	request = &UpdateLoadBalancerEditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "UpdateLoadBalancerEdition", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLoadBalancerEditionResponse creates a response to parse from UpdateLoadBalancerEdition response
func CreateUpdateLoadBalancerEditionResponse() (response *UpdateLoadBalancerEditionResponse) {
	response = &UpdateLoadBalancerEditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
