package das

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

// GetDasProServiceUsage invokes the das.GetDasProServiceUsage API synchronously
func (client *Client) GetDasProServiceUsage(request *GetDasProServiceUsageRequest) (response *GetDasProServiceUsageResponse, err error) {
	response = CreateGetDasProServiceUsageResponse()
	err = client.DoAction(request, response)
	return
}

// GetDasProServiceUsageWithChan invokes the das.GetDasProServiceUsage API asynchronously
func (client *Client) GetDasProServiceUsageWithChan(request *GetDasProServiceUsageRequest) (<-chan *GetDasProServiceUsageResponse, <-chan error) {
	responseChan := make(chan *GetDasProServiceUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDasProServiceUsage(request)
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

// GetDasProServiceUsageWithCallback invokes the das.GetDasProServiceUsage API asynchronously
func (client *Client) GetDasProServiceUsageWithCallback(request *GetDasProServiceUsageRequest, callback func(response *GetDasProServiceUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDasProServiceUsageResponse
		var err error
		defer close(result)
		response, err = client.GetDasProServiceUsage(request)
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

// GetDasProServiceUsageRequest is the request struct for api GetDasProServiceUsage
type GetDasProServiceUsageRequest struct {
	*requests.RpcRequest
	Signature      string `position:"Query" name:"Signature"`
	UserId         string `position:"Query" name:"UserId"`
	Uid            string `position:"Query" name:"Uid"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	AccessKey      string `position:"Query" name:"AccessKey"`
}

// GetDasProServiceUsageResponse is the response struct for api GetDasProServiceUsage
type GetDasProServiceUsageResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int64  `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      []Unit `json:"Data" xml:"Data"`
}

// CreateGetDasProServiceUsageRequest creates a request to invoke GetDasProServiceUsage API
func CreateGetDasProServiceUsageRequest() (request *GetDasProServiceUsageRequest) {
	request = &GetDasProServiceUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetDasProServiceUsage", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDasProServiceUsageResponse creates a response to parse from GetDasProServiceUsage response
func CreateGetDasProServiceUsageResponse() (response *GetDasProServiceUsageResponse) {
	response = &GetDasProServiceUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
