# Installation

You must first install [Node](https://nodejs.org/en/download/). Then run the command `npm install` in the command line. This will install the Node packages necessary to run the program. Then run `node index.js` on the command line. This will open up a chrome instance that will take input from our AI and perform the moves on the board. The way this works is the AI will make a move and write it to `move.txt`. Then the Node program will see that the file was changed and will parse the move. It then executes the move on the page.

You can also manually save a move in `move.txt` and the Node script will try to do the move on the page. For example, start with the initial board state and save `move.txt` after typing in the file "1 7 1 6" and saving, this will move our side's pawn up one tile.

Note: still has issues with edge cases like taking pieces and when pieces move backwards.
