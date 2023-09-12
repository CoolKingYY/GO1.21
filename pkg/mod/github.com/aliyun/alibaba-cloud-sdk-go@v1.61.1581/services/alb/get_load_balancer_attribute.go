package alb

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

// GetLoadBalancerAttribute invokes the alb.GetLoadBalancerAttribute API synchronously
func (client *Client) GetLoadBalancerAttribute(request *GetLoadBalancerAttributeRequest) (response *GetLoadBalancerAttributeResponse, err error) {
	response = CreateGetLoadBalancerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// GetLoadBalancerAttributeWithChan invokes the alb.GetLoadBalancerAttribute API asynchronously
func (client *Client) GetLoadBalancerAttributeWithChan(request *GetLoadBalancerAttributeRequest) (<-chan *GetLoadBalancerAttributeResponse, <-chan error) {
	responseChan := make(chan *GetLoadBalancerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLoadBalancerAttribute(request)
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

// GetLoadBalancerAttributeWithCallback invokes the alb.GetLoadBalancerAttribute API asynchronously
func (client *Client) GetLoadBalancerAttributeWithCallback(request *GetLoadBalancerAttributeRequest, callback func(response *GetLoadBalancerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLoadBalancerAttributeResponse
		var err error
		defer close(result)
		response, err = client.GetLoadBalancerAttribute(request)
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

// GetLoadBalancerAttributeRequest is the request struct for api GetLoadBalancerAttribute
type GetLoadBalancerAttributeRequest struct {
	*requests.RpcRequest
	LoadBalancerId string `position:"Query" name:"LoadBalancerId"`
}

// GetLoadBalancerAttributeResponse is the response struct for api GetLoadBalancerAttribute
type GetLoadBalancerAttributeResponse struct {
	*responses.BaseResponse
	AddressAllocatedMode         string                       `json:"AddressAllocatedMode" xml:"AddressAllocatedMode"`
	AddressType                  string                       `json:"AddressType" xml:"AddressType"`
	BandwidthCapacity            int                          `json:"BandwidthCapacity" xml:"BandwidthCapacity"`
	BandwidthPackageId           string                       `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
	CreateTime                   string                       `json:"CreateTime" xml:"CreateTime"`
	DNSName                      string                       `json:"DNSName" xml:"DNSName"`
	LoadBalancerBussinessStatus  string                       `json:"LoadBalancerBussinessStatus" xml:"LoadBalancerBussinessStatus"`
	LoadBalancerEdition          string                       `json:"LoadBalancerEdition" xml:"LoadBalancerEdition"`
	LoadBalancerId               string                       `json:"LoadBalancerId" xml:"LoadBalancerId"`
	LoadBalancerName             string                       `json:"LoadBalancerName" xml:"LoadBalancerName"`
	ServiceManagedEnabled        bool                         `json:"ServiceManagedEnabled" xml:"ServiceManagedEnabled"`
	ServiceManagedMode           string                       `json:"ServiceManagedMode" xml:"ServiceManagedMode"`
	LoadBalancerStatus           string                       `json:"LoadBalancerStatus" xml:"LoadBalancerStatus"`
	RegionId                     string                       `json:"RegionId" xml:"RegionId"`
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	ResourceGroupId              string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	VpcId                        string                       `json:"VpcId" xml:"VpcId"`
	ConfigManagedEnabled         bool                         `json:"ConfigManagedEnabled" xml:"ConfigManagedEnabled"`
	AddressIpVersion             string                       `json:"AddressIpVersion" xml:"AddressIpVersion"`
	Ipv6AddressType              string                       `json:"Ipv6AddressType" xml:"Ipv6AddressType"`
	FeatureLabels                []string                     `json:"FeatureLabels" xml:"FeatureLabels"`
	AccessLogConfig              AccessLogConfig              `json:"AccessLogConfig" xml:"AccessLogConfig"`
	DeletionProtectionConfig     DeletionProtectionConfig     `json:"DeletionProtectionConfig" xml:"DeletionProtectionConfig"`
	LoadBalancerBillingConfig    LoadBalancerBillingConfig    `json:"LoadBalancerBillingConfig" xml:"LoadBalancerBillingConfig"`
	ModificationProtectionConfig ModificationProtectionConfig `json:"ModificationProtectionConfig" xml:"ModificationProtectionConfig"`
	LoadBalancerOperationLocks   []LoadBalancerOperationLock  `json:"LoadBalancerOperationLocks" xml:"LoadBalancerOperationLocks"`
	Tags                         []Tag                        `json:"Tags" xml:"Tags"`
	ZoneMappings                 []ZoneMapping                `json:"ZoneMappings" xml:"ZoneMappings"`
}

// CreateGetLoadBalancerAttributeRequest creates a request to invoke GetLoadBalancerAttribute API
func CreateGetLoadBalancerAttributeRequest() (request *GetLoadBalancerAttributeRequest) {
	request = &GetLoadBalancerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "GetLoadBalancerAttribute", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLoadBalancerAttributeResponse creates a response to parse from GetLoadBalancerAttribute response
func CreateGetLoadBalancerAttributeResponse() (response *GetLoadBalancerAttributeResponse) {
	response = &GetLoadBalancerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
