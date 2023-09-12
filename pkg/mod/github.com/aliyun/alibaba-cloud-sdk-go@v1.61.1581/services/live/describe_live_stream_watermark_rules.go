package live

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

// DescribeLiveStreamWatermarkRules invokes the live.DescribeLiveStreamWatermarkRules API synchronously
func (client *Client) DescribeLiveStreamWatermarkRules(request *DescribeLiveStreamWatermarkRulesRequest) (response *DescribeLiveStreamWatermarkRulesResponse, err error) {
	response = CreateDescribeLiveStreamWatermarkRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamWatermarkRulesWithChan invokes the live.DescribeLiveStreamWatermarkRules API asynchronously
func (client *Client) DescribeLiveStreamWatermarkRulesWithChan(request *DescribeLiveStreamWatermarkRulesRequest) (<-chan *DescribeLiveStreamWatermarkRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamWatermarkRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamWatermarkRules(request)
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

// DescribeLiveStreamWatermarkRulesWithCallback invokes the live.DescribeLiveStreamWatermarkRules API asynchronously
func (client *Client) DescribeLiveStreamWatermarkRulesWithCallback(request *DescribeLiveStreamWatermarkRulesRequest, callback func(response *DescribeLiveStreamWatermarkRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamWatermarkRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamWatermarkRules(request)
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

// DescribeLiveStreamWatermarkRulesRequest is the request struct for api DescribeLiveStreamWatermarkRules
type DescribeLiveStreamWatermarkRulesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamWatermarkRulesResponse is the response struct for api DescribeLiveStreamWatermarkRules
type DescribeLiveStreamWatermarkRulesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	RuleInfoList RuleInfoList `json:"RuleInfoList" xml:"RuleInfoList"`
}

// CreateDescribeLiveStreamWatermarkRulesRequest creates a request to invoke DescribeLiveStreamWatermarkRules API
func CreateDescribeLiveStreamWatermarkRulesRequest() (request *DescribeLiveStreamWatermarkRulesRequest) {
	request = &DescribeLiveStreamWatermarkRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamWatermarkRules", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamWatermarkRulesResponse creates a response to parse from DescribeLiveStreamWatermarkRules response
func CreateDescribeLiveStreamWatermarkRulesResponse() (response *DescribeLiveStreamWatermarkRulesResponse) {
	response = &DescribeLiveStreamWatermarkRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
