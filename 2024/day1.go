package challenges2024

import (
	"fmt"
	"math"
	"slices"

	"github.com/dulranga/advent-of-code/helpers"
)

/*
Day 1: Historian Hysteria

The Chief Historian is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.

As each location is checked, they will mark it on their list with a star. They figure the Chief Historian must be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get fifty stars on their list before Santa takes off on December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently empty. Eventually, someone decides that the best place to check first would be the Chief Historian's office.

Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?

Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the location ID. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.

There's just one problem: by holding the two lists up side by side (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?

For example:
```
3   4
4   3
2   5
1   3
3   9
3   3
```
Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example, if you pair up a 3 from the left list with a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.

In the example list above, the pairs and distances would be as follows:

- The smallest number in the left list is 1, and the smallest number in the right list is 3. The distance between them is 2.
- The second-smallest number in the left list is 2, and the second-smallest number in the right list is another 3. The distance between them is 1.
- The third-smallest number in both lists is 3, so the distance between them is 0.
- The next numbers to pair up are 3 and 4, a distance of 1.
- The fifth-smallest numbers in each list are 3 and 5, a distance of 2.
- Finally, the largest number in the left list is 4, while the largest number in the right list is 9; these are a distance 5 apart.
- To find the total distance between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is 2 + 1 + 0 + 1 + 2 + 5, a total distance of 11!

Your actual left and right lists contain many location IDs. What is the total distance between your lists?
*/
func Day1() {

	input := helpers.ParseInputFile("2024/inputs/day1.txt")
	inputSize := len(input)

	// pre occupy two arrays to speed things up
	left := make([]int, inputSize)
	right := make([]int, inputSize)

	// parsing location ids from each statement
	for i, row := range input {
		var l int
		var r int

		fmt.Sscanf(row, "%d   %d", &l, &r)

		// we can only do this cuz the array is pre defined with correct size
		left[i] = l
		right[i] = r

	}

	slices.Sort(left)
	slices.Sort(right)

	fmt.Printf("sum: %v\n", calculateDistancePart2(left, right))
}

func calculateDistancePart1(left, right []int) int {
	var sum int

	for i := range left {
		l := left[i]
		r := right[i]

		sum += int(math.Abs(float64(r - l)))

	}

	return sum
}

/*
--- Part Two ---
Your analysis only confirmed what everyone feared: the two lists of location IDs are indeed very different.

Or are they?

The Historians can't agree on which group made the mistakes or how to read most of the Chief's handwriting, but in the commotion you notice an interesting detail: a lot of location IDs appear in both lists! Maybe the other numbers aren't location IDs at all but rather misinterpreted handwriting.

This time, you'll need to figure out exactly how often each number from the left list appears in the right list. Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.

Here are the same example lists again:

3   4
4   3
2   5
1   3
3   9
3   3
For these example lists, here is the process of finding the similarity score:

The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
The fourth number, 1, also does not appear in the right list.
The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
The last number, 3, appears in the right list three times; the similarity score again increases by 9.
So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).

Once again consider your left and right lists. What is their similarity score?
*/
func calculateDistancePart2(left, right []int) int {

	previousIterationLookupEnd := 0
	totalSimilarityScore := 0

	for _, item := range left {

		itemRepeatCount := 0
		lookup := BinaryLookup(right, item, 0)

		if lookup == -1 {
			// item not found in the list
			continue
		}
		searchableArray := right[previousIterationLookupEnd : lookup+1]

		if right[lookup] != item {
			// if the array ends with similar items we need the items to be counted in the next iteration as well
			// so we do NOT truncate the array in this case
			previousIterationLookupEnd = lookup

		}

		for _, searchItem := range searchableArray {
			if item == searchItem {
				itemRepeatCount++
			}
		}

		totalSimilarityScore += item * itemRepeatCount

	}
	return totalSimilarityScore
}

// this function gives the last index of a sorted array where you can find the lookupItem
// this uses binary search to look for the item
func BinaryLookup(list []int, lookupItem int, __startingIdx int) int {

	size := len(list)

	midItemIdx := int(size / 2)
	midItem := list[midItemIdx]
	lastIdx := size - 1

	if lookupItem > list[lastIdx] || lookupItem < list[0] {
		// not found
		// lookup item is either smaller or greater than the list
		return -1
	}
	if lookupItem == list[lastIdx] {
		// last item in the array
		return __startingIdx + lastIdx
	}

	if lookupItem < midItem {
		return BinaryLookup(list[:midItemIdx], lookupItem, __startingIdx)

	} else if lookupItem > midItem {
		return BinaryLookup(list[midItemIdx:], lookupItem, __startingIdx+midItemIdx)

	} else {
		// lookup item is exactly equal
		var current = midItemIdx
		for {
			// looping to find if multiple copies included
			current++
			if lookupItem < list[current] {
				return current + __startingIdx - 1
			}
		}
	}

}
