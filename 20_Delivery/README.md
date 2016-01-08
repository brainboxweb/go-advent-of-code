--- Day 20: Infinite Elves and Infinite Houses ---
===

Notes
--

Full disclosure: I got the right answers for this one... but I got them very, very slowly ( - I'm talking hours run time!).

I've gone back and improved the solution based on the ideas on this page https://www.reddit.com/r/adventofcode/comments/3xjpp2/day_20_solutions/

Hats off to those who figured out the correct algorithm on their own. I'm very impressed ;)


Tests
---
        go test

Part One
---

        presentCount := 36000000
        presentsPerElf := 10
        limitPerElf := -1

        result := GetHouseNumber(presentCount, presentsPerElf, limitPerElf)  //Expect 831600



Part Two
---
        presentCount := 36000000
        presentsPerElf := 11
        limitPerElf := 50

        result := GetHouseNumber(presentCount, presentsPerElf, limitPerElf)  //Expect 884520

