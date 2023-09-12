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

// DeleteDBLink invokes the polardb.DeleteDBLink API synchronously
func (client *Client) DeleteDBLink(request *DeleteDBLinkRequest) (response *DeleteDBLinkResponse, err error) {
	response = CreateDeleteDBLinkResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDBLinkWithChan invokes the polardb.DeleteDBLink API asynchronously
func (client *Client) DeleteDBLinkWithChan(request *DeleteDBLinkRequest) (<-chan *DeleteDBLinkResponse, <-chan error) {
	responseChan := make(chan *DeleteDBLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDBLink(request)
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

// DeleteDBLinkWithCallback invokes the polardb.DeleteDBLink API asynchronously
func (client *Client) DeleteDBLinkWithCallback(request *DeleteDBLinkRequest, callback func(response *DeleteDBLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDBLinkResponse
		var err error
		defer close(result)
		response, err = client.DeleteDBLink(request)
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

// DeleteDBLinkRequest is the request struct for api DeleteDBLink
type DeleteDBLinkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBLinkName           string           `position:"Query" name:"DBLinkName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteDBLinkResponse is the response struct for api DeleteDBLink
type DeleteDBLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDBLinkRequest creates a request to invoke DeleteDBLink API
func CreateDeleteDBLinkRequest() (request *DeleteDBLinkRequest) {
	request = &DeleteDBLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DeleteDBLink", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDBLinkResponse creates a response to parse from DeleteDBLink response
func CreateDeleteDBLinkResponse() (response *DeleteDBLinkResponse) {
	response = &DeleteDBLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
