/*
LEET CODE PROBLEM
56. Merge Intervals

Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.

APPROACH:
Sort the intervals on Start
Then merge
*/

package main

import ("fmt"
        "sort"
)

// Defining a tuple type to hold the intervals as go doesn't have a tuple by default
type Interval struct {
    Start,End int
}

func SortIntervalsByStart(intervals []Interval) []Interval {
    var indexKeyMap = make(map[int]Interval) //Each item is mapped to an index which is the key
    var keyIndexMap = make(map[int][]int) //Each key is mapped to multiple indices, the values at these indices contain the intervals
    var sortedIntervals []Interval
    var keyArray sort.IntSlice
    for idx,interval := range intervals {
        indexKeyMap[idx] = interval
        keyIndexMap[interval.Start] = append(keyIndexMap[interval.Start],idx)
        keyArray = append(keyArray,interval.Start)
    }
    // Sort the key array and iterate over it to create the sorted intervals array
    keyArray.Sort()
    for _,val := range keyArray {
        intervalIdx := keyIndexMap[val][len(keyIndexMap[val])-1]
        interval := indexKeyMap[intervalIdx]
        keyIndexMap[val] = keyIndexMap[val][:len(keyIndexMap[val])-1]
        sortedIntervals = append(sortedIntervals,interval)
    }
    return sortedIntervals
}

func MergeOverlappingIntervals(intervals []Interval) []Interval {
    // Sort the array based on Start value of each interval
    intervals = SortIntervalsByStart(intervals)
    if len(intervals) == 0 {
        return intervals
    }
    var master Interval = intervals[0]
    var intervalsArray []Interval
    for _,val := range intervals {
        if (val.Start <= master.End && val.Start >= master.Start && val.End >= master.End)  {
            master.End = val.End
        } else if (val.Start > master.End) {
            intervalsArray = append(intervalsArray,master)
            master = val
        }
    }
    intervalsArray = append(intervalsArray,master)
    return intervalsArray
}


func main() {
    array := []Interval{{1,5},{0,7},{2,6},{5,8},{4,12},{0,15}}
    //fmt.Println(array)
    mergedIntervals := MergeOverlappingIntervals(array)
    fmt.Println(mergedIntervals)
}