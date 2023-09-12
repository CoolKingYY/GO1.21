package vs

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

// DescribeCluster invokes the vs.DescribeCluster API synchronously
func (client *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	response = CreateDescribeClusterResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterWithChan invokes the vs.DescribeCluster API asynchronously
func (client *Client) DescribeClusterWithChan(request *DescribeClusterRequest) (<-chan *DescribeClusterResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCluster(request)
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

// DescribeClusterWithCallback invokes the vs.DescribeCluster API asynchronously
func (client *Client) DescribeClusterWithCallback(request *DescribeClusterRequest, callback func(response *DescribeClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterResponse
		var err error
		defer close(result)
		response, err = client.DescribeCluster(request)
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

// DescribeClusterRequest is the request struct for api DescribeCluster
type DescribeClusterRequest struct {
	*requests.RpcRequest
	ShowLog   string           `position:"Query" name:"ShowLog"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeClusterResponse is the response struct for api DescribeCluster
type DescribeClusterResponse struct {
	*responses.BaseResponse
	Status        string         `json:"Status" xml:"Status"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	Description   string         `json:"Description" xml:"Description"`
	MaintainTime  string         `json:"MaintainTime" xml:"MaintainTime"`
	Name          string         `json:"Name" xml:"Name"`
	ClusterId     string         `json:"ClusterId" xml:"ClusterId"`
	InternalPorts []InternalPort `json:"InternalPorts" xml:"InternalPorts"`
}

// CreateDescribeClusterRequest creates a request to invoke DescribeCluster API
func CreateDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterResponse creates a response to parse from DescribeCluster response
func CreateDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
