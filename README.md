# Advent of Code 2024 🎄🎁
This repo contains my solutions for different years.
- 2023 solutions in Golang [here](./2023/).
- 2024 solutions in Rust [here](./2024/).

## Days, stars and solutions 
> Benchmarks are run on an "Intel i7-10700K @ 5.100Ghz CPU".
> Use the `--release` flag when benchmarking a day.

| Day | Solution | Stars | Parse benchmark | Part 1 benchmark | Part 2 benchmark |
|-----|----------|-------|-------------------|-----------------|------------------|
| [Day 1:](https://adventofcode.com/2024/day/1) Historian Hysteria | [Here](./2024/solutions/day01.rs) | ⭐⭐ | 63µ | 264ns | 50µs |
| [Day 2:](https://adventofcode.com/2024/day/2) Red-Nosed Reports | [Here](./2024/solutions/day02.rs) | ⭐⭐ | 142µs | 7µs | 57µs |
| [Day 3:](https://adventofcode.com/2024/day/3) Mull It Over | [Here](./2024/solutions/day03.rs) | ⭐⭐ | 65µs | 345ns | 466ns |
| [Day 4:](https://adventofcode.com/2024/day/4) Ceres Search | [Here](./2024/solutions/day04.rs) | ⭐⭐ | 7µs | 1ms | 360µs |
| [Day 5:](https://adventofcode.com/2024/day/5) Print Queue | [Here](./2024/solutions/day05.rs) | ⭐⭐ | 201µs | 395µs | 797µs |
| [Day 6:](https://adventofcode.com/2024/day/6) Guard Gallivant | [Here](./2024/solutions/day06.rs) | ⭐⭐ | 36µs | 251µs | 1s |
| [Day 7:](https://adventofcode.com/2024/day/7) Bridge Repair | [Here](./2024/solutions/day07.rs) | ⭐⭐ | 455µs | 578µs | 84ms |
| [Day 8:](https://adventofcode.com/2024/day/8) Resonant Collinearity | [Here](./2024/solutions/day08.rs) | ⭐⭐ | 21µs | 116µs | 92µs |
| [Day 9:](https://adventofcode.com/2024/day/9) Disk Fragmenter | [Here](./2024/solutions/day09.rs) | ⭐⭐ | 1ms | 520ms | 332ms |
| [Day 10:](https://adventofcode.com/2024/day/10) Hoof It | [Here](./2024/solutions/day10.rs) | ⭐⭐ | 65µs | 489µs | 436µs |
| [Day 11:](https://adventofcode.com/2024/day/11) Plutonian Pebbles | [Here](./2024/solutions/day11.rs) | ⭐⭐ | 1µs | 582µs | 32ms |
| [Day 12:](https://adventofcode.com/2024/day/12) Garden Groups | [Here](./2024/solutions/day12.rs) | ⭐⭐ | 221µs | 7ms | 11ms |
| [Day 13:](https://adventofcode.com/2024/day/13) Claw Contraption | [Here](./2024/solutions/day13.rs) | ⭐⭐ | 158µs | 3µs | 6µs |
| [Day 14:](https://adventofcode.com/2024/day/14) Restroom Redoubt | [Here](./2024/solutions/day14.rs) | ⭐⭐ | 94µs | 180µs | 790ms |
| [Day 15:](https://adventofcode.com/2024/day/15) Warehouse Woes | [Here](./2024/solutions/day15.rs) | ⭐⭐ | 210µs | 130µs | 3ms |
| [Day 16:](https://adventofcode.com/2024/day/16) Reindeer Maze | [Here](./2024/solutions/day16.rs) | ⭐⭐ | 20µs | 3s | 6s |
| [Day 17:](https://adventofcode.com/2024/day/17) Chronospatial Computer | [Here](./2024/solutions/day17.rs) | ⭐ | 532ns | 224ns | - |
| [Day 18:](https://adventofcode.com/2024/day/18) RAM Run | [Here](./2024/solutions/day18.rs) | ⭐⭐ | 255µs | 113µs | 257ms |
| [Day 19:](https://adventofcode.com/2024/day/19) Linen Layout | [Here](./2024/solutions/day19.rs) | ⭐⭐ | 64µs | 4ms | 30ms |
| [Day 20:](https://adventofcode.com/2024/day/20) Race Condition | [Here](./2024/solutions/day20.rs) | ⭐⭐ | 76µs | 1ms | 42ms |
| [Day 21:](https://adventofcode.com/2024/day/21) Keypad Conundrum | [Here](./2024/solutions/day21.rs) | ⭐ | 1µs | 6ms | - |
| [Day 22:](https://adventofcode.com/2024/day/22) Monkey Market | [Here](./2024/solutions/day22.rs) | ⭐⭐ | 30µs | 5ms | 394ms |
| [Day 23:](https://adventofcode.com/2024/day/23) LAN Party | [Here](./2024/solutions/day23.rs) | ⭐⭐ | 584µs | 2ms | 26ms |
| [Day 24:](https://adventofcode.com/2024/day/24) Crossed Wires | [Here](./2024/solutions/day24.rs) | ⭐ | 244µs | 1ms | - |
| [Day 25:](https://adventofcode.com/2024/day/25) Code Chronicle | [Here](./2024/solutions/day25.rs) | ⭐ | 293µs | 220µs | - |

## Run it yourself
### Run a day 
While being at the [2024](./2024/) directory:
```
cargo run -- [--day <day>] [--sample] [--bench <n>]
```
> --day=<1-25> Optional (default:`1`). Specifies the day to execute the puzzle `(1-25)`

> --sample: Optional (default: `false`). If set, the input file will be the `sample.txt` located in the [inputs](./2024/inputs) folder.

> --bench=\<n\>: Optional (default: `0`). If set, the puzzles will run `n` times, outputing only the average timings.

#### Example output
![Output screenshot](https://github.com/user-attachments/assets/072b854a-4e15-4284-a5c5-3745c6bd0f76)
