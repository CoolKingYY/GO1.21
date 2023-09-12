package clickhouse

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

// DescribeProcessList invokes the clickhouse.DescribeProcessList API synchronously
func (client *Client) DescribeProcessList(request *DescribeProcessListRequest) (response *DescribeProcessListResponse, err error) {
	response = CreateDescribeProcessListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProcessListWithChan invokes the clickhouse.DescribeProcessList API asynchronously
func (client *Client) DescribeProcessListWithChan(request *DescribeProcessListRequest) (<-chan *DescribeProcessListResponse, <-chan error) {
	responseChan := make(chan *DescribeProcessListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProcessList(request)
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

// DescribeProcessListWithCallback invokes the clickhouse.DescribeProcessList API asynchronously
func (client *Client) DescribeProcessListWithCallback(request *DescribeProcessListRequest, callback func(response *DescribeProcessListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProcessListResponse
		var err error
		defer close(result)
		response, err = client.DescribeProcessList(request)
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

// DescribeProcessListRequest is the request struct for api DescribeProcessList
type DescribeProcessListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InitialQueryId       string           `position:"Query" name:"InitialQueryId"`
	QueryDurationMs      requests.Integer `position:"Query" name:"QueryDurationMs"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InitialUser          string           `position:"Query" name:"InitialUser"`
	Keyword              string           `position:"Query" name:"Keyword"`
	Order                string           `position:"Query" name:"Order"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeProcessListResponse is the response struct for api DescribeProcessList
type DescribeProcessListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ProcessList ProcessList `json:"ProcessList" xml:"ProcessList"`
}

// CreateDescribeProcessListRequest creates a request to invoke DescribeProcessList API
func CreateDescribeProcessListRequest() (request *DescribeProcessListRequest) {
	request = &DescribeProcessListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeProcessList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeProcessListResponse creates a response to parse from DescribeProcessList response
func CreateDescribeProcessListResponse() (response *DescribeProcessListResponse) {
	response = &DescribeProcessListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
