module Main.Base where

import Control.Monad (mfilter)
import Data.Function (on)
import Data.List (sortBy)
import Data.Maybe (fromMaybe)
import Data.Time (UTCTime, defaultTimeLocale, formatTime)

-- Calendar settings, overridable through environment variables
data Config = Config
  { calendarName :: String
  , productId :: String
  , writeIndex :: Bool
  , nameInDescription :: Bool
  }
  deriving (Eq, Show)

defaultConfig :: Config
defaultConfig =
  Config
    { calendarName = "中国节假日安排"
    , productId = "-//Rank Technology//Chinese Holidays//EN"
    , writeIndex = True
    , nameInDescription = False
    }

configFromEnv :: [(String, String)] -> Config
configFromEnv env =
  Config
    { calendarName = setting "CALENDAR_NAME" calendarName
    , productId = setting "CALENDAR_PRODID" productId
    , writeIndex = flag "CALENDAR_WRITE_INDEX" writeIndex
    , nameInDescription = flag "CALENDAR_NAME_IN_DESCRIPTION" nameInDescription
    }
  where
    value key = mfilter (not . null) $ lookup key env
    setting key field = fromMaybe (field defaultConfig) $ value key
    flag key field = maybe (field defaultConfig) (`elem` ["1", "true"]) $ value key

data Status = Both | Rest | Work deriving (Enum)

instance Show Status where
  show Both = ""
  show Rest = "假期"
  show Work = "补班"

-- Title of output ics file
titleStatus :: Config -> Status -> String
titleStatus config Both = calendarName config
titleStatus config kind = calendarName config ++ "（" ++ show kind ++ "）"

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
