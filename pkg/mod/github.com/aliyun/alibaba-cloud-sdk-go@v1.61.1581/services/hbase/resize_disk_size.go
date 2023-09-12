package hbase

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

// ResizeDiskSize invokes the hbase.ResizeDiskSize API synchronously
func (client *Client) ResizeDiskSize(request *ResizeDiskSizeRequest) (response *ResizeDiskSizeResponse, err error) {
	response = CreateResizeDiskSizeResponse()
	err = client.DoAction(request, response)
	return
}

// ResizeDiskSizeWithChan invokes the hbase.ResizeDiskSize API asynchronously
func (client *Client) ResizeDiskSizeWithChan(request *ResizeDiskSizeRequest) (<-chan *ResizeDiskSizeResponse, <-chan error) {
	responseChan := make(chan *ResizeDiskSizeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResizeDiskSize(request)
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

// ResizeDiskSizeWithCallback invokes the hbase.ResizeDiskSize API asynchronously
func (client *Client) ResizeDiskSizeWithCallback(request *ResizeDiskSizeRequest, callback func(response *ResizeDiskSizeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResizeDiskSizeResponse
		var err error
		defer close(result)
		response, err = client.ResizeDiskSize(request)
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

// ResizeDiskSizeRequest is the request struct for api ResizeDiskSize
type ResizeDiskSizeRequest struct {
	*requests.RpcRequest
	NodeDiskSize requests.Integer `position:"Query" name:"NodeDiskSize"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
}

// ResizeDiskSizeResponse is the response struct for api ResizeDiskSize
type ResizeDiskSizeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateResizeDiskSizeRequest creates a request to invoke ResizeDiskSize API
func CreateResizeDiskSizeRequest() (request *ResizeDiskSizeRequest) {
	request = &ResizeDiskSizeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "ResizeDiskSize", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResizeDiskSizeResponse creates a response to parse from ResizeDiskSize response
func CreateResizeDiskSizeResponse() (response *ResizeDiskSizeResponse) {
	response = &ResizeDiskSizeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
