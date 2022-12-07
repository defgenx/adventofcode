import gleam/erlang/file
import gleam/io
import gleam/int
import gleam/string
import gleam/list
import gleam/result
import gleam/option.{None, Option, Some}
import gleam/map.{Map}

pub opaque type FileSystem {
  FileSystem(pwd: List(String), files: Map(List(String), Int))
}

// Idea inspired from elixir solution...
pub fn main() {
  assert Ok(f) = file.read("./input.txt")

  let computed_log =
    f
    |> string.split("\n")
    |> list.fold(
      FileSystem(pwd: [], files: map.new()),
      fn(acc, x) { process_row(x, acc) },
    )

  io.debug(computed_log)
  io.debug(part_a(computed_log))
  io.debug(part_b(computed_log))
}

pub fn part_a(fs: FileSystem) -> Int {
  fs.files
  |> map.fold(
    0,
    fn(acc, _, value) {
      case value <= 100000 {
        True -> acc + value
        False -> acc
      }
    },
  )
}

pub fn part_b(fs: FileSystem) -> Int {
  let total_available = 70000000
  let unused_requirement = 30000000

  assert Ok(root_size) = map.get(point_point(fs).files, ["/"])

  let need = unused_requirement - { total_available - root_size }

  map.values(fs.files)
  |> list.sort(int.compare)
  |> list.find(fn(x) { x >= need })
  |> result.unwrap(0)
}

fn point_point(fs: FileSystem) -> FileSystem {
  let base = list.take(fs.pwd, list.length(fs.pwd) - 1)
  let dir_size =
    map.get(fs.files, fs.pwd)
    |> result.unwrap(0)
  FileSystem(pwd: base, files: map.update(fs.files, base, increment(dir_size)))
}

fn in_in(fs: FileSystem, dir: String) -> FileSystem {
  FileSystem(..fs, pwd: list.append(fs.pwd, [dir]))
}

fn process_row(r: String, fs: FileSystem) -> FileSystem {
  case string.split(r, " ") {
    ["$", "ls"] -> fs
    ["$", "cd", ".."] -> point_point(fs)
    ["$", "cd", "/"] -> FileSystem(..fs, pwd: ["/"])
    ["$", "cd", dir] -> in_in(fs, dir)
    ["dir", dir] ->
      FileSystem(
        ..fs,
        files: map.update(fs.files, list.append(fs.pwd, [dir]), increment(0)),
      )
    [size, _] -> {
      assert Ok(file_size) = int.parse(size)
      FileSystem(
        ..fs,
        files: map.update(fs.files, fs.pwd, increment(file_size)),
      )
    }
  }
}

fn increment(size: Int) {
  fn(osize: Option(Int)) {
    case osize {
      Some(o) -> o + size
      None -> 0 + size
    }
  }
}
