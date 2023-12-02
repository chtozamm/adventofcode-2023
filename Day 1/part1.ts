// const input = `1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`.split("\n");

const input = await Deno.readTextFile("input.txt").then((text) =>
  text.split("\r\n")
);

console.log(input);

const calibrationValues: number[] = [];

for (let i = 0; i < input.length; i++) {
  const digits: string[] = [];

  for (let j = 0; j < input[i].length; j++) {
    const char = input[i][j];
    if (!isNaN(+char)) {
      digits.push(char);
    }
  }

  const first = digits[0];
  const second = digits[digits.length - 1];

  const calibrationValue = first + second;
  calibrationValues.push(Number(calibrationValue));
}

console.log(calibrationValues);

const sum = calibrationValues.reduce((a, b) => a + b);
console.log(sum);
