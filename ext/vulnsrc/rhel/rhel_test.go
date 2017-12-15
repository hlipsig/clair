// Copyright 2017 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rhel

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/coreos/clair/database"
	"github.com/coreos/clair/ext/versionfmt/rpm"
	"github.com/stretchr/testify/assert"
)

func TestRHELParser(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(filename))

	// Test parsing testdata/fetcher_rhel_test.1.xml
	testFile, _ := os.Open(path + "/testdata/fetcher_rhel_test.1.xml")
	vulnerabilities, err := parseRHSA(testFile)
	if assert.Nil(t, err) && assert.Len(t, vulnerabilities, 1) {
		assert.Equal(t, "CVE-2015-0252", vulnerabilities[0].Name)
		assert.Equal(t, "https://access.redhat.com/security/cve/CVE-2015-0252", vulnerabilities[0].Link)
		assert.Equal(t, database.MediumSeverity, vulnerabilities[0].Severity)
		assert.Equal(t, `Xerces-C is a validating XML parser written in a portable subset of C++. A flaw was found in the way the Xerces-C XML parser processed certain XML documents. A remote attacker could provide specially crafted XML input that, when parsed by an application using Xerces-C, would cause that application to crash.`, vulnerabilities[0].Description)

		expectedFeatures := []database.AffectedFeature{
			{
				Namespace: database.Namespace{
					Name:          "centos:7",
					VersionFormat: rpm.ParserName,
				},
				FeatureName:     "xerces-c",
				AffectedVersion: "0:3.1.1-7.el7_1",
				FixedInVersion:  "0:3.1.1-7.el7_1",
			},
			{
				Namespace: database.Namespace{
					Name:          "centos:7",
					VersionFormat: rpm.ParserName,
				},
				FeatureName:     "xerces-c-devel",
				AffectedVersion: "0:3.1.1-7.el7_1",
				FixedInVersion:  "0:3.1.1-7.el7_1",
			},
			{
				Namespace: database.Namespace{
					Name:          "centos:7",
					VersionFormat: rpm.ParserName,
				},
				FeatureName:     "xerces-c-doc",
				AffectedVersion: "0:3.1.1-7.el7_1",
				FixedInVersion:  "0:3.1.1-7.el7_1",
			},
		}

		for _, expectedFeature := range expectedFeatures {
			assert.Contains(t, vulnerabilities[0].Affected, expectedFeature)
		}
	}

	// Test parsing testdata/fetcher_rhel_test.2.xml
	testFile, _ = os.Open(path + "/testdata/fetcher_rhel_test.2.xml")
	vulnerabilities, err = parseRHSA(testFile)
	if assert.Nil(t, err) && assert.Len(t, vulnerabilities, 17) {
		assert.Equal(t, "CVE-2015-2722", vulnerabilities[0].Name)
		assert.Equal(t, "https://access.redhat.com/security/cve/CVE-2015-2722", vulnerabilities[0].Link)
		assert.Equal(t, database.CriticalSeverity, vulnerabilities[0].Severity)
		assert.Equal(t, `Mozilla Firefox is an open source web browser. XULRunner provides the XUL Runtime environment for Mozilla Firefox. Several flaws were found in the processing of malformed web content. A web page containing malicious content could cause Firefox to crash or, potentially, execute arbitrary code with the privileges of the user running Firefox.`, vulnerabilities[0].Description)

		expectedFeatures := []database.AffectedFeature{
			{
				Namespace: database.Namespace{
					Name:          "centos:6",
					VersionFormat: rpm.ParserName,
				},
				FeatureName:     "firefox",
				FixedInVersion:  "0:38.1.0-1.el6_6",
				AffectedVersion: "0:38.1.0-1.el6_6",
			},
			{
				Namespace: database.Namespace{
					Name:          "centos:7",
					VersionFormat: rpm.ParserName,
				},
				FeatureName:     "firefox",
				FixedInVersion:  "0:38.1.0-1.el7_1",
				AffectedVersion: "0:38.1.0-1.el7_1",
			},
		}

		for _, expectedFeature := range expectedFeatures {
			assert.Contains(t, vulnerabilities[0].Affected, expectedFeature)
		}
	}
}
