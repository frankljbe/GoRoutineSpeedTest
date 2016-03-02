package tests

import (
  "github.com/frankljbe/speed_comparison_test/parallelslice"
  "testing"
  "time"
  "fmt"
);
var arraySize = 10000000;
func complexFunction(element float64) float64{
  for i :=0; i<1000; i++ {
    element = element/.9879;
  }
  return element;
}

func simpleFunction(element float64) float64{
  element = element/.9879;
  return element;
}


func TestSimpleFunctionSpeed(t *testing.T) {
  testSlice := parallelslice.NewParallelSlice(arraySize, simpleFunction);
  start := time.Now();
  testSlice.RunSerial();
  fmt.Println(time.Since(start).Seconds(), "Serial: ");
  start = time.Now();
  testSlice.RunParallelManyRoutines();
  fmt.Println(time.Since(start).Seconds(), "Parallel Many Routines: ");
  start = time.Now();
  testSlice.RunParallelOneRoutinePerCore();
  fmt.Println(time.Since(start).Seconds(), "Parallel One Routine Per Core: ");
}

func TestComplexFunctionSpeed(t *testing.T) {
  testSlice := parallelslice.NewParallelSlice(arraySize, complexFunction);
  start := time.Now();
  testSlice.RunSerial();
  fmt.Println(time.Since(start).Seconds(), "Serial: ");
  start = time.Now();
  testSlice.RunParallelManyRoutines();
  fmt.Println(time.Since(start).Seconds(), "Parallel Many Routines: ");
  start = time.Now();
  testSlice.RunParallelOneRoutinePerCore();
  fmt.Println(time.Since(start).Seconds(), "Parallel One Routine Per Core: ");
}
