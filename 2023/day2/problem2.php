<?php

if ($argc === 2) {
    driver($argv[1]);
} else {
    driver('sample.txt');
}

function driver($inputFilePath)
{
    $inputData = file_get_contents($inputFilePath);
    $inputArray = explode("\n", $inputData);
    $sumOfPowers = 0;

    foreach ($inputArray as $row) {

        $countMap = [
            'red' => 0,
            'green' => 0,
            'blue' => 0,
        ];

        $power = 1;

        [$gameNumberPart, $gamePart] = explode(':', $row);
        $games = str_replace(';', ',', $gamePart);
        $ballCounts = explode(',', $games);

        foreach ($ballCounts as $ballCount) {
            [$count, $colour] = explode(' ', trim($ballCount));

            if ($countMap[$colour] < $count) {
                $countMap[$colour] = $count;
            }
        }

        foreach (['red', 'green', 'blue'] as $c) {
            $power = $power * $countMap[$c];
        }

        $sumOfPowers += $power;
    }

    echo ($sumOfPowers);
}
