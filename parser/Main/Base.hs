module Main.Base where

import Data.List (intercalate, sortBy)
import Data.Time (UTCTime, defaultTimeLocale, formatTime)

data DateDataType = Both | Rest | Work deriving (Enum)

instance Show DateDataType where
  show Both = ""
  show Rest = "假期"
  show Work = "补班"

-- Title of output ics file
titleDateDataType :: DateDataType -> String
titleDateDataType Both = "节假日"
titleDateDataType flag = "节假日（" ++ show flag ++ "）"

-- Index of input txt file
indexDateDataType :: DateDataType -> Int
indexDateDataType Both = 0
indexDateDataType Rest = 1
indexDateDataType Work = 2

data Date = Date
  { name :: String,
    time :: UTCTime,
    flag :: DateDataType,
    index :: Int,
    total :: Int
  }

instance Show Date where
  show (Date name time flag index total) =
    intercalate " " $
      [ formatTime defaultTimeLocale "%Y-%m-%d" time,
        name,
        show flag,
        show index ++ "/" ++ show total
      ]

constructDate :: String -> DateDataType -> (Int, Int, UTCTime) -> Date
constructDate name flag (index, total, time) = Date name time flag index total

sortByDate :: [Date] -> [Date]
sortByDate = sortBy (\(Date _ t1 _ _ _) (Date _ t2 _ _ _) -> compare t1 t2)

filterByType :: DateDataType -> [Date] -> [Date]
filterByType Both = id
filterByType flag = filter (\(Date _ _ f _ _) -> indexDateDataType f == indexDateDataType flag)
