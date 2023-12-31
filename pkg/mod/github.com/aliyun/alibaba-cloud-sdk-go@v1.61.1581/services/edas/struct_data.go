package edas

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

// Data is a nested struct in edas response
type Data struct {
	Region                string         `json:"Region" xml:"Region"`
	Name                  string         `json:"Name" xml:"Name"`
	ExtraJson             string         `json:"ExtraJson" xml:"ExtraJson"`
	ChangeOrderId         string         `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Group                 string         `json:"Group" xml:"Group"`
	DubboMockItemJson     string         `json:"DubboMockItemJson" xml:"DubboMockItemJson"`
	TotalPages            int            `json:"TotalPages" xml:"TotalPages"`
	PageSize              int            `json:"PageSize" xml:"PageSize"`
	OversoldFactor        int            `json:"OversoldFactor" xml:"OversoldFactor"`
	EdasAppName           string         `json:"EdasAppName" xml:"EdasAppName"`
	Source                string         `json:"Source" xml:"Source"`
	ConsumerAppId         string         `json:"ConsumerAppId" xml:"ConsumerAppId"`
	Size                  int            `json:"Size" xml:"Size"`
	SlbPort               int            `json:"SlbPort" xml:"SlbPort"`
	Metadata              string         `json:"Metadata" xml:"Metadata"`
	PageNumber            int            `json:"PageNumber" xml:"PageNumber"`
	ConsumerAppName       string         `json:"ConsumerAppName" xml:"ConsumerAppName"`
	VpcId                 string         `json:"VpcId" xml:"VpcId"`
	ServiceType           string         `json:"ServiceType" xml:"ServiceType"`
	ServiceName           string         `json:"ServiceName" xml:"ServiceName"`
	ExtSlbIp              string         `json:"ExtSlbIp" xml:"ExtSlbIp"`
	UpdateTime            int64          `json:"UpdateTime" xml:"UpdateTime"`
	Version               string         `json:"Version" xml:"Version"`
	ExtSlbId              string         `json:"ExtSlbId" xml:"ExtSlbId"`
	TotalElements         int            `json:"TotalElements" xml:"TotalElements"`
	ProviderAppName       string         `json:"ProviderAppName" xml:"ProviderAppName"`
	Id                    int64          `json:"Id" xml:"Id"`
	ExtSlbName            string         `json:"ExtSlbName" xml:"ExtSlbName"`
	SlbName               string         `json:"SlbName" xml:"SlbName"`
	VServerGroupId        string         `json:"VServerGroupId" xml:"VServerGroupId"`
	ScMockItemJson        string         `json:"ScMockItemJson" xml:"ScMockItemJson"`
	AccountId             string         `json:"AccountId" xml:"AccountId"`
	ExtVServerGroupId     string         `json:"ExtVServerGroupId" xml:"ExtVServerGroupId"`
	RegistryType          string         `json:"RegistryType" xml:"RegistryType"`
	NamespaceId           string         `json:"NamespaceId" xml:"NamespaceId"`
	Enable                bool           `json:"Enable" xml:"Enable"`
	SpringApplicationName string         `json:"SpringApplicationName" xml:"SpringApplicationName"`
	SlbId                 string         `json:"SlbId" xml:"SlbId"`
	ClusterType           int            `json:"ClusterType" xml:"ClusterType"`
	ProviderAppId         string         `json:"ProviderAppId" xml:"ProviderAppId"`
	TotalSize             int            `json:"TotalSize" xml:"TotalSize"`
	DubboApplicationName  string         `json:"DubboApplicationName" xml:"DubboApplicationName"`
	SlbIp                 string         `json:"SlbIp" xml:"SlbIp"`
	Result                []MseMockRules `json:"Result" xml:"Result"`
	Content               []Provider     `json:"Content" xml:"Content"`
	RuleList              RuleList       `json:"RuleList" xml:"RuleList"`
	Methods               []Method       `json:"Methods" xml:"Methods"`
}
