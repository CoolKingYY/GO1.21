package quickbi_public

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

// QueryShareList invokes the quickbi_public.QueryShareList API synchronously
func (client *Client) QueryShareList(request *QueryShareListRequest) (response *QueryShareListResponse, err error) {
	response = CreateQueryShareListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryShareListWithChan invokes the quickbi_public.QueryShareList API asynchronously
func (client *Client) QueryShareListWithChan(request *QueryShareListRequest) (<-chan *QueryShareListResponse, <-chan error) {
	responseChan := make(chan *QueryShareListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryShareList(request)
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

// QueryShareListWithCallback invokes the quickbi_public.QueryShareList API asynchronously
func (client *Client) QueryShareListWithCallback(request *QueryShareListRequest, callback func(response *QueryShareListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryShareListResponse
		var err error
		defer close(result)
		response, err = client.QueryShareList(request)
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

// QueryShareListRequest is the request struct for api QueryShareList
type QueryShareListRequest struct {
	*requests.RpcRequest
	ReportId    string `position:"Query" name:"ReportId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryShareListResponse is the response struct for api QueryShareList
type QueryShareListResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Data `json:"Result" xml:"Result"`
}

// CreateQueryShareListRequest creates a request to invoke QueryShareList API
func CreateQueryShareListRequest() (request *QueryShareListRequest) {
	request = &QueryShareListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryShareList", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryShareListResponse creates a response to parse from QueryShareList response
func CreateQueryShareListResponse() (response *QueryShareListResponse) {
	response = &QueryShareListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
