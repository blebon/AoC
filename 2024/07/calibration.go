package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/blebon/AoC/2024/util"
	log "github.com/sirupsen/logrus"
)

type cal struct {
	Key  int64
	Vals []int64
}

func readCalibrationFile(f string) *[]cal {
	s, err := util.ReadFile(f)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	r := []cal{}
	for _, l := range s {
		kv := strings.Split(l, ":")
		k, _ := strconv.Atoi(kv[0])
		var v []int64
		for _, sv := range strings.Fields(kv[1]) {
			i, _ := strconv.Atoi(sv)
			v = append(v, int64(i))
		}
		r = append(r, cal{Key: int64(k), Vals: v})
	}

	return &r
}

func eval[V int64](k V, v []V, withConcat bool) V {
	if len(v) == 1 {
		if v[0] == k {
			return k
		}
		return 0
	}

	type fn func(i, j V) V
	add := func(i, j V) V { return i + j }
	mul := func(i, j V) V { return i * j }
	con := func(i, j V) V {
		if withConcat {
			concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", v[0], v[1]))
			return V(concat)
		}
		return 0
	}
	ops := []fn{add, mul, con}

	for _, op := range ops {
		n := eval(k, append([]V{op(v[0], v[1])}, v[2:]...), withConcat)
		if n != 0 {
			return n
		}
	}

	return 0
}

func getCalibration(f string, withConcat bool) int64 {
	c := readCalibrationFile(f)

	ch := make(chan int64)
	go func() {
		defer close(ch)
		for _, v := range *c {
			ch <- eval(v.Key, v.Vals, withConcat)
		}
	}()

	var ans atomic.Int64
	for v := range ch {
		ans.Add(v)
	}
	return ans.Load()
}
