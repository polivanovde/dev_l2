package main

import (
	"testing"
)

func TestArgs(t *testing.T) {
	arg := []string{"", "log.log", "образом", "-i", "-n", "-C1"}
	file, word, isRev, isIgn, isCnt, isFull, isNum, after, before := args(arg)
	if file != "log.log" {
		t.Errorf("file expected is `log.log`. return is %v", file)
	}
	if word != "образом" {
		t.Errorf("word expected is `образом`. return is %v", word)
	}
	if !isIgn {
		t.Errorf("isIgn expected is `true`. return is %v", isIgn)
	}
	if !isNum {
		t.Errorf("isNum expected is `true`. return is %v", isNum)
	}
	if after != 1 {
		t.Errorf("after expected is `1`. return is %v", after)
	}
	if before != 1 {
		t.Errorf("before expected is `1`. return is %v", before)
	}
	//false
	if isRev {
		t.Errorf("isRev expected is `false`. return is %v", isRev)
	}
	if isCnt {
		t.Errorf("isIgn expected is `false`. return is %v", isCnt)
	}
	if isFull {
		t.Errorf("isFull expected is `false`. return is %v", isFull)
	}
}
func TestGrep(t *testing.T) {
	var (
		tCh  = make(chan map[int]string)
		rCh  = make(chan map[int]string)
		from = make(map[int]string)
	)
	from[1] = "Банальные, но неопровержимые выводы, а также базовые сценарии поведения пользователей освещают"
	from[2] = "чрезвычайным Образом интересные особенности картины в целом, однако конкретные выводы, разумеется,"
	from[3] = "объединены в целые кластеры себе подобных. Равным Образом, современная методология разработки позволяет"
	from[4] = "оценить значение форм воздействия. И нет сомнений, что независимые государства"
	from[5] = "функционально разнесены на независимые элементы."

	tCh <- from
	close(tCh)
	grep("образом", tCh, rCh, false, true, false, 0, 0)

	res := <-rCh
	close(rCh)
	if res == nil {
		t.Error("before expected is map. return is nill")
	}
}
