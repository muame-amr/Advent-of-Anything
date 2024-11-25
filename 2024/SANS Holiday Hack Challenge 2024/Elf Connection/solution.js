const wordSets = {
	1: [
		"Tinsel",
		"Sleigh",
		"Belafonte",
		"Bag",
		"Comet",
		"Garland",
		"Jingle Bells",
		"Mittens",
		"Vixen",
		"Gifts",
		"Star",
		"Crosby",
		"White Christmas",
		"Prancer",
		"Lights",
		"Blitzen",
	],
	2: [
		"Nmap",
		"burp",
		"Frida",
		"OWASP Zap",
		"Metasploit",
		"netcat",
		"Cycript",
		"Nikto",
		"Cobalt Strike",
		"wfuzz",
		"Wireshark",
		"AppMon",
		"apktool",
		"HAVOC",
		"Nessus",
		"Empire",
	],
	3: [
		"AES",
		"WEP",
		"Symmetric",
		"WPA2",
		"Caesar",
		"RSA",
		"Asymmetric",
		"TKIP",
		"One-time Pad",
		"LEAP",
		"Blowfish",
		"hash",
		"hybrid",
		"Ottendorf",
		"3DES",
		"Scytale",
	],
	4: [
		"IGMP",
		"TLS",
		"Ethernet",
		"SSL",
		"HTTP",
		"IPX",
		"PPP",
		"IPSec",
		"FTP",
		"SSH",
		"IP",
		"IEEE 802.11",
		"ARP",
		"SMTP",
		"ICMP",
		"DNS",
	],
};

let correctSets = [
	[0, 5, 10, 14], // Set 1
	[1, 3, 7, 9], // Set 2
	[2, 6, 11, 12], // Set 3
	[4, 8, 13, 15], // Set 4
];

// Easy: run this to get the answers
for (let i = 0; i < 4; i++) {
	for (let j = 0; j < 4; j++) {
		let set = correctSets[j];
		let wordSet = wordSets[i + 1];
		let correctWords = set.map((index) => wordSet[index]);
		console.log(correctWords.join(", "));
	}
}

// Hard: run this to get the answers
score = 100000; // any value bigger than 500000