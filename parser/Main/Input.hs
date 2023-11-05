module Main.Input where

import Data.List.Split (splitOn)
import Data.Time (UTCTime, addUTCTime, defaultTimeLocale, formatTime, parseTimeOrError)
import Main.Base
import System.FilePath (takeBaseName)

-- Join data from each year and each type
join :: [(String, [Date], [Date])] -> [Date]
join dataByFile = concatMap (\(_, rest, work) -> rest ++ work) dataByFile

-- Parse holiday data
-- Organized by year
parseByFile :: (FilePath, String) -> (String, [Date], [Date])
parseByFile (file, content) = (year, rest, work)
  where
    year = takeBaseName file
    rest = parse year content Rest
    work = parse year content Work

-- Convert data to Date
parse :: String -> String -> DateDataType -> [Date]
parse year content flag = concatMap constructor $ zip (map head raw) dates
  where
    constructor (name, dates) = map (constructDate name flag) dates
    dates = map parseDate $ map (!! indexDateDataType flag) raw
    raw = parseFile content

-- Parse data from file
-- Result: [[Name, RestDays, WorkDays]]
parseFile :: String -> [[String]]
parseFile content = holidays
  where
    holidays = filter ((== 3) . length) . map (splitOn ";") $ eachData
    eachData = filter (not . null) . map (head . words) $ eachLine
    eachLine = filter (not . null) . map (head . splitOn "//") $ lines content

-- Expand date ranges to UTCTime list
-- Support multiple date ranges separated by comma
parseDate :: String -> [(Int, Int, UTCTime)]
parseDate "" = []
parseDate range = zip3 [1 ..] (repeat $ length dates) dates
  where
    dates = concatMap (parseDate' . splitOn "-") $ splitOn "," range

-- Parse date range
-- 1. like "2020.1.1"
-- 2. like "2020.1.1-2020.1.3"
parseDate' :: [String] -> [UTCTime]
parseDate' [date] = [parseTime date]
parseDate' [start, end]
  | start == end = parseDate' [end]
  | otherwise = first : parseDate' [second, end]
  where
    first = parseTime start
    second = printTime $ addUTCTime 86400 first

-- Parse date in format "2020.1.1"
parseTime :: String -> UTCTime
parseTime date = parseTimeOrError True defaultTimeLocale "%Y.%-m.%-d" date :: UTCTime

-- Format date in format "2020.1.1"
printTime :: UTCTime -> String
printTime = formatTime defaultTimeLocale "%Y.%-m.%-d"
