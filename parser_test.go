package redis_protocol_parser

import "testing"

func Test_Decode(t *testing.T) {
	if err, rType, val := decode(":30\r\n"); err != nil {
		t.Fatal(err)
	} else {
		if rType != IntegerReply || val.getInt() != 30 {
			t.Fatal("unit testing failed")
		} else {
			t.Log(rType, val)
		}
	}

	if err, rType, val := decode("+OK\r\n"); err != nil {
		t.Fatal(err)
	} else {
		if rType != StatusReply || val.getString() != "OK" {
			t.Fatal("unit testing failed")
		} else {
			t.Log(rType, val)
		}
	}
}
