package smartag

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

// Acr is a nested struct in smartag response
type Acr struct {
	Direction       string                                `json:"Direction" xml:"Direction"`
	Type            string                                `json:"Type" xml:"Type"`
	IpProtocol      string                                `json:"IpProtocol" xml:"IpProtocol"`
	Priority        int                                   `json:"Priority" xml:"Priority"`
	AclId           string                                `json:"AclId" xml:"AclId"`
	Policy          string                                `json:"Policy" xml:"Policy"`
	Description     string                                `json:"Description" xml:"Description"`
	GmtCreate       int64                                 `json:"GmtCreate" xml:"GmtCreate"`
	DestCidr        string                                `json:"DestCidr" xml:"DestCidr"`
	DestPortRange   string                                `json:"DestPortRange" xml:"DestPortRange"`
	Name            string                                `json:"Name" xml:"Name"`
	AcrId           string                                `json:"AcrId" xml:"AcrId"`
	SourceCidr      string                                `json:"SourceCidr" xml:"SourceCidr"`
	SourcePortRange string                                `json:"SourcePortRange" xml:"SourcePortRange"`
	DpiSignatureIds DpiSignatureIdsInDescribeACLAttribute `json:"DpiSignatureIds" xml:"DpiSignatureIds"`
	DpiGroupIds     DpiGroupIdsInDescribeACLAttribute     `json:"DpiGroupIds" xml:"DpiGroupIds"`
}
