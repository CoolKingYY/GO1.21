package iot

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

// QueryEdgeInstanceSceneRule invokes the iot.QueryEdgeInstanceSceneRule API synchronously
func (client *Client) QueryEdgeInstanceSceneRule(request *QueryEdgeInstanceSceneRuleRequest) (response *QueryEdgeInstanceSceneRuleResponse, err error) {
	response = CreateQueryEdgeInstanceSceneRuleResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEdgeInstanceSceneRuleWithChan invokes the iot.QueryEdgeInstanceSceneRule API asynchronously
func (client *Client) QueryEdgeInstanceSceneRuleWithChan(request *QueryEdgeInstanceSceneRuleRequest) (<-chan *QueryEdgeInstanceSceneRuleResponse, <-chan error) {
	responseChan := make(chan *QueryEdgeInstanceSceneRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEdgeInstanceSceneRule(request)
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

// QueryEdgeInstanceSceneRuleWithCallback invokes the iot.QueryEdgeInstanceSceneRule API asynchronously
func (client *Client) QueryEdgeInstanceSceneRuleWithCallback(request *QueryEdgeInstanceSceneRuleRequest, callback func(response *QueryEdgeInstanceSceneRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEdgeInstanceSceneRuleResponse
		var err error
		defer close(result)
		response, err = client.QueryEdgeInstanceSceneRule(request)
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

// QueryEdgeInstanceSceneRuleRequest is the request struct for api QueryEdgeInstanceSceneRule
type QueryEdgeInstanceSceneRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryEdgeInstanceSceneRuleResponse is the response struct for api QueryEdgeInstanceSceneRule
type QueryEdgeInstanceSceneRuleResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	Success      bool                             `json:"Success" xml:"Success"`
	Code         string                           `json:"Code" xml:"Code"`
	ErrorMessage string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryEdgeInstanceSceneRule `json:"Data" xml:"Data"`
}

// CreateQueryEdgeInstanceSceneRuleRequest creates a request to invoke QueryEdgeInstanceSceneRule API
func CreateQueryEdgeInstanceSceneRuleRequest() (request *QueryEdgeInstanceSceneRuleRequest) {
	request = &QueryEdgeInstanceSceneRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryEdgeInstanceSceneRule", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryEdgeInstanceSceneRuleResponse creates a response to parse from QueryEdgeInstanceSceneRule response
func CreateQueryEdgeInstanceSceneRuleResponse() (response *QueryEdgeInstanceSceneRuleResponse) {
	response = &QueryEdgeInstanceSceneRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}