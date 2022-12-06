import gleam/erlang/file
import gleam/io
import gleam/string
import gleam/list

pub fn main() {
  assert Ok(f) = file.read("./input.txt")

  io.debug(find_first(list_markers(f, 4)))
  io.debug(find_first(list_markers(f, 14)))
}

fn find_first(res: List(#(Int, String, Bool))) -> #(Int, String, Bool) {
  let res =
    res
    |> list.filter(fn(x) {
      let #(_, _, valid) = x
      valid == True
    })

  assert Ok(res_first) = list.first(res)
  res_first
}

fn list_markers(l: String, size: Int) -> List(#(Int, String, Bool)) {
  let list =
    l
    |> string.split("")
    |> list.window(size)

  list
  |> list.index_map(fn(index, char_window) {
    let found = list.unique(char_window)
    #(
      index + size,
      string.join(char_window, ""),
      list.length(found) == list.length(char_window),
    )
  })
}
