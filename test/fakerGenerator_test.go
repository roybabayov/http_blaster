package test

import (
	"fmt"
	"github.com/v3io/http_blaster/httpblaster/data_generator"
	"testing"
)



func Test_facker_generator(t *testing.T) {
	gen := data_generator.Faker{}
	gen.GenerateRandomData()
}

func Test_facker_generator_to_igzemditem(t *testing.T) {
	gen := data_generator.Faker{}
	gen.GenerateRandomData()
	str:=  gen.ConvertToIgzEmdItemJson()
	fmt.Println(str)
}