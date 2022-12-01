import gleam/erlang/file
import gleam/io
import gleam/string
import gleam/list
import gleam/int
import gleam/pair
import gleam/order

pub fn main() {
  assert Ok(f) = file.read("./input.txt")
  let parsed_file = parse_input(f)
  sum_elfs(parsed_file, 1)
  sum_elfs(parsed_file, 3)
}

fn sum_elfs(elf_list: List(Int), take: Int) {
  elf_list
  |> list.sort(fn(a, b) {
    int.compare(a, b)
    |> order.reverse
  })
  |> list.take(take)
  |> int.sum()
  |> int.to_string()
  |> io.println()
}

fn parse_input(input: String) -> List(Int) {
  input
  |> string.split("\n\n")
  |> list.map(fn(x) {
    x
    |> string.split("\n")
    |> list.map(fn(x) {
      assert Ok(val_int) =
        x
        |> int.parse()
      val_int
    })
    |> int.sum()
  })
}
