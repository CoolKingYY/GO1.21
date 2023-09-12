package mse

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

// UpdateMessageQueueRoute invokes the mse.UpdateMessageQueueRoute API synchronously
func (client *Client) UpdateMessageQueueRoute(request *UpdateMessageQueueRouteRequest) (response *UpdateMessageQueueRouteResponse, err error) {
	response = CreateUpdateMessageQueueRouteResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMessageQueueRouteWithChan invokes the mse.UpdateMessageQueueRoute API asynchronously
func (client *Client) UpdateMessageQueueRouteWithChan(request *UpdateMessageQueueRouteRequest) (<-chan *UpdateMessageQueueRouteResponse, <-chan error) {
	responseChan := make(chan *UpdateMessageQueueRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMessageQueueRoute(request)
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

// UpdateMessageQueueRouteWithCallback invokes the mse.UpdateMessageQueueRoute API asynchronously
func (client *Client) UpdateMessageQueueRouteWithCallback(request *UpdateMessageQueueRouteRequest, callback func(response *UpdateMessageQueueRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMessageQueueRouteResponse
		var err error
		defer close(result)
		response, err = client.UpdateMessageQueueRoute(request)
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

// UpdateMessageQueueRouteRequest is the request struct for api UpdateMessageQueueRoute
type UpdateMessageQueueRouteRequest struct {
	*requests.RpcRequest
	Tags           *[]string        `position:"Query" name:"Tags"  type:"Json"`
	Enable         requests.Boolean `position:"Query" name:"Enable"`
	AppId          string           `position:"Query" name:"AppId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Region         string           `position:"Query" name:"Region"`
}

// UpdateMessageQueueRouteResponse is the response struct for api UpdateMessageQueueRoute
type UpdateMessageQueueRouteResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	Data           string `json:"Data" xml:"Data"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdateMessageQueueRouteRequest creates a request to invoke UpdateMessageQueueRoute API
func CreateUpdateMessageQueueRouteRequest() (request *UpdateMessageQueueRouteRequest) {
	request = &UpdateMessageQueueRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateMessageQueueRoute", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMessageQueueRouteResponse creates a response to parse from UpdateMessageQueueRoute response
func CreateUpdateMessageQueueRouteResponse() (response *UpdateMessageQueueRouteResponse) {
	response = &UpdateMessageQueueRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
