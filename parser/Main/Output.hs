module Main.Output where

import Data.Time (defaultTimeLocale, formatTime, nominalDiffTimeToSeconds)
import Data.Time.Clock.POSIX (utcTimeToPOSIXSeconds)
import Data.UUID (fromWords64, toString)
import Main.Base
import Text.Printf (printf)

-- Generate ics files
icsByType :: DateDataType -> [Date] -> String
icsByType flag dates = unlines [icsHead flag, icsBody, icsTail]
  where
    icsBody = unlines $ map icsEvent $ sortByDate $ filterByType flag dates

-- Standard ics format for the beginning
icsHead :: DateDataType -> String
icsHead flag =
  unlines
    [ "BEGIN:VCALENDAR",
      "VERSION:2.0",
      "PRODID:-//Rank Technology//Chinese Holidays//EN",
      "X-WR-CALNAME:" ++ titleDateDataType flag
      -- "X-WR-TIMEZONE:Asia/Shanghai",
    ]

-- Standard ics format for each event
icsEvent :: Date -> String
icsEvent (Date name time flag index total) =
  unlines
    [ "BEGIN:VEVENT",
      "UID:" ++ uuid,
      "DTSTART;VALUE=DATE:" ++ formatTime defaultTimeLocale "%Y%m%d" time,
      "SUMMARY:" ++ name ++ show flag,
      "DESCRIPTION:" ++ printf "%s 第%d天/共%d天" (show flag) index total,
      "END:VEVENT"
    ]
  where
    uuid = toString $ fromWords64 1 t
    t = floor $ nominalDiffTimeToSeconds $ utcTimeToPOSIXSeconds time

-- Standard ics format for the ending
icsTail :: String
icsTail = "END:VCALENDAR"
