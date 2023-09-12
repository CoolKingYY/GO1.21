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

// DescribeLiveCertificateDetail invokes the live.DescribeLiveCertificateDetail API synchronously
func (client *Client) DescribeLiveCertificateDetail(request *DescribeLiveCertificateDetailRequest) (response *DescribeLiveCertificateDetailResponse, err error) {
	response = CreateDescribeLiveCertificateDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveCertificateDetailWithChan invokes the live.DescribeLiveCertificateDetail API asynchronously
func (client *Client) DescribeLiveCertificateDetailWithChan(request *DescribeLiveCertificateDetailRequest) (<-chan *DescribeLiveCertificateDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveCertificateDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveCertificateDetail(request)
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

// DescribeLiveCertificateDetailWithCallback invokes the live.DescribeLiveCertificateDetail API asynchronously
func (client *Client) DescribeLiveCertificateDetailWithCallback(request *DescribeLiveCertificateDetailRequest, callback func(response *DescribeLiveCertificateDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveCertificateDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveCertificateDetail(request)
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

// DescribeLiveCertificateDetailRequest is the request struct for api DescribeLiveCertificateDetail
type DescribeLiveCertificateDetailRequest struct {
	*requests.RpcRequest
	CertName      string           `position:"Query" name:"CertName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeLiveCertificateDetailResponse is the response struct for api DescribeLiveCertificateDetail
type DescribeLiveCertificateDetailResponse struct {
	*responses.BaseResponse
	CertName  string `json:"CertName" xml:"CertName"`
	Cert      string `json:"Cert" xml:"Cert"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	CertId    int64  `json:"CertId" xml:"CertId"`
}

// CreateDescribeLiveCertificateDetailRequest creates a request to invoke DescribeLiveCertificateDetail API
func CreateDescribeLiveCertificateDetailRequest() (request *DescribeLiveCertificateDetailRequest) {
	request = &DescribeLiveCertificateDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveCertificateDetail", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveCertificateDetailResponse creates a response to parse from DescribeLiveCertificateDetail response
func CreateDescribeLiveCertificateDetailResponse() (response *DescribeLiveCertificateDetailResponse) {
	response = &DescribeLiveCertificateDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
