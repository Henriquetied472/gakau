package gakau

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	pb "github.com/schollz/progressbar/v3"
)

func average(xy []int32, i int) int64 {
	total := 0.0
	xz := xy[i:]
	for _, v := range xz {
		total += float64(v)
	}
	total /= float64(len(xy))
	return int64(math.Floor(total))
}

func EncryptMessage(msg string) string {
	bar := pb.Default(100)
	smsg := strings.Split(msg, "")
	unit := len(smsg) / 10
	var res string

	for i, v := range smsg {
		rand.Seed(average([]rune(strings.Join(smsg, "")), i))
		jmp := rand.Intn(100)
		rn := []rune(v)
		res += strconv.FormatInt(int64(int(rn[0])+jmp), 36)

		bar.Add(int(unit))
		time.Sleep(20 * time.Millisecond)
	}

	res += "%#"

	return res
}
