package dyplsapi

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

// DeleteSecretBlacklist invokes the dyplsapi.DeleteSecretBlacklist API synchronously
func (client *Client) DeleteSecretBlacklist(request *DeleteSecretBlacklistRequest) (response *DeleteSecretBlacklistResponse, err error) {
	response = CreateDeleteSecretBlacklistResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSecretBlacklistWithChan invokes the dyplsapi.DeleteSecretBlacklist API asynchronously
func (client *Client) DeleteSecretBlacklistWithChan(request *DeleteSecretBlacklistRequest) (<-chan *DeleteSecretBlacklistResponse, <-chan error) {
	responseChan := make(chan *DeleteSecretBlacklistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSecretBlacklist(request)
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

// DeleteSecretBlacklistWithCallback invokes the dyplsapi.DeleteSecretBlacklist API asynchronously
func (client *Client) DeleteSecretBlacklistWithCallback(request *DeleteSecretBlacklistRequest, callback func(response *DeleteSecretBlacklistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSecretBlacklistResponse
		var err error
		defer close(result)
		response, err = client.DeleteSecretBlacklist(request)
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

// DeleteSecretBlacklistRequest is the request struct for api DeleteSecretBlacklist
type DeleteSecretBlacklistRequest struct {
	*requests.RpcRequest
	BlackType            string           `position:"Query" name:"BlackType"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Remark               string           `position:"Query" name:"Remark"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
	BlackNo              string           `position:"Query" name:"BlackNo"`
	WayControl           string           `position:"Query" name:"WayControl"`
}

// DeleteSecretBlacklistResponse is the response struct for api DeleteSecretBlacklist
type DeleteSecretBlacklistResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSecretBlacklistRequest creates a request to invoke DeleteSecretBlacklist API
func CreateDeleteSecretBlacklistRequest() (request *DeleteSecretBlacklistRequest) {
	request = &DeleteSecretBlacklistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "DeleteSecretBlacklist", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSecretBlacklistResponse creates a response to parse from DeleteSecretBlacklist response
func CreateDeleteSecretBlacklistResponse() (response *DeleteSecretBlacklistResponse) {
	response = &DeleteSecretBlacklistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}