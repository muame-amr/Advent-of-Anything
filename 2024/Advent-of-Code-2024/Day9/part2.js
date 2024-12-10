const path = "input.in";
const file = Bun.file(path);

const text = await file.text();
const data = text.split("").map(Number);

const disk = [];

for (let i = 0; i < data.length; i++) {
	const file = [];
	for (let j = data[i]; j > 0; j--) {
		if (i % 2 == 0) {
			file.push(i / 2);
		} else {
			file.push(".");
		}
	}
	if (file.length) disk.push(file);
}

console.log(disk);

const movedIds = [];
outer: for (let i = disk.length - 1; i >= 0; i--) {
	if (disk[i][0] !== "." && !movedIds.includes(disk[i][0])) {
		for (let j = 0; j <= i; j++) {
			if (disk[j][0] === "." && disk[j].length >= disk[i].length) {
				movedIds.push(disk[i][0]);
				if (disk[j].length === disk[i].length) {
					let temp = [...disk[j]];
					disk[j] = disk[i];
					disk[i] = temp;
					continue outer;
				} else {
					const temp = [...disk[i]];
					disk[i].fill(".");
					disk.splice(j, 1, temp, disk[j].slice(disk[i].length));
					i++;
				}
				continue outer;
			}
		}
	}
}

const defrag = disk.flat();

console.log(defrag);

let checksum = 0;
defrag.forEach((block, id) => {
	if (block !== ".") checksum += block * id;
});

console.log(checksum); // 6398096697992
