#!/bin/bash



for u in "16" "32"
do
  for e in "B" "L"
  do
    filepath="test_utf${u}_${e}E"
    cp test_utf8_bom.csv "${filepath}.csv"
    cp test_utf8_bom.csv "${filepath}_BOM.csv"

    nkf --overwrite -w${u}${e}0 "${filepath}.csv"
    nkf --overwrite -w${u}${e} "${filepath}_BOM.csv"
  done
done
