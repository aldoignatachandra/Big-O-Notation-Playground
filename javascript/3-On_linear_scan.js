/* 
    On_linear_scan.js
    Big-O: O(n) – linear
    Iterates once over the entire collection (sum). 
*/

import { performance } from "node:perf_hooks";

const sum = (arr) => {
  let total = 0;
  for (const n of arr) total += n;
  return total;
};

const main = () => {
  const N = 10000000;
  const big = Array.from({ length: N }, (_, i) => i);
  const t0 = performance.now();
  const total = sum(big);
  const t1 = performance.now();
  console.log("sum:", total);
  console.log(`O(n) demo with N=${N} took ${(t1 - t0).toFixed(3)} ms`);
};

main();
