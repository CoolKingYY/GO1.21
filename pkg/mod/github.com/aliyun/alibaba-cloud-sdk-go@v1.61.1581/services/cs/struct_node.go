package cs

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

// Node is a nested struct in cs response
type Node struct {
	ErrorMessage       string   `json:"error_message" xml:"error_message"`
	CreationTime       string   `json:"creation_time" xml:"creation_time"`
	NodeStatus         string   `json:"node_status" xml:"node_status"`
	InstanceName       string   `json:"instance_name" xml:"instance_name"`
	IsAliyunNode       bool     `json:"is_aliyun_node" xml:"is_aliyun_node"`
	NodeName           string   `json:"node_name" xml:"node_name"`
	SpotStrategy       string   `json:"spot_strategy" xml:"spot_strategy"`
	ExpiredTime        string   `json:"expired_time" xml:"expired_time"`
	Source             string   `json:"source" xml:"source"`
	InstanceTypeFamily string   `json:"instance_type_family" xml:"instance_type_family"`
	InstanceId         string   `json:"instance_id" xml:"instance_id"`
	InstanceChargeType string   `json:"instance_charge_type" xml:"instance_charge_type"`
	InstanceRole       string   `json:"instance_role" xml:"instance_role"`
	State              string   `json:"state" xml:"state"`
	InstanceStatus     string   `json:"instance_status" xml:"instance_status"`
	ImageId            string   `json:"image_id" xml:"image_id"`
	NodepoolId         string   `json:"nodepool_id" xml:"nodepool_id"`
	InstanceType       string   `json:"instance_type" xml:"instance_type"`
	HostName           string   `json:"host_name" xml:"host_name"`
	IpAddress          []string `json:"ip_address" xml:"ip_address"`
}
