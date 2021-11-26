package gamespinmodel

import (
	"math/rand"
	"time"
)

type GameSpinPlay struct {
	Index int     `json:"index"`
	Award float32 `json:"award"`
}

func (GameSpinPlay) Awards() [8]float32 {
	return [8]float32{0, 1, 2, 0, 1.5, 0, 3, 0.5}
}

func (g *GameSpinPlay) RandomAward() {
	awards := g.Awards()
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(awards))

	g.Index = randIdx
	g.Award = awards[randIdx]
}
