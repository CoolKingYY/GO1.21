package qualitycheck

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

// TaskResultReviewInfo is a nested struct in qualitycheck response
type TaskResultReviewInfo struct {
	Status              int        `json:"Status" xml:"Status"`
	DataType            int        `json:"DataType" xml:"DataType"`
	HitNumber           int        `json:"HitNumber" xml:"HitNumber"`
	HitRule             bool       `json:"HitRule" xml:"HitRule"`
	NextVid             string     `json:"NextVid" xml:"NextVid"`
	PreVid              string     `json:"PreVid" xml:"PreVid"`
	ReviewAccuracyRate  float64    `json:"ReviewAccuracyRate" xml:"ReviewAccuracyRate"`
	RealViolationNumber int        `json:"RealViolationNumber" xml:"RealViolationNumber"`
	IsHitRule           bool       `json:"IsHitRule" xml:"IsHitRule"`
	FileName            string     `json:"FileName" xml:"FileName"`
	TotalScore          int        `json:"TotalScore" xml:"TotalScore"`
	CheckNumber         int        `json:"CheckNumber" xml:"CheckNumber"`
	FileMergeName       string     `json:"FileMergeName" xml:"FileMergeName"`
	BucketName          string     `json:"BucketName" xml:"BucketName"`
	HandTaskFile        bool       `json:"HandTaskFile" xml:"HandTaskFile"`
	TaskId              string     `json:"TaskId" xml:"TaskId"`
	Vid                 string     `json:"Vid" xml:"Vid"`
	HitRuleSet          HitRuleSet `json:"HitRuleSet" xml:"HitRuleSet"`
}
