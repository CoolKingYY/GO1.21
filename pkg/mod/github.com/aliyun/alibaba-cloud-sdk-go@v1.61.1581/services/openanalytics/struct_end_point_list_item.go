package openanalytics

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

// EndPointListItem is a nested struct in openanalytics response
type EndPointListItem struct {
	EndPointID      string `json:"endPointID" xml:"endPointID"`
	Zone            string `json:"zone" xml:"zone"`
	VSwitch         string `json:"vSwitch" xml:"vSwitch"`
	Status          string `json:"status" xml:"status"`
	VpcID           string `json:"vpcID" xml:"vpcID"`
	Host            string `json:"host" xml:"host"`
	DomainURL       string `json:"domainURL" xml:"domainURL"`
	NetworkType     string `json:"networkType" xml:"networkType"`
	AllowIP         string `json:"allowIP" xml:"allowIP"`
	Port            string `json:"port" xml:"port"`
	CloudInstanceID string `json:"cloudInstanceID" xml:"cloudInstanceID"`
}
