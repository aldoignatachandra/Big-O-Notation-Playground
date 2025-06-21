<?php

/* 
    OlogN_binary_search.php
    Big-O: O(log n) – logarithmic
    Binary search halves the remaining search space each step. 
*/

function binarySearch($arr, $target)
{
    $low = 0;
    $high = count($arr) - 1;

    while ($low <= $high) {
        $mid = intval(($low + $high) / 2);  // equivalent to >> 1 in JS
        if ($arr[$mid] === $target) return $mid;
        $arr[$mid] < $target ? $low = $mid + 1 : $high = $mid - 1;
    }

    return -1;
}

function main()
{
    $N = 1000000;
    $bigSorted = range(0, $N - 1);

    $t0 = microtime(true);
    $idx = binarySearch($bigSorted, $N - 1); // worst-case (near end)
    $t1 = microtime(true);

    echo "found at index: $idx\n";
    echo sprintf("O(log n) demo with N=%d took %.3f ms\n", $N, ($t1 - $t0) * 1000);
}

main();
