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

// GetServiceList invokes the mse.GetServiceList API synchronously
func (client *Client) GetServiceList(request *GetServiceListRequest) (response *GetServiceListResponse, err error) {
	response = CreateGetServiceListResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceListWithChan invokes the mse.GetServiceList API asynchronously
func (client *Client) GetServiceListWithChan(request *GetServiceListRequest) (<-chan *GetServiceListResponse, <-chan error) {
	responseChan := make(chan *GetServiceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceList(request)
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

// GetServiceListWithCallback invokes the mse.GetServiceList API asynchronously
func (client *Client) GetServiceListWithCallback(request *GetServiceListRequest, callback func(response *GetServiceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceListResponse
		var err error
		defer close(result)
		response, err = client.GetServiceList(request)
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

// GetServiceListRequest is the request struct for api GetServiceList
type GetServiceListRequest struct {
	*requests.RpcRequest
	ServiceName    string `position:"Query" name:"ServiceName"`
	Ip             string `position:"Query" name:"Ip"`
	ServiceType    string `position:"Query" name:"ServiceType"`
	AppId          string `position:"Query" name:"AppId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Region         string `position:"Query" name:"Region"`
}

// GetServiceListResponse is the response struct for api GetServiceList
type GetServiceListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                     `json:"Message" xml:"Message"`
	RequestId      string                     `json:"RequestId" xml:"RequestId"`
	Code           int                        `json:"Code" xml:"Code"`
	Success        bool                       `json:"Success" xml:"Success"`
	Data           []MscServiceDetailResponse `json:"Data" xml:"Data"`
}

// CreateGetServiceListRequest creates a request to invoke GetServiceList API
func CreateGetServiceListRequest() (request *GetServiceListRequest) {
	request = &GetServiceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetServiceList", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetServiceListResponse creates a response to parse from GetServiceList response
func CreateGetServiceListResponse() (response *GetServiceListResponse) {
	response = &GetServiceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
