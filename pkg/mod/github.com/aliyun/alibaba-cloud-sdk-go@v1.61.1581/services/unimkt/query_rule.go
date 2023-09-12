package unimkt

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

// QueryRule invokes the unimkt.QueryRule API synchronously
func (client *Client) QueryRule(request *QueryRuleRequest) (response *QueryRuleResponse, err error) {
	response = CreateQueryRuleResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRuleWithChan invokes the unimkt.QueryRule API asynchronously
func (client *Client) QueryRuleWithChan(request *QueryRuleRequest) (<-chan *QueryRuleResponse, <-chan error) {
	responseChan := make(chan *QueryRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRule(request)
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

// QueryRuleWithCallback invokes the unimkt.QueryRule API asynchronously
func (client *Client) QueryRuleWithCallback(request *QueryRuleRequest, callback func(response *QueryRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRuleResponse
		var err error
		defer close(result)
		response, err = client.QueryRule(request)
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

// QueryRuleRequest is the request struct for api QueryRule
type QueryRuleRequest struct {
	*requests.RpcRequest
	Business         string `position:"Query" name:"Business"`
	UserId           string `position:"Query" name:"UserId"`
	OriginSiteUserId string `position:"Query" name:"OriginSiteUserId"`
	Environment      string `position:"Query" name:"Environment"`
	AppName          string `position:"Query" name:"AppName"`
	TenantId         string `position:"Query" name:"TenantId"`
	UserSite         string `position:"Query" name:"UserSite"`
	RuleId           string `position:"Query" name:"RuleId"`
}

// QueryRuleResponse is the response struct for api QueryRule
type QueryRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Model     Model  `json:"Model" xml:"Model"`
}

// CreateQueryRuleRequest creates a request to invoke QueryRule API
func CreateQueryRuleRequest() (request *QueryRuleRequest) {
	request = &QueryRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "QueryRule", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRuleResponse creates a response to parse from QueryRule response
func CreateQueryRuleResponse() (response *QueryRuleResponse) {
	response = &QueryRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
