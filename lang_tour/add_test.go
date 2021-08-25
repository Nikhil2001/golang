package main

import "testing"
import "fmt"

func TestAddNumber(t *testing.T) {
   var tests = [][]int{{1,2,3},{1,22,23},{11,2,33}}
   for _,test := range tests {
      testName := fmt.Sprintf("test %v + %v",test[0],test[1])
      t.Run(testName,func(t *testing.T) {
         if test[2] != add2(test[0],test[1]){
          t.Errorf("test %v + %v != %v",test[0],test[1],test[2])
         }
      })
   }

}