package outboundbot

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

// QueryScriptsByStatus invokes the outboundbot.QueryScriptsByStatus API synchronously
func (client *Client) QueryScriptsByStatus(request *QueryScriptsByStatusRequest) (response *QueryScriptsByStatusResponse, err error) {
	response = CreateQueryScriptsByStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryScriptsByStatusWithChan invokes the outboundbot.QueryScriptsByStatus API asynchronously
func (client *Client) QueryScriptsByStatusWithChan(request *QueryScriptsByStatusRequest) (<-chan *QueryScriptsByStatusResponse, <-chan error) {
	responseChan := make(chan *QueryScriptsByStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryScriptsByStatus(request)
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

// QueryScriptsByStatusWithCallback invokes the outboundbot.QueryScriptsByStatus API asynchronously
func (client *Client) QueryScriptsByStatusWithCallback(request *QueryScriptsByStatusRequest, callback func(response *QueryScriptsByStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryScriptsByStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryScriptsByStatus(request)
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

// QueryScriptsByStatusRequest is the request struct for api QueryScriptsByStatus
type QueryScriptsByStatusRequest struct {
	*requests.RpcRequest
	StatusList *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// QueryScriptsByStatusResponse is the response struct for api QueryScriptsByStatus
type QueryScriptsByStatusResponse struct {
	*responses.BaseResponse
	HttpStatusCode int     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string  `json:"Code" xml:"Code"`
	Message        string  `json:"Message" xml:"Message"`
	RequestId      string  `json:"RequestId" xml:"RequestId"`
	Success        bool    `json:"Success" xml:"Success"`
	Scripts        Scripts `json:"Scripts" xml:"Scripts"`
}

// CreateQueryScriptsByStatusRequest creates a request to invoke QueryScriptsByStatus API
func CreateQueryScriptsByStatusRequest() (request *QueryScriptsByStatusRequest) {
	request = &QueryScriptsByStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "QueryScriptsByStatus", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryScriptsByStatusResponse creates a response to parse from QueryScriptsByStatus response
func CreateQueryScriptsByStatusResponse() (response *QueryScriptsByStatusResponse) {
	response = &QueryScriptsByStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
