module SonarSweep.Challenges where
  import Text.Printf

  readInput :: FilePath -> IO [Int]
  readInput path = do
    contents <- readFile path
    let input = map read  . lines $ contents
    return input

  challenge1Loop :: [Int] -> Int -> Int -> Int
  challenge1Loop inputData 0 count = count
  challenge1Loop inputData index count = if inputData!!index > inputData!!(index - 1) then challenge1Loop inputData (index - 1) (count + 1)
    else challenge1Loop inputData (index - 1) (count)

  challenge1 :: [Int] -> Int
  challenge1 inputData = challenge1Loop inputData (length inputData - 1) 0

  challenge2Loop :: [Int] -> Int -> Int -> Int
  challenge2Loop inputData 0 count = count
  challenge2Loop inputData index count = if (inputData!!index + inputData!!(index + 1) + inputData!!(index + 2)) > (inputData!!(index - 1) + inputData!!index + inputData!!(index + 1)) then challenge2Loop inputData (index - 1) (count + 1)
    else challenge2Loop inputData (index - 1) (count)

  challenge2 :: [Int] -> Int
  challenge2 inputData = challenge2Loop inputData (length inputData - 3) 0

  run :: IO ()
  run = do
    printf "Day 1 - Sonar Sweep\n"
    let fileName = "SonarSweep/input.txt"
    inputData <- readInput fileName

    let challenge1Result = challenge1 inputData
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData
    printf "Challenge 2: %d\n" challenge2Result
