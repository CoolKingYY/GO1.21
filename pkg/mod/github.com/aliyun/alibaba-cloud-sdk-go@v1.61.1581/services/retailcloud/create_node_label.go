package retailcloud

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

// CreateNodeLabel invokes the retailcloud.CreateNodeLabel API synchronously
func (client *Client) CreateNodeLabel(request *CreateNodeLabelRequest) (response *CreateNodeLabelResponse, err error) {
	response = CreateCreateNodeLabelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNodeLabelWithChan invokes the retailcloud.CreateNodeLabel API asynchronously
func (client *Client) CreateNodeLabelWithChan(request *CreateNodeLabelRequest) (<-chan *CreateNodeLabelResponse, <-chan error) {
	responseChan := make(chan *CreateNodeLabelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNodeLabel(request)
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

// CreateNodeLabelWithCallback invokes the retailcloud.CreateNodeLabel API asynchronously
func (client *Client) CreateNodeLabelWithCallback(request *CreateNodeLabelRequest, callback func(response *CreateNodeLabelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNodeLabelResponse
		var err error
		defer close(result)
		response, err = client.CreateNodeLabel(request)
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

// CreateNodeLabelRequest is the request struct for api CreateNodeLabel
type CreateNodeLabelRequest struct {
	*requests.RpcRequest
	LabelKey   string `position:"Query" name:"LabelKey"`
	LabelValue string `position:"Query" name:"LabelValue"`
	ClusterId  string `position:"Query" name:"ClusterId"`
}

// CreateNodeLabelResponse is the response struct for api CreateNodeLabel
type CreateNodeLabelResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateNodeLabelRequest creates a request to invoke CreateNodeLabel API
func CreateCreateNodeLabelRequest() (request *CreateNodeLabelRequest) {
	request = &CreateNodeLabelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "CreateNodeLabel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateNodeLabelResponse creates a response to parse from CreateNodeLabel response
func CreateCreateNodeLabelResponse() (response *CreateNodeLabelResponse) {
	response = &CreateNodeLabelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
