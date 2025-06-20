/* 
    O1_constant.js
    Big-O: O(1) – constant time
    We do the same amount of work regardless of N. 
*/

import { performance } from "node:perf_hooks";

const getMiddleElement = (arr) => {
  // Single arithmetic operation + array access ⇒ constant
  return arr[Math.floor(arr.length / 2)];
};

const main = () => {
  const big = Array.from({ length: 10000000 }, (_, i) => i);
  const t0 = performance.now();
  const mid = getMiddleElement(big);
  const t1 = performance.now();
  console.log("middle element:", mid);
  console.log(`O(1) demo took ${(t1 - t0).toFixed(3)} ms`);
};

main();
