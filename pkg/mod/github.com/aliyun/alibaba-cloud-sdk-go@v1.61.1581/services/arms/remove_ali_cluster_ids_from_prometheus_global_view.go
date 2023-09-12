package arms

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

// RemoveAliClusterIdsFromPrometheusGlobalView invokes the arms.RemoveAliClusterIdsFromPrometheusGlobalView API synchronously
func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalView(request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) (response *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, err error) {
	response = CreateRemoveAliClusterIdsFromPrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveAliClusterIdsFromPrometheusGlobalViewWithChan invokes the arms.RemoveAliClusterIdsFromPrometheusGlobalView API asynchronously
func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalViewWithChan(request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) (<-chan *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveAliClusterIdsFromPrometheusGlobalView(request)
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

// RemoveAliClusterIdsFromPrometheusGlobalViewWithCallback invokes the arms.RemoveAliClusterIdsFromPrometheusGlobalView API asynchronously
func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalViewWithCallback(request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest, callback func(response *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveAliClusterIdsFromPrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.RemoveAliClusterIdsFromPrometheusGlobalView(request)
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

// RemoveAliClusterIdsFromPrometheusGlobalViewRequest is the request struct for api RemoveAliClusterIdsFromPrometheusGlobalView
type RemoveAliClusterIdsFromPrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	GlobalViewClusterId string `position:"Query" name:"GlobalViewClusterId"`
	ClusterIds          string `position:"Query" name:"ClusterIds"`
	GroupName           string `position:"Query" name:"GroupName"`
}

// RemoveAliClusterIdsFromPrometheusGlobalViewResponse is the response struct for api RemoveAliClusterIdsFromPrometheusGlobalView
type RemoveAliClusterIdsFromPrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRemoveAliClusterIdsFromPrometheusGlobalViewRequest creates a request to invoke RemoveAliClusterIdsFromPrometheusGlobalView API
func CreateRemoveAliClusterIdsFromPrometheusGlobalViewRequest() (request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) {
	request = &RemoveAliClusterIdsFromPrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "RemoveAliClusterIdsFromPrometheusGlobalView", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveAliClusterIdsFromPrometheusGlobalViewResponse creates a response to parse from RemoveAliClusterIdsFromPrometheusGlobalView response
func CreateRemoveAliClusterIdsFromPrometheusGlobalViewResponse() (response *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) {
	response = &RemoveAliClusterIdsFromPrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
