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

// CreateDBLink invokes the polardb.CreateDBLink API synchronously
func (client *Client) CreateDBLink(request *CreateDBLinkRequest) (response *CreateDBLinkResponse, err error) {
	response = CreateCreateDBLinkResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBLinkWithChan invokes the polardb.CreateDBLink API asynchronously
func (client *Client) CreateDBLinkWithChan(request *CreateDBLinkRequest) (<-chan *CreateDBLinkResponse, <-chan error) {
	responseChan := make(chan *CreateDBLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBLink(request)
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

// CreateDBLinkWithCallback invokes the polardb.CreateDBLink API asynchronously
func (client *Client) CreateDBLinkWithCallback(request *CreateDBLinkRequest, callback func(response *CreateDBLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBLinkResponse
		var err error
		defer close(result)
		response, err = client.CreateDBLink(request)
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

// CreateDBLinkRequest is the request struct for api CreateDBLink
type CreateDBLinkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceDBName         string           `position:"Query" name:"SourceDBName"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	TargetDBName         string           `position:"Query" name:"TargetDBName"`
	TargetIp             string           `position:"Query" name:"TargetIp"`
	DBLinkName           string           `position:"Query" name:"DBLinkName"`
	TargetPort           string           `position:"Query" name:"TargetPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TargetDBInstanceName string           `position:"Query" name:"TargetDBInstanceName"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	TargetDBPasswd       string           `position:"Query" name:"TargetDBPasswd"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetDBAccount      string           `position:"Query" name:"TargetDBAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// CreateDBLinkResponse is the response struct for api CreateDBLink
type CreateDBLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDBLinkRequest creates a request to invoke CreateDBLink API
func CreateCreateDBLinkRequest() (request *CreateDBLinkRequest) {
	request = &CreateDBLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateDBLink", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDBLinkResponse creates a response to parse from CreateDBLink response
func CreateCreateDBLinkResponse() (response *CreateDBLinkResponse) {
	response = &CreateDBLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
