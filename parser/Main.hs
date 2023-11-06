module Main where

import Data.List (isPrefixOf, sortBy)
import Main.Base
import Main.Input
import Main.Output
import System.Directory (listDirectory)
import System.FilePath ((</>))

-- Parse holiday data
-- Generate ics files
main :: IO ()
main = do
  -- read files
  filesInDir <- listDirectory "./data"
  let files = filter ("20" `isPrefixOf`) filesInDir
  contents <- mapM (readFile . ("./data" </>)) files

  -- parse data
  let dataByFile = zip files contents
  let dataByYear = map parseByFile dataByFile
  let dataMixed = join dataByYear

  -- debug log
  debug dataByYear

  -- write files
  writeFile "./docs/index.html" $ icsByType Both dataMixed

  writeFile "./docs/main.ics" $ icsByType Both dataMixed
  writeFile "./docs/rest.ics" $ icsByType Rest dataMixed
  writeFile "./docs/work.ics" $ icsByType Work dataMixed

-- Log holiday data ordered by each year
debug :: [(String, [Date], [Date])] -> IO ()
debug dataByYear = mapM_ showByYear $ sortByYear data'
  where
    data' = map (\(y, r, w) -> (y, sortByDate $ r ++ w)) dataByYear
    sortByYear = sortBy (\(y1, _) (y2, _) -> compare y1 y2)
    showByYear (year, dates) = do
      putStrLn $ "\nYear" <> year
      mapM_ print dates
