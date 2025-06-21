<?php

/* 
    OnLogN_merge_sort.php
    Big-O: O(n log n)
    Classic merge-sort: log n levels of splitting × n work each level. 
*/

function mergeSort($arr)
{
    if (count($arr) <= 1) return $arr;

    $mid = intval(count($arr) / 2);
    $left = mergeSort(array_slice($arr, 0, $mid));
    $right = mergeSort(array_slice($arr, $mid));

    return merge($left, $right);
}

function merge($a, $b)
{
    $out = [];

    while (count($a) > 0 && count($b) > 0) {
        if ($a[0] <= $b[0]) {
            $out[] = array_shift($a);
        } else {
            $out[] = array_shift($b);
        }
    }

    return array_merge($out, $a, $b);
}

function main()
{
    $N = 100000; // use smaller N, merge-sort is heavy
    $nums = [];
    for ($i = 0; $i < $N; $i++) {
        $nums[] = mt_rand() / mt_getrandmax();
    }

    $t0 = microtime(true);
    mergeSort($nums);
    $t1 = microtime(true);

    echo sprintf("O(n log n) demo with N=%d took %.3f ms\n", $N, ($t1 - $t0) * 1000);
}

main();
