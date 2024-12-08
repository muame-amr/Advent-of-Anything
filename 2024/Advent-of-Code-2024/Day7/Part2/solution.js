const path = "test.in";
const file = Bun.file(path);

const text = await file.text();
const ops = ["+", "*", "|"];

function checkEquation(numbers, target) {
	let combinations = ops.map((op) => [op]);
	for (let i = 1; i < numbers.length - 1; i++) {
		combinations = combinations.flatMap((combination) =>
			ops.map((op) => [...combination, op])
		);
	}

	for (let combination of combinations) {
		let result = numbers[0];
		for (let i = 0; i < combination.length; i++) {
			if (combination[i] === "+") {
				result += numbers[i + 1];
			} else if (combination[i] === "*") {
				result *= numbers[i + 1];
			} else if (combination[i] === "|") {
				result = Number(result + "" + numbers[i + 1]);
			}
		}
		if (result === target) {
			return true;
		}
	}
	return false;
}

const list = text
	.split("\n")
	.map((row) => row.replace(":", ""))
	.map((row) => row.split(" ").map(Number));

let result = 0;
list.forEach((row) => {
	const target = row[0];
	const numbers = row.slice(1);
	if (checkEquation(numbers, target)) {
		result += target;
	}
});
console.log(result); // 348360680516005
