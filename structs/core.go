package structs

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

const (
	XSAC_NAME        = "xsacname"
	XSAC_PROP        = "xsacprop"
	XSAC_REF         = "xsacref"
	XSAC_KEY         = "xsackey"
	XSAC_REF_INSPECT = "inspect"
	XSAC_CHILD       = "xsacchild"
	XSAC_FIELD       = "xsacfield"

	XHTTP_OPT = "xhttpopt"

	XSAC_NULL = "NULL"
	XSAC_YES  = "YES"
	XSAC_NO   = "NO"
)

func FindMetaType(t reflect.Type) reflect.Type {
	metaType := t
	for metaType.Kind() == reflect.Pointer || metaType.Kind() == reflect.Slice {
		metaType = metaType.Elem()
	}
	return metaType
}

func FindStructFieldMetaType(fields reflect.StructField) reflect.Type {
	return FindMetaType(fields.Type)
}

type ZeroCoreStructs struct {
	ZeroMeta

	ID         string                 `json:"id,omitempty" xhttpopt:"OX"`
	CreateTime *Time                  `json:"createTime,omitempty" xhttpopt:"XX"`
	UpdateTime *Time                  `json:"updateTime,omitempty" xhttpopt:"XX"`
	Features   map[string]interface{} `json:"features,omitempty" xhttpopt:"OO"`
	Flag       int                    `json:"-"`
}

func (e *ZeroCoreStructs) XsacPrimaryType() string { return "UUID" }
func (e *ZeroCoreStructs) XsacDataSource() string  { return "" }
func (e *ZeroCoreStructs) XsacDbName() string      { return "" }
func (e *ZeroCoreStructs) XsacTableName() string   { panic("not implemented") }
func (e *ZeroCoreStructs) XsacDeleteOpt() byte     { return 0b10000000 }
func (e *ZeroCoreStructs) XsacPartition() string   { return XSAC_PARTITION_NONE }

func (e *ZeroCoreStructs) findXsacEntry(fields reflect.StructField) []*ZeroXsacEntry {
	entries := make([]*ZeroXsacEntry, 0)

	xrProp := fields.Tag.Get(XSAC_PROP)
	if len(xrProp) > 0 {
		xrPropItems := strings.Split(xrProp, ",")
		if len(xrPropItems) == 3 {
			columnName := fields.Tag.Get(XSAC_NAME)
			if len(columnName) <= 0 {
				columnName = exHumpToLine(fields.Name)
			}

			entries = append(entries, NewColumn(
				e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName(),
				columnName, xrPropItems[0], xrPropItems[1], xrPropItems[2]))

			xsacKey := fields.Tag.Get(XSAC_KEY)
			if len(xsacKey) > 0 {
				if strings.HasPrefix(xsacKey, "foreign") {
					xrKeyItems := strings.Split(xsacKey, ",")
					if len(xrKeyItems) == 3 {
						entries = append(entries, NewForeignKey(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName(), columnName, xrKeyItems[1], xrKeyItems[2]))
					}
				} else {
					switch xsacKey {
					case "primary":
						entries = append(entries, NewPrimaryKey(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName(), columnName))
					case "key":
						entries = append(entries, NewKey(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName(), columnName))
					case "unique":
						entries = append(entries, NewUniqueKey(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName(), columnName))
					}
				}
			}
		}
	}
	return entries
}

func (e *ZeroCoreStructs) readXsacEntries(xrType reflect.Type) []*ZeroXsacEntry {
	entries := make([]*ZeroXsacEntry, 0)
	for i := 0; i < xrType.NumField(); i++ {
		if xrType.Field(i).Anonymous {
			entries = append(entries, e.readXsacEntries(xrType.Field(i).Type)...)
		} else {
			entries = append(entries, e.findXsacEntry(xrType.Field(i))...)
		}
	}
	return entries
}

func (e *ZeroCoreStructs) XsacDeclares() ZeroXsacEntrySet {
	entries := make([]*ZeroXsacEntry, 0)
	if e.This().(ZeroXsacDeclares).XsacDeleteOpt()&0b10000000 == 0b10000000 {
		entries = append(entries, NewTable0s(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName()))
	} else {
		entries = append(entries, NewTable0fs(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName()))
	}
	entries = append(entries, e.readXsacEntries(reflect.TypeOf(e.This()).Elem())...)
	return entries
}

func (e *ZeroCoreStructs) findXsacRefEntry(fields reflect.StructField) []*ZeroXsacEntry {
	entries := make([]*ZeroXsacEntry, 0)
	xrRefProp := fields.Tag.Get(XSAC_REF)
	metaType := FindStructFieldMetaType(fields)
	if len(xrRefProp) > 0 {
		xrRefProppItems := strings.Split(xrRefProp, ",")
		if len(xrRefProppItems) == 4 && xrRefProppItems[3] == XSAC_REF_INSPECT {
			entries = append(entries, NewTable(e.This().(ZeroXsacDeclares).XsacDbName(), xrRefProppItems[0]))
			entries = append(entries, NewColumn(e.This().(ZeroXsacDeclares).XsacDbName(), xrRefProppItems[0], xrRefProppItems[1], XSAC_NO, e.XsacPrimaryType(), XSAC_NULL))
			entries = append(entries, NewColumn(e.This().(ZeroXsacDeclares).XsacDbName(), xrRefProppItems[0], xrRefProppItems[2], XSAC_NO, e.XsacPrimaryType(), XSAC_NULL))
			entries = append(entries, NewForeignKey(e.This().(ZeroXsacDeclares).XsacDbName(), xrRefProppItems[0], xrRefProppItems[1], e.This().(ZeroXsacDeclares).XsacTableName(), "id"))
			entries = append(entries, NewForeignKey(e.This().(ZeroXsacDeclares).XsacDbName(), xrRefProppItems[0], xrRefProppItems[2], reflect.New(metaType).Interface().(ZeroXsacDeclares).XsacTableName(), "id"))
		}
	}
	return entries
}

