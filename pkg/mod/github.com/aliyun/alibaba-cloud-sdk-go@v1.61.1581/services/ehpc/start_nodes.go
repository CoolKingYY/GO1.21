package ehpc

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

// StartNodes invokes the ehpc.StartNodes API synchronously
func (client *Client) StartNodes(request *StartNodesRequest) (response *StartNodesResponse, err error) {
	response = CreateStartNodesResponse()
	err = client.DoAction(request, response)
	return
}

// StartNodesWithChan invokes the ehpc.StartNodes API asynchronously
func (client *Client) StartNodesWithChan(request *StartNodesRequest) (<-chan *StartNodesResponse, <-chan error) {
	responseChan := make(chan *StartNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartNodes(request)
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

// StartNodesWithCallback invokes the ehpc.StartNodes API asynchronously
func (client *Client) StartNodesWithCallback(request *StartNodesRequest, callback func(response *StartNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartNodesResponse
		var err error
		defer close(result)
		response, err = client.StartNodes(request)
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

// StartNodesRequest is the request struct for api StartNodes
type StartNodesRequest struct {
	*requests.RpcRequest
	Role      string                `position:"Query" name:"Role"`
	Instance  *[]StartNodesInstance `position:"Query" name:"Instance"  type:"Repeated"`
	ClusterId string                `position:"Query" name:"ClusterId"`
}

// StartNodesInstance is a repeated param struct in StartNodesRequest
type StartNodesInstance struct {
	Id string `name:"Id"`
}

// StartNodesResponse is the response struct for api StartNodes
type StartNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateStartNodesRequest creates a request to invoke StartNodes API
func CreateStartNodesRequest() (request *StartNodesRequest) {
	request = &StartNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StartNodes", "", "")
	request.Method = requests.GET
	return
}

// CreateStartNodesResponse creates a response to parse from StartNodes response
func CreateStartNodesResponse() (response *StartNodesResponse) {
	response = &StartNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}