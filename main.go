package main

import "fmt"

func main() {
	priceTokenA := 449.8
	priceTokenB := 1.0
	// priceTokenA * amntTokenA + priceTokenB * amntTokenB = totalAmntUsd
	totalAmntUsd := 2050.0
	shareAmntTokenA := 0.85148
	shareAmntTokenB := 1592.004629
	// share * (priceTokenA * shareAmntTokenA + priceTokenB * shareAmntTokenB) = totalAmntUsd
	share := totalAmntUsd / (priceTokenA*shareAmntTokenA + priceTokenB*shareAmntTokenB)
	amntTokenA := share * shareAmntTokenA
	amntTokenB := share * shareAmntTokenB
	amntTokenAUsd := amntTokenA * priceTokenA
	amntTokenBUsd := amntTokenB * priceTokenB

	fmt.Printf("amntTokenA: %v\n", amntTokenA)
	fmt.Printf("amntTokenB: %v\n", amntTokenB)
	fmt.Printf("amntTokenAUsd: %v\n", amntTokenAUsd)
	fmt.Printf("amntTokenBUsd: %v\n", amntTokenBUsd)
}
