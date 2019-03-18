package redis_protocol_parser

type ReplyValue struct {
	intValue    int
	stringValue string
}

func (val *ReplyValue) getInt() int {
	return val.intValue
}

func (val *ReplyValue) getString() string {
	return val.stringValue
}
