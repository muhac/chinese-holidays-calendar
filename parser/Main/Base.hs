module Main.Base where

import Data.Function (on)
import Data.List (sortBy)
import Data.Time (UTCTime, defaultTimeLocale, formatTime)

data Status = Both | Rest | Work deriving (Enum)

instance Show Status where
  show Both = ""
  show Rest = "假期"
  show Work = "补班"

-- Title of output ics file
titleStatus :: Status -> String
titleStatus Both = "中国节假日安排"
titleStatus kind = "中国节假日安排（" ++ show kind ++ "）"

-- Index of input txt file
indexStatus :: Status -> Int
indexStatus Both = 0
indexStatus Rest = 1
indexStatus Work = 2

instance Eq Status where
  (==) = (==) `on` indexStatus

data Yearly = Yearly
  { year :: String
  , rest :: [Holiday]
  , work :: [Holiday]
  }

join :: Yearly -> [Holiday]
join y = rest y ++ work y

data HolidayRaw = HolidayRaw
  { rawName :: String
  , rawRest :: String
  , rawWork :: String
  }

rawDate :: Status -> HolidayRaw -> String
rawDate Rest = rawRest
rawDate Work = rawWork
rawDate _ = return ""

toHolidayRaw :: [String] -> Maybe HolidayRaw
toHolidayRaw [n, r, w] = Just $ HolidayRaw n r w
toHolidayRaw [n, r] = toHolidayRaw [n, r, ""]
toHolidayRaw _ = Nothing

data Holiday = Holiday
  { holidayGroup :: Group
  , holidayDate :: Date
  }

instance Show Holiday where
  show (Holiday group date) = unwords [show date, show group]

toHolidays :: Group -> [Date] -> [Holiday]
toHolidays group dates = Holiday group <$> dates

data Group = Group
  { holidayStatus :: Status
  , holidayName :: String
  }

instance Show Group where
  show (Group status name) = unwords [name, show status]

data Date = Date
  { holidayIndex :: Int
  , holidayTotal :: Int
  , holidayTime :: UTCTime
  }

instance Show Date where
  show (Date index total time) =
    unwords
      [ formatTime defaultTimeLocale "%Y-%m-%d" time
      , show index ++ "/" ++ show total
      ]

sortByDate :: [Holiday] -> [Holiday]
sortByDate = sortBy (compare `on` holidayTime . holidayDate)

filterByStatus :: Status -> [Holiday] -> [Holiday]
filterByStatus Both = id
filterByStatus kind = filter ((== kind) . holidayStatus . holidayGroup)
