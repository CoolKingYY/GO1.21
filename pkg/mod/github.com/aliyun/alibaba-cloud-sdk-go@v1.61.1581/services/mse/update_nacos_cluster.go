package mse

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

// UpdateNacosCluster invokes the mse.UpdateNacosCluster API synchronously
func (client *Client) UpdateNacosCluster(request *UpdateNacosClusterRequest) (response *UpdateNacosClusterResponse, err error) {
	response = CreateUpdateNacosClusterResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNacosClusterWithChan invokes the mse.UpdateNacosCluster API asynchronously
func (client *Client) UpdateNacosClusterWithChan(request *UpdateNacosClusterRequest) (<-chan *UpdateNacosClusterResponse, <-chan error) {
	responseChan := make(chan *UpdateNacosClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNacosCluster(request)
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

// UpdateNacosClusterWithCallback invokes the mse.UpdateNacosCluster API asynchronously
func (client *Client) UpdateNacosClusterWithCallback(request *UpdateNacosClusterRequest, callback func(response *UpdateNacosClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNacosClusterResponse
		var err error
		defer close(result)
		response, err = client.UpdateNacosCluster(request)
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

// UpdateNacosClusterRequest is the request struct for api UpdateNacosCluster
type UpdateNacosClusterRequest struct {
	*requests.RpcRequest
	ClusterName             string           `position:"Query" name:"ClusterName"`
	CheckPort               requests.Integer `position:"Query" name:"CheckPort"`
	GroupName               string           `position:"Query" name:"GroupName"`
	InstanceId              string           `position:"Query" name:"InstanceId"`
	NamespaceId             string           `position:"Query" name:"NamespaceId"`
	HealthChecker           string           `position:"Query" name:"HealthChecker"`
	AcceptLanguage          string           `position:"Query" name:"AcceptLanguage"`
	ServiceName             string           `position:"Query" name:"ServiceName"`
	UseInstancePortForCheck requests.Boolean `position:"Query" name:"UseInstancePortForCheck"`
}

// UpdateNacosClusterResponse is the response struct for api UpdateNacosCluster
type UpdateNacosClusterResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateUpdateNacosClusterRequest creates a request to invoke UpdateNacosCluster API
func CreateUpdateNacosClusterRequest() (request *UpdateNacosClusterRequest) {
	request = &UpdateNacosClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateNacosCluster", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateNacosClusterResponse creates a response to parse from UpdateNacosCluster response
func CreateUpdateNacosClusterResponse() (response *UpdateNacosClusterResponse) {
	response = &UpdateNacosClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}