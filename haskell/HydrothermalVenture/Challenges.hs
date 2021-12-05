module HydrothermalVenture.Challenges where
  import Text.Printf

  data Point = Point { x :: Int
                     , y :: Int
                     } deriving (Show, Read)

  data Vector = Vector { from :: Point
                       , to :: Point
                       } deriving (Show, Read)

  readInput :: FilePath -> IO [Vector]
  readInput path = do
    contents <- readFile path
    -- let input = lines contents
    return []

  challenge1 :: [Vector] -> Int
  challenge1 x = 0

  challenge2 :: [Vector] -> Int
  challenge2 x = 0

  run :: IO ()
  run = do
    printf "Day 5 - Hydrothermal Venture\n"
    let fileName = "HydrothermalVenture/input.txt"
    inputData <- readInput fileName

    let challenge1Result = challenge1 inputData
    printf "Challenge 1: %d\n" challenge1Result

    let challenge2Result = challenge2 inputData
    printf "Challenge 2: %d\n" challenge2Result
