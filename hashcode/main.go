package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type slide []Picture
type Result []slide

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Picture, len(input))
	perm := r.Perm(len(input))
	for i, randIndex := range perm {
		ret[i] = input[randIndex]
	}

	var res Result
	var lastVert Picture
	var haveLastVert bool
	for _, i := range ret {
		var sl slide
		if i.vert {
			if !haveLastVert {
				lastVert = i
				haveLastVert = true
			} else {
				sl = slide{lastVert, i}
				haveLastVert = false
			}
		} else {
			sl = slide{i}
		}
		if len(sl) != 0 {
			res = append(res, sl)
		}
	}

	score := Score(res)

	//var finalRes Result

	for i := 0; i < 60000000; i++ {

		if i%50000 == 0 {
			fmt.Printf("%d %d\n", i, score)
		}
		if i%100000 == 0 {
			Dump(res, os.Args[1])
		}

		r := rand.New(rand.NewSource(int64(i)))
		r1 := r.Intn(len(res)-2) + 1
		r2 := r.Intn(len(res)-2) + 1

		if r1 == r2 {
			continue
		}
		if r2 < r1 {
			r1, r2 = r2, r1
		}

		new1 := getSingleScore(res[r1], res[r2+1])
		new2 := getSingleScore(res[r2], res[r1+1])
		new3 := getSingleScore(res[r1-1], res[r2])
		new4 := getSingleScore(res[r2-1], res[r1])
		new := new1 + new2 + new3 + new4

		old1 := getSingleScore(res[r1], res[r1+1])
		old2 := getSingleScore(res[r2], res[r2+1])
		old3 := getSingleScore(res[r1-1], res[r1])
		old4 := getSingleScore(res[r2-1], res[r2])
		old := old1 + old2 + old3 + old4

		diff := new - old

		reverseNew := new1 + new3
		reverseOld := old3 + old2
		reverseDiff := reverseNew - reverseOld

		if reverseDiff > diff && reverseDiff > 0 {
			for left, right := r1, r2; left < right; left, right = left+1, right-1 {
				res[left], res[right] = res[right], res[left]
			}
			score += reverseDiff
		} else if diff > 0 {
			res[r1], res[r2] = res[r2], res[r1]
			score += diff
		}
	}

	fmt.Println("final score")
	score = Score(res)
	fmt.Println(score)
}

func Dump(out Result, file string) {
	output := fmt.Sprintf("./%s.out", strings.TrimSuffix(file, ".in"))
	f, _ := os.Create(output)
	defer f.Close()

	w := bufio.NewWriter(f)

	num := len(out)
	w.WriteString(fmt.Sprintf("%d", num))
	w.WriteString("\n")
	for _, c := range out {
		if len(c) == 1 {
			w.WriteString(fmt.Sprintf("%d", c[0].index))
		} else {
			w.WriteString(fmt.Sprintf("%d %d", c[0].index, c[1].index))
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func ParseInt(v string) int {
	x, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return int(x)
}

type Picture struct {
	index int
	vert  bool
	tags  map[string]bool
}

func ParseInput(path string) []Picture {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	numPictures := ParseInt(scanner.Text())

	pictures := []Picture{}

	for i := 0; i < numPictures; i++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")
		var picture Picture
		if tmp[0] == "H" {
			picture.vert = false
		} else {
			picture.vert = true
		}

		picture.index = i
		slice := tmp[2:]
		s := make(map[string]bool, len(slice))
		for _, t := range slice {
			s[t] = true
		}
		picture.tags = s
		pictures = append(pictures, picture)
	}

	return pictures
}

func Score(res Result) int {
	score := 0
	for i := 0; i < len(res)-1; i++ {

		min := getSingleScore(res[i], res[i+1])
		score += min
	}
	return score
}

func getSingleScore(left slide, right slide) int {
	thisTags := getTagSet(left)
	nextTags := getTagSet(right)

	intersect := thisTags.intersect(nextTags)

	if intersect == 0 {
		return 0
	}

	l := len(thisTags) - intersect

	if l == 0 {
		return 0
	}

	r := len(nextTags) - intersect

	min := l
	if intersect < min {
		min = intersect
	}

	if r < min {
		min = r
	}

	return min
}

type tagSet map[string]bool

func getTagSet(sl slide) tagSet {
	res := map[string]bool{}
	res = sl[0].tags
	if len(sl) == 2 {
		for t := range sl[1].tags {
			res[t] = true
		}
	}
	return res
}

func (ts tagSet) intersect(right tagSet) int {
	score := 0
	for k := range ts {
		_, ok := right[k]
		if ok {
			score++
		}
	}
	return score
}
