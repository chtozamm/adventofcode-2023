// const input = `two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen`.split("\n");

const input = await Deno.readTextFile("input.txt").then((text) =>
  text.split("\r\n")
);

// console.log(input);

const calibrationValues: number[] = [];

function findDigits(s: string, digits: string[] = []) {
  if (!s || s.length === 0) return digits;
  if (!isNaN(+s[0])) {
    digits.push(s[0]);
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("one")) {
    digits.push("1");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("two")) {
    digits.push("2");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("three")) {
    digits.push("3");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("four")) {
    digits.push("4");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("five")) {
    digits.push("5");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("six")) {
    digits.push("6");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("seven")) {
    digits.push("7");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("eight")) {
    digits.push("8");
    return findDigits(s.slice(1), digits);
  } else if (s.startsWith("nine")) {
    digits.push("9");
    return findDigits(s.slice(1), digits);
  } else {
    return findDigits(s.slice(1), digits);
  }
}

for (let i = 0; i < input.length; i++) {
  const digits: string[] = findDigits(input[i]);

  const first = digits[0];
  const second = digits[digits.length - 1];

  const calibrationValue = first + second;
  calibrationValues.push(Number(calibrationValue));
}

const sum = calibrationValues.reduce((a, b) => a + b);
console.log(sum);
