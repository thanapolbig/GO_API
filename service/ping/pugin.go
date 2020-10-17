package ping

import (
	mssql "golang-101/database/mssql"
)

func logHeartbeat(hb heartbeat) (err error) {

	if err = mssql.DB.Table("log_heartbeat").Save(hb).Error; err != nil {
		return
	}

	return
}

