package redis_protocol_parser

import (
	"errors"
	"strconv"
)

type ReplyType uint8

// redis返回值类型
const (
	ErrorReply ReplyType = iota
	StatusReply
	IntegerReply
	BulkReply
	MultiBulkReply
	OthersReply
)

// 解析字符串函数
func decode(protocol string) (error, ReplyType, ReplyValue) {
	val := ReplyValue{
		intValue:    0,
		stringValue: "",
	}

	if err := validityChecker(protocol); err != nil {
		return err, OthersReply, val
	}
	err, replyType := replyTypeChecker(protocol)
	if err != nil {
		return err, OthersReply, val
	}
	replyVal := protocol[1 : len(protocol)-2]
	switch replyType {
	case ErrorReply:
		return errors.New(replyVal), replyType, val
	case StatusReply:
		val.stringValue = replyVal
		return nil, replyType, val
	case IntegerReply:
		if intVal, err := strconv.Atoi(replyVal); err != nil {
			return err, replyType, val
		} else {
			val.intValue = intVal
			return nil, replyType, val
		}
	}
	return nil, replyType, val
}

// 判断返回值类型
func replyTypeChecker(protocol string) (error, ReplyType) {
	//状态回复（status reply）的第一个字节是 "+"
	//错误回复（error reply）的第一个字节是 "-"
	//整数回复（integer reply）的第一个字节是 ":"
	//批量回复（bulk reply）的第一个字节是 "$"
	//多条批量回复（multi bulk reply）的第一个字节是 "*"
	firstLetter := protocol[0:1]
	rType := StatusReply
	switch firstLetter {
	case "+":
		rType = StatusReply
	case "-":
		rType = ErrorReply
	case ":":
		rType = IntegerReply
	default:
		return errors.New("unsupported reply type"), OthersReply
	}
	return nil, rType
}

// 合法性验证
func validityChecker(protocol string) error {
	rType := protocol[:1]
	if rType != "+" && rType != "-" && rType != ":" && rType != "$" && rType != "*" {
		return errors.New("protocol parse failed")
	}
	if protocol[len(protocol)-2:] != "\r\n" {
		return errors.New("protocol parse failed")
	}
	return nil
}
