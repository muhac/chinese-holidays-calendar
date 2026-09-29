module Main.Output where

import Data.List.Split (splitOn)
import Data.Time (defaultTimeLocale, formatTime)
import Data.UUID (fromWords, toString)
import Main.Base
import Numeric (readHex)
import Text.Printf (printf)

-- Generate ics files
generate :: Config -> [Holiday] -> Status -> String
generate config dates status = unlines [icsHead config status, icsBody, icsTail]
  where
    icsBody = unlines $ map (icsEvent config) events
    events = sortByDate $ filterByStatus status dates

-- Standard ics format for the beginning
icsHead :: Config -> Status -> String
icsHead config status =
  unlines
    [ "BEGIN:VCALENDAR"
    , "VERSION:2.0"
    , "PRODID:" ++ productId config
    , "X-WR-CALNAME:" ++ titleStatus config status
    ]

-- Standard ics format for each event
-- Names may carry an English title: "中文名_English_Name"
icsEvent :: Config -> Holiday -> String
icsEvent config (Holiday (Group status name) (Date index total time)) =
  unlines
    [ "BEGIN:VEVENT"
    , "UID:" ++ uuid
    , "DTSTART;VALUE=DATE:" ++ formatTime defaultTimeLocale "%Y%m%d" time
    , "SUMMARY:" ++ summary
    , "DESCRIPTION:" ++ prefix ++ show status ++ printf "第%d天 / 共%d天" index total
    , "END:VEVENT"
    ]
  where
    (nameCn, nameEn) = case splitOn "_" name of
      cn : en -> (cn, unwords en)
      [] -> (name, "")
    summary = if null nameEn then nameCn ++ show status else nameEn
    prefix = if nameInDescription config then nameCn ++ " " else ""
    uuid = toString $ fromWords a b c d
    a = fst . head . readHex $ formatTime defaultTimeLocale "%Y%m%d" time
    b = fromIntegral $ shift index + total
    c = fromIntegral $ shift $ indexStatus status
    d = 0xa95511fe -- 955.WLB
    shift = (*) 0x10000

-- Standard ics format for the ending
icsTail :: String
icsTail = "END:VCALENDAR"
