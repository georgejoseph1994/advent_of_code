<?php

if ($argc === 2) {
    driver($argv[1]);
} else {
    driver('sample.txt');
}

function driver($inputFilePath)
{
    $inputData = file_get_contents($inputFilePath);

    // create a 2d matrix of the input
    $matrix = explode("\n", $inputData);

    $total = 0;

    // iterate over each row
    foreach ($matrix as $rowNumber => $row) {
        $startPointer = $endPointer = null;
        $isNumberInInput = false;

        // iterate over each column
        foreach (str_split($row) as $columnNumber => $character) {
            if (is_numeric($character)) {
                $isNumberInInput = true;

                if ($startPointer === null) {
                    $startPointer = $columnNumber;
                }
                $endPointer = $columnNumber;

                if ($columnNumber !== sizeof(str_split($row)) - 1) {
                    continue;
                }
            }

            if ($isNumberInInput) {
                if (hasSymbolInSurroundingCells($startPointer, $endPointer, $rowNumber,  $matrix)) {
                    $total += (int) getSubString($startPointer, $endPointer, $row);
                }
            }

            $startPointer = $columnNumber + 1;
            $endPointer = $columnNumber;
            $isNumberInInput = false;
        }
    }

    print($total);
}


function isSymbol($input): bool
{
    if ($input === '.' || is_numeric($input)) {
        return false;
    }

    return true;
}

function stringHasSymbol(string $input): bool
{
    foreach (str_split($input) as $char) {
        if (isSymbol($char)) {
            return true;
        }
    }
    return false;
}

function getSubString($startPointer, $endPointer, $inputString): string
{
    return substr($inputString, $startPointer, $endPointer - $startPointer + 1);
}

function hasSymbolInSurroundingCells($startPointer, $endPointer, $rowNumber,  $matrix)
{
    $startPos = $startPointer > 0 ? $startPointer - 1 : $startPointer;
    $endPos = $endPointer < (sizeof($matrix) - 1) ? $endPointer + 1 : $endPointer;

    // check row above
    if ($rowNumber > 0) {
        if (stringHasSymbol(getSubString($startPos, $endPos,  $matrix[$rowNumber - 1]))) {
            return true;
        }
    }

    // check row below
    if ($rowNumber !== sizeof($matrix) - 1) {

        if (stringHasSymbol(getSubString($startPos, $endPos,  $matrix[$rowNumber + 1]))) {
            return true;
        }
    }

    // check cell to the left
    if ($startPointer != 0 && isSymbol($matrix[$rowNumber][$startPointer - 1])) {
        return true;
    }

    // check cell to the right
    if ($endPointer != sizeof(str_split($matrix[0])) - 1 && isSymbol($matrix[$rowNumber][$endPointer + 1])) {
        return true;
    }

    return false;
}
