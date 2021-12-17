module BinaryDiagnostic.Challenges where
  import Text.Printf

  readInput :: FilePath -> IO [String]
  readInput path = do
    contents <- readFile path
    let input = lines contents
    return input

  challenge1 :: [String] -> Int
  challenge1 x = 0

  challenge2 :: [String] -> Int
  challenge2 x = 0

  run :: IO ()
  run = do
    printf "Day 3 - Binary Diagnostic\n"
    let fileName = "BinaryDiagnostic/input.txt"
    inputData <- readInput fileName

    let challenge1Result = challenge1 inputData
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData
    printf "Challenge 2: %d\n" challenge2Result
