module Main where

import Main.Input
import System.Exit (exitFailure, exitSuccess)
import Test.HUnit

testA1 =
  TestCase
    ( assertEqual
        "A1 - case 1. like 2020.1.1 - basic"
        ["2020.1.1"]
        $ map printTime
        $ parseDate' ["2020.1.1"]
    )

testA2 =
  TestCase
    ( assertEqual
        "A2 - case 1. like 2020.1.1 - robust"
        ["2021.11.2"]
        (map printTime $ parseDate' ["2021.11.02"])
    )

testA3 =
  TestCase
    ( assertEqual
        "A3 - case 2. like 2020.1.1-2020.1.3 - basic"
        ["2022.1.1", "2022.1.2", "2022.1.3"]
        (map printTime $ parseDate' ["2022.1.1", "2022.1.3"])
    )

testA4 =
  TestCase
    ( assertEqual
        "A4 - case 2. like 2020.1.1-2020.1.3 - cross month"
        ["2023.2.27", "2023.2.28", "2023.3.1", "2023.3.2"]
        (map printTime $ parseDate' ["2023.2.27", "2023.3.2"])
    )

testA5 =
  TestCase
    ( assertEqual
        "A5 - case 2. like 2020.1.1-2020.1.3 - cross month 2"
        ["2024.2.28", "2024.2.29", "2024.3.1"]
        (map printTime $ parseDate' ["2024.2.28", "2024.3.1"])
    )

testA6 =
  TestCase
    ( assertEqual
        "A6 - case 2. like 2020.1.1-2020.1.3 - cross year"
        ["2025.12.30", "2025.12.31", "2026.1.1", "2026.1.2"]
        (map printTime $ parseDate' ["2025.12.30", "2026.1.2"])
    )

testB1 =
  TestCase
    ( assertEqual
        "B1 - case 1. like 2020.1.1 - a1 2"
        [(1, 2, "2020.1.1"), (2, 2, "2020.1.3")]
        $ map (\(a, b, c) -> (a, b, printTime c))
        $ parseDate "2020.1.1,2020.1.3"
    )

testB2 =
  TestCase
    ( assertEqual
        "B2 - case 1. like 2020.1.1 - a1 3"
        [(1, 3, "2020.1.1"), (2, 3, "2020.1.3"), (3, 3, "2020.1.5")]
        $ map (\(a, b, c) -> (a, b, printTime c))
        $ parseDate "2020.1.1,2020.1.3,2020.1.5"
    )

testB3 =
  TestCase
    ( assertEqual
        "B3 - case 2. like 2020.1.1-2020.1.3 - a3 2"
        [(1, 4, "2020.1.1"), (2, 4, "2020.1.2"), (3, 4, "2020.1.6"), (4, 4, "2020.1.7")]
        $ map (\(a, b, c) -> (a, b, printTime c))
        $ parseDate "2020.1.1-2020.1.2,2020.1.6-2020.1.7"
    )

testB4 =
  TestCase
    ( assertEqual
        "B4 - case 2. like 2020.1.1-2020.1.3 - a1 a3"
        [(1, 4, "2020.1.1"), (2, 4, "2020.1.2"), (3, 4, "2020.1.3"), (4, 4, "2020.1.6")]
        $ map (\(a, b, c) -> (a, b, printTime c))
        $ parseDate "2020.1.1-2020.1.3,2020.1.6"
    )

testB5 =
  TestCase
    ( assertEqual
        "B5 - case 2. like 2020.1.1-2020.1.3 - a3 a1"
        [(1, 5, "2019.12.6"), (2, 5, "2020.1.1"), (3, 5, "2020.1.2"), (4, 5, "2020.1.3"), (5, 5, "2021.11.11")]
        $ map (\(a, b, c) -> (a, b, printTime c))
        $ parseDate "2019.12.6,2020.1.1-2020.1.3,2021.11.11"
    )

testC1 =
  TestCase
    ( assertEqual
        "C1 - data - basic"
        [["a", "b", "c"]]
        $ parseFile "a;b;c"
    )

testC2 =
  TestCase
    ( assertEqual
        "C2 - data - robust"
        [["a12", "b345", "c678"]]
        $ parseFile "   a12;b345;c678    "
    )

testC3 =
  TestCase
    ( assertEqual
        "C3 - comments - basic"
        []
        $ parseFile "//comment"
    )

testC4 =
  TestCase
    ( assertEqual
        "C4 - comments - robust"
        []
        $ parseFile "   // a;b;c  123  "
    )

testC5 =
  TestCase
    ( assertEqual
        "C5 - hybrid - basic"
        [["a", "b", "c"]]
        $ parseFile "a;b;c  // d;e;f  456  "
    )

testC6 =
  TestCase
    ( assertEqual
        "C6 - hybrid - robust"
        [["a", "b", "c"]]
        $ parseFile "   a;b;c 123 // d;e;f  456  "
    )

testC7 =
  TestCase
    ( assertEqual
        "C7 - empty - basic"
        []
        $ parseFile ""
    )

testC8 =
  TestCase
    ( assertEqual
        "C8 - empty - robust"
        []
        $ parseFile "  //   "
    )

testC9 =
  TestCase
    ( assertEqual
        "C9 - multi-line"
        [["a", "b", "c"], ["d", "e", "f"]]
        $ parseFile . unlines
        $ [ "// INTRO",
            "",
            "a;b;c // dp1",
            "// comment",
            "d;e;f // dp2",
            "END  // TEST"
          ]
    )

tests :: Test
tests =
  TestList
    [ -- A. parseDate'
      TestLabel "Test parseDate' 1" testA1,
      TestLabel "Test parseDate' 2" testA2,
      TestLabel "Test parseDate' 3" testA3,
      TestLabel "Test parseDate' 4" testA4,
      TestLabel "Test parseDate' 5" testA5,
      TestLabel "Test parseDate' 6" testA6,
      -- B. parseDate
      TestLabel "Test parseDate 1" testB1,
      TestLabel "Test parseDate 2" testB2,
      TestLabel "Test parseDate 3" testB3,
      TestLabel "Test parseDate 4" testB4,
      TestLabel "Test parseDate 5" testB5,
      -- C. parseFile
      TestLabel "Test parseFile 1" testC1,
      TestLabel "Test parseFile 2" testC2,
      TestLabel "Test parseFile 3" testC3,
      TestLabel "Test parseFile 4" testC4,
      TestLabel "Test parseFile 5" testC5,
      TestLabel "Test parseFile 6" testC6,
      TestLabel "Test parseFile 7" testC7,
      TestLabel "Test parseFile 8" testC8,
      TestLabel "Test parseFile 9" testC9
    ]

main :: IO ()
main = do
  counts <- runTestTT tests
  if errors counts + failures counts == 0
    then exitSuccess
    else exitFailure
