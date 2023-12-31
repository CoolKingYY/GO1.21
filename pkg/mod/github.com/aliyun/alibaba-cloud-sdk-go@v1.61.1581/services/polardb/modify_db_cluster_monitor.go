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

// ModifyDBClusterMonitor invokes the polardb.ModifyDBClusterMonitor API synchronously
func (client *Client) ModifyDBClusterMonitor(request *ModifyDBClusterMonitorRequest) (response *ModifyDBClusterMonitorResponse, err error) {
	response = CreateModifyDBClusterMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterMonitorWithChan invokes the polardb.ModifyDBClusterMonitor API asynchronously
func (client *Client) ModifyDBClusterMonitorWithChan(request *ModifyDBClusterMonitorRequest) (<-chan *ModifyDBClusterMonitorResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterMonitor(request)
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

// ModifyDBClusterMonitorWithCallback invokes the polardb.ModifyDBClusterMonitor API asynchronously
func (client *Client) ModifyDBClusterMonitorWithCallback(request *ModifyDBClusterMonitorRequest, callback func(response *ModifyDBClusterMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterMonitorResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterMonitor(request)
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

// ModifyDBClusterMonitorRequest is the request struct for api ModifyDBClusterMonitor
type ModifyDBClusterMonitorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Period               string           `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBClusterMonitorResponse is the response struct for api ModifyDBClusterMonitor
type ModifyDBClusterMonitorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterMonitorRequest creates a request to invoke ModifyDBClusterMonitor API
func CreateModifyDBClusterMonitorRequest() (request *ModifyDBClusterMonitorRequest) {
	request = &ModifyDBClusterMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterMonitor", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterMonitorResponse creates a response to parse from ModifyDBClusterMonitor response
func CreateModifyDBClusterMonitorResponse() (response *ModifyDBClusterMonitorResponse) {
	response = &ModifyDBClusterMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
