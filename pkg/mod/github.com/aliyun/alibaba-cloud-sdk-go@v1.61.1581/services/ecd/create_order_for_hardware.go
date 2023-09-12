package ecd

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

// CreateOrderForHardware invokes the ecd.CreateOrderForHardware API synchronously
func (client *Client) CreateOrderForHardware(request *CreateOrderForHardwareRequest) (response *CreateOrderForHardwareResponse, err error) {
	response = CreateCreateOrderForHardwareResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrderForHardwareWithChan invokes the ecd.CreateOrderForHardware API asynchronously
func (client *Client) CreateOrderForHardwareWithChan(request *CreateOrderForHardwareRequest) (<-chan *CreateOrderForHardwareResponse, <-chan error) {
	responseChan := make(chan *CreateOrderForHardwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrderForHardware(request)
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

// CreateOrderForHardwareWithCallback invokes the ecd.CreateOrderForHardware API asynchronously
func (client *Client) CreateOrderForHardwareWithCallback(request *CreateOrderForHardwareRequest, callback func(response *CreateOrderForHardwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrderForHardwareResponse
		var err error
		defer close(result)
		response, err = client.CreateOrderForHardware(request)
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

// CreateOrderForHardwareRequest is the request struct for api CreateOrderForHardware
type CreateOrderForHardwareRequest struct {
	*requests.RpcRequest
	StreetName      string           `position:"Query" name:"StreetName"`
	CountryCode     string           `position:"Query" name:"CountryCode"`
	ContactName     string           `position:"Query" name:"ContactName"`
	ProvName        string           `position:"Query" name:"ProvName"`
	CityCode        string           `position:"Query" name:"CityCode"`
	MobilePhone     string           `position:"Query" name:"MobilePhone"`
	HardwareType    string           `position:"Query" name:"HardwareType"`
	DistrictCode    string           `position:"Query" name:"DistrictCode"`
	Email           string           `position:"Query" name:"Email"`
	HardwareVersion string           `position:"Query" name:"HardwareVersion"`
	Amount          requests.Integer `position:"Query" name:"Amount"`
	DetailAddress   string           `position:"Query" name:"DetailAddress"`
	AutoPay         requests.Boolean `position:"Query" name:"AutoPay"`
	CityName        string           `position:"Query" name:"CityName"`
	CountryName     string           `position:"Query" name:"CountryName"`
	PromotionId     string           `position:"Query" name:"PromotionId"`
	ZipCode         string           `position:"Query" name:"ZipCode"`
	DistrictName    string           `position:"Query" name:"DistrictName"`
	Phone           string           `position:"Query" name:"Phone"`
	StreetCode      string           `position:"Query" name:"StreetCode"`
	ProvCode        string           `position:"Query" name:"ProvCode"`
}

// CreateOrderForHardwareResponse is the response struct for api CreateOrderForHardware
type CreateOrderForHardwareResponse struct {
	*responses.BaseResponse
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrderForHardwareRequest creates a request to invoke CreateOrderForHardware API
func CreateCreateOrderForHardwareRequest() (request *CreateOrderForHardwareRequest) {
	request = &CreateOrderForHardwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateOrderForHardware", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateOrderForHardwareResponse creates a response to parse from CreateOrderForHardware response
func CreateCreateOrderForHardwareResponse() (response *CreateOrderForHardwareResponse) {
	response = &CreateOrderForHardwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}