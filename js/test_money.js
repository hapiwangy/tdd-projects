const assert = require('assert');
class Dollar {
    constructor(amount){
        this.amount = amount;
    }
    times(multiplier) {
        return new Dollar(this.amount * multiplier);
    }
}
let fiver = new Dollar(5);
let tenner = fiver.times(2);
// 不輸出任何東西是正常的，只有在出錯的時候會顯示
assert.strictEqual(tenner.amount, 10);