# Snowball Fight

## Findings

1. Based on the hint, you need to enter single player mode

## Solution

1. Enter private room
2. Get your `windows.location.href`

   ```js
   windows.location.href =
   	"https://hhc23-snowball.holidayhackchallenge.com/room/?username={username}&roomId={roomId}&roomType=private&gameType=co-op&id={id}&dna={dnaSeq}&singlePlayer=false";
   ```

3. Set the parameter `singlePlayer` to `True`
4. A mighty elf will join you to win the game !
