package qualitycheck

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

// ListSkillGroupConfig invokes the qualitycheck.ListSkillGroupConfig API synchronously
func (client *Client) ListSkillGroupConfig(request *ListSkillGroupConfigRequest) (response *ListSkillGroupConfigResponse, err error) {
	response = CreateListSkillGroupConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupConfigWithChan invokes the qualitycheck.ListSkillGroupConfig API asynchronously
func (client *Client) ListSkillGroupConfigWithChan(request *ListSkillGroupConfigRequest) (<-chan *ListSkillGroupConfigResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroupConfig(request)
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

// ListSkillGroupConfigWithCallback invokes the qualitycheck.ListSkillGroupConfig API asynchronously
func (client *Client) ListSkillGroupConfigWithCallback(request *ListSkillGroupConfigRequest, callback func(response *ListSkillGroupConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupConfigResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroupConfig(request)
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

// ListSkillGroupConfigRequest is the request struct for api ListSkillGroupConfig
type ListSkillGroupConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// ListSkillGroupConfigResponse is the response struct for api ListSkillGroupConfig
type ListSkillGroupConfigResponse struct {
	*responses.BaseResponse
	Code      string                     `json:"Code" xml:"Code"`
	Message   string                     `json:"Message" xml:"Message"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Success   bool                       `json:"Success" xml:"Success"`
	Data      DataInListSkillGroupConfig `json:"Data" xml:"Data"`
}

// CreateListSkillGroupConfigRequest creates a request to invoke ListSkillGroupConfig API
func CreateListSkillGroupConfigRequest() (request *ListSkillGroupConfigRequest) {
	request = &ListSkillGroupConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "ListSkillGroupConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateListSkillGroupConfigResponse creates a response to parse from ListSkillGroupConfig response
func CreateListSkillGroupConfigResponse() (response *ListSkillGroupConfigResponse) {
	response = &ListSkillGroupConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
