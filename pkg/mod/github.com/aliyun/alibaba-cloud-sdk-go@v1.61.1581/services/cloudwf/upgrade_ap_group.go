package cloudwf

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

// UpgradeAPGroup invokes the cloudwf.UpgradeAPGroup API synchronously
// api document: https://help.aliyun.com/api/cloudwf/upgradeapgroup.html
func (client *Client) UpgradeAPGroup(request *UpgradeAPGroupRequest) (response *UpgradeAPGroupResponse, err error) {
	response = CreateUpgradeAPGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeAPGroupWithChan invokes the cloudwf.UpgradeAPGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/upgradeapgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeAPGroupWithChan(request *UpgradeAPGroupRequest) (<-chan *UpgradeAPGroupResponse, <-chan error) {
	responseChan := make(chan *UpgradeAPGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeAPGroup(request)
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

// UpgradeAPGroupWithCallback invokes the cloudwf.UpgradeAPGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/upgradeapgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeAPGroupWithCallback(request *UpgradeAPGroupRequest, callback func(response *UpgradeAPGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeAPGroupResponse
		var err error
		defer close(result)
		response, err = client.UpgradeAPGroup(request)
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

// UpgradeAPGroupRequest is the request struct for api UpgradeAPGroup
type UpgradeAPGroupRequest struct {
	*requests.RpcRequest
	Ids *[]string `position:"Query" name:"Ids"  type:"Repeated"`
}

// UpgradeAPGroupResponse is the response struct for api UpgradeAPGroup
type UpgradeAPGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateUpgradeAPGroupRequest creates a request to invoke UpgradeAPGroup API
func CreateUpgradeAPGroupRequest() (request *UpgradeAPGroupRequest) {
	request = &UpgradeAPGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "UpgradeAPGroup", "cloudwf", "openAPI")
	return
}

// CreateUpgradeAPGroupResponse creates a response to parse from UpgradeAPGroup response
func CreateUpgradeAPGroupResponse() (response *UpgradeAPGroupResponse) {
	response = &UpgradeAPGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
