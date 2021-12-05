module Dive.Challenges where
  import Text.Printf

  data Command = Command { cmd :: String
                         , val :: Int
                         } deriving (Show, Read)

  readCommand :: String -> Command
  readCommand line =
    let [cmd, val] = words line
    in Command {cmd = cmd, val = read val}

  readInput :: FilePath -> IO [Command]
  readInput path = do
    contents <- readFile path
    let input = map readCommand . lines $ contents
    return $ input

  challenge1ApplyCommand :: (Int, Int, Command) -> (Int, Int)
  challenge1ApplyCommand (h, d, Command "forward" x) = (h + x, d)
  challenge1ApplyCommand (h, d, Command "down" x) = (h, d + x)
  challenge1ApplyCommand (h, d, Command "up" x) = (h, d - x)

  challenge1loop :: ([Command], Int, Int, Int) -> Int
  challenge1loop (inputData, index, h, d) = do
    if index >= length inputData
      then h * d
    else do
        let (newH, newD) = challenge1ApplyCommand (h, d, inputData!!index)
        challenge1loop (inputData, index + 1, newH, newD)

  challenge1 :: [Command] -> Int
  challenge1 inputData = challenge1loop (inputData, 0, 0, 0)

  challenge2ApplyCommand :: (Int, Int, Int, Command) -> (Int, Int, Int)
  challenge2ApplyCommand (h, d, a, Command "forward" x) = (h + x, d + a * x, a)
  challenge2ApplyCommand (h, d, a, Command "down" x) = (h, d, a + x)
  challenge2ApplyCommand (h, d, a, Command "up" x) = (h, d, a - x)

  challenge2loop :: ([Command], Int, Int, Int, Int) -> Int
  challenge2loop (inputData, index, h, d, a) = do
    if index >= length inputData
      then h * d
    else do
        let (newH, newD, newA) = challenge2ApplyCommand (h, d, a, inputData!!index)
        challenge2loop (inputData, index + 1, newH, newD, newA)

  challenge2 :: [Command] -> Int
  challenge2 inputData = challenge2loop (inputData, 0, 0, 0, 0)

  run :: IO ()
  run = do
    printf "Day 2 - Dive!\n"
    let fileName = "Dive/input.txt"
    inputData <- readInput fileName

    let challenge1Result = challenge1 inputData
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData
    printf "Challenge 2: %d\n" challenge2Result
