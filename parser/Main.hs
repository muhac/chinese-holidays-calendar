module Main where

import Control.Monad (when)
import Data.Function (on)
import Data.List (isPrefixOf, sortBy)
import Main.Base
import Main.Input
import Main.Output
import System.Directory (listDirectory)
import System.Environment (getEnvironment)
import System.FilePath ((</>))

-- Parse holiday data
-- Generate ics files
main :: IO ()
main = do
  config <- configFromEnv <$> getEnvironment

  -- read files
  filesInDir <- listDirectory "./data"
  let files = filter ("20" `isPrefixOf`) filesInDir
  contents <- mapM (readFile . ("./data" </>)) files

  -- parse data
  let calendarYearly = zipWith parseFile files contents
  debug calendarYearly

  -- write files
  let calendar = calendarYearly >>= join
  when (writeIndex config) $
    writeFile "./docs/index.html" $ generate config calendar Both

  writeFile "./docs/main.ics" $ generate config calendar Both
  writeFile "./docs/rest.ics" $ generate config calendar Rest
  writeFile "./docs/work.ics" $ generate config calendar Work

-- Log holiday data ordered by each year
debug :: [Yearly] -> IO ()
debug yearly =
  let orderByYear = sortBy (compare `on` year) yearly
      printByDate holiday = do
        putStrLn $ "\nYear " ++ year holiday
        mapM_ print $ sortByDate $ join holiday
   in mapM_ printByDate orderByYear
