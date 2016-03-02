package parallelslice;

import (
    "time"
    "math/rand"
    "runtime"
    "sync"
);

type ParallelSlice struct {
  size int;
  fakeData []float64;
  operationFunc func(float64) float64;
}

func NewParallelSlice(size int, operationFunc func(float64) float64) *ParallelSlice {
  source := rand.NewSource(time.Now().UnixNano())
  randomGen := rand.New(source);
  fakeData := make([]float64, size, size);
  for i, _ := range fakeData {
    fakeData[i] = randomGen.Float64();
  }
  return &ParallelSlice{
    size : size,
    fakeData: fakeData,
    operationFunc : operationFunc,
  }
}

func (this *ParallelSlice) RunSerial() {
  for _, element := range this.fakeData {
    this.operationFunc(element);
  }
}
func (this *ParallelSlice) RunParallelManyRoutines() {
  var wg sync.WaitGroup;
  wg.Add(this.size);
  for _, element := range this.fakeData {
    go func() {
      defer wg.Done();
      this.operationFunc(element);
    }()
  }
  wg.Wait();
}

func (this *ParallelSlice) RunParallelOneRoutinePerCore() {
  var numCores = runtime.NumCPU();
  var wg sync.WaitGroup;
  wg.Add(numCores);
  var elementsPerCore = this.size/numCores;
  for i := 0; i < numCores; i++ {
    go func(core int) {
      defer wg.Done();
      for j := core * elementsPerCore; j < ((core + 1) * elementsPerCore); j++ {
        this.operationFunc(this.fakeData[j]);
      }
    }(i)
  }
  wg.Wait();
}
