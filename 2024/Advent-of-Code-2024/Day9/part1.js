const path = "test.in";
const file = Bun.file(path);

const text = await file.text();
const data = text.split("").map(Number);

const disk = [];

for (let i = 0; i < data.length; i++) {
	for (let j = data[i]; j > 0; j--) {
		if (i % 2 == 0) {
			disk.push(i / 2);
		} else {
			disk.push(".");
		}
	}
}

console.log(disk);

disk.forEach((block, idx) => {
	if (block === ".") {
		while (true) {
			const temp = disk.pop();
			if (temp === ".") {
				continue;
			} else {
				disk[idx] = temp;
				break;
			}
		}
	}
});

console.log(disk);

let checksum = 0;
disk.forEach((block, id) => {
	checksum += block * id;
});

console.log(checksum);
