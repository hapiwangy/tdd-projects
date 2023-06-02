const assert = require('assert');
// class Dollar {
//     constructor(amount){
//         this.amount = amount;
//     }
//     times(multiplier) {
//         return new Dollar(this.amount * multiplier);
//     }
// }

class Money {
    constructor(amount, currency){
        this.amount = amount;
        this.currency = currency;
    }
    times (multiplier){
        return new Money(this.amount * multiplier, this.currency);
    }
    divide (divisor){
        return new Money(this.amount / divisor, this.currency)
    }
}
let fiver = new Money(5, "USD");
let tenner = new Money(10, "USD");
// 不輸出任何東西是正常的，只有在出錯的時候會顯示
// assert.strictEqual(tenner.amount, 10);
// assert.strictEqual(tenner.currency, "USD");
assert.deepStrictEqual(fiver.times(2), tenner)

let tenEuros = new Money(10, "EUR");
let twentyEuros = new Money(20, "EUR");
// assert.strictEqual(twentyEuros.amount, 20);
// assert.strictEqual(twentyEuros.currency, "EUR");
assert.deepStrictEqual(tenEuros.times(2), twentyEuros)

let originalMoney = new Money(4002, "KRW");
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMOneyAfterDivision = new Money(1000.5, "KRW");
// 這個方法可以一次比較物件裡面的所有東西
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMOneyAfterDivision)
