/*
* Copyright (C) 2021 THL A29 Limited, a Tencent company.  All rights reserved.
* This source code is licensed under the Apache License Version 2.0.
 */

package db

import (
	"nocalhost/internal/nhctl/dbutils"
	"nocalhost/internal/nhctl/nocalhost_path"
	"path/filepath"
)

func OpenApplicationLevelDB(ns, app, nid string, readonly bool) (*dbutils.LevelDBUtils, error) {
	path := nocalhost_path.GetAppDbDir(ns, app, nid)
	return dbutils.OpenLevelDB(path, readonly)
}

func CreateApplicationLevelDB(ns, app, nid string, errorIfExist bool) error {
	path := nocalhost_path.GetAppDbDir(ns, app, nid)
	return dbutils.CreateLevelDB(path, errorIfExist)
}

func GetPortForwardLevelDB() (*dbutils.LevelDBUtils, error) {
	path := filepath.Join(nocalhost_path.GetNhctlHomeDir(), nocalhost_path.DefaultNhctlPortForward)
	return dbutils.OpenLevelDB(path, false)
}
