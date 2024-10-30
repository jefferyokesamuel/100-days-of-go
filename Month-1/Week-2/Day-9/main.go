package main

import "fmt"

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    n := len(height)
    leftMax := make([]int, n)
    rightMax := make([]int, n)
    totalWater := 0

    // Fill leftMax array
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }

    // Fill rightMax array
    rightMax[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }

    // Calculate total trapped water
    for i := 0; i < n; i++ {
        waterLevel := min(leftMax[i], rightMax[i])
        if waterLevel > height[i] {
            totalWater += waterLevel - height[i]
        }
    }

    return totalWater
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    fmt.Println("Trapped water:", trap(height)) // Output: 6
}
