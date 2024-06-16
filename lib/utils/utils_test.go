package utils

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_randomSlice(t *testing.T) {
    var test1 []int=[]int{1,2,3,4,5}

    pp.Println(RandomSliceArray(test1,3))
}

func Test_getDate(t *testing.T) {
    fmt.Println(GetCurrentDateSpecial())
}