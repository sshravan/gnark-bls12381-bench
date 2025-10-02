package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	curve "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr/fft"
)

var logSamples = []int{9, 10, 11, 12}

// var logSamples = []int{9, 10, 11, 12, 18, 20, 22, 24}

func BenchmarkMultiExpG1(b *testing.B) {

	maxlogSample := slices.Max(logSamples)
	maxSample := 1 << maxlogSample

	sampleScalars := make([]fr.Element, maxSample)
	samplePoints := make([]curve.G1Affine, maxSample)

	fillBenchScalars(sampleScalars[:])
	fillBenchBasesG1(samplePoints[:])

	for _, logSample := range logSamples {
		sampleSize := 1 << logSample

		b.Run(fmt.Sprintf("logsize-%02d", logSample), func(b *testing.B) {
			var testPoint curve.G1Affine

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				testPoint.MultiExp(
					samplePoints[:sampleSize],
					sampleScalars[:sampleSize],
					ecc.MultiExpConfig{},
				)
			}
		})
	}
}

func BenchmarkMultiExpG2(b *testing.B) {

	maxlogSample := slices.Max(logSamples)
	maxSample := 1 << maxlogSample

	sampleScalars := make([]fr.Element, maxSample)
	samplePoints := make([]curve.G2Affine, maxSample)

	fillBenchScalars(sampleScalars[:])
	fillBenchBasesG2(samplePoints[:])

	for _, logSample := range logSamples {
		sampleSize := 1 << logSample

		b.Run(fmt.Sprintf("logsize-%02d", logSample), func(b *testing.B) {
			var testPoint curve.G2Affine

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				testPoint.MultiExp(
					samplePoints[:sampleSize],
					sampleScalars[:sampleSize],
					ecc.MultiExpConfig{},
				)
			}
		})
	}
}

func BenchmarkFFT(b *testing.B) {

	maxlogSample := slices.Max(logSamples)
	maxSample := 1 << maxlogSample

	sampleScalars := make([]fr.Element, maxSample)
	fillBenchScalars(sampleScalars[:])

	for _, logSample := range logSamples {
		sampleSize := 1 << logSample

		b.Run(fmt.Sprintf("logsize-%02d-basic", logSample), func(b *testing.B) {
			domain := fft.NewDomain(uint64(sampleSize))
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				domain.FFT(sampleScalars[:sampleSize], fft.DIT)
			}
		})
		b.Run(fmt.Sprintf("logsize-%02d-coset", logSample), func(b *testing.B) {
			domain := fft.NewDomain(uint64(sampleSize))
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				domain.FFT(sampleScalars[:sampleSize], fft.DIT, fft.OnCoset())
			}
		})
	}
}

func BenchmarkMillerLoop(b *testing.B) {

	var maxsize = 30

	pointsG1 := make([]curve.G1Affine, maxsize)
	pointsG2 := make([]curve.G2Affine, maxsize)

	fillBenchBasesG1(pointsG1[:])
	fillBenchBasesG2(pointsG2[:])

	for size := 1; size <= maxsize; size++ {
		b.Run(fmt.Sprintf("size-%02d", size), func(b *testing.B) {
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				_, _ = curve.MillerLoop(pointsG1[:size], pointsG2[:size])
			}
		})
	}
}

func BenchmarkFinalExp(b *testing.B) {
	var a curve.GT
	a.MustSetRandom()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.FinalExponentiation(&a)
	}
}
