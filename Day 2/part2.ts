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

const powersOfGames: number[] = [];

function countRequiredCubes(rounds: string[]) {
  const requiredCubes = {
    reds: 0,
    greens: 0,
    blues: 0,
  };

  const reds: number[] = [0];
  const greens: number[] = [0];
  const blues: number[] = [0];

  for (let i = 0; i < rounds.length; i++) {
    const roundCubes = rounds[i].split(",").map((s) => s.trim());
    // console.log(roundCubes);

    for (let j = 0; j < roundCubes.length; j++) {
      if (roundCubes[j].includes("red")) {
        reds.push(Number(roundCubes[j].split(" ")[0]));
      }
      if (roundCubes[j].includes("green")) {
        greens.push(Number(roundCubes[j].split(" ")[0]));
      }
      if (roundCubes[j].includes("blue")) {
        blues.push(Number(roundCubes[j].split(" ")[0]));
      }
    }

    requiredCubes.reds = Math.max(...reds);
    requiredCubes.greens = Math.max(...greens);
    requiredCubes.blues = Math.max(...blues);
  }
  // console.log(requiredCubes);

  return requiredCubes;
}

for (let i = 0; i < input.length; i++) {
  const rounds = input[i].split(";").map((s) => s.trim());

  const requiredCubes = countRequiredCubes(rounds);
  const powerOfGame =
    requiredCubes.reds * requiredCubes.greens * requiredCubes.blues;
  powersOfGames.push(powerOfGame);
}

const sum = powersOfGames.reduce((a, b) => a + b);
console.log("The answer: " + sum);
