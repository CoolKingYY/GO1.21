package rds

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

// DescribeDBInstanceSSL invokes the rds.DescribeDBInstanceSSL API synchronously
func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (response *DescribeDBInstanceSSLResponse, err error) {
	response = CreateDescribeDBInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceSSLWithChan invokes the rds.DescribeDBInstanceSSL API asynchronously
func (client *Client) DescribeDBInstanceSSLWithChan(request *DescribeDBInstanceSSLRequest) (<-chan *DescribeDBInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceSSL(request)
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

// DescribeDBInstanceSSLWithCallback invokes the rds.DescribeDBInstanceSSL API asynchronously
func (client *Client) DescribeDBInstanceSSLWithCallback(request *DescribeDBInstanceSSLRequest, callback func(response *DescribeDBInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceSSL(request)
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

// DescribeDBInstanceSSLRequest is the request struct for api DescribeDBInstanceSSL
type DescribeDBInstanceSSLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceSSLResponse is the response struct for api DescribeDBInstanceSSL
type DescribeDBInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	ConnectionString         string `json:"ConnectionString" xml:"ConnectionString"`
	SSLExpireTime            string `json:"SSLExpireTime" xml:"SSLExpireTime"`
	SSLEnabled               string `json:"SSLEnabled" xml:"SSLEnabled"`
	RequireUpdateReason      string `json:"RequireUpdateReason" xml:"RequireUpdateReason"`
	CAType                   string `json:"CAType" xml:"CAType"`
	ServerCert               string `json:"ServerCert" xml:"ServerCert"`
	ServerKey                string `json:"ServerKey" xml:"ServerKey"`
	SSLCreateTime            string `json:"SSLCreateTime" xml:"SSLCreateTime"`
	ClientCACert             string `json:"ClientCACert" xml:"ClientCACert"`
	ClientCACertExpireTime   string `json:"ClientCACertExpireTime" xml:"ClientCACertExpireTime"`
	ClientCertRevocationList string `json:"ClientCertRevocationList" xml:"ClientCertRevocationList"`
	ACL                      string `json:"ACL" xml:"ACL"`
	ReplicationACL           string `json:"ReplicationACL" xml:"ReplicationACL"`
	ServerCAUrl              string `json:"ServerCAUrl" xml:"ServerCAUrl"`
	RequireUpdate            string `json:"RequireUpdate" xml:"RequireUpdate"`
	RequireUpdateItem        string `json:"RequireUpdateItem" xml:"RequireUpdateItem"`
	LastModifyStatus         string `json:"LastModifyStatus" xml:"LastModifyStatus"`
	ModifyStatusReason       string `json:"ModifyStatusReason" xml:"ModifyStatusReason"`
}

// CreateDescribeDBInstanceSSLRequest creates a request to invoke DescribeDBInstanceSSL API
func CreateDescribeDBInstanceSSLRequest() (request *DescribeDBInstanceSSLRequest) {
	request = &DescribeDBInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceSSL", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceSSLResponse creates a response to parse from DescribeDBInstanceSSL response
func CreateDescribeDBInstanceSSLResponse() (response *DescribeDBInstanceSSLResponse) {
	response = &DescribeDBInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
