package qualitycheck

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

// GetRuleDimension invokes the qualitycheck.GetRuleDimension API synchronously
func (client *Client) GetRuleDimension(request *GetRuleDimensionRequest) (response *GetRuleDimensionResponse, err error) {
	response = CreateGetRuleDimensionResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleDimensionWithChan invokes the qualitycheck.GetRuleDimension API asynchronously
func (client *Client) GetRuleDimensionWithChan(request *GetRuleDimensionRequest) (<-chan *GetRuleDimensionResponse, <-chan error) {
	responseChan := make(chan *GetRuleDimensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleDimension(request)
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

// GetRuleDimensionWithCallback invokes the qualitycheck.GetRuleDimension API asynchronously
func (client *Client) GetRuleDimensionWithCallback(request *GetRuleDimensionRequest, callback func(response *GetRuleDimensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleDimensionResponse
		var err error
		defer close(result)
		response, err = client.GetRuleDimension(request)
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

// GetRuleDimensionRequest is the request struct for api GetRuleDimension
type GetRuleDimensionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetRuleDimensionResponse is the response struct for api GetRuleDimension
type GetRuleDimensionResponse struct {
	*responses.BaseResponse
	CurrentPage       int                    `json:"CurrentPage" xml:"CurrentPage"`
	DataSize          int                    `json:"DataSize" xml:"DataSize"`
	RequestId         string                 `json:"RequestId" xml:"RequestId"`
	Success           bool                   `json:"Success" xml:"Success"`
	ReviewStatus      int                    `json:"ReviewStatus" xml:"ReviewStatus"`
	CompSubTaskCount  int                    `json:"CompSubTaskCount" xml:"CompSubTaskCount"`
	Code              string                 `json:"Code" xml:"Code"`
	TotalSubTaskCount int                    `json:"TotalSubTaskCount" xml:"TotalSubTaskCount"`
	Message           string                 `json:"Message" xml:"Message"`
	PageSize          int                    `json:"PageSize" xml:"PageSize"`
	TotalCount        int                    `json:"TotalCount" xml:"TotalCount"`
	Data              DataInGetRuleDimension `json:"Data" xml:"Data"`
}

// CreateGetRuleDimensionRequest creates a request to invoke GetRuleDimension API
func CreateGetRuleDimensionRequest() (request *GetRuleDimensionRequest) {
	request = &GetRuleDimensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRuleDimension", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleDimensionResponse creates a response to parse from GetRuleDimension response
func CreateGetRuleDimensionResponse() (response *GetRuleDimensionResponse) {
	response = &GetRuleDimensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
