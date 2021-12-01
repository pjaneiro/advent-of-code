import SonarSweep.Challenges

import Test.Hspec

main :: IO ()
main = hspec $ do
  let inputData = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]

  describe "Challenge1" $ do
    let result = 7
    it "returns the right value for the sample input" $ do
      SonarSweep.Challenges.challenge1 inputData `shouldBe` result

  describe "Challenge2" $ do
    let result = 5
    it "returns the right value for the sample input" $ do
      SonarSweep.Challenges.challenge2 inputData `shouldBe` result
