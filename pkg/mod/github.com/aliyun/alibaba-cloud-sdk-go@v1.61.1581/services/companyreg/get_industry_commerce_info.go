package companyreg

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

// GetIndustryCommerceInfo invokes the companyreg.GetIndustryCommerceInfo API synchronously
func (client *Client) GetIndustryCommerceInfo(request *GetIndustryCommerceInfoRequest) (response *GetIndustryCommerceInfoResponse, err error) {
	response = CreateGetIndustryCommerceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetIndustryCommerceInfoWithChan invokes the companyreg.GetIndustryCommerceInfo API asynchronously
func (client *Client) GetIndustryCommerceInfoWithChan(request *GetIndustryCommerceInfoRequest) (<-chan *GetIndustryCommerceInfoResponse, <-chan error) {
	responseChan := make(chan *GetIndustryCommerceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIndustryCommerceInfo(request)
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

// GetIndustryCommerceInfoWithCallback invokes the companyreg.GetIndustryCommerceInfo API asynchronously
func (client *Client) GetIndustryCommerceInfoWithCallback(request *GetIndustryCommerceInfoRequest, callback func(response *GetIndustryCommerceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIndustryCommerceInfoResponse
		var err error
		defer close(result)
		response, err = client.GetIndustryCommerceInfo(request)
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

// GetIndustryCommerceInfoRequest is the request struct for api GetIndustryCommerceInfo
type GetIndustryCommerceInfoRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// GetIndustryCommerceInfoResponse is the response struct for api GetIndustryCommerceInfo
type GetIndustryCommerceInfoResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	CompanyType         string `json:"CompanyType" xml:"CompanyType"`
	EstablishmentDate   string `json:"EstablishmentDate" xml:"EstablishmentDate"`
	BusinessStatus      string `json:"BusinessStatus" xml:"BusinessStatus"`
	TaxNo               string `json:"TaxNo" xml:"TaxNo"`
	RegisteredCaptial   string `json:"RegisteredCaptial" xml:"RegisteredCaptial"`
	Name                string `json:"Name" xml:"Name"`
	LegalRepresentative string `json:"LegalRepresentative" xml:"LegalRepresentative"`
	OperatingPeriod     string `json:"OperatingPeriod" xml:"OperatingPeriod"`
	BizScope            string `json:"BizScope" xml:"BizScope"`
	CorpAddress         string `json:"CorpAddress" xml:"CorpAddress"`
}

// CreateGetIndustryCommerceInfoRequest creates a request to invoke GetIndustryCommerceInfo API
func CreateGetIndustryCommerceInfoRequest() (request *GetIndustryCommerceInfoRequest) {
	request = &GetIndustryCommerceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "GetIndustryCommerceInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateGetIndustryCommerceInfoResponse creates a response to parse from GetIndustryCommerceInfo response
func CreateGetIndustryCommerceInfoResponse() (response *GetIndustryCommerceInfoResponse) {
	response = &GetIndustryCommerceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
