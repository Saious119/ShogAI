const puppeteer = require('puppeteer');
const fs = require('fs');
const Jimp = require('jimp');
const chokidar = require('chokidar');

/* Global variables */
var isReady = false
var board = []
var page
var globalMoveArr = []
const tileWidth = 38
const numTilesX = 9
const startX = 142
const tileHeight = 49
const numTilesY = 9
const startY = 8

/* Makes a move when our AI sends next move to file */
const watcher = chokidar.watch('./move.txt', {usePolling: true, interval: 100})
watcher.on('change', makeMove)

/* Web scraping and initialization for interacting with the web shogi board */
;(async () => {
  // Makes the browser to be launched in a headful way
  const browser = await puppeteer.launch({headless: false, devtools: true});

  // Creates a new page on the default browser context and sets it as global variable
  page = await browser.newPage();

  // Instructs the blank page to navigate a URL
  await page.goto('https://japanesechess.org/shogi2014/');
  await page.bringToFront()

  // Waits until the canvas is rendered
  await page.waitForSelector('canvas');
  await page.waitFor(1000)

  // Get image data in URL format from the shogi board on the web page
  const dataURL = await getBoardImg()

  // Gets initial board state from the canvas image data
  const boardState = await getBoardState(dataURL)
  console.log(`INITIAL BOARD STATE:`)
  console.log(`--------------------------------------\n${getBoardStr(boardState)}\n--------------------------------------`)
  
  // Initialization over, get ready to start playing game
  isReady = true
})();

/* Uses image data from the board and returns the board state as a string */
async function getBoardState(dataURL) {
  const base64String = dataURL.substr(dataURL.indexOf(',') + 1)
  const imgBuffer = Buffer.from(base64String, 'base64')
  fs.writeFileSync('image.png', imgBuffer)

  // Read in image data to get pixel data
  const img = await Jimp.read('./image.png')
  .then(image => {
    return image
  }).catch((err) => {
    console.log(err)
  })

  // Process image and get board state
  board = []
  for (let j = startY; j < numTilesY*tileHeight+startY; j+=tileHeight) { // Go through each row
    const boardRow = []
    for (let i = startX; i < numTilesX*tileWidth+startX; i+=tileWidth) { // Check each tile in the row
      boardRow.push(getPiece(img, i, j))
    }
    board.push(boardRow)
  }

  // Return board state as a string
  return board
}

