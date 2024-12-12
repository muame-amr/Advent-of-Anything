const file = Bun.file("test.in");
const text = await file.text();
let stones = text.split(" ").map(Number);

for (let i = 0; i < 25; ++i) {
	const newStones = [];
	stones.forEach((stone) => {
		if (stone == 0) {
			newStones.push(1);
		} else if (stone.toString().length % 2 === 0) {
			const str = stone.toString();
			const midpoint = str.length / 2;
			const left = parseInt(str.slice(0, midpoint), 10);
			const right = parseInt(str.slice(midpoint), 10);
			newStones.push(left, right);
		} else {
			newStones.push(stone * 2024);
		}
	});
	stones = newStones;
}

console.log(stones.length);
