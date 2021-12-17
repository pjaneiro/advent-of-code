module GiantSquid.Challenges where
  import Text.Printf

  data House = House { number :: Int
                     , marked :: Bool
                     } deriving (Show, Read)

  readInput :: FilePath -> IO ([Int], [[[House]]])
  readInput path = do
    contents <- readFile path
    -- let input = lines contents
    return ([], [[[]]])

  challenge1 :: ([Int], [[[House]]]) -> Int
  challenge1 x = 0

  challenge2 :: ([Int], [[[House]]]) -> Int
  challenge2 x = 0

  run :: IO ()
  run = do
    printf "Day 4 - Giant Squid\n"
    let fileName = "GiantSquid/input.txt"
    inputData <- readInput fileName

    let challenge1Result = challenge1 inputData
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData
    printf "Challenge 2: %d\n" challenge2Result
