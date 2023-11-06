module Main.Output where

import Data.Time (defaultTimeLocale, formatTime, nominalDiffTimeToSeconds)
import Data.Time.Clock.POSIX (utcTimeToPOSIXSeconds)
import Data.UUID (fromWords, toString)
import Main.Base
import Text.Printf (printf)

-- Generate ics files
icsByType :: DateType -> [Date] -> String
icsByType flag dates = unlines [icsHead flag, icsBody, icsTail]
  where
    icsBody = unlines $ map icsEvent $ sortByDate $ filterByType flag dates

-- Standard ics format for the beginning
icsHead :: DateType -> String
icsHead flag =
  unlines
    [ "BEGIN:VCALENDAR",
      "VERSION:2.0",
      "PRODID:-//Rank Technology//Chinese Holidays//EN",
      "X-WR-CALNAME:" <> titleDateType flag
      -- "X-WR-TIMEZONE:Asia/Shanghai",
    ]

-- Standard ics format for each event
icsEvent :: Date -> String
icsEvent (Date name time flag index total) =
  unlines
    [ "BEGIN:VEVENT",
      "UID:" <> uuid,
      "DTSTART;VALUE=DATE:" <> formatTime defaultTimeLocale "%Y%m%d" time,
      "SUMMARY:" <> name <> show flag,
      "DESCRIPTION:" <> show flag <> printf "第%d天 / 共%d天" index total,
      "END:VEVENT"
    ]
  where
    uuid = toString $ fromWords a b c d
    a = floor . nominalDiffTimeToSeconds . utcTimeToPOSIXSeconds $ time
    b = fromIntegral $ indexDateType flag
    c = fromIntegral total
    d = fromIntegral index

-- Standard ics format for the ending
icsTail :: String
icsTail = "END:VCALENDAR"
