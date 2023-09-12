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

// DescribeGlobalDatabaseNetwork invokes the polardb.DescribeGlobalDatabaseNetwork API synchronously
func (client *Client) DescribeGlobalDatabaseNetwork(request *DescribeGlobalDatabaseNetworkRequest) (response *DescribeGlobalDatabaseNetworkResponse, err error) {
	response = CreateDescribeGlobalDatabaseNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGlobalDatabaseNetworkWithChan invokes the polardb.DescribeGlobalDatabaseNetwork API asynchronously
func (client *Client) DescribeGlobalDatabaseNetworkWithChan(request *DescribeGlobalDatabaseNetworkRequest) (<-chan *DescribeGlobalDatabaseNetworkResponse, <-chan error) {
	responseChan := make(chan *DescribeGlobalDatabaseNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGlobalDatabaseNetwork(request)
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

// DescribeGlobalDatabaseNetworkWithCallback invokes the polardb.DescribeGlobalDatabaseNetwork API asynchronously
func (client *Client) DescribeGlobalDatabaseNetworkWithCallback(request *DescribeGlobalDatabaseNetworkRequest, callback func(response *DescribeGlobalDatabaseNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGlobalDatabaseNetworkResponse
		var err error
		defer close(result)
		response, err = client.DescribeGlobalDatabaseNetwork(request)
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

// DescribeGlobalDatabaseNetworkRequest is the request struct for api DescribeGlobalDatabaseNetwork
type DescribeGlobalDatabaseNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	GDNId                string           `position:"Query" name:"GDNId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeGlobalDatabaseNetworkResponse is the response struct for api DescribeGlobalDatabaseNetwork
type DescribeGlobalDatabaseNetworkResponse struct {
	*responses.BaseResponse
	GDNStatus      string                                     `json:"GDNStatus" xml:"GDNStatus"`
	DBVersion      string                                     `json:"DBVersion" xml:"DBVersion"`
	RequestId      string                                     `json:"RequestId" xml:"RequestId"`
	GDNId          string                                     `json:"GDNId" xml:"GDNId"`
	CreateTime     string                                     `json:"CreateTime" xml:"CreateTime"`
	DBType         string                                     `json:"DBType" xml:"DBType"`
	GDNDescription string                                     `json:"GDNDescription" xml:"GDNDescription"`
	Connections    []Connection                               `json:"Connections" xml:"Connections"`
	DBClusters     []DBClusterInDescribeGlobalDatabaseNetwork `json:"DBClusters" xml:"DBClusters"`
}

// CreateDescribeGlobalDatabaseNetworkRequest creates a request to invoke DescribeGlobalDatabaseNetwork API
func CreateDescribeGlobalDatabaseNetworkRequest() (request *DescribeGlobalDatabaseNetworkRequest) {
	request = &DescribeGlobalDatabaseNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeGlobalDatabaseNetwork", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeGlobalDatabaseNetworkResponse creates a response to parse from DescribeGlobalDatabaseNetwork response
func CreateDescribeGlobalDatabaseNetworkResponse() (response *DescribeGlobalDatabaseNetworkResponse) {
	response = &DescribeGlobalDatabaseNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
