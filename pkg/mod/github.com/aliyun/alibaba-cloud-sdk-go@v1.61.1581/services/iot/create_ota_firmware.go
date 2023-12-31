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

// CreateOTAFirmware invokes the iot.CreateOTAFirmware API synchronously
func (client *Client) CreateOTAFirmware(request *CreateOTAFirmwareRequest) (response *CreateOTAFirmwareResponse, err error) {
	response = CreateCreateOTAFirmwareResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOTAFirmwareWithChan invokes the iot.CreateOTAFirmware API asynchronously
func (client *Client) CreateOTAFirmwareWithChan(request *CreateOTAFirmwareRequest) (<-chan *CreateOTAFirmwareResponse, <-chan error) {
	responseChan := make(chan *CreateOTAFirmwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOTAFirmware(request)
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

// CreateOTAFirmwareWithCallback invokes the iot.CreateOTAFirmware API asynchronously
func (client *Client) CreateOTAFirmwareWithCallback(request *CreateOTAFirmwareRequest, callback func(response *CreateOTAFirmwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOTAFirmwareResponse
		var err error
		defer close(result)
		response, err = client.CreateOTAFirmware(request)
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

// CreateOTAFirmwareRequest is the request struct for api CreateOTAFirmware
type CreateOTAFirmwareRequest struct {
	*requests.RpcRequest
	SignMethod    string                         `position:"Query" name:"SignMethod"`
	MultiFiles    *[]CreateOTAFirmwareMultiFiles `position:"Query" name:"MultiFiles"  type:"Repeated"`
	NeedToVerify  requests.Boolean               `position:"Query" name:"NeedToVerify"`
	Type          requests.Integer               `position:"Query" name:"Type"`
	FirmwareUrl   string                         `position:"Query" name:"FirmwareUrl"`
	IotInstanceId string                         `position:"Query" name:"IotInstanceId"`
	FirmwareDesc  string                         `position:"Query" name:"FirmwareDesc"`
	ModuleName    string                         `position:"Query" name:"ModuleName"`
	FirmwareSign  string                         `position:"Query" name:"FirmwareSign"`
	FirmwareSize  requests.Integer               `position:"Query" name:"FirmwareSize"`
	FirmwareName  string                         `position:"Query" name:"FirmwareName"`
	ProductKey    string                         `position:"Query" name:"ProductKey"`
	SrcVersion    string                         `position:"Query" name:"SrcVersion"`
	ApiProduct    string                         `position:"Body" name:"ApiProduct"`
	ApiRevision   string                         `position:"Body" name:"ApiRevision"`
	Udi           string                         `position:"Query" name:"Udi"`
	DestVersion   string                         `position:"Query" name:"DestVersion"`
}

// CreateOTAFirmwareMultiFiles is a repeated param struct in CreateOTAFirmwareRequest
type CreateOTAFirmwareMultiFiles struct {
	Size      string `name:"Size"`
	Name      string `name:"Name"`
	SignValue string `name:"SignValue"`
	FileMd5   string `name:"FileMd5"`
	Url       string `name:"Url"`
}

// CreateOTAFirmwareResponse is the response struct for api CreateOTAFirmware
type CreateOTAFirmwareResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateOTAFirmwareRequest creates a request to invoke CreateOTAFirmware API
func CreateCreateOTAFirmwareRequest() (request *CreateOTAFirmwareRequest) {
	request = &CreateOTAFirmwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateOTAFirmware", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateOTAFirmwareResponse creates a response to parse from CreateOTAFirmware response
func CreateCreateOTAFirmwareResponse() (response *CreateOTAFirmwareResponse) {
	response = &CreateOTAFirmwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
