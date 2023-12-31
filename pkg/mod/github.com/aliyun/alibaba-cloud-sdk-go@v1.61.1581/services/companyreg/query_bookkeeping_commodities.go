package companyreg

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

// QueryBookkeepingCommodities invokes the companyreg.QueryBookkeepingCommodities API synchronously
func (client *Client) QueryBookkeepingCommodities(request *QueryBookkeepingCommoditiesRequest) (response *QueryBookkeepingCommoditiesResponse, err error) {
	response = CreateQueryBookkeepingCommoditiesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBookkeepingCommoditiesWithChan invokes the companyreg.QueryBookkeepingCommodities API asynchronously
func (client *Client) QueryBookkeepingCommoditiesWithChan(request *QueryBookkeepingCommoditiesRequest) (<-chan *QueryBookkeepingCommoditiesResponse, <-chan error) {
	responseChan := make(chan *QueryBookkeepingCommoditiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBookkeepingCommodities(request)
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

// QueryBookkeepingCommoditiesWithCallback invokes the companyreg.QueryBookkeepingCommodities API asynchronously
func (client *Client) QueryBookkeepingCommoditiesWithCallback(request *QueryBookkeepingCommoditiesRequest, callback func(response *QueryBookkeepingCommoditiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBookkeepingCommoditiesResponse
		var err error
		defer close(result)
		response, err = client.QueryBookkeepingCommodities(request)
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

// QueryBookkeepingCommoditiesRequest is the request struct for api QueryBookkeepingCommodities
type QueryBookkeepingCommoditiesRequest struct {
	*requests.RpcRequest
	ServiceType string `position:"Query" name:"ServiceType"`
	City        string `position:"Query" name:"City"`
	CompanyType string `position:"Query" name:"CompanyType"`
	AreaType    string `position:"Query" name:"AreaType"`
}

// QueryBookkeepingCommoditiesResponse is the response struct for api QueryBookkeepingCommodities
type QueryBookkeepingCommoditiesResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      []BookkeepingCommodity `json:"Data" xml:"Data"`
}

// CreateQueryBookkeepingCommoditiesRequest creates a request to invoke QueryBookkeepingCommodities API
func CreateQueryBookkeepingCommoditiesRequest() (request *QueryBookkeepingCommoditiesRequest) {
	request = &QueryBookkeepingCommoditiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "QueryBookkeepingCommodities", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryBookkeepingCommoditiesResponse creates a response to parse from QueryBookkeepingCommodities response
func CreateQueryBookkeepingCommoditiesResponse() (response *QueryBookkeepingCommoditiesResponse) {
	response = &QueryBookkeepingCommoditiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
