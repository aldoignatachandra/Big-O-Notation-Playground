/* 
    OnLogN_merge_sort.js
    Big-O: O(n log n)
    Classic merge-sort: log n levels of splitting × n work each level. 
*/

import { performance } from "node:perf_hooks";

const mergeSort = (arr) => {
  if (arr.length <= 1) return arr;
  const mid = arr.length >> 1;
  const left = mergeSort(arr.slice(0, mid));
  const right = mergeSort(arr.slice(mid));
  return merge(left, right);
};

const merge = (a, b) => {
  const out = [];
  while (a.length && b.length) {
    out.push(a[0] <= b[0] ? a.shift() : b.shift());
  }
  return [...out, ...a, ...b];
};

const main = () => {
  const N = 200000; // use smaller N, merge-sort is heavy
  const nums = Array.from({ length: N }, () => Math.random());
  const t0 = performance.now();
  mergeSort(nums);
  const t1 = performance.now();
  console.log(`O(n log n) demo with N=${N} took ${(t1 - t0).toFixed(3)} ms`);
};

main();
