module Main.Base where

import Data.Function (on)
import Data.List (sortBy)
import Data.Time (UTCTime, defaultTimeLocale, formatTime)

data DateType = Both | Rest | Work deriving (Enum)

instance Show DateType where
  show Both = ""
  show Rest = "假期"
  show Work = "补班"

-- Title of output ics file
titleDateType :: DateType -> String
titleDateType Both = "中国节假日安排"
titleDateType flag = "中国节假日安排（" <> show flag <> "）"

-- Index of input txt file
indexDateType :: DateType -> Int
indexDateType Both = 0
indexDateType Rest = 1
indexDateType Work = 2

data Date = Date
  { name :: String,
    time :: UTCTime,
    flag :: DateType,
    index :: Int,
    total :: Int
  }

instance Show Date where
  show (Date name time flag index total) =
    unwords
      [ formatTime defaultTimeLocale "%Y-%m-%d" time,
        name,
        show flag,
        show index <> "/" <> show total
      ]

constructDate :: String -> DateType -> (Int, Int, UTCTime) -> Date
constructDate name flag (index, total, time) = Date name time flag index total

sortByDate :: [Date] -> [Date]
sortByDate = sortBy (compare `on` time)

filterByType :: DateType -> [Date] -> [Date]
filterByType Both = id
filterByType flag = filter (\(Date _ _ f _ _) -> show f == show flag)
