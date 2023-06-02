import unittest

# class Dollar:
#     def __init__(self, amount) -> None:
#         self.amount = amount
#     def times(self, multiplier):
#         return Dollar(self.amount * multiplier)
class Money:
    def __init__(self,amount, currency) -> None:
        self.amount = amount
        self.currency = currency
    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)
    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)
    # 複寫__eq__方法，之後比較簡單很多(建議的做法)
    def __eq__(self, other):
        return self.amount == other.amount and self.currency == other.currency
# 測試類別要是unittest.TestCase的子類別
class TestMoney(unittest.TestCase):
    # 方法要以test作為開頭(不知道是規定或是偏好)
    def testMultiplicationInMoney(self):
        fiver = Money(5, "USD")
        tenner = Money(10, "USD")
        # assertEqual將實際值和期望值做比較
        self.assertEqual(tenner, fiver.times(2))

    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = Money(20, "EUR")
        # self.assertEqual(20, twentyEuros.amount)
        # self.assertEqual("EUR", twentyEuros.currency)
        self.assertEqual(tenEuros.times(2), twentyEuros)

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        actualMoneyAfterDivision = originalMoney.divide(4)
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(expectedMoneyAfterDivision, actualMoneyAfterDivision)
        # self.assertEqual(expectedMoneyAfterDivision.amount, 
        #                  actualMoneyAfterDivision.amount)
        # self.assertEqual(expectedMoneyAfterDivision.currency,
        #                  actualMoneyAfterDivision.currency)


if __name__ == '__main__':
    unittest.main()