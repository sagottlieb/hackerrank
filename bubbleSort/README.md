**Sorting: Bubble Sort**

Consider the following version of Bubble Sort:

```
for (int i = 0; i < n; i++) {
    
    for (int j = 0; j < n - 1; j++) {
        // Swap adjacent elements if they are in decreasing order
        if (a[j] > a[j + 1]) {
            swap(a[j], a[j + 1]);
        }
    }
    
}
```

**Task** 
Given an _n_-element array, _A_=_a[0]_,_a[1]_,...,_a[n-1]_ of distinct elements,
 sort array _A_ in ascending order using the Bubble Sort algorithm above. 
 Once sorted, print the following three lines:
 
1. Array is sorted in numSwaps swaps, where _numSwaps_ is the number of swaps that took place.
2. First Element: firstElement, where _firstElement_ is the first element in the sorted array.
3. Last Element: lastElement, where _lastElement_ is the last element in the sorted array.

**Hint:** To complete this challenge, you must add a variable that keeps a running tally of all swaps that occur during execution.

**Input Format**

The first line contains an integer, _n_, denoting the number of elements in array _A_. 
The second line contains _n_ space-separated integers describing the respective values of _a[0]_,_a[1]_,...,_a[n-1]_.

**Constraints**
- 2 <= _n_ <= 600
- 1 <= _a[i]_ <= 2 x 10^6, for all _i_ in [0, _n-1_]

**Output Format**
1. Array is sorted in numSwaps swaps, where _numSwaps_ is the number of swaps that took place.
2. First Element: firstElement, where _firstElement_ is the first element in the sorted array.
3. Last Element: lastElement, where _lastElement_ is the last element in the sorted array.
