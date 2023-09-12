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

// GetPredictiveJobGroup invokes the cloudcallcenter.GetPredictiveJobGroup API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getpredictivejobgroup.html
func (client *Client) GetPredictiveJobGroup(request *GetPredictiveJobGroupRequest) (response *GetPredictiveJobGroupResponse, err error) {
	response = CreateGetPredictiveJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetPredictiveJobGroupWithChan invokes the cloudcallcenter.GetPredictiveJobGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getpredictivejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPredictiveJobGroupWithChan(request *GetPredictiveJobGroupRequest) (<-chan *GetPredictiveJobGroupResponse, <-chan error) {
	responseChan := make(chan *GetPredictiveJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPredictiveJobGroup(request)
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

// GetPredictiveJobGroupWithCallback invokes the cloudcallcenter.GetPredictiveJobGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getpredictivejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPredictiveJobGroupWithCallback(request *GetPredictiveJobGroupRequest, callback func(response *GetPredictiveJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPredictiveJobGroupResponse
		var err error
		defer close(result)
		response, err = client.GetPredictiveJobGroup(request)
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

// GetPredictiveJobGroupRequest is the request struct for api GetPredictiveJobGroup
type GetPredictiveJobGroupRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	JobGroupId   string `position:"Query" name:"JobGroupId"`
}

// GetPredictiveJobGroupResponse is the response struct for api GetPredictiveJobGroup
type GetPredictiveJobGroupResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateGetPredictiveJobGroupRequest creates a request to invoke GetPredictiveJobGroup API
func CreateGetPredictiveJobGroupRequest() (request *GetPredictiveJobGroupRequest) {
	request = &GetPredictiveJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetPredictiveJobGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPredictiveJobGroupResponse creates a response to parse from GetPredictiveJobGroup response
func CreateGetPredictiveJobGroupResponse() (response *GetPredictiveJobGroupResponse) {
	response = &GetPredictiveJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
