# Elf Connection

## Easy Solution
In the source code, you will find `wordSets` and `correctSets` variables that store
list of words and its group respectively. Map each elements in `wordSets` to `correctSets`
to find the answer.

```js
for (let i = 0; i < 4; i++) {
	for (let j = 0; j < 4; j++) {
		let set = correctSets[j];
		let wordSet = wordSets[i + 1];
		let correctWords = set.map((index) => wordSet[index]);
		console.log(correctWords.join(", "));
	}
}
```

## Hard Solution
You will also find `score` variable in the source code that store your score. To beat the high score,
you just need to assign `score` to a value that higher than current high score (50000).

```js
// Enter this in the browser's console
score = 100000;
```