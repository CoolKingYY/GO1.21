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

// DescribeSchemas invokes the clickhouse.DescribeSchemas API synchronously
func (client *Client) DescribeSchemas(request *DescribeSchemasRequest) (response *DescribeSchemasResponse, err error) {
	response = CreateDescribeSchemasResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSchemasWithChan invokes the clickhouse.DescribeSchemas API asynchronously
func (client *Client) DescribeSchemasWithChan(request *DescribeSchemasRequest) (<-chan *DescribeSchemasResponse, <-chan error) {
	responseChan := make(chan *DescribeSchemasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSchemas(request)
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

// DescribeSchemasWithCallback invokes the clickhouse.DescribeSchemas API asynchronously
func (client *Client) DescribeSchemasWithCallback(request *DescribeSchemasRequest, callback func(response *DescribeSchemasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSchemasResponse
		var err error
		defer close(result)
		response, err = client.DescribeSchemas(request)
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

// DescribeSchemasRequest is the request struct for api DescribeSchemas
type DescribeSchemasRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TableName            string           `position:"Query" name:"TableName"`
	SchemaName           string           `position:"Query" name:"SchemaName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSchemasResponse is the response struct for api DescribeSchemas
type DescribeSchemasResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeSchemas `json:"Items" xml:"Items"`
}

// CreateDescribeSchemasRequest creates a request to invoke DescribeSchemas API
func CreateDescribeSchemasRequest() (request *DescribeSchemasRequest) {
	request = &DescribeSchemasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeSchemas", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSchemasResponse creates a response to parse from DescribeSchemas response
func CreateDescribeSchemasResponse() (response *DescribeSchemasResponse) {
	response = &DescribeSchemasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
