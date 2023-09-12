package ccc

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

// AttemptList is a nested struct in ccc response
type AttemptList struct {
	AttemptId               string `json:"AttemptId" xml:"AttemptId"`
	ContactId               string `json:"ContactId" xml:"ContactId"`
	CaseId                  string `json:"CaseId" xml:"CaseId"`
	CampaignId              string `json:"CampaignId" xml:"CampaignId"`
	InstanceId              string `json:"InstanceId" xml:"InstanceId"`
	QueueId                 string `json:"QueueId" xml:"QueueId"`
	Caller                  string `json:"Caller" xml:"Caller"`
	Callee                  string `json:"Callee" xml:"Callee"`
	AgentId                 string `json:"AgentId" xml:"AgentId"`
	DialTime                int64  `json:"DialTime" xml:"DialTime"`
	DialDuration            int64  `json:"DialDuration" xml:"DialDuration"`
	CustomerEstablishedTime int64  `json:"CustomerEstablishedTime" xml:"CustomerEstablishedTime"`
	CustomerReleasedTime    int64  `json:"CustomerReleasedTime" xml:"CustomerReleasedTime"`
	EnterIvrTime            int64  `json:"EnterIvrTime" xml:"EnterIvrTime"`
	IvrDuration             int64  `json:"IvrDuration" xml:"IvrDuration"`
	EnqueueTime             int64  `json:"EnqueueTime" xml:"EnqueueTime"`
	QueueDuration           int64  `json:"QueueDuration" xml:"QueueDuration"`
	AssignAgentTime         int64  `json:"AssignAgentTime" xml:"AssignAgentTime"`
	AgentRingDuration       int64  `json:"AgentRingDuration" xml:"AgentRingDuration"`
	AgentEstablishedTime    int64  `json:"AgentEstablishedTime" xml:"AgentEstablishedTime"`
}