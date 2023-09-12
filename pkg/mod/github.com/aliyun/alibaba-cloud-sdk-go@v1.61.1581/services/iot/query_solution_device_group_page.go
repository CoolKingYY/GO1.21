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

// QuerySolutionDeviceGroupPage invokes the iot.QuerySolutionDeviceGroupPage API synchronously
func (client *Client) QuerySolutionDeviceGroupPage(request *QuerySolutionDeviceGroupPageRequest) (response *QuerySolutionDeviceGroupPageResponse, err error) {
	response = CreateQuerySolutionDeviceGroupPageResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySolutionDeviceGroupPageWithChan invokes the iot.QuerySolutionDeviceGroupPage API asynchronously
func (client *Client) QuerySolutionDeviceGroupPageWithChan(request *QuerySolutionDeviceGroupPageRequest) (<-chan *QuerySolutionDeviceGroupPageResponse, <-chan error) {
	responseChan := make(chan *QuerySolutionDeviceGroupPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySolutionDeviceGroupPage(request)
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

// QuerySolutionDeviceGroupPageWithCallback invokes the iot.QuerySolutionDeviceGroupPage API asynchronously
func (client *Client) QuerySolutionDeviceGroupPageWithCallback(request *QuerySolutionDeviceGroupPageRequest, callback func(response *QuerySolutionDeviceGroupPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySolutionDeviceGroupPageResponse
		var err error
		defer close(result)
		response, err = client.QuerySolutionDeviceGroupPage(request)
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

// QuerySolutionDeviceGroupPageRequest is the request struct for api QuerySolutionDeviceGroupPage
type QuerySolutionDeviceGroupPageRequest struct {
	*requests.RpcRequest
	FuzzyGroupName string           `position:"Query" name:"FuzzyGroupName"`
	ProjectCode    string           `position:"Query" name:"ProjectCode"`
	PageId         requests.Integer `position:"Query" name:"PageId"`
	IotInstanceId  string           `position:"Query" name:"IotInstanceId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ApiProduct     string           `position:"Body" name:"ApiProduct"`
	ApiRevision    string           `position:"Body" name:"ApiRevision"`
}

// QuerySolutionDeviceGroupPageResponse is the response struct for api QuerySolutionDeviceGroupPage
type QuerySolutionDeviceGroupPageResponse struct {
	*responses.BaseResponse
	RequestId    string                             `json:"RequestId" xml:"RequestId"`
	Success      bool                               `json:"Success" xml:"Success"`
	Code         string                             `json:"Code" xml:"Code"`
	ErrorMessage string                             `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySolutionDeviceGroupPage `json:"Data" xml:"Data"`
}

// CreateQuerySolutionDeviceGroupPageRequest creates a request to invoke QuerySolutionDeviceGroupPage API
func CreateQuerySolutionDeviceGroupPageRequest() (request *QuerySolutionDeviceGroupPageRequest) {
	request = &QuerySolutionDeviceGroupPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySolutionDeviceGroupPage", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySolutionDeviceGroupPageResponse creates a response to parse from QuerySolutionDeviceGroupPage response
func CreateQuerySolutionDeviceGroupPageResponse() (response *QuerySolutionDeviceGroupPageResponse) {
	response = &QuerySolutionDeviceGroupPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
