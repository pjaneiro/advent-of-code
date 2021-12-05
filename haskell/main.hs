import BinaryDiagnostic.Challenges
import Dive.Challenges
import GiantSquid.Challenges
import HydrothermalVenture.Challenges
import SonarSweep.Challenges
import Text.Printf

main = do
  SonarSweep.Challenges.run
  printf "\n"
  Dive.Challenges.run
  printf "\n"
  BinaryDiagnostic.Challenges.run
  printf "\n"
  GiantSquid.Challenges.run
  printf "\n"
  HydrothermalVenture.Challenges.run
  printf "\n"
