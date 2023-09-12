package dms_enterprise

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

// GetDBTopology invokes the dms_enterprise.GetDBTopology API synchronously
func (client *Client) GetDBTopology(request *GetDBTopologyRequest) (response *GetDBTopologyResponse, err error) {
	response = CreateGetDBTopologyResponse()
	err = client.DoAction(request, response)
	return
}

// GetDBTopologyWithChan invokes the dms_enterprise.GetDBTopology API asynchronously
func (client *Client) GetDBTopologyWithChan(request *GetDBTopologyRequest) (<-chan *GetDBTopologyResponse, <-chan error) {
	responseChan := make(chan *GetDBTopologyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDBTopology(request)
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

// GetDBTopologyWithCallback invokes the dms_enterprise.GetDBTopology API asynchronously
func (client *Client) GetDBTopologyWithCallback(request *GetDBTopologyRequest, callback func(response *GetDBTopologyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDBTopologyResponse
		var err error
		defer close(result)
		response, err = client.GetDBTopology(request)
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

// GetDBTopologyRequest is the request struct for api GetDBTopology
type GetDBTopologyRequest struct {
	*requests.RpcRequest
	LogicDbId requests.Integer `position:"Query" name:"LogicDbId"`
	Tid       requests.Integer `position:"Query" name:"Tid"`
}

// GetDBTopologyResponse is the response struct for api GetDBTopology
type GetDBTopologyResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	Success      bool       `json:"Success" xml:"Success"`
	ErrorMessage string     `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string     `json:"ErrorCode" xml:"ErrorCode"`
	DBTopology   DBTopology `json:"DBTopology" xml:"DBTopology"`
}

// CreateGetDBTopologyRequest creates a request to invoke GetDBTopology API
func CreateGetDBTopologyRequest() (request *GetDBTopologyRequest) {
	request = &GetDBTopologyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDBTopology", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDBTopologyResponse creates a response to parse from GetDBTopology response
func CreateGetDBTopologyResponse() (response *GetDBTopologyResponse) {
	response = &GetDBTopologyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}