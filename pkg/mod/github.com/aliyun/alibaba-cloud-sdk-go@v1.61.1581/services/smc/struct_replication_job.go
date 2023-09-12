package smc

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

// ReplicationJob is a nested struct in smc response
type ReplicationJob struct {
	JobId                  string                                   `json:"JobId" xml:"JobId"`
	SourceId               string                                   `json:"SourceId" xml:"SourceId"`
	Name                   string                                   `json:"Name" xml:"Name"`
	Description            string                                   `json:"Description" xml:"Description"`
	RegionId               string                                   `json:"RegionId" xml:"RegionId"`
	TargetType             string                                   `json:"TargetType" xml:"TargetType"`
	ScheduledStartTime     string                                   `json:"ScheduledStartTime" xml:"ScheduledStartTime"`
	ImageName              string                                   `json:"ImageName" xml:"ImageName"`
	InstanceId             string                                   `json:"InstanceId" xml:"InstanceId"`
	ImageId                string                                   `json:"ImageId" xml:"ImageId"`
	Status                 string                                   `json:"Status" xml:"Status"`
	BusinessStatus         string                                   `json:"BusinessStatus" xml:"BusinessStatus"`
	ErrorCode              string                                   `json:"ErrorCode" xml:"ErrorCode"`
	Progress               float64                                  `json:"Progress" xml:"Progress"`
	CreationTime           string                                   `json:"CreationTime" xml:"CreationTime"`
	ValidTime              string                                   `json:"ValidTime" xml:"ValidTime"`
	StartTime              string                                   `json:"StartTime" xml:"StartTime"`
	EndTime                string                                   `json:"EndTime" xml:"EndTime"`
	NetMode                int                                      `json:"NetMode" xml:"NetMode"`
	SystemDiskSize         int                                      `json:"SystemDiskSize" xml:"SystemDiskSize"`
	VpcId                  string                                   `json:"VpcId" xml:"VpcId"`
	VSwitchId              string                                   `json:"VSwitchId" xml:"VSwitchId"`
	TransitionInstanceId   string                                   `json:"TransitionInstanceId" xml:"TransitionInstanceId"`
	StatusInfo             string                                   `json:"StatusInfo" xml:"StatusInfo"`
	ReplicationParameters  string                                   `json:"ReplicationParameters" xml:"ReplicationParameters"`
	RunOnce                bool                                     `json:"RunOnce" xml:"RunOnce"`
	Frequency              int                                      `json:"Frequency" xml:"Frequency"`
	MaxNumberOfImageToKeep int                                      `json:"MaxNumberOfImageToKeep" xml:"MaxNumberOfImageToKeep"`
	InstanceType           string                                   `json:"InstanceType" xml:"InstanceType"`
	LaunchTemplateId       string                                   `json:"LaunchTemplateId" xml:"LaunchTemplateId"`
	LaunchTemplateVersion  string                                   `json:"LaunchTemplateVersion" xml:"LaunchTemplateVersion"`
	InstanceRamRole        string                                   `json:"InstanceRamRole" xml:"InstanceRamRole"`
	ContainerNamespace     string                                   `json:"ContainerNamespace" xml:"ContainerNamespace"`
	ContainerRepository    string                                   `json:"ContainerRepository" xml:"ContainerRepository"`
	ContainerTag           string                                   `json:"ContainerTag" xml:"ContainerTag"`
	LicenseType            string                                   `json:"LicenseType" xml:"LicenseType"`
	SystemDiskParts        SystemDiskPartsInDescribeReplicationJobs `json:"SystemDiskParts" xml:"SystemDiskParts"`
	DataDisks              DataDisksInDescribeReplicationJobs       `json:"DataDisks" xml:"DataDisks"`
	ReplicationJobRuns     ReplicationJobRuns                       `json:"ReplicationJobRuns" xml:"ReplicationJobRuns"`
}
