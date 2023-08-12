function greedyCoinChange(coins, amount) {
  coins.sort((a, b) => b - a); // Sort the coins in descending order

  let coinCount = 0;
  let remainingAmount = amount;
  let list = [];

  for (let coin of coins) {
    while (remainingAmount >= coin) {
      remainingAmount -= coin;
      coinCount++;
      list.push(coin);
    }

    if (remainingAmount === 0) {
      break;
    }
  }

  if (remainingAmount > 0) {
    return "Cannot make exact change.";
  }

  //   return [coinCount, list];
  return { first: coinCount, second: list };
}

const coinDenominations = [25, 10, 5, 1];
const targetAmount = 321321;

const { first, second } = greedyCoinChange(coinDenominations, targetAmount);
console.log(`Minimum number of coins needed: ${first} with content ${""}`);
