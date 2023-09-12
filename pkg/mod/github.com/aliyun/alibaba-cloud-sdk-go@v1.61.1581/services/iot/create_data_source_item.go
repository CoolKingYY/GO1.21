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

// CreateDataSourceItem invokes the iot.CreateDataSourceItem API synchronously
func (client *Client) CreateDataSourceItem(request *CreateDataSourceItemRequest) (response *CreateDataSourceItemResponse, err error) {
	response = CreateCreateDataSourceItemResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataSourceItemWithChan invokes the iot.CreateDataSourceItem API asynchronously
func (client *Client) CreateDataSourceItemWithChan(request *CreateDataSourceItemRequest) (<-chan *CreateDataSourceItemResponse, <-chan error) {
	responseChan := make(chan *CreateDataSourceItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataSourceItem(request)
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

// CreateDataSourceItemWithCallback invokes the iot.CreateDataSourceItem API asynchronously
func (client *Client) CreateDataSourceItemWithCallback(request *CreateDataSourceItemRequest, callback func(response *CreateDataSourceItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataSourceItemResponse
		var err error
		defer close(result)
		response, err = client.CreateDataSourceItem(request)
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

// CreateDataSourceItemRequest is the request struct for api CreateDataSourceItem
type CreateDataSourceItemRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	DataSourceId  requests.Integer `position:"Query" name:"DataSourceId"`
	Topic         string           `position:"Query" name:"Topic"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// CreateDataSourceItemResponse is the response struct for api CreateDataSourceItem
type CreateDataSourceItemResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCreateDataSourceItemRequest creates a request to invoke CreateDataSourceItem API
func CreateCreateDataSourceItemRequest() (request *CreateDataSourceItemRequest) {
	request = &CreateDataSourceItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateDataSourceItem", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDataSourceItemResponse creates a response to parse from CreateDataSourceItem response
func CreateCreateDataSourceItemResponse() (response *CreateDataSourceItemResponse) {
	response = &CreateDataSourceItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}