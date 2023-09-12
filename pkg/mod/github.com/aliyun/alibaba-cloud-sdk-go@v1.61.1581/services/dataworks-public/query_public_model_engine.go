package dataworks_public

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

// QueryPublicModelEngine invokes the dataworks_public.QueryPublicModelEngine API synchronously
func (client *Client) QueryPublicModelEngine(request *QueryPublicModelEngineRequest) (response *QueryPublicModelEngineResponse, err error) {
	response = CreateQueryPublicModelEngineResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPublicModelEngineWithChan invokes the dataworks_public.QueryPublicModelEngine API asynchronously
func (client *Client) QueryPublicModelEngineWithChan(request *QueryPublicModelEngineRequest) (<-chan *QueryPublicModelEngineResponse, <-chan error) {
	responseChan := make(chan *QueryPublicModelEngineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPublicModelEngine(request)
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

// QueryPublicModelEngineWithCallback invokes the dataworks_public.QueryPublicModelEngine API asynchronously
func (client *Client) QueryPublicModelEngineWithCallback(request *QueryPublicModelEngineRequest, callback func(response *QueryPublicModelEngineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPublicModelEngineResponse
		var err error
		defer close(result)
		response, err = client.QueryPublicModelEngine(request)
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

// QueryPublicModelEngineRequest is the request struct for api QueryPublicModelEngine
type QueryPublicModelEngineRequest struct {
	*requests.RpcRequest
	Text      string `position:"Body" name:"Text"`
	ProjectId string `position:"Body" name:"ProjectId"`
}

// QueryPublicModelEngineResponse is the response struct for api QueryPublicModelEngine
type QueryPublicModelEngineResponse struct {
	*responses.BaseResponse
	RequestId   string                   `json:"RequestId" xml:"RequestId"`
	ReturnValue []map[string]interface{} `json:"ReturnValue" xml:"ReturnValue"`
}

// CreateQueryPublicModelEngineRequest creates a request to invoke QueryPublicModelEngine API
func CreateQueryPublicModelEngineRequest() (request *QueryPublicModelEngineRequest) {
	request = &QueryPublicModelEngineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "QueryPublicModelEngine", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryPublicModelEngineResponse creates a response to parse from QueryPublicModelEngine response
func CreateQueryPublicModelEngineResponse() (response *QueryPublicModelEngineResponse) {
	response = &QueryPublicModelEngineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
