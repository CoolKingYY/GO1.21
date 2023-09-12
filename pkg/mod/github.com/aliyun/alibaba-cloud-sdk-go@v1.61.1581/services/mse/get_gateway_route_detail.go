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

// GetGatewayRouteDetail invokes the mse.GetGatewayRouteDetail API synchronously
func (client *Client) GetGatewayRouteDetail(request *GetGatewayRouteDetailRequest) (response *GetGatewayRouteDetailResponse, err error) {
	response = CreateGetGatewayRouteDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetGatewayRouteDetailWithChan invokes the mse.GetGatewayRouteDetail API asynchronously
func (client *Client) GetGatewayRouteDetailWithChan(request *GetGatewayRouteDetailRequest) (<-chan *GetGatewayRouteDetailResponse, <-chan error) {
	responseChan := make(chan *GetGatewayRouteDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGatewayRouteDetail(request)
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

// GetGatewayRouteDetailWithCallback invokes the mse.GetGatewayRouteDetail API asynchronously
func (client *Client) GetGatewayRouteDetailWithCallback(request *GetGatewayRouteDetailRequest, callback func(response *GetGatewayRouteDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGatewayRouteDetailResponse
		var err error
		defer close(result)
		response, err = client.GetGatewayRouteDetail(request)
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

// GetGatewayRouteDetailRequest is the request struct for api GetGatewayRouteDetail
type GetGatewayRouteDetailRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	RouteId         requests.Integer `position:"Query" name:"RouteId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
}

// GetGatewayRouteDetailResponse is the response struct for api GetGatewayRouteDetail
type GetGatewayRouteDetailResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetGatewayRouteDetailRequest creates a request to invoke GetGatewayRouteDetail API
func CreateGetGatewayRouteDetailRequest() (request *GetGatewayRouteDetailRequest) {
	request = &GetGatewayRouteDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetGatewayRouteDetail", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetGatewayRouteDetailResponse creates a response to parse from GetGatewayRouteDetail response
func CreateGetGatewayRouteDetailResponse() (response *GetGatewayRouteDetailResponse) {
	response = &GetGatewayRouteDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
