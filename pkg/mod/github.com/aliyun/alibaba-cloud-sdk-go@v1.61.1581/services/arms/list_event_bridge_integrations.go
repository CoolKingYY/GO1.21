package arms

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

// ListEventBridgeIntegrations invokes the arms.ListEventBridgeIntegrations API synchronously
func (client *Client) ListEventBridgeIntegrations(request *ListEventBridgeIntegrationsRequest) (response *ListEventBridgeIntegrationsResponse, err error) {
	response = CreateListEventBridgeIntegrationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEventBridgeIntegrationsWithChan invokes the arms.ListEventBridgeIntegrations API asynchronously
func (client *Client) ListEventBridgeIntegrationsWithChan(request *ListEventBridgeIntegrationsRequest) (<-chan *ListEventBridgeIntegrationsResponse, <-chan error) {
	responseChan := make(chan *ListEventBridgeIntegrationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEventBridgeIntegrations(request)
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

// ListEventBridgeIntegrationsWithCallback invokes the arms.ListEventBridgeIntegrations API asynchronously
func (client *Client) ListEventBridgeIntegrationsWithCallback(request *ListEventBridgeIntegrationsRequest, callback func(response *ListEventBridgeIntegrationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEventBridgeIntegrationsResponse
		var err error
		defer close(result)
		response, err = client.ListEventBridgeIntegrations(request)
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

// ListEventBridgeIntegrationsRequest is the request struct for api ListEventBridgeIntegrations
type ListEventBridgeIntegrationsRequest struct {
	*requests.RpcRequest
	Size requests.Integer `position:"Query" name:"Size"`
	Name string           `position:"Query" name:"Name"`
	Page requests.Integer `position:"Query" name:"Page"`
}

// ListEventBridgeIntegrationsResponse is the response struct for api ListEventBridgeIntegrations
type ListEventBridgeIntegrationsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateListEventBridgeIntegrationsRequest creates a request to invoke ListEventBridgeIntegrations API
func CreateListEventBridgeIntegrationsRequest() (request *ListEventBridgeIntegrationsRequest) {
	request = &ListEventBridgeIntegrationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListEventBridgeIntegrations", "", "")
	request.Method = requests.GET
	return
}

// CreateListEventBridgeIntegrationsResponse creates a response to parse from ListEventBridgeIntegrations response
func CreateListEventBridgeIntegrationsResponse() (response *ListEventBridgeIntegrationsResponse) {
	response = &ListEventBridgeIntegrationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}