package r_kvstore

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

// DeleteShardingNode invokes the r_kvstore.DeleteShardingNode API synchronously
func (client *Client) DeleteShardingNode(request *DeleteShardingNodeRequest) (response *DeleteShardingNodeResponse, err error) {
	response = CreateDeleteShardingNodeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteShardingNodeWithChan invokes the r_kvstore.DeleteShardingNode API asynchronously
func (client *Client) DeleteShardingNodeWithChan(request *DeleteShardingNodeRequest) (<-chan *DeleteShardingNodeResponse, <-chan error) {
	responseChan := make(chan *DeleteShardingNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteShardingNode(request)
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

// DeleteShardingNodeWithCallback invokes the r_kvstore.DeleteShardingNode API asynchronously
func (client *Client) DeleteShardingNodeWithCallback(request *DeleteShardingNodeRequest, callback func(response *DeleteShardingNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteShardingNodeResponse
		var err error
		defer close(result)
		response, err = client.DeleteShardingNode(request)
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

// DeleteShardingNodeRequest is the request struct for api DeleteShardingNode
type DeleteShardingNodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ShardCount           requests.Integer `position:"Query" name:"ShardCount"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeleteShardingNodeResponse is the response struct for api DeleteShardingNode
type DeleteShardingNodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteShardingNodeRequest creates a request to invoke DeleteShardingNode API
func CreateDeleteShardingNodeRequest() (request *DeleteShardingNodeRequest) {
	request = &DeleteShardingNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteShardingNode", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteShardingNodeResponse creates a response to parse from DeleteShardingNode response
func CreateDeleteShardingNodeResponse() (response *DeleteShardingNodeResponse) {
	response = &DeleteShardingNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
