package compare

func MergeRecords(arr1, arr2 []Record) []Record {
	merged := make([]Record, 0, len(arr1)+len(arr2))
	merged = append(merged, arr1...)
	merged = append(merged, arr2...)

	return merged
}

func IsRecordInList(record Record, recordList []Record) bool {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return true
		}
	}
	return false
}

func GetRecordInList(record Record, recordList []Record) Record {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return r
		}
	}
	return Record{}
}

func IsRecordsEqual(informedRecord, record Record) bool {
	return informedRecord.Player == record.Player && informedRecord.Sport == record.Sport && informedRecord.StatType == record.StatType && informedRecord.Team == record.Team &&
		informedRecord.Opponent == record.Opponent
}

func IsUnderdogRecordInList(record Record, recordList []Record) bool {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return true
		}
	}
	return false
}

func IsPrizepricksRecordInList(record Record, recordList []Record) bool {
	for _, r := range recordList {
		if IsRecordsEqual(r, record) {
			return true
		}
	}
	return false
}
