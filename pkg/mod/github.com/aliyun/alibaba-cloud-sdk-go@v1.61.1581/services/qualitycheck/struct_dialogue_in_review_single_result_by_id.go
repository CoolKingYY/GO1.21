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

// DialogueInReviewSingleResultById is a nested struct in qualitycheck response
type DialogueInReviewSingleResultById struct {
	Words           string `json:"Words" xml:"Words"`
	Identity        string `json:"Identity" xml:"Identity"`
	Begin           int64  `json:"Begin" xml:"Begin"`
	BeginTime       int64  `json:"BeginTime" xml:"BeginTime"`
	EmotionValue    int    `json:"EmotionValue" xml:"EmotionValue"`
	End             int64  `json:"End" xml:"End"`
	SpeechRate      int    `json:"SpeechRate" xml:"SpeechRate"`
	Role            string `json:"Role" xml:"Role"`
	SilenceDuration int    `json:"SilenceDuration" xml:"SilenceDuration"`
	HourMinSec      string `json:"HourMinSec" xml:"HourMinSec"`
}
