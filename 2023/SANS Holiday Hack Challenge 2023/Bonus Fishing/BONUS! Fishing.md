# BONUS! Fishing

1.  Automate fishing via browser console:

    ```js
    function startAutoFishing() {
    	// Click the castReelBtn to start the process
    	castReelBtn.click();

    	// Set up an interval to check for the 'gotone' class after a delay
    	var autoFishing = setInterval(function () {
    		if (reelItInBtn.classList.contains("gotone")) {
    			console.log("Reel In!");
    			reelItInBtn.click();

    			// Click the castReelBtn again after reeling in
    			setTimeout(function () {
    				castReelBtn.click();
    			}, 1000); // Adjust the delay as needed

    			// Clear the interval
    			// clearInterval(autoFishing);
    		}
    	}, 200); // Check every 200 milliseconds (adjust as needed)
    }
    ```

2.  Find the heatmap URL spot for fish in HTML commented code: `https://2023.holidayhackchallenge.com/sea/fishdensityref.html`
3.  [Find](./Bonus%20Fishing/fish_list.js) the remaining fish that hasn't been caught.
4.  Use heatmap to catch the remaining fish!
