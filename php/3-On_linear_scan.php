<?php

/* 
    On_linear_scan.php
    Big-O: O(n) – linear
    Iterates once over the entire collection (sum). 
*/

function sum($arr)
{
    $total = 0;
    foreach ($arr as $n) {
        $total += $n;
    }
    return $total;
}

function main()
{
    $N = 1000000;
    $big = range(0, $N - 1);

    $t0 = microtime(true);
    $total = sum($big);
    $t1 = microtime(true);

    echo "sum: $total\n";
    echo sprintf("O(n) demo with N=%d took %.3f ms\n", $N, ($t1 - $t0) * 1000);
}

main();
