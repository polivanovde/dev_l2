package main

import (
    "testing"
)

func TestUnPack(t *testing.T) {
    var got interface{}
    got = UnPack(`qwe\452/f2\\\\\\2`)
    if got != `qwe444444/ff\\\\` {
        t.Errorf("got expected is `qwe44444/ff\\\\`. return is %v", got)
    }
}