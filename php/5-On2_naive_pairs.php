<?php

/* 
    On2_naive_pairs.php
    Big-O: O(n²) – quadratic
    Compare every pair of elements (e.g., brute-force duplicate search). 
*/

function hasDuplicate($arr)
{
    $n = count($arr);
    for ($i = 0; $i < $n; $i++) {
        for ($j = $i + 1; $j < $n; $j++) {
            if ($arr[$i] === $arr[$j]) return true;
        }
    }
    return false;
}

function main()
{
    $N = 30000; // 30k² ≈ 900M comparisons ⇒ slow
    $nums = range(0, $N - 1);
    $nums[] = $N - 1; // guarantee duplicate at the tail (worst-case)

    $t0 = microtime(true);
    $dup = hasDuplicate($nums);
    $t1 = microtime(true);

    echo "duplicate found: " . ($dup ? "true" : "false") . "\n";
    echo sprintf("O(n²) demo with N=%d took %.3f ms\n", $N, ($t1 - $t0) * 1000);
}

main();
