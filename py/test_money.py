import unittest
class Dollar:
    def __init__(self, amount) -> None:
        self.amount = amount
    def times(self, multiplier):
        return Dollar(self.amount * multiplier)
# 測試類別要是unittest.TestCase的子類別
class TestMoney(unittest.TestCase):
    # 方法要以test作為開頭(不知道是規定或是偏好)
    def testMultiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)
        # assertEqual將實際值和期望值做比較
        self.assertEqual(10, tenner.amount)

if __name__ == '__main__':
    unittest.main()