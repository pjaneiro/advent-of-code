module SonarSweep.Main where
  import Text.Printf

  readInput :: FilePath -> IO [Int]
  readInput path = do
    contents <- readFile path
    let input = map read  . lines $ contents
    return input

  challenge1 :: [Int] -> Int -> Int -> Int
  challenge1 inputData 0 count = count
  challenge1 inputData index count = if inputData!!index > inputData!!(index - 1) then challenge1 inputData (index - 1) (count + 1)
    else challenge1 inputData (index - 1) (count)

  challenge2 :: [Int] -> Int -> Int -> Int
  challenge2 _ _ _ = 0

  run :: IO ()
  run = do
    printf "Day 1 - Sonar Sweep\n"
    let fileName = "SonarSweep/input_sonarsweep.txt"
    inputData <- readInput fileName
    let inputLength = length inputData

    let challenge1Result = challenge1 inputData (inputLength - 1) 0
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData (inputLength - 1) 0
    printf "Challenge 2: %d\n" challenge2Result