/* Determines what piece occupies the tile being checked and returns it as a string */
function getPiece(img, x, y) {
  const p1 = Jimp.intToRGBA(img.getPixelColor(x+16,y+10)) // Top left pixel
  // console.log(p1)
  // console.log(`Coords: (${x},${y})`)

  const p2 = Jimp.intToRGBA(img.getPixelColor(x+16,y+tileHeight+2-10)) // Bottom left pixel
  // console.log(p2)
  // console.log(`Coords: (${x},${y+tileHeight})`)

  // Check for pixels from top left corner
  if (p1.r === 87 && p1.g === 78 && p1.b === 0) { // The opponent lancelot is #574e00 or rgb(87,78,0)
    return "L1"
  } else if (p1.r === 37 && p1.g === 33 && p1.b === 0) { // The opponent silver general is #252100 or rgb(37,33,0)
    return "S1"
  } else if (p1.r === 12 && p1.g === 11 && p1.b === 0) { // The opponent gold general is #0c0b00 or rgb(12,11,0)
    return "G1"
  } else if (p1.r === 62 && p1.g === 56 && p1.b === 0) { // The opponent king is #3e3800 or rgb(62,56,0)
    return "K1"
  } else if (p1.r === 197 && p1.g === 177 && p1.b === 0) { // The opponent pawn is #c5b100 or rgb(197,177,0)
    return "P1"
  } else if (p1.r === 0 && p1.g === 0 && p1.b === 0) { // The opponent piece has a black pixel first
    const newP1 = Jimp.intToRGBA(img.getPixelColor(x+16+1,y+10))
    if (newP1.r === 12 && newP1.g === 11 && newP1.b === 0) { // The opponent knight is #0c0b00 or rgb(12,11,0)
      return "N1"
    } else if (newP1.r === 98 && newP1.g === 88 && newP1.b === 0) { // The opponent rook is #625800 or rgb(98,88,0)
      return "R1"
    } else if (newP1.r === 0 && newP1.g === 0 && newP1.b === 0) { // The opponent bishop has a second black pixel
      return "B1"
    }
  }

  // Check pixels from bottom left corner
  if (p2.r === 97 && p2.g === 86 && p2.b === 2) { // Our lancelot is #615602 or rgb(97,86,2)
    return "L2"
  } else if (p2.r === 182 && p2.g === 162 && p2.b === 4) { // Our silver general is #b6a204 or rgb(182,162,4)
    return "S2"
  } else if (p2.r === 85 && p2.g === 76 && p2.b === 2) { // Our gold general is #554c02 or rgb(85,76,2)
    return "G2"
  } else if (p2.r === 0 && p2.g === 0 && p2.b === 0) { // Our piece has a black pixel first
    const newP2 = Jimp.intToRGBA(img.getPixelColor(x+16+1,y+tileHeight+2-10))
    if (newP2.r === 96 && newP2.g === 86 && newP2.b === 0) { // Our bishop is #605600 or rgb(96,86,0)
      return "B2"
    } else if (newP2.r === 0 && newP2.g === 0 && newP2.b === 0) { // Our piece has a black pixel second
      const newestP2 = Jimp.intToRGBA(img.getPixelColor(x+16+1,y+tileHeight+2-9))
      if (newestP2.r === 84 && newestP2.g === 76 && newestP2.b === 0) { // Our king is #544c00 or rgb(84,76,0)
        return "K2"
      } else if (newestP2.r === 132 && newestP2.g === 119 && newestP2.b === 0) { // Our rook is #847700 or rgb(132,119,0)
        return "R2"
      }
    } else if (newP2.r === 48 && newP2.g === 43 && newP2.b === 0) { // Our piece has a second pixel with color #302b00 or rgb(48,43,0)
      const newestP2 = Jimp.intToRGBA(img.getPixelColor(x+16+1,y+tileHeight+2-9))
      if (newestP2.r === 120 && newestP2.g === 108 && newestP2.b === 0) { // Our knight is #786c00 or rgb(120,108,0)
        return "N2"
      } else if (newestP2.r === 156 && newestP2.g === 140 && newestP2.b === 0) { // Our pawn is #9c8c00 or rgb(156,140,0)
        return "P2"
      }
    }
  }

  // A blank tile is the only option left
  return "O"
}

/*
Initial board state:

L1 N1 S1 G1 K1 G1 S1 N1 L1
 O R1 O O O O O B1 O
 P1 P1 P1 P1 P1 P1 P1 P1 P1
 O O O O O O O O
 O O O O O O O O
 O O O O O O O O
 P2 P2 P2 P2 P2 P2 P2 P2 P2
 O B2 O O O O O R2 O
 L2 N2 S2 G2 K2 G2 S2 N2 L2
 */
function getBoardStr(board) {
  let boardState = ""
  for (let j = 0; j < numTilesY; j++) {
    const boardRow = board[j].join(" ")
    if (j != numTilesY-1) {
      boardState += boardRow + "\n"
    } else {
      boardState += boardRow
    }
  }
  return boardState
}

