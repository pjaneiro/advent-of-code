import SonarSweep.Challenges
import Dive.Challenges

import Test.Hspec

main :: IO ()
main = hspec $ do
  describe "Day 1" $ do
    let inputData = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]

    describe "Challenge1" $ do
      let result = 7
      it "returns the right value for the sample input" $ do
        SonarSweep.Challenges.challenge1 inputData `shouldBe` result

    describe "Challenge2" $ do
      let result = 5
      it "returns the right value for the sample input" $ do
        SonarSweep.Challenges.challenge2 inputData `shouldBe` result

  describe "Day 2" $ do
    let inputData = [Dive.Challenges.Command "forward" 5,Dive.Challenges.Command "down" 5,Dive.Challenges.Command "forward" 8,Dive.Challenges.Command "up" 3,Dive.Challenges.Command "down" 8,Dive.Challenges.Command "forward" 2]

    describe "Challenge 1" $ do
      let result = 150
      it "returns the right value for the sample input" $ do
        Dive.Challenges.challenge1 inputData `shouldBe` result

    describe "Challenge 2" $ do
      let result = 900
      it "returns the right value for the sample input" $ do
        Dive.Challenges.challenge2 inputData `shouldBe` result
