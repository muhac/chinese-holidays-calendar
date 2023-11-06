module Main.Input where

import Data.Char (isSpace)
import Data.List.Split (splitOn)
import Data.Time (UTCTime, addUTCTime, defaultTimeLocale, formatTime, parseTimeOrError)
import Main.Base
import System.FilePath (takeBaseName)

-- Join data from each year and each type
join :: [(String, [Date], [Date])] -> [Date]
join = concatMap (\ (_, rest, work) -> rest ++ work)

-- Parse holiday data
-- Organized by year
parseByFile :: (FilePath, String) -> (String, [Date], [Date])
parseByFile (file, content) = (year, rest, work)
  where
    year = takeBaseName file
    rest = parse content Rest
    work = parse content Work

-- Convert data to Date
parse :: String -> DateType -> [Date]
parse content flag = concatMap constructor $ zip (map head raw) dates
  where
    constructor (name, dates) = constructDate name flag <$> dates
    dates = parseDate <$> map (!! indexDateType flag) raw
    raw = parseFile content

-- Parse data from file
-- Result: [[Name, RestDays, WorkDays]]
parseFile :: String -> [[String]]
parseFile content = filter ((== 3) . length) $ splitOn ";" <$> rawEvents
  where
    rawEvents = filter (not . null) . map (head . words) $ uncomment
    uncomment = filter (not . null) . map (head . splitOn "//") $ unindent
    unindent = dropWhile isSpace <$> lines content

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
parseDate' [single] = [parseTime single]
parseDate' [start, end]
  | start == end = parseDate' [end]
  | otherwise = first : parseDate' [second, end]
  where
    first = parseTime start
    second = printTime $ addUTCTime day first
    day = 24 * 60 * 60

-- Parse date in format "2020.1.1"
parseTime :: String -> UTCTime
parseTime = parseTimeOrError True defaultTimeLocale "%Y.%-m.%-d"

-- Format date in format "2020.1.1"
printTime :: UTCTime -> String
printTime = formatTime defaultTimeLocale "%Y.%-m.%-d"
