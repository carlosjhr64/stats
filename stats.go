package stats

import "math"

const VERSION string = "0.0.0"

type Summary struct {
  Last float64
  N    float64
  Sum  float64
  Sum2 float64
}

func NewSummary() *Summary {
  return &Summary{0.0, 0.0, 0.0, 0.0}
}

func (s *Summary) Add(x, w float64) {
  wx := w*x
  s.Last = x
  s.N    += w
  s.Sum  += wx
  s.Sum2 += wx*x
}

func (s *Summary) RD(x, a, w float64) {
  d := (a - x) / a
  s.Add(d, w)
}

func (s *Summary) LRD(x, a, w float64) {
  lx, la := math.Log(x), math.Log(a)
  s.RD(lx, la, w)
}

func (s *Summary) Mean() float64 {
  return s.Sum / s.N
}

func (s *Summary) RMS() float64 {
  return math.Sqrt(s.Sum2/s.N)
}

func (s *Summary) ERMS() float64 {
  return math.Exp(s.RMS())
}
