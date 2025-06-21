<?php

/* 
    O1_constant.php
    Big-O: O(1) – constant time
    We do the same amount of work regardless of N. 
*/

function getMiddleElement($arr)
{
    // Single arithmetic operation + array access ⇒ constant
    return $arr[floor(count($arr) / 2)];
}

function main()
{
    $N = 1000000;
    $big = range(0, $N - 1);

    $t0 = microtime(true);
    $mid = getMiddleElement($big);
    $t1 = microtime(true);

    echo "middle element: $mid\n";
    echo sprintf("O(1) demo took %.3f ms\n", ($t1 - $t0) * 1000);
}

main();
