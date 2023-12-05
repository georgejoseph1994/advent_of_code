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
    $doubleDigitNumber = 0;

    foreach ($inputArray as $row) {
        $firstNumber = null;
        $lastNumber = null;

        foreach (str_split($row) as $character) {
            if (is_numeric($character) && $firstNumber === null) {
                $firstNumber = (int) $character;
            }

            if (is_numeric($character)) {
                $lastNumber = (int) $character;
            }
        }
        $doubleDigitNumber += $firstNumber * 10 + $lastNumber;
    }

    echo $doubleDigitNumber;
}
