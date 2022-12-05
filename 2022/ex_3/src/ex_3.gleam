import gleam/erlang/file
import gleam/erlang/charlist
import gleam/io
import gleam/int
import gleam/string
import gleam/list
import gleam/result

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
  |> list.map(fn(x) {
    let sub_string =
      x
      |> string.split("")

    let #(first, second) = list.split(sub_string, list.length(sub_string) / 2)

    first
    |> list.filter(list.contains(second, _))
    |> list.first
    |> result.unwrap("")
  })
  |> sum_result
}

fn part_b(r: List(String)) -> Int {
  let chunked_list = list.sized_chunk(r, 3)

  chunked_list
  |> list.map(fn(x) {
    let [first, second, third] = x

    first
    |> string.split("")
    |> list.filter(fn(x) {
      string.contains(second, x) && string.contains(third, x)
    })
    |> list.first
    |> result.unwrap("")
  })
  |> sum_result
}

fn priority(item: String) -> Int {
  case <<item:utf8>> {
    <<p:int>> if p > 97 -> p - 96
    <<p:int>> -> p - 38
  }
}

fn sum_result(l: List(String)) -> Int {
  l
  |> list.map(fn(x) { priority(x) })
  |> int.sum
}
