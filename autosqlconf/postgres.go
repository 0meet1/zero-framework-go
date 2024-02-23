package autosqlconf

import (
	"errors"
	"fmt"
	"strings"

	x0meet1 "github.com/0meet1/zero-framework"
)

type ZeroDbAutoPostgresProcessor struct {
	x0meet1.ZeroCoreProcessor
}

func (processor *ZeroDbAutoPostgresProcessor) ColumnExists(tableSchema string, tableName string, columName string) (int, error) {
	const COLUMN_EXISTS_SQL = "SELECT COLUMN_EXISTS($1 ,$2 ,$3)"
	rows, err := processor.PreparedStmt(COLUMN_EXISTS_SQL).Query(tableSchema, tableName, columName)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
	}
	var _state int64
	err = rows.Scan(&_state)
	if err != nil {
		return 0, err
	}
	return int(_state), nil
}

func (processor *ZeroDbAutoPostgresProcessor) ColumnDiff(
	tableSchema string,
	tableName string,
	columName string,
	isNullable string,
	columnType string,
	columnDefault string) (int, error) {
	const COLUMN_EXISTS_SQL = "SELECT COLUMN_DIFF($1 ,$2 ,$3, $4, $5, $6)"
	if strings.ToUpper(columnDefault) == ZDA_NULL {
		rows, err := processor.PreparedStmt(COLUMN_EXISTS_SQL).Query(tableSchema, tableName, columName, isNullable, columnType, nil)
		defer func() {
			if rows != nil {
				rows.Close()
			}
		}()
		if err != nil {
			return 0, err
		}
		if !rows.Next() {
			return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
		}
		var _state int64
		err = rows.Scan(&_state)
		if err != nil {
			return 0, err
		}
		return int(_state), nil
	} else {
		rows, err := processor.PreparedStmt(COLUMN_EXISTS_SQL).Query(tableSchema, tableName, columName, isNullable, columnType, columnDefault)
		defer func() {
			if rows != nil {
				rows.Close()
			}
		}()
		if err != nil {
			return 0, err
		}
		if !rows.Next() {
			return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
		}
		var _state int64
		err = rows.Scan(&_state)
		if err != nil {
			return 0, err
		}
		return int(_state), nil
	}
}

func (processor *ZeroDbAutoPostgresProcessor) DMLColumn(
	tableSchema string,
	tableName string,
	columName string,
	isNullable string,
	columnType string,
	columnDefault string) error {
	const DML_COLUMN_SQL = "SELECT DML_COLUMN($1 ,$2 ,$3, $4, $5, $6)"
	if strings.ToUpper(columnDefault) == "NULL" {
		_, err := processor.PreparedStmt(DML_COLUMN_SQL).Exec(tableSchema, tableName, columName, isNullable, columnType, nil)
		return err
	} else {
		_, err := processor.PreparedStmt(DML_COLUMN_SQL).Exec(tableSchema, tableName, columName, isNullable, columnType, columnDefault)
		return err
	}
}

func (processor *ZeroDbAutoPostgresProcessor) DropColumn(tableSchema string, tableName string, columName string) error {
	const DROP_COLUMN_SQL = "SELECT DROP_COLUMN($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_COLUMN_SQL).Exec(tableSchema, tableName, columName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) IndexExists(tableSchema string, tableName string, indexName string) (int, error) {
	const INDEX_EXISTS_SQL = "SELECT INDEX_EXISTS($1 ,$2 ,$3)"
	rows, err := processor.PreparedStmt(INDEX_EXISTS_SQL).Query(tableSchema, tableName, indexName)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
	}
	var _state int64
	err = rows.Scan(&_state)
	if err != nil {
		return 0, err
	}
	return int(_state), nil
}

func (processor *ZeroDbAutoPostgresProcessor) DMLConstraint(
	tableSchema string,
	tableName string,
	indexName string,
	defineIndexSQL string) error {
	const DML_CONSTRAINT_SQL = "SELECT DML_CONSTRAINT($1 ,$2 ,$3, $4)"
	_, err := processor.PreparedStmt(DML_CONSTRAINT_SQL).Exec(tableSchema, tableName, indexName, defineIndexSQL)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropConstraint(tableSchema string, tableName string, indexName string) error {
	const DROP_CONSTRAINT_SQL = "SELECT DROP_CONSTRAINT($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_CONSTRAINT_SQL).Exec(tableSchema, tableName, indexName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DMLIndex(
	tableSchema string,
	tableName string,
	indexName string) error {
	const DML_INDEX_SQL = "SELECT DML_INDEX($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DML_INDEX_SQL).Exec(tableSchema, tableName, indexName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropIndex(tableSchema string, tableName string, indexName string) error {
	const DROP_INDEX_SQL = "SELECT DROP_INDEX($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_INDEX_SQL).Exec(tableSchema, tableName, indexName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) TriggerExists(
	tableSchema string,
	tableName string,
	triggerTiming string,
	triggerEvent string,
	triggerName string,
	triggerAction string) (int, error) {
	const TRIGGER_EXISTS_SQL = "SELECT TRIGGER_EXISTS($1 ,$2 ,$3 ,$4 ,$5 ,$6)"
	rows, err := processor.PreparedStmt(TRIGGER_EXISTS_SQL).Query(tableSchema, tableName, triggerTiming, triggerEvent, triggerName, triggerAction)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
	}
	var _state int64
	err = rows.Scan(&_state)
	if err != nil {
		return 0, err
	}
	return int(_state), nil
}

