package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mRand "math/rand"
	"strconv"
	"time"
)

func Shuffle(slice [][3]interface{}) {
	r := mRand.New(mRand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func RandInt(n int64) int64 {
	result, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		panic(err)
	}
	return result.Int64()
}

func RandFloat32() float32 {
	result, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", float32(result.Int64())*0.1), 32)
	if err != nil {
		panic(err)
	}

	return float32(value)
}
