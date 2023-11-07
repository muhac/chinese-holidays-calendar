module Main.Input where

import Data.List (uncons)
import Data.List.Split (splitOn)
import Data.Maybe (mapMaybe)
import Data.Time (UTCTime, addUTCTime, defaultTimeLocale, formatTime, parseTimeOrError)
import Main.Base
import System.FilePath (takeBaseName)

-- Parse holiday data
-- Organized by year
parseFile :: FilePath -> String -> Yearly
parseFile file content = Yearly yearName daysRest daysWork
  where
    yearName = takeBaseName file
    daysRest = parse content Rest
    daysWork = parse content Work

-- Convert data to Date
parse :: String -> Status -> [Holiday]
parse content status = zip names dates >>= uncurry toHolidays
  where
    names = Group status . rawName <$> raw
    dates = parseDates . rawDate status <$> raw
    raw = parseRaw content

-- Parse raw data from file
-- Format: name;rest;work
parseRaw :: String -> [HolidayRaw]
parseRaw content = mapMaybe (toHolidayRaw . splitOn ";") eventsRaw
  where
    eventsRaw = mapMaybe (fmap fst . uncons . words) uncomment
    uncomment = mapMaybe (fmap fst . uncons . splitOn "//") $ lines content

-- Expand date ranges to UTCTime list
-- Support multiple date ranges separated by comma
parseDates :: String -> [Date]
parseDates ranges = zipWith3 Date [1 ..] (repeat $ length dates) dates
  where
    dates = splitOn "," ranges >>= parseDate . splitOn "-"

-- Parse date range
-- 1. like "2020.1.1"
-- 2. like "2020.1.1-2020.1.3"
parseDate :: [String] -> [UTCTime]
parseDate [""] = []
parseDate [single] = [parseTime single]
parseDate [start, end]
  | start == end = parseDate [end]
  | otherwise = first : parseDate [second, end]
  where
    first = parseTime start
    second = printTime $ addUTCTime day first
    day = 24 * 60 * 60
parseDate _ = []

-- Parse date in format "2020.1.1"
parseTime :: String -> UTCTime
parseTime = parseTimeOrError True defaultTimeLocale "%Y.%-m.%-d"

-- Format date in format "2020.1.1"
printTime :: UTCTime -> String
printTime = formatTime defaultTimeLocale "%Y.%-m.%-d"
