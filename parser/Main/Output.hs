module Main.Output where

import Data.Time (defaultTimeLocale, formatTime)
import Data.UUID (fromWords, toString)
import Main.Base
import Numeric (readHex)
import Text.Printf (printf)

-- Generate ics files
generate :: [Holiday] -> Status -> String
generate dates status = unlines [icsHead status, icsBody, icsTail]
  where
    icsBody = unlines $ map icsEvent events
    events = sortByDate $ filterByStatus status dates

-- Standard ics format for the beginning
icsHead :: Status -> String
icsHead status =
  unlines
    [ "BEGIN:VCALENDAR"
    , "VERSION:2.0"
    , "PRODID:-//Rank Technology//Chinese Holidays//EN"
    , "X-WR-CALNAME:" ++ titleStatus status
    ]

-- Standard ics format for each event
icsEvent :: Holiday -> String
icsEvent (Holiday (Group status name) (Date index total time)) =
  unlines
    [ "BEGIN:VEVENT"
    , "UID:" ++ uuid
    , "DTSTART;VALUE=DATE:" ++ formatTime defaultTimeLocale "%Y%m%d" time
    , "SUMMARY:" ++ name ++ show status
    , "DESCRIPTION:" ++ show status ++ printf "第%d天 / 共%d天" index total
    , "END:VEVENT"
    ]
  where
    uuid = toString $ fromWords a b c d
    a = fst . head . readHex $ formatTime defaultTimeLocale "%Y%m%d" time
    b = fromIntegral $ shift index + total
    c = fromIntegral $ shift $ indexStatus status
    d = 0xa95511fe -- 955.WLB
    shift = (*) 0x10000

-- Standard ics format for the ending
icsTail :: String
icsTail = "END:VCALENDAR"
