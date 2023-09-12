package dataworks_public

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

// GetDataServiceFolder invokes the dataworks_public.GetDataServiceFolder API synchronously
func (client *Client) GetDataServiceFolder(request *GetDataServiceFolderRequest) (response *GetDataServiceFolderResponse, err error) {
	response = CreateGetDataServiceFolderResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataServiceFolderWithChan invokes the dataworks_public.GetDataServiceFolder API asynchronously
func (client *Client) GetDataServiceFolderWithChan(request *GetDataServiceFolderRequest) (<-chan *GetDataServiceFolderResponse, <-chan error) {
	responseChan := make(chan *GetDataServiceFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataServiceFolder(request)
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

// GetDataServiceFolderWithCallback invokes the dataworks_public.GetDataServiceFolder API asynchronously
func (client *Client) GetDataServiceFolderWithCallback(request *GetDataServiceFolderRequest, callback func(response *GetDataServiceFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataServiceFolderResponse
		var err error
		defer close(result)
		response, err = client.GetDataServiceFolder(request)
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

// GetDataServiceFolderRequest is the request struct for api GetDataServiceFolder
type GetDataServiceFolderRequest struct {
	*requests.RpcRequest
	TenantId  requests.Integer `position:"Body" name:"TenantId"`
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
	FolderId  requests.Integer `position:"Body" name:"FolderId"`
}

// GetDataServiceFolderResponse is the response struct for api GetDataServiceFolder
type GetDataServiceFolderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Folder    Folder `json:"Folder" xml:"Folder"`
}

// CreateGetDataServiceFolderRequest creates a request to invoke GetDataServiceFolder API
func CreateGetDataServiceFolderRequest() (request *GetDataServiceFolderRequest) {
	request = &GetDataServiceFolderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetDataServiceFolder", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDataServiceFolderResponse creates a response to parse from GetDataServiceFolder response
func CreateGetDataServiceFolderResponse() (response *GetDataServiceFolderResponse) {
	response = &GetDataServiceFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
