const file = Bun.file("input.in");
const text = await file.text();
let stones = text.split(" ").reduce((acc, val) => {
	acc[val] = 1;
	return acc;
}, {});

for (let i = 0; i < 75; ++i) {
	const newStones = {};
	for (const [stone, count] of Object.entries(stones)) {
		if (stone == 0) {
			newStones[1] = (newStones[1] || 0) + count;
		} else if (stone.length % 2 === 0) {
			const midpoint = stone.length / 2;
			const left = parseInt(stone.slice(0, midpoint), 10);
			const right = parseInt(stone.slice(midpoint), 10);
			newStones[left] = (newStones[left] || 0) + count;
			newStones[right] = (newStones[right] || 0) + count;
		} else {
			newStones[stone * 2024] = (newStones[stone * 2024] || 0) + count;
		}
	}
	stones = newStones;
}

console.log(Object.values(stones).reduce((acc, val) => acc + val, 0)); //259755538429618
