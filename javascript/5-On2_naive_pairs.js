/* 
    On2_naive_pairs.js
    Big-O: O(n²) – quadratic
    Compare every pair of elements (e.g., brute-force duplicate search). 
*/

import { performance } from "node:perf_hooks";

const hasDuplicate = (arr) => {
  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[i] === arr[j]) return true;
    }
  }
  return false;
};

const main = () => {
  const N = 30000; // 30k² ≈ 900M comparisons ⇒ slow
  const nums = Array.from({ length: N }, (_, i) => i);
  nums.push(N - 1); // guarantee duplicate at the tail (worst-case)
  const t0 = performance.now();
  const dup = hasDuplicate(nums);
  const t1 = performance.now();
  console.log("duplicate found:", dup);
  console.log(`O(n²) demo with N=${N} took ${(t1 - t0).toFixed(3)} ms`);
};

main();
