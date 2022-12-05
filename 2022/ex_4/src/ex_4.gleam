import gleam/erlang/file
import gleam/io
import gleam/int
import gleam/string
import gleam/order
import gleam/list
import gleam/result
import gleam/iterator.{Done, Next}

pub fn main() {
  assert Ok(f) = file.read("./input.txt")
  let rows =
    f
    |> string.split("\n")

  io.debug(part_a(rows))
  io.debug(part_b(rows))
}

fn part_a(r: List(String)) -> Int {
  r
  |> list.filter(fn(x) {
    let [first, second] = string.split(x, ",")

    parse_range_a(first, second) == 1
  })
  |> list.length()
}

fn part_b(r: List(String)) -> Int {
  r
  |> list.filter(fn(x) {
    let [first, second] = string.split(x, ",")

    parse_range_b(first, second) == 1
  })
  |> list.length()
}

fn parse_range_a(l: String, r: String) -> Int {
  let [lfirst, lsecond] = string.split(l, "-")
  assert Ok(lfi) = int.parse(lfirst)
  assert Ok(lsi) = int.parse(lsecond)
  let [rfirst, rsecond] = string.split(r, "-")
  assert Ok(rfi) = int.parse(rfirst)
  assert Ok(rsi) = int.parse(rsecond)

  case #(int.compare(lfi, rfi), int.compare(lsi, rsi)) {
    #(order.Lt, order.Gt) -> 1
    #(order.Lt, order.Eq) -> 1
    #(order.Eq, order.Gt) -> 1
    #(order.Gt, order.Lt) -> 1
    #(order.Eq, order.Lt) -> 1
    #(order.Gt, order.Eq) -> 1
    #(order.Eq, order.Eq) -> 1
    _ -> 0
  }
}

fn parse_range_b(l: String, r: String) -> Int {
  let [lfirst, lsecond] = string.split(l, "-")
  assert Ok(lfi) = int.parse(lfirst)
  assert Ok(lsi) = int.parse(lsecond)
  let [rfirst, rsecond] = string.split(r, "-")
  assert Ok(rfi) = int.parse(rfirst)
  assert Ok(rsi) = int.parse(rsecond)
  case #(int.compare(lsi, rfi), int.compare(lfi, rsi)) {
    #(order.Lt, _) -> 0
    #(_, order.Gt) -> 0
    _ -> 1
  }
}