/* Performs two mouse clicks on the page in order to move a piece on the board */
function performMove(moveArr) {
  // Headless chrome instance always ensures (relative to page) that the canvas has a top value of 100px and a left value of 13px
  const top = 100
  const left = 13
  const clickCenterBuffer = 20
  globalMoveArr = [[moveArr[0],moveArr[1]], [moveArr[2],moveArr[3]]]
  const initPosX = clickCenterBuffer + left + startX + (moveArr[0]-1)*tileWidth
  const initPosY = clickCenterBuffer + top + startY + (moveArr[1]-1)*tileHeight
  const nextPosX = clickCenterBuffer + left + startX + (moveArr[2]-1)*tileWidth
  const nextPosY = clickCenterBuffer + top + startY + (moveArr[3]-1)*tileHeight
  // console.log(`Mouse x: ${initPosX} Mouse y: ${initPosY}`)
  // console.log(`Mouse2 x: ${nextPosX} Mouse2 y: ${nextPosY}`)
  page.mouse.click(initPosX, initPosY)
  page.mouse.click(nextPosX, nextPosY)
}

/* Pulls canvas data from page and returns the image data */
async function getBoardImg() {
  const dataURL = await page.evaluate(() => {
    const canvas = document.querySelector('canvas')
    return dataURL = canvas.toDataURL()
  })
  return dataURL
}

/* Listens for when our AI makes a move, then sends the result of the opponent AI to a file and waits again */
async function makeMove() {
  if (isReady) {
    // Parse move from file
    const data = fs.readFileSync('./move.txt', 'utf8')
    const moveData = data.split(/\s+/)

    // Perform move on the board
    performMove(moveData)

    // Get board state from canvas before opponent moves
    const prevBoardImg = await getBoardImg()
    const prevBoardState = await getBoardState(prevBoardImg)
    // console.log(`\n${getBoardStr(prevBoardState)}\n`)

    // Wait for opponent to make move and canvas to update
    await page.waitFor(1000)
    
    // Get board state from canvas after opponent moves
    const currBoardImg = await getBoardImg()
    const currBoardState = await getBoardState(currBoardImg)
    // console.log(`\n${getBoardStr(currBoardState)}\n`)

    // Write opponent move to file
    const enemyMove = await computeEnemyMove(prevBoardState, currBoardState, currBoardImg)
    fs.writeFileSync('./board.txt', new Uint8Array(Buffer.from(enemyMove)))
  }
}

/* Gets the enemy move and returns properly formatted string for our AI to parse */
async function computeEnemyMove(prev, curr, imgURL) {
  // Get Jimp processed image
  const base64String = imgURL.substr(imgURL.indexOf(',') + 1)
  const imgBuffer = Buffer.from(base64String, 'base64')
  fs.writeFileSync('image.png', imgBuffer)

  // Read in image data to get pixel data
  const image = await Jimp.read('./image.png')
  .then(img => {
    return img
  }).catch((err) => {
    console.log(err)
  })
  let coords = []
  for (let i = 0; i < numTilesX; i++) { // Go through each row
    for (let j = 0; j < numTilesY; j++) { // Go through each column of the row
      if (prev[i][j] !== curr[i][j]) {
        coords.push([j+1,i+1]) // Push the position of the current tile (because it changed from prev board state)
      }
    }
  }

  console.log("Coordinates that are different after the AI moved:")
  console.log(coords)
  
  // Gets the piece that is on the first pair of coordinates (enemy piece moving foward means this function returns "O" cuz its blank)
  const piece = getPiece(image, ...coords[0])
  // console.log(`Piece from the coords "${coords[0].toString()}": ${piece}`)
  if (piece === "O") {
    // console.log("PIECE WAS AN O")
    return getMoveStr(coords)
  } else {
    // console.log("PIECE WAS NOT O")
    return getMoveStr(coords.reverse())
  }
}

/* Takes two pairs of coords and compiles a string in the format for our AI to parse */
function getMoveStr(coords) {
  console.log(`Enemy move: (${coords[0][0]},${coords[0][1]}) -> (${coords[1][0]},${coords[1][1]})`)
  let moveArr = []
  coords.forEach(pair => {
    moveArr.push(...pair)
  })
  return moveArr.join(" ")
}