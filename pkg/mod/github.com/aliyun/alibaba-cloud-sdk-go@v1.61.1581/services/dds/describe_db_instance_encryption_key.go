package dds

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

// DescribeDBInstanceEncryptionKey invokes the dds.DescribeDBInstanceEncryptionKey API synchronously
func (client *Client) DescribeDBInstanceEncryptionKey(request *DescribeDBInstanceEncryptionKeyRequest) (response *DescribeDBInstanceEncryptionKeyResponse, err error) {
	response = CreateDescribeDBInstanceEncryptionKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceEncryptionKeyWithChan invokes the dds.DescribeDBInstanceEncryptionKey API asynchronously
func (client *Client) DescribeDBInstanceEncryptionKeyWithChan(request *DescribeDBInstanceEncryptionKeyRequest) (<-chan *DescribeDBInstanceEncryptionKeyResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceEncryptionKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceEncryptionKey(request)
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

// DescribeDBInstanceEncryptionKeyWithCallback invokes the dds.DescribeDBInstanceEncryptionKey API asynchronously
func (client *Client) DescribeDBInstanceEncryptionKeyWithCallback(request *DescribeDBInstanceEncryptionKeyRequest, callback func(response *DescribeDBInstanceEncryptionKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceEncryptionKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceEncryptionKey(request)
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

// DescribeDBInstanceEncryptionKeyRequest is the request struct for api DescribeDBInstanceEncryptionKey
type DescribeDBInstanceEncryptionKeyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EncryptionKey        string           `position:"Query" name:"EncryptionKey"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceEncryptionKeyResponse is the response struct for api DescribeDBInstanceEncryptionKey
type DescribeDBInstanceEncryptionKeyResponse struct {
	*responses.BaseResponse
	Origin              string `json:"Origin" xml:"Origin"`
	Description         string `json:"Description" xml:"Description"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
	EncryptionKeyStatus string `json:"EncryptionKeyStatus" xml:"EncryptionKeyStatus"`
	MaterialExpireTime  string `json:"MaterialExpireTime" xml:"MaterialExpireTime"`
	KeyUsage            string `json:"KeyUsage" xml:"KeyUsage"`
	EncryptionKey       string `json:"EncryptionKey" xml:"EncryptionKey"`
	Creator             string `json:"Creator" xml:"Creator"`
	DeleteDate          string `json:"DeleteDate" xml:"DeleteDate"`
}

// CreateDescribeDBInstanceEncryptionKeyRequest creates a request to invoke DescribeDBInstanceEncryptionKey API
func CreateDescribeDBInstanceEncryptionKeyRequest() (request *DescribeDBInstanceEncryptionKeyRequest) {
	request = &DescribeDBInstanceEncryptionKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstanceEncryptionKey", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceEncryptionKeyResponse creates a response to parse from DescribeDBInstanceEncryptionKey response
func CreateDescribeDBInstanceEncryptionKeyResponse() (response *DescribeDBInstanceEncryptionKeyResponse) {
	response = &DescribeDBInstanceEncryptionKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
