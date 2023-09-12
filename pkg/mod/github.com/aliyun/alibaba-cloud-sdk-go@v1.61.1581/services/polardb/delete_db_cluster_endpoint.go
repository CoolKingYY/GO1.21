package polardb

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

// DeleteDBClusterEndpoint invokes the polardb.DeleteDBClusterEndpoint API synchronously
func (client *Client) DeleteDBClusterEndpoint(request *DeleteDBClusterEndpointRequest) (response *DeleteDBClusterEndpointResponse, err error) {
	response = CreateDeleteDBClusterEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDBClusterEndpointWithChan invokes the polardb.DeleteDBClusterEndpoint API asynchronously
func (client *Client) DeleteDBClusterEndpointWithChan(request *DeleteDBClusterEndpointRequest) (<-chan *DeleteDBClusterEndpointResponse, <-chan error) {
	responseChan := make(chan *DeleteDBClusterEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDBClusterEndpoint(request)
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

// DeleteDBClusterEndpointWithCallback invokes the polardb.DeleteDBClusterEndpoint API asynchronously
func (client *Client) DeleteDBClusterEndpointWithCallback(request *DeleteDBClusterEndpointRequest, callback func(response *DeleteDBClusterEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDBClusterEndpointResponse
		var err error
		defer close(result)
		response, err = client.DeleteDBClusterEndpoint(request)
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

// DeleteDBClusterEndpointRequest is the request struct for api DeleteDBClusterEndpoint
type DeleteDBClusterEndpointRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBEndpointId         string           `position:"Query" name:"DBEndpointId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteDBClusterEndpointResponse is the response struct for api DeleteDBClusterEndpoint
type DeleteDBClusterEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDBClusterEndpointRequest creates a request to invoke DeleteDBClusterEndpoint API
func CreateDeleteDBClusterEndpointRequest() (request *DeleteDBClusterEndpointRequest) {
	request = &DeleteDBClusterEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DeleteDBClusterEndpoint", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDBClusterEndpointResponse creates a response to parse from DeleteDBClusterEndpoint response
func CreateDeleteDBClusterEndpointResponse() (response *DeleteDBClusterEndpointResponse) {
	response = &DeleteDBClusterEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}