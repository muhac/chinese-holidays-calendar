module Main where

import Data.Function (on)
import Main.Base
import Main.Input
import System.Exit (exitFailure, exitSuccess)
import Test.HUnit

instance Eq HolidayRaw where
  (==) = (==) `on` raw

raw a = (rawName a, rawRest a, rawWork a)

instance Show HolidayRaw where
  show raw = unwords [rawName raw, rawRest raw, rawWork raw]

testA1 =
  TestCase
    ( assertEqual
        "A1 - data - basic"
        [HolidayRaw "a" "b" "c"]
        $ parseRaw "a;b;c"
    )

testA2 =
  TestCase
    ( assertEqual
        "A2 - data - robust"
        [HolidayRaw "a12" "b345" "c678"]
        $ parseRaw "   a12;b345;c678    "
    )

testA3 =
  TestCase
    ( assertEqual
        "A3 - comments - basic"
        []
        $ parseRaw "//comment"
    )

testA4 =
  TestCase
    ( assertEqual
        "A4 - comments - robust"
        []
        $ parseRaw "   // a;b;c  123  "
    )

testA5 =
  TestCase
    ( assertEqual
        "A5 - hybrid - basic"
        [HolidayRaw "a" "b" "c"]
        $ parseRaw "a;b;c  // d;e;f  456  "
    )

testA6 =
  TestCase
    ( assertEqual
        "A6 - hybrid - robust"
        [HolidayRaw "a" "b" "c"]
        $ parseRaw "   a;b;c 123 // d;e;f  456  "
    )

testA7 =
  TestCase
    ( assertEqual
        "A7 - empty - basic"
        []
        $ parseRaw ""
    )

testA8 =
  TestCase
    ( assertEqual
        "A8 - empty - robust"
        []
        $ parseRaw "  //   "
    )

testA9 =
  TestCase
    ( assertEqual
        "A9 - multi-line"
        [HolidayRaw "a" "b" "c", HolidayRaw "d" "e" "f"]
        $ parseRaw . unlines
        $ [ "// INTRO"
          , ""
          , "a;b;c // dp1"
          , "// comment"
          , "d;e;f // dp2"
          , "END  // TEST"
          ]
    )

testB1 =
  TestCase
    ( assertEqual
        "B1 - case 1. like 2020.1.1 - basic"
        ["2020.1.1"]
        $ map printTime
        $ parseDate ["2020.1.1"]
    )

testB2 =
  TestCase
    ( assertEqual
        "B2 - case 1. like 2020.1.1 - robust"
        ["2021.11.2"]
        (map printTime $ parseDate ["2021.11.02"])
    )

testB3 =
  TestCase
    ( assertEqual
        "B3 - case 2. like 2020.1.1-2020.1.3 - basic"
        ["2022.1.1", "2022.1.2", "2022.1.3"]
        (map printTime $ parseDate ["2022.1.1", "2022.1.3"])
    )

testB4 =
  TestCase
    ( assertEqual
        "B4 - case 2. like 2020.1.1-2020.1.3 - cross month"
        ["2023.2.27", "2023.2.28", "2023.3.1", "2023.3.2"]
        (map printTime $ parseDate ["2023.2.27", "2023.3.2"])
    )

testB5 =
  TestCase
    ( assertEqual
        "B5 - case 2. like 2020.1.1-2020.1.3 - cross month 2"
        ["2024.2.28", "2024.2.29", "2024.3.1"]
        (map printTime $ parseDate ["2024.2.28", "2024.3.1"])
    )

testB6 =
  TestCase
    ( assertEqual
        "B6 - case 2. like 2020.1.1-2020.1.3 - cross year"
        ["2025.12.30", "2025.12.31", "2026.1.1", "2026.1.2"]
        (map printTime $ parseDate ["2025.12.30", "2026.1.2"])
    )

testC1 =
  TestCase
    ( assertEqual
        "C1 - case 1. like 2020.1.1 - a1 2"
        [(1, 2, "2020.1.1"), (2, 2, "2020.1.3")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2020.1.1,2020.1.3"
    )

testC2 =
  TestCase
    ( assertEqual
        "C2 - case 1. like 2020.1.1 - a1 3"
        [(1, 3, "2020.1.1"), (2, 3, "2020.1.3"), (3, 3, "2020.1.5")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2020.1.1,2020.1.3,2020.1.5"
    )

testC3 =
  TestCase
    ( assertEqual
        "C3 - case 2. like 2020.1.1-2020.1.3 - a3 2"
        [(1, 4, "2020.1.1"), (2, 4, "2020.1.2"), (3, 4, "2020.1.6"), (4, 4, "2020.1.7")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2020.1.1-2020.1.2,2020.1.6-2020.1.7"
    )

testC4 =
  TestCase
    ( assertEqual
        "C4 - case 2. like 2020.1.1-2020.1.3 - a1 a3"
        [(1, 4, "2020.1.1"), (2, 4, "2020.1.2"), (3, 4, "2020.1.3"), (4, 4, "2020.1.6")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2020.1.1-2020.1.3,2020.1.6"
    )

testC5 =
  TestCase
    ( assertEqual
        "C5 - case 2. like 2020.1.1-2020.1.3 - a3 a1"
        [(1, 5, "2019.12.6"), (2, 5, "2020.1.1"), (3, 5, "2020.1.2"), (4, 5, "2020.1.3"), (5, 5, "2021.11.11")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2019.12.6,2020.1.1-2020.1.3,2021.11.11"
    )

testC6 =
  TestCase
    ( assertEqual
        "C6 - case 2. like 2020.1.1-2020.1.3 - a3 a1"
        [(1, 3, "2019.12.6"), (2, 3, "2020.1.1"), (3, 3, "2021.11.11")]
        $ map (\(Date a b c) -> (a, b, printTime c))
        $ parseDates "2019.12.6,2020.1.1-2020.1.1,,,2021.11.11"
    )

tests :: Test
tests =
  TestList
    [ -- A. parseRaw
      TestLabel "Test parseRaw 1" testA1
    , TestLabel "Test parseRaw 2" testA2
    , TestLabel "Test parseRaw 3" testA3
    , TestLabel "Test parseRaw 4" testA4
    , TestLabel "Test parseRaw 5" testA5
    , TestLabel "Test parseRaw 6" testA6
    , TestLabel "Test parseRaw 7" testA7
    , TestLabel "Test parseRaw 8" testA8
    , TestLabel "Test parseRaw 9" testA9
    , -- B. parseDate
      TestLabel "Test parseDate 1" testB1
    , TestLabel "Test parseDate 2" testB2
    , TestLabel "Test parseDate 3" testB3
    , TestLabel "Test parseDate 4" testB4
    , TestLabel "Test parseDate 5" testB5
    , TestLabel "Test parseDate 6" testB6
    , -- C. parseDates
      TestLabel "Test parseDates 1" testC1
    , TestLabel "Test parseDates 2" testC2
    , TestLabel "Test parseDates 3" testC3
    , TestLabel "Test parseDates 4" testC4
    , TestLabel "Test parseDates 5" testC5
    , TestLabel "Test parseDates 6" testC6
    ]

main :: IO ()
main = do
  counts <- runTestTT tests
  if errors counts + failures counts == 0
    then exitSuccess
    else exitFailure
