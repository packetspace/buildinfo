/*
Copyright © 2022 Mike Joseph <mike@mjoseph.org>

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

package buildinfo

import "runtime/debug"

type BuildInfo debug.BuildInfo

const (
	unknownVersion = "(unknown)"
)

var (
	version string
	Default *BuildInfo
)

func init() {
	Default, _ = GetBuildInfo()
}

func GetBuildInfo() (*BuildInfo, bool) {
	bi, ok := debug.ReadBuildInfo()
	return (*BuildInfo)(bi), ok
}

func (bi *BuildInfo) AppVersion() string {
	if version != "" {
		return version
	} else if bi != nil {
		return bi.Main.Version
	} else {
		return unknownVersion
	}
}

func (bi *BuildInfo) IsVersionKnown() bool {
	if version != "" {
		return true
	} else if bi != nil {
		return true
	} else {
		return false
	}
}

func (bi *BuildInfo) IsVersionInjected() bool {
	if version != "" {
		return true
	} else {
		return false
	}
}
