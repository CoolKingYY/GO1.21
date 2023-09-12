package live

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

// SetLiveStreamsNotifyUrlConfig invokes the live.SetLiveStreamsNotifyUrlConfig API synchronously
func (client *Client) SetLiveStreamsNotifyUrlConfig(request *SetLiveStreamsNotifyUrlConfigRequest) (response *SetLiveStreamsNotifyUrlConfigResponse, err error) {
	response = CreateSetLiveStreamsNotifyUrlConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetLiveStreamsNotifyUrlConfigWithChan invokes the live.SetLiveStreamsNotifyUrlConfig API asynchronously
func (client *Client) SetLiveStreamsNotifyUrlConfigWithChan(request *SetLiveStreamsNotifyUrlConfigRequest) (<-chan *SetLiveStreamsNotifyUrlConfigResponse, <-chan error) {
	responseChan := make(chan *SetLiveStreamsNotifyUrlConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLiveStreamsNotifyUrlConfig(request)
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

// SetLiveStreamsNotifyUrlConfigWithCallback invokes the live.SetLiveStreamsNotifyUrlConfig API asynchronously
func (client *Client) SetLiveStreamsNotifyUrlConfigWithCallback(request *SetLiveStreamsNotifyUrlConfigRequest, callback func(response *SetLiveStreamsNotifyUrlConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLiveStreamsNotifyUrlConfigResponse
		var err error
		defer close(result)
		response, err = client.SetLiveStreamsNotifyUrlConfig(request)
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

// SetLiveStreamsNotifyUrlConfigRequest is the request struct for api SetLiveStreamsNotifyUrlConfig
type SetLiveStreamsNotifyUrlConfigRequest struct {
	*requests.RpcRequest
	AuthKey       string           `position:"Query" name:"AuthKey"`
	AuthType      string           `position:"Query" name:"AuthType"`
	NotifyReqAuth string           `position:"Query" name:"NotifyReqAuth"`
	NotifyUrl     string           `position:"Query" name:"NotifyUrl"`
	NotifyType    string           `position:"Query" name:"NotifyType"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	NotifyAuthKey string           `position:"Query" name:"NotifyAuthKey"`
}

// SetLiveStreamsNotifyUrlConfigResponse is the response struct for api SetLiveStreamsNotifyUrlConfig
type SetLiveStreamsNotifyUrlConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLiveStreamsNotifyUrlConfigRequest creates a request to invoke SetLiveStreamsNotifyUrlConfig API
func CreateSetLiveStreamsNotifyUrlConfigRequest() (request *SetLiveStreamsNotifyUrlConfigRequest) {
	request = &SetLiveStreamsNotifyUrlConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetLiveStreamsNotifyUrlConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLiveStreamsNotifyUrlConfigResponse creates a response to parse from SetLiveStreamsNotifyUrlConfig response
func CreateSetLiveStreamsNotifyUrlConfigResponse() (response *SetLiveStreamsNotifyUrlConfigResponse) {
	response = &SetLiveStreamsNotifyUrlConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
