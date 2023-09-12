package retailadvqa_public

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

// LoadDataToExternalAudience invokes the retailadvqa_public.LoadDataToExternalAudience API synchronously
func (client *Client) LoadDataToExternalAudience(request *LoadDataToExternalAudienceRequest) (response *LoadDataToExternalAudienceResponse, err error) {
	response = CreateLoadDataToExternalAudienceResponse()
	err = client.DoAction(request, response)
	return
}

// LoadDataToExternalAudienceWithChan invokes the retailadvqa_public.LoadDataToExternalAudience API asynchronously
func (client *Client) LoadDataToExternalAudienceWithChan(request *LoadDataToExternalAudienceRequest) (<-chan *LoadDataToExternalAudienceResponse, <-chan error) {
	responseChan := make(chan *LoadDataToExternalAudienceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LoadDataToExternalAudience(request)
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

// LoadDataToExternalAudienceWithCallback invokes the retailadvqa_public.LoadDataToExternalAudience API asynchronously
func (client *Client) LoadDataToExternalAudienceWithCallback(request *LoadDataToExternalAudienceRequest, callback func(response *LoadDataToExternalAudienceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LoadDataToExternalAudienceResponse
		var err error
		defer close(result)
		response, err = client.LoadDataToExternalAudience(request)
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

// LoadDataToExternalAudienceRequest is the request struct for api LoadDataToExternalAudience
type LoadDataToExternalAudienceRequest struct {
	*requests.RpcRequest
	AccessId   string `position:"Query" name:"AccessId"`
	AudienceId string `position:"Query" name:"AudienceId"`
	OssPath    string `position:"Query" name:"OssPath"`
}

// LoadDataToExternalAudienceResponse is the response struct for api LoadDataToExternalAudience
type LoadDataToExternalAudienceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	ExStack   string `json:"ExStack" xml:"ExStack"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateLoadDataToExternalAudienceRequest creates a request to invoke LoadDataToExternalAudience API
func CreateLoadDataToExternalAudienceRequest() (request *LoadDataToExternalAudienceRequest) {
	request = &LoadDataToExternalAudienceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "LoadDataToExternalAudience", "", "")
	request.Method = requests.POST
	return
}

// CreateLoadDataToExternalAudienceResponse creates a response to parse from LoadDataToExternalAudience response
func CreateLoadDataToExternalAudienceResponse() (response *LoadDataToExternalAudienceResponse) {
	response = &LoadDataToExternalAudienceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
