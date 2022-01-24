/*
    Copyright (C) 2022 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import "github.com/awslabs/goformation/v4/cloudformation/sns"

// SnsTopicConfig hold config for SnsTopic
type SnsTopicConfig struct {
	Config
	Name        string `json:"name"`
	KmsMasterID string `json:"kms_master_id"`
}

// GetSnsTopicConfig returns config for SnsTopic
func GetSnsTopicConfig(t *sns.Topic) []AWSResourceConfig {
	cf := SnsTopicConfig{
		Config: Config{
			Name: t.TopicName,
		},
		Name:        t.TopicName,
		KmsMasterID: t.KmsMasterKeyId,
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: t.AWSCloudFormationMetadata,
	}}
}