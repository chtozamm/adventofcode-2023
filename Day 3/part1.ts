// const input = `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..`.split("\n");

const input = await Deno.readTextFile("input.txt").then((text) =>
  text.split("\r\n")
);

const ROWS = input.length;
const LINE_LENGTH = input[0].length;

const partNumbers: number[] = [];

const isDigit = (char: string) => /[0-9]/.test(char);
const isSymbol = (char: string) => char !== "." && !isDigit(char);

for (let y = 0; y < ROWS; y++) {
  let value = "";
  let hasSymbolAttached = false;

  for (let x = 0; x < LINE_LENGTH; x++) {
    // if current character is *not* a digit
    if (value.length > 0 && !isDigit(input[y][x])) {
      if (hasSymbolAttached) {
        partNumbers.push(+value);
      }
      value = "";
      hasSymbolAttached = false;
      continue;
    }

    // if current character *is* a digit
    if (isDigit(input[y][x])) {
      value += input[y][x];

      if (y > 0) {
        if (x > 0) {
          // *--
          // -v-
          // ---
          if (isSymbol(input[y - 1][x - 1])) {
            hasSymbolAttached = true;
          }
        }
        // -*-
        // -v-
        // ---
        if (isSymbol(input[y - 1][x])) hasSymbolAttached = true;
        // --*
        // -v-
        // ---
        if (x < LINE_LENGTH - 1) {
          if (isSymbol(input[y - 1][x + 1])) {
            hasSymbolAttached = true;
          }
        }
      }
      // ---
      // *v-
      // ---
      if (x > 0) {
        if (isSymbol(input[y][x - 1])) {
          hasSymbolAttached = true;
        }
      }
      // ---
      // -v*
      // ---
      if (x < LINE_LENGTH - 1) {
        if (isSymbol(input[y][x + 1])) {
          hasSymbolAttached = true;
        }
      }
      // ---
      // -v-
      // *--
      if (y < ROWS - 1) {
        if (x > 0) {
          if (isSymbol(input[y + 1][x - 1])) {
            hasSymbolAttached = true;
          }
        }
        // ---
        // -v-
        // -*-
        if (isSymbol(input[y + 1][x])) {
          hasSymbolAttached = true;
        }
        // ---
        // -v-
        // --*
        if (x < LINE_LENGTH - 1) {
          if (isSymbol(input[y + 1][x + 1])) {
            hasSymbolAttached = true;
          }
        }
      }
    }

    // if current character *is* a digit and is the last char in line
    if (value.length > 0 && x === LINE_LENGTH - 1 && hasSymbolAttached) {
      partNumbers.push(+value);
    }
  }
}

// console.log(partNumbers);
// console.log(partNumbers.slice(700, 900));

const sum = partNumbers.reduce((a, b) => a + b);
console.log("The answer: " + sum);
