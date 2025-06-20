/* 
    OlogN_binary_search.js
    Big-O: O(log n) – logarithmic
    Binary search halves the remaining search space each step. 
*/

import { performance } from "node:perf_hooks";

const binarySearch = (arr, target) => {
  let low = 0;
  let high = arr.length - 1;
  while (low <= high) {
    const mid = (low + high) >> 1;  // use bitwise right shift to divide by 2
    if (arr[mid] === target) return mid;
    arr[mid] < target ? (low = mid + 1) : (high = mid - 1);
  }
  return -1;
};

const main = () => {
  const N = 10000000;
  const bigSorted = Array.from({ length: N }, (_, i) => i);
  const t0 = performance.now();
  const idx = binarySearch(bigSorted, N - 1); // worst-case (near end)
  const t1 = performance.now();
  console.log("found at index:", idx);
  console.log(`O(log n) demo with N=${N} took ${(t1 - t0).toFixed(3)} ms`);
};

main();
