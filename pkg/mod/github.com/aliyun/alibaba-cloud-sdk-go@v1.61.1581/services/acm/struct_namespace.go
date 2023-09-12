package acm

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

// Namespace is a nested struct in acm response
type Namespace struct {
	Name          string `json:"Name" xml:"Name"`
	NamespaceName string `json:"NamespaceName" xml:"NamespaceName"`
	Endpoint      string `json:"Endpoint" xml:"Endpoint"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	SecretKey     string `json:"SecretKey" xml:"SecretKey"`
	ConfigCount   int    `json:"ConfigCount" xml:"ConfigCount"`
	NamespaceId   string `json:"NamespaceId" xml:"NamespaceId"`
	AccessKey     string `json:"AccessKey" xml:"AccessKey"`
	Quota         int    `json:"Quota" xml:"Quota"`
	Type          int    `json:"Type" xml:"Type"`
}