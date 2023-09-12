package cloudcallcenter

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

// DeleteContactFlow invokes the cloudcallcenter.DeleteContactFlow API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecontactflow.html
func (client *Client) DeleteContactFlow(request *DeleteContactFlowRequest) (response *DeleteContactFlowResponse, err error) {
	response = CreateDeleteContactFlowResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteContactFlowWithChan invokes the cloudcallcenter.DeleteContactFlow API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecontactflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteContactFlowWithChan(request *DeleteContactFlowRequest) (<-chan *DeleteContactFlowResponse, <-chan error) {
	responseChan := make(chan *DeleteContactFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteContactFlow(request)
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

// DeleteContactFlowWithCallback invokes the cloudcallcenter.DeleteContactFlow API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecontactflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteContactFlowWithCallback(request *DeleteContactFlowRequest, callback func(response *DeleteContactFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteContactFlowResponse
		var err error
		defer close(result)
		response, err = client.DeleteContactFlow(request)
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

// DeleteContactFlowRequest is the request struct for api DeleteContactFlow
type DeleteContactFlowRequest struct {
	*requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// DeleteContactFlowResponse is the response struct for api DeleteContactFlow
type DeleteContactFlowResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteContactFlowRequest creates a request to invoke DeleteContactFlow API
func CreateDeleteContactFlowRequest() (request *DeleteContactFlowRequest) {
	request = &DeleteContactFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteContactFlow", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteContactFlowResponse creates a response to parse from DeleteContactFlow response
func CreateDeleteContactFlowResponse() (response *DeleteContactFlowResponse) {
	response = &DeleteContactFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
