const file = Bun.file("test.in");
const input = await file.text();

const machines = input
	.split("\r\n\r\n")
	.map((m) => m.match(/\d+/g).map(Number));

function tokens([aX, aY, bX, bY, pX, pY]) {
    // Cramer's Rule 2 * 2 Matrix
	// aX * a + bX * b = pX
	// aY * a + bY * b = pY
	const a = (pX * bY - pY * bX) / (aX * bY - aY * bX);
	const b = (aX * pY - aY * pX) / (aX * bY - aY * bX);
	return Number.isInteger(a) && Number.isInteger(b) ? a * 3 + b : 0;
}

const total = machines.reduce((acc, machine) => acc + tokens(machine), 0);
console.log(total);
