# Nearest Larger Number

Given an array of numbers and an index i, return the index of the nearest larger number of the number at index i, where distance is measured in array indices.

If two distances to larger numbers are the equal, then return the one on the right. 

If the array at i doesn't have a nearest larger integer, then return -1.

<br/>

## Examples
Given [4, 1, 3, 5, 6] and index 0, you should return 3.

Given [13, 7, 5, 9, 4, 3, 20, 50, 1] and index 7, you should return -1.

<br/>

## Follow-up
If you can preprocess the array, can you do this in constant time?