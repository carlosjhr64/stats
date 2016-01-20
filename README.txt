package stats // import "github.com/carlosjhr64/stats"

const VERSION string = "0.0.0"

func NewSummary() *Summary

type Summary struct {
	Last float64
	N    float64
	Sum  float64
	Sum2 float64
}

func (s *Summary) Add(x, w float64)
func (s *Summary) RD(x, a, w float64)
func (s *Summary) LRD(x, a, w float64)
func (s *Summary) Mean() float64
func (s *Summary) RMS() float64
func (s *Summary) ERMS() float64
