/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package community

import (
	"github.com/chaosblade-io/chaosblade-operator/pkg/runtime/chaosblade"
)

const Community = "community"

func init() {
	chaosblade.Products[Community] = &chaosblade.ProductConstant{
		Home:          "/opt/chaosblade",
		BladeBin:      "/opt/chaosblade/blade",
		PodName:       "chaosblade-tool",
		ImageRepoFunc: ImageRepoForCommunity,
		PodLabels:     map[string]string{"app": "chaosblade-tool"},
	}
}

var ImageRepoForCommunity = func() string {
	return chaosblade.ImageRepository
}