func (processor *ZeroDbAutoPostgresProcessor) DMLTrigger(
	tableSchema string,
	tableName string,
	triggerTiming string,
	triggerEvent string,
	triggerName string,
	triggerAction string) error {
	const DML_TRIGGER_SQL = "SELECT DML_TRIGGER($1 ,$2 ,$3 ,$4 ,$5 ,$6)"
	_, err := processor.PreparedStmt(DML_TRIGGER_SQL).Exec(tableSchema, tableName, triggerTiming, triggerEvent, triggerName, triggerAction)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropTrigger(
	tableSchema string,
	tableName string,
	triggerName string) error {
	const DROP_TRIGGER_SQL = "SELECT DROP_TRIGGER($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_TRIGGER_SQL).Exec(tableSchema, tableName, triggerName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DMLPrimary(
	tableSchema string,
	tableName string,
	columnName string) error {
	const DML_PRIMARY_SQL = "SELECT DML_PRIMARY($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DML_PRIMARY_SQL).Exec(tableSchema, tableName, columnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropPrimary(
	tableSchema string,
	tableName string,
	columnName string) error {
	const DROP_PRIMARY_SQL = "SELECT DROP_PRIMARY($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_PRIMARY_SQL).Exec(tableSchema, tableName, columnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DMLUnique(
	tableSchema string,
	tableName string,
	columnName string) error {
	const DML_UNIQUE_SQL = "SELECT DML_UNIQUE($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DML_UNIQUE_SQL).Exec(tableSchema, tableName, columnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropUnique(
	tableSchema string,
	tableName string,
	columnName string) error {
	const DROP_UNIQUE_SQL = "SELECT DROP_UNIQUE($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_UNIQUE_SQL).Exec(tableSchema, tableName, columnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DMLForeign(
	tableSchema string,
	tableName string,
	columnName string,
	relTableName string,
	relColumnName string) error {
	const DML_FOREIGN_SQL = "SELECT DML_FOREIGN($1 ,$2 ,$3 ,$4 ,$5)"
	_, err := processor.PreparedStmt(DML_FOREIGN_SQL).Exec(tableSchema, tableName, columnName, relTableName, relColumnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropForeign(
	tableSchema string,
	tableName string,
	columnName string) error {
	const DROP_FOREIGN_SQL = "SELECT DROP_FOREIGN($1 ,$2 ,$3)"
	_, err := processor.PreparedStmt(DROP_FOREIGN_SQL).Exec(tableSchema, tableName, columnName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) TableExists(tableSchema string, tableName string) (int, error) {
	const TABLE_EXISTS_SQL = "SELECT TABLE_EXISTS($1 ,$2)"
	rows, err := processor.PreparedStmt(TABLE_EXISTS_SQL).Query(tableSchema, tableName)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, errors.New(fmt.Sprintf("query `COLUMN_EXISTS_SQL` failed"))
	}
	var _state int64
	err = rows.Scan(&_state)
	if err != nil {
		return 0, err
	}
	return int(_state), nil
}

func (processor *ZeroDbAutoPostgresProcessor) DMLTable(tableSchema string, tableName string) error {
	const DML_TABLE_SQL = "SELECT DML_TABLE($1 ,$2)"
	_, err := processor.PreparedStmt(DML_TABLE_SQL).Exec(tableSchema, tableName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) Create0Struct(tableSchema string, tableName string) error {
	const CREATE_0STRUCT_SQL = "SELECT create_0struct($1 ,$2)"
	_, err := processor.PreparedStmt(CREATE_0STRUCT_SQL).Exec(tableSchema, tableName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) Create0FlagStruct(tableSchema string, tableName string) error {
	const CREATE_0FLAGSTRUCT_SQL = "SELECT create_0flagstruct($1 ,$2)"
	_, err := processor.PreparedStmt(CREATE_0FLAGSTRUCT_SQL).Exec(tableSchema, tableName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DML0SPart(tableSchema string, tableName string) error {
	const DML_0SPART_SQL = "SELECT DML_0SPART($1 ,$2)"
	_, err := processor.PreparedStmt(DML_0SPART_SQL).Exec(tableSchema, tableName)
	return err
}

func (processor *ZeroDbAutoPostgresProcessor) DropPartitionTable(tableSchema string, tableName string) error {
	const DROP_PARTITION_TABLE_SQL = "SELECT DROP_PARTITION_TABLE($1 ,$2)"
	_, err := processor.PreparedStmt(DROP_PARTITION_TABLE_SQL).Exec(tableSchema, tableName)
	return err
}
