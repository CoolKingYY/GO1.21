package ga

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

// EndpointGroupsItem is a nested struct in ga response
type EndpointGroupsItem struct {
	HealthCheckIntervalSeconds int                          `json:"HealthCheckIntervalSeconds" xml:"HealthCheckIntervalSeconds"`
	TrafficPercentage          int                          `json:"TrafficPercentage" xml:"TrafficPercentage"`
	Description                string                       `json:"Description" xml:"Description"`
	EndpointGroupId            string                       `json:"EndpointGroupId" xml:"EndpointGroupId"`
	HealthCheckPath            string                       `json:"HealthCheckPath" xml:"HealthCheckPath"`
	ThresholdCount             int                          `json:"ThresholdCount" xml:"ThresholdCount"`
	EndpointRequestProtocol    string                       `json:"EndpointRequestProtocol" xml:"EndpointRequestProtocol"`
	Name                       string                       `json:"Name" xml:"Name"`
	EndpointGroupRegion        string                       `json:"EndpointGroupRegion" xml:"EndpointGroupRegion"`
	State                      string                       `json:"State" xml:"State"`
	HealthCheckProtocol        string                       `json:"HealthCheckProtocol" xml:"HealthCheckProtocol"`
	HealthCheckPort            int                          `json:"HealthCheckPort" xml:"HealthCheckPort"`
	AcceleratorId              string                       `json:"AcceleratorId" xml:"AcceleratorId"`
	EndpointGroupType          string                       `json:"EndpointGroupType" xml:"EndpointGroupType"`
	ListenerId                 string                       `json:"ListenerId" xml:"ListenerId"`
	EndpointGroupIpList        []string                     `json:"EndpointGroupIpList" xml:"EndpointGroupIpList"`
	ForwardingRuleIds          []string                     `json:"ForwardingRuleIds" xml:"ForwardingRuleIds"`
	PortOverrides              []PortOverridesItem          `json:"PortOverrides" xml:"PortOverrides"`
	EndpointConfigurations     []EndpointConfigurationsItem `json:"EndpointConfigurations" xml:"EndpointConfigurations"`
}