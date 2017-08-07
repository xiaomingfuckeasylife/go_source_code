package tar

import (
	"testing"
	"math"
)

func TestFitsInBase256(t *testing.T) {

	vectors := []struct{
		int int64
		width int
		ok bool
	}{
		{+1,8,true},
		{0,8,true},
		{-1,8,true},
		{1<<56 , 8 , false},
		{1<<56-1,8,true},
		{-1<<56-1,8, false},
		{121654,8,true},
		{-9849849,8,true},
		{math.MaxInt64,9,true},
		{math.MaxInt64,8,false},
		{0,9,true},
		{math.MinInt64,9,true},
		{math.MaxInt64,12,true},
		{0,12,true},
		{math.MinInt64,12,true},
	}

	for _ , v :=range vectors{
		ok := fitsInBase256(v.width,v.int)
		if ok != v.ok {
			t.Errorf("fitsInBase256(%d,%d):got %v , want %v",v.int, v.width , ok , v.ok)
		}
	}
}

func TestParseNumeric(t *testing.T) {

	//vector := []struct{
	//	in string
	//	want int64
	//	ok bool
	//}{
	//	//TODO not quite understand
	//	// Test basee-256(binary) encoded values.
	//	{"",0,true},
	//	{"0x80",0,true},
	//	{"0x80\x00",0,true},
	//	{"0x80\x00\x00",0,true},
	//	{"\xbf",1<<6-1,true},
	//
	//	// Test base-8(octal) encoded values .
	//}

}