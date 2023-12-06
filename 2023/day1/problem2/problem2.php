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

    $numberMap = [
        '1' => 1,
        '2' => 2,
        '3' => 3,
        '4' => 4,
        '5' => 5,
        '6' => 6,
        '7' => 7,
        '8' => 8,
        '9' => 9,
        'one' => 1,
        'two' => 2,
        'three' => 3,
        'four' => 4,
        'five' => 5,
        'six' => 6,
        'seven' => 7,
        'eight' => 8,
        'nine' => 9,
    ];

    foreach ($inputArray as $row) {
        $firstNumber = null;
        $lastNumber = null;

        // find first number
        for ($i = 0; $i <= strlen($row); $i++) {
            $subString = substr($row, 0, $i);
            foreach (array_keys($numberMap) as $number) {
                if (str_contains($subString, $number)) {
                    $firstNumber = $number;
                    break;
                }
            }
            if ($firstNumber !== null) {
                break;
            }
        }

        // find last number
        for ($i = 1; $i <= strlen($row); $i++) {
            $subString = substr($row, (-1 * $i));

            foreach (array_keys($numberMap) as $number) {
                if (str_contains($subString, $number)) {
                    $lastNumber = $number;
                    break;
                }
            }
            if ($lastNumber !== null) {
                break;
            }
        }

        $doubleDigitNumber += $numberMap[$firstNumber] * 10 + $numberMap[$lastNumber];
    }
    print_r($doubleDigitNumber);
}
