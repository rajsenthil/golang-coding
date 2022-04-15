package algorithms

import (
	"log"
	"testing"
)

func TestGridTravel(t *testing.T) {
	gridTravel2_3 := gridTravel(2, 3)
	log.Printf("Grid travel on 2x3 = %v", gridTravel2_3)

	gridTravel3_3 := gridTravel(3, 3)
	log.Printf("Grid travel on 3x3 = %v", gridTravel3_3)

	gridTravel18_18 := gridTravel(18, 18)
	log.Printf("Grid travel on 18x18 = %v", gridTravel18_18)
}