func (e *ZeroCoreStructs) readXsacRefEntries(xrType reflect.Type) []*ZeroXsacEntry {
	entries := make([]*ZeroXsacEntry, 0)
	for i := 0; i < xrType.NumField(); i++ {
		if xrType.Field(i).Anonymous {
			entries = append(entries, e.readXsacRefEntries(xrType.Field(i).Type)...)
		} else {
			entries = append(entries, e.findXsacRefEntry(xrType.Field(i))...)
		}
	}
	return entries
}

func (e *ZeroCoreStructs) XsacRefDeclares() ZeroXsacEntrySet {
	entries := e.readXsacRefEntries(reflect.TypeOf(e.This()).Elem())
	switch e.This().(ZeroXsacDeclares).XsacPartition() {
	case XSAC_PARTITION_YEAR:
		entries = append(entries, NewYearPartition(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName()))
	case XSAC_PARTITION_MONTH:
		entries = append(entries, NewMonthPartition(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName()))
	case XSAC_PARTITION_DAY:
		entries = append(entries, NewDayPartition(e.This().(ZeroXsacDeclares).XsacDbName(), e.This().(ZeroXsacDeclares).XsacTableName()))
	}
	return entries
}

func (e *ZeroCoreStructs) findXopFields(xrType reflect.Type, ignore bool) ZeroXsacFieldSet {
	fields := make([]*ZeroXsacField, 0)
	for i := 0; i < xrType.NumField(); i++ {
		if xrType.Field(i).Anonymous {
			fields = append(fields, e.findXopFields(xrType.Field(i).Type, ignore)...)
		} else if len(xrType.Field(i).Tag.Get(XHTTP_OPT)) > 0 {
			fields = append(fields, NewXsacField(xrType.Field(i), ignore))
		}
	}
	return fields
}

func (e *ZeroCoreStructs) XsacFields(xm ...int) ZeroXsacFieldSet {
	fields := make([]*ZeroXsacField, 0)
	fields = append(fields, e.findXopFields(reflect.TypeOf(e.This()).Elem(), len(xm) > 0)...)
	return fields
}

func (e *ZeroCoreStructs) InitDefault() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	newDate := Time(time.Now())
	e.ID = uid.String()
	e.CreateTime = &newDate
	e.UpdateTime = &newDate
	if e.Features == nil {
		e.Features = make(map[string]interface{})
	}
	return nil
}

func (e *ZeroCoreStructs) JSONFeature() string {
	if e.Features == nil {
		e.Features = make(map[string]interface{})
	}
	mjson, _ := json.Marshal(e.Features)
	return string(mjson)
}

func (e *ZeroCoreStructs) JSONFeatureWithString(jsonString string) {
	var jsonMap map[string]interface{}
	_ = json.Unmarshal([]byte(jsonString), &jsonMap)
	e.Features = jsonMap
}

func (e *ZeroCoreStructs) LoadRowData(rowmap map[string]interface{}) {
	e.ID = ParseStringField(rowmap, "id")
	e.CreateTime = ParseDateField(rowmap, "create_time")
	e.UpdateTime = ParseDateField(rowmap, "update_time")
	e.Features = ParseJSONField(rowmap, "features")
	e.Flag = ParseIntField(rowmap, "flag")
}

func (e *ZeroCoreStructs) String() string {
	mjson, _ := json.Marshal(e)
	return string(mjson)
}

func (e *ZeroCoreStructs) Map() map[string]interface{} {
	mjson, _ := json.Marshal(e)
	var jsonMap map[string]interface{}
	_ = json.Unmarshal([]byte(mjson), &jsonMap)
	return jsonMap
}

func ParseStringField(rowmap map[string]interface{}, fieldName string) string {
	v, ok := rowmap[fieldName]
	if ok {
		if reflect.TypeOf(v).Kind() == reflect.String {
			return v.(string)
		} else {
			return string(v.([]uint8))
		}
	}
	return ""
}

func ParseDateField(rowmap map[string]interface{}, fieldName string) *Time {
	_, ok := rowmap[fieldName]
	if ok {
		rowdata := Time(rowmap[fieldName].(time.Time))
		return &rowdata
	}
	return nil
}

func ParseJSONField(rowmap map[string]interface{}, fieldName string) map[string]interface{} {
	datastr := ParseStringField(rowmap, fieldName)
	if len(datastr) > 0 {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(datastr), &jsonMap)
		return jsonMap
	}
	return nil
}

func ParseIntField(rowmap map[string]interface{}, fieldName string) int {
	_, ok := rowmap[fieldName]
	if ok {
		return int(rowmap[fieldName].(int64))
	}
	return 0
}

func ParseFloatField(rowmap map[string]interface{}, fieldName string) float64 {
	_, ok := rowmap[fieldName]
	if ok {
		return rowmap[fieldName].(float64)
	}
	return 0
}

func ParseBytesField(rowmap map[string]interface{}, fieldName string) []byte {
	_, ok := rowmap[fieldName]
	if ok {
		return rowmap[fieldName].([]uint8)
	}
	return nil
}
