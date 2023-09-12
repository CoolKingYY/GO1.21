package iot

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

// CreateOTADynamicUpgradeJob invokes the iot.CreateOTADynamicUpgradeJob API synchronously
func (client *Client) CreateOTADynamicUpgradeJob(request *CreateOTADynamicUpgradeJobRequest) (response *CreateOTADynamicUpgradeJobResponse, err error) {
	response = CreateCreateOTADynamicUpgradeJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOTADynamicUpgradeJobWithChan invokes the iot.CreateOTADynamicUpgradeJob API asynchronously
func (client *Client) CreateOTADynamicUpgradeJobWithChan(request *CreateOTADynamicUpgradeJobRequest) (<-chan *CreateOTADynamicUpgradeJobResponse, <-chan error) {
	responseChan := make(chan *CreateOTADynamicUpgradeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOTADynamicUpgradeJob(request)
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

// CreateOTADynamicUpgradeJobWithCallback invokes the iot.CreateOTADynamicUpgradeJob API asynchronously
func (client *Client) CreateOTADynamicUpgradeJobWithCallback(request *CreateOTADynamicUpgradeJobRequest, callback func(response *CreateOTADynamicUpgradeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOTADynamicUpgradeJobResponse
		var err error
		defer close(result)
		response, err = client.CreateOTADynamicUpgradeJob(request)
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

// CreateOTADynamicUpgradeJobRequest is the request struct for api CreateOTADynamicUpgradeJob
type CreateOTADynamicUpgradeJobRequest struct {
	*requests.RpcRequest
	DynamicMode      requests.Integer                 `position:"Query" name:"DynamicMode"`
	MultiModuleMode  requests.Boolean                 `position:"Query" name:"MultiModuleMode"`
	RetryCount       requests.Integer                 `position:"Query" name:"RetryCount"`
	TimeoutInMinutes requests.Integer                 `position:"Query" name:"TimeoutInMinutes"`
	NeedConfirm      requests.Boolean                 `position:"Query" name:"NeedConfirm"`
	GroupType        string                           `position:"Query" name:"GroupType"`
	NeedPush         requests.Boolean                 `position:"Query" name:"NeedPush"`
	IotInstanceId    string                           `position:"Query" name:"IotInstanceId"`
	DownloadProtocol string                           `position:"Query" name:"DownloadProtocol"`
	Tag              *[]CreateOTADynamicUpgradeJobTag `position:"Query" name:"Tag"  type:"Repeated"`
	GroupId          string                           `position:"Query" name:"GroupId"`
	FirmwareId       string                           `position:"Query" name:"FirmwareId"`
	ProductKey       string                           `position:"Query" name:"ProductKey"`
	RetryInterval    requests.Integer                 `position:"Query" name:"RetryInterval"`
	SrcVersion       *[]string                        `position:"Query" name:"SrcVersion"  type:"Repeated"`
	OverwriteMode    requests.Integer                 `position:"Query" name:"OverwriteMode"`
	ApiProduct       string                           `position:"Body" name:"ApiProduct"`
	ApiRevision      string                           `position:"Body" name:"ApiRevision"`
	MaximumPerMinute requests.Integer                 `position:"Query" name:"MaximumPerMinute"`
}

// CreateOTADynamicUpgradeJobTag is a repeated param struct in CreateOTADynamicUpgradeJobRequest
type CreateOTADynamicUpgradeJobTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateOTADynamicUpgradeJobResponse is the response struct for api CreateOTADynamicUpgradeJob
type CreateOTADynamicUpgradeJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateOTADynamicUpgradeJobRequest creates a request to invoke CreateOTADynamicUpgradeJob API
func CreateCreateOTADynamicUpgradeJobRequest() (request *CreateOTADynamicUpgradeJobRequest) {
	request = &CreateOTADynamicUpgradeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateOTADynamicUpgradeJob", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateOTADynamicUpgradeJobResponse creates a response to parse from CreateOTADynamicUpgradeJob response
func CreateCreateOTADynamicUpgradeJobResponse() (response *CreateOTADynamicUpgradeJobResponse) {
	response = &CreateOTADynamicUpgradeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}