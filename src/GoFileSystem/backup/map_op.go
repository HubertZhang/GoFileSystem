package main

const (
	GREETING       = 0
	KV_INSERT      = 1
	KV_DELETE      = 2
	KV_GET         = 3
	KV_UPDATE      = 4
	KVMAN_COUNTKEY = 5
	KVMAN_DUMP     = 6
	KVMAN_SHUTDOWN = 7
)

var table = make(map[string]string)

func Perform(op *Op) {
	switch op.OpCode {
	case KV_INSERT:
		table[op.Key] = op.Value
	case KV_UPDATE:
		table[op.Key] = op.Value
	case KV_DELETE:
		delete(table, op.Key)
	}
}
