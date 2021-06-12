/*
Copyright 2021 The Vitess Authors.

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

package mysqlctlproto

import (
	"strings"
	"time"

	"vitess.io/vitess/go/protoutil"
	"vitess.io/vitess/go/vt/log"
	"vitess.io/vitess/go/vt/mysqlctl"
	"vitess.io/vitess/go/vt/mysqlctl/backupstorage"
	"vitess.io/vitess/go/vt/topo/topoproto"

	mysqlctlpb "vitess.io/vitess/go/vt/proto/mysqlctl"
)

// BackupHandleToProto returns a BackupInfo proto from a BackupHandle.
func BackupHandleToProto(bh backupstorage.BackupHandle) *mysqlctlpb.BackupInfo {
	bi := &mysqlctlpb.BackupInfo{
		Name:      bh.Name(),
		Directory: bh.Directory(),
	}

	if parts := strings.Split(bi.Name, "."); len(parts) == 3 {
		// parts[0]: date part of mysqlctl.BackupTimestampFormat
		// parts[1]: time part of mysqlctl.BackupTimestampFormat
		// parts[2]: tablet alias
		timestamp := strings.Join(parts[:2], ".")
		aliasStr := parts[2]

		backupTime, err := time.Parse(mysqlctl.BackupTimestampFormat, timestamp)
		if err != nil {
			log.Errorf("error parsing backup time for %s/%s: %s", bi.Directory, bi.Name, err)
		} else {
			bi.Time = protoutil.TimeToProto(backupTime)
		}

		alias, err := topoproto.ParseTabletAlias(aliasStr)
		if err != nil {
			log.Errorf("error parsing tablet alias for %s/%s: %s", bi.Directory, bi.Name, err)
		} else {
			bi.TabletAlias = alias
		}
	}

	return bi
}
