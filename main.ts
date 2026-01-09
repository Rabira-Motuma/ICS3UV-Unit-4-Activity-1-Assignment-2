/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-05
* @fileoverview This program prints a pattern of numbers in a right-angled triangle.
*/

// variables
let rowsAsString: string;
let rowsAsNumber:number;

rowsAsString = prompt("How many rows would you like?.") || "0";
rowsAsNumber = parseInt(rowsAsString);

// for statements
for (let i = 1; i <= rowsAsNumber; i++) {
  let row = ""
  for (let j = 1; j <= i; j++) {
    row = row + ` ${j}`;
  }
  console.log(row);
}

console.log("\nDone.")
