package main

import (
"testing"
"time"
)

func TestNtpTime(t *testing.T) {
    var got interface{}
    got = NtpTime()
    switch v := got.(type) {
    	case time.Time:
        default:
            t.Errorf("got Type is %T", v)
    }
}