const fishArray = [
	"Flutterfin Rainbow-Roll",
	"Gelatina Ringletfin",
	"Funfetti Flick-Flick",
	"Fluffle-Muffin Sparklefin",
	"Jovian Jamboree Jellydonut Jellyfish",
	"Gumball Glooperfish",
	"Jester Jellyfin",
	"The Chocolate Star Gingo Guppy",
	"Whirly Snuffleback Trout",
	"Jolly Jellyjam Fish",
	"Lounging Liquorice Crustacean-Nosed Berryfin",
	"Jinglefin Jellyfrizzle",
	"The Polka-Dot-Propeller Puffling Fish",
	"JubiliFLOPinear Snorkeldonut",
	"Glittering Gummy Guppy",
	"Whiskered Whizzler",
	"Flamango-Buzzling Sushi Swimmer",
	"Flutterfin Bubblegum Gumball",
	"ChocoSeahorsefly",
	"Speckled Toastfin Snorkelback",
	"Jamboree Jellywing",
	"Fantabulous Fry-Sherbert Aquapine",
	"Jamboree Jellofish",
	"Bubblegum Blowfish-Bee",
	"Frizzle Fringe Flutterfin",
	"Whiskered Jumblefish",
	"Bumblefin Toffee Torpedo",
	"Jolly Jellypeanut Fish",
	"Gumbubble Guppy",
	"Glittering Gummy Whipray",
	"Flippity Flan Flopper",
	"Biscuit Bugle-Tail Fish",
	"Strudel Scuttle Scalefish",
	"Bubblegum Bumblefin",
	"Flutterfin Falafeluncher",
	"The Splendiferous Spaghetti Starfin",
	"The Rainbow Jelibelly Floatfish",
	"Rhinoceros Beetle Bumble Tuna",
	"Hatwearing Hippofish",
	"Frizzle Fish",
	"Polka-Pop CandyFloss Fish",
	"Bumbleberry Floatfish",
	"Caramelotus Humming Float",
	"Dandy Candy Goby",
	"The Polka-Dot Pudding Puff",
	"Frosted Donut Jellyfluff Puffer",
	"Flamingo Flapjack Finaticus",
	"The Splendiferous Spaghetti Seahorsicle",
	"Flutterfin Scoopscale",
	"Frizzle Frazzle Fly-n-Fish",
	"Frizzle-Frizzled Jambalaya Jellyfish",
	"Sprinkfish",
	"Fantail Flutterfin",
	"JellyChip CuddleSwimmer",
	"Whiskered Lollipop Loonfish",
	"Jester Gumball Pufferfish",
	"The Hummingbrewster BumbleFlish",
	"Jangleroo Snackfin",
	"Blibbering Blubberwing",
	"The Bubblegum Confeetish",
	"Fantastical Fusilloni Flounderfish",
	"Pizzafin Flutterbub",
	"The Whiskered Watermelon Pufferfish",
	"The Bumblebee Doughnut Delphin",
	"Pistachio Pizzafin Puffinfly",
	"Aquatic JellyPuff Doughnut Shark",
	"Gumball Guppygator",
	"The Burgerwing Seahorse",
	"Bellychuckle Balloonfish",
	"FizzleWing PuffleGill",
	"Bumbleberry Rainbow Flicorn Fish",
	"Whistlefin Wafflegill",
	"Pizzadillo Glitter-Guppy",
	"Jamboree Jellydonut Jellyfish Trout",
	"The Bubblegum Bumblefin",
	"Gelatino Floatyfin",
	"The Frambuzzle Flickerfin",
	"The Speckled Pizzafin Fizzflyer",
	"Sparkling Pizzafin Pixie-fish",
	"Bumblecado Finstache Hybridsail",
	"Pizzamarine Popcorn Puffer",
	"Laughter Ligrolomia",
	"Frosted Jelly Doughnut Pegasus Finfish",
	"The Whirling Donut Jellygator",
	"Flutterfin Cupcake Goby",
	"The Gumball Guppy",
	"Bubblegum Blowfish Beetle Bug",
	"Sparkling Gumbubble Piscadot",
	"The Flamboyant Flutter-fish",
	"Twinkling Tortellini Trouterfly",
	"Beatleberry Fluff Guppy.",
	"Glaze Meringuelle",
	"The Whiskered Blubberberry Flapper",
	"Sherbet Swooshfin",
	"Marzipoisson Popsicala",
	"Bubblegum Ballistic Barracuda",
	"Puzzletail Splashcake",
	"Fantasia Fluffernutter Finfish",
	"Rainbow Jelly-Dough Fish",
	"Flutterfin Pizzapuffer",
	"BugBrella Aquacake",
	"Twirly Finny Cakeling",
	"Frizzleberry Flapjack Fish",
	"Whiskered Sprinkle Glider",
	"The Pristimaela Parfait Pengu-Angel",
	"Bubblerooni WhiskerWaffle",
	"The Speckled Whisker-Spoon Puffer",
	"BumbleSquid Donutella",
	"Sparkleberry Gobblefin",
	"Fizzgiggle Frizzlefin",
	"JibberJelly Sundae Swimmer",
	"The Flutterfin Pastry Puffer",
	"Rainbow Gummy Scalefish",
	"Jingle JellyFroth Fish",
	"The Spotted Flutterfin Pastrytetra",
	"Flutterfin Hotcheeto Penguinfish",
	"Piscis Cyberneticus Skodo",
	"Oreo OctoPufferRock",
	"Fluffernutter Pufferpine",
	"Whirlygig Polka-Dotted Jelly-Donut Pufferfish",
	"Bumbleberry Gilled Glider",
	"Polkadot Pancake Puffer",
	"Mermacorn Fish",
	"Sprinkle Starfish Sardine",
	"Choco-Bumblefin Parrot Trout",
	"The Fantabulous Gala Glazed-Guppy",
	"Pudding Puff ParrotMoth Fish",
	"Fantastical Flapjack Flipperfin",
	"TruffleBugle ZephyrFish",
	"Bumbleberry Glitterfin",
	"The Jester Jellycarafe",
	"The Flamingotuna McSprinklefin",
	"Whiskerfroth Flutterfin",
	"Spotted Sprinkledonut Puffer",
	"Stripe-tailed Pepperoni Puffer",
	"Jelly-Feather Macaroon Guppy",
	"Flutterfin Pancake Puffer.",
	"Whiskered Rainbow Glidleberry",
	"Chucklefin Clownfish",
	"Bumbleberry Snorkelsnout",
	"Jolly Jellydozer",
	"The Polka Dotted Jello-fish",
	"The Bumbleberry Guppiesaurus",
	"Flutterfin Pizzacrust Glimmertail",
	"Bumblebee, Pizza-fin Jamboree",
	"Whizzbizzle Poptuckle",
	"Candyfloss Clownphino",
	"Flutterglaze Bumblefin",
	"Bumbleberry Poptarticus",
	"Plaid Zephyr Cuddlefin",
	"Jolly Jambalaya Jubilee Fish",
	"Confetti Clownfrippery Fish",
	"Rainbow Jelly-Bumble Shark",
	"Marshmallow Pogo-Starfish",
	"The Spangled Jelly-Tortle Ripplefin",
	"Fantabulous Rainbow Polka Poptartfish",
	"The ChocoChandelier Goldnipper",
	"Gummybrella Anemofin",
	"Gummy Fizzler",
	"The Bumblebelly Polkadot Glaze-fish",
	"Fantaray Flakefin",
	"Splendiferous Ribbontail",
	"The Butterfleagleberry Seahorse",
	"Sushinano Sweetsquid",
	"The Whiskered Melonfin",
	"The Fantastical Fizzbopper",
	"Splashtastic Bagelback Rainbownose",
	"Pizzafly Rainbowgill",
	"Frizzling Bubblehopper",
	"Cuckoo Bubblegum Unicornfish",
	"The Lucid Lollyscale",
];

function checkFishName(fishName) {
	const index = fishArray.indexOf(fishName);
	if (index !== -1) {
		// Fish is in the array, remove it
		fishArray.splice(index, 1);
		console.log(`${fishName} is in the array. Removed from the list.`);
	}
}

function clickNextButton() {
	const nextButton = document.querySelector(".nextBtn"); // Replace with the actual selector
	if (nextButton) {
		nextButton.click();
		return true;
	} else {
		return false;
	}
}

async function automateFishChecking() {
	let hasNextFish = true;

	while (hasNextFish) {
		// Extract the fish name from the h3 element
		let fishName = document
			.querySelector(".cotd-info > h3:nth-child(1)")
			.textContent.trim(); // Adjust the selector based on your website structure

		// Check if the fish name is in the array
		if (fishName) {
			checkFishName(fishName);
		}

		// Click the button to display the next fish
		hasNextFish = clickNextButton();

		// Add a delay to ensure the page has updated before continuing
		await new Promise((resolve) => setTimeout(resolve, 1500)); // Adjust the delay as needed
	}
}

await automateFishChecking();

console.log(fishArray);

// [ "Frizzle Fringe Flutterfin", "Piscis Cyberneticus Skodo", "Bumbleberry Snorkelsnout" ]
