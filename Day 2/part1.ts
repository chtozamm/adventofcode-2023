// const input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
//   .split("\n")
//   .map((line) => line.slice(line.indexOf(":") + 2));

const input = await Deno.readTextFile("input.txt").then((text) =>
  text.split("\r\n").map((line) => line.slice(line.indexOf(":") + 2))
);

// console.log(input);

const cubes = {
  red: 12,
  green: 13,
  blue: 14,
};

const possibleGames: number[] = [];

function checkPossibility(rounds: string[]) {
  for (let i = 0; i < rounds.length; i++) {
    const roundCubesAmount = {
      red: 0,
      green: 0,
      blue: 0,
    };

    const roundCubes = rounds[i].split(",").map((s) => s.trim());
    // console.log(roundCubes);

    for (let j = 0; j < roundCubes.length; j++) {
      if (roundCubes[j].includes("red")) {
        roundCubesAmount.red = Number(roundCubes[j].split(" ")[0]);
      }
      if (roundCubes[j].includes("green")) {
        roundCubesAmount.green = Number(roundCubes[j].split(" ")[0]);
      }
      if (roundCubes[j].includes("blue")) {
        roundCubesAmount.blue = Number(roundCubes[j].split(" ")[0]);
      }
    }

    if (
      roundCubesAmount.red > cubes.red ||
      roundCubesAmount.green > cubes.green ||
      roundCubesAmount.blue > cubes.blue
    )
      return false;
  }
  return true;
}

for (let i = 0; i < input.length; i++) {
  const rounds = input[i].split(";").map((s) => s.trim());

  const possible = checkPossibility(rounds);
  if (!possible) {
    continue;
  } else {
    possibleGames.push(i + 1);
  }
}

const sum = possibleGames.reduce((a, b) => a + b);
console.log("The answer: " + sum);
