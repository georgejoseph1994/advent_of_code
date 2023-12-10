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

    $availableBallMap = [
        'red' => 12,
        'green' => 13,
        'blue' => 14,
    ];

    $sumOfPossibleGameIds = 0;

    foreach ($inputArray as $row) {
        $isInvalidGame = false;
        [$gameNumberPart, $gamePart] = explode(':', $row);
        $gameNumber = explode(' ', $gameNumberPart)[1];

        $games = explode(';', $gamePart);
        foreach ($games as $game) {
            $rounds = explode(',', trim($game));
            foreach ($rounds as $round) {
                [$count, $colour] = explode(' ', trim($round));
                if ($count > $availableBallMap[$colour]) {
                    $isInvalidGame = true;
                    break;
                }
            }
        }
        if (!$isInvalidGame) {
            $sumOfPossibleGameIds += $gameNumber;
            // print($gameNumber . "\n");
        }
    }
    print($sumOfPossibleGameIds);
}
