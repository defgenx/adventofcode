import gleam/erlang/file
import gleam/io
import gleam/string
import gleam/list
import gleam/int
import gleam/pair
import gleam/order

pub fn main() {
  assert Ok(f) = file.read("./input.txt")
  io.debug(solution_a(f))
  io.debug(solution_b(f))
}

fn solution_a(input: String) -> Int {
  input
  |> string.split("\n")
  |> list.map(fn(x) {
    x
    |> string.split(" ")
    |> translate_points()
  })
  |> sum_points_result
}

fn solution_b(input: String) -> Int {
  input
  |> string.split("\n")
  |> list.map(fn(x) {
    x
    |> string.split(" ")
  })
  |> list.map(fn(x) {
    map_secret(x)
    |> translate_points()
  })
  |> sum_points_result
}

fn translate_points(l: List(String)) -> List(Int) {
  l
  |> list.map(fn(x) {
    case x {
      i if i == "A" || i == "X" -> 1
      i if i == "B" || i == "Y" -> 2
      i if i == "C" || i == "Z" -> 3
    }
  })
}

fn map_secret(l: List(String)) -> List(String) {
  assert Ok(first) = list.first(l)
  assert Ok(second) = list.last(l)
  case #(first, second) {
    #(i, "Y") -> [i, i]
    #("A", "X") -> ["A", "C"]
    #("B", "X") -> ["B", "A"]
    #("C", "X") -> ["C", "B"]
    #("A", "Z") -> ["A", "B"]
    #("B", "Z") -> ["B", "C"]
    #("C", "Z") -> ["C", "A"]
  }
}

fn sum_points_result(l: List(List(Int))) -> Int {
  l
  |> list.map(fn(x) {
    assert Ok(first) = list.first(x)
    assert Ok(second) = list.last(x)
    case #(first, second) {
      #(i, j) if i == j -> i + 3
      #(i, j) if i == 1 && j == 2 || i == 2 && j == 3 || i == 3 && j == 1 ->
        j + 6
      #(_, j) -> j
    }
  })
  |> int.sum
}
