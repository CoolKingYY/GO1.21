package dataworks_public

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

// TestNetworkConnection invokes the dataworks_public.TestNetworkConnection API synchronously
func (client *Client) TestNetworkConnection(request *TestNetworkConnectionRequest) (response *TestNetworkConnectionResponse, err error) {
	response = CreateTestNetworkConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// TestNetworkConnectionWithChan invokes the dataworks_public.TestNetworkConnection API asynchronously
func (client *Client) TestNetworkConnectionWithChan(request *TestNetworkConnectionRequest) (<-chan *TestNetworkConnectionResponse, <-chan error) {
	responseChan := make(chan *TestNetworkConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestNetworkConnection(request)
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

// TestNetworkConnectionWithCallback invokes the dataworks_public.TestNetworkConnection API asynchronously
func (client *Client) TestNetworkConnectionWithCallback(request *TestNetworkConnectionRequest, callback func(response *TestNetworkConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestNetworkConnectionResponse
		var err error
		defer close(result)
		response, err = client.TestNetworkConnection(request)
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

// TestNetworkConnectionRequest is the request struct for api TestNetworkConnection
type TestNetworkConnectionRequest struct {
	*requests.RpcRequest
	ResourceGroup  string           `position:"Query" name:"ResourceGroup"`
	EnvType        string           `position:"Query" name:"EnvType"`
	DatasourceName string           `position:"Query" name:"DatasourceName"`
	ProjectId      requests.Integer `position:"Query" name:"ProjectId"`
}

// TestNetworkConnectionResponse is the response struct for api TestNetworkConnection
type TestNetworkConnectionResponse struct {
	*responses.BaseResponse
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	TaskList  TaskList `json:"TaskList" xml:"TaskList"`
}

// CreateTestNetworkConnectionRequest creates a request to invoke TestNetworkConnection API
func CreateTestNetworkConnectionRequest() (request *TestNetworkConnectionRequest) {
	request = &TestNetworkConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "TestNetworkConnection", "", "")
	request.Method = requests.POST
	return
}

// CreateTestNetworkConnectionResponse creates a response to parse from TestNetworkConnection response
func CreateTestNetworkConnectionResponse() (response *TestNetworkConnectionResponse) {
	response = &TestNetworkConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
