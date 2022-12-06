import gleam/erlang/file
import gleam/io
import gleam/option.{Some}
import gleam/int
import gleam/string
import gleam/regex
import gleam/list
import gleam/map.{Map}

pub fn main() {
  assert Ok(f) = file.read("./input.txt")
  let [stack, moves] =
    f
    |> string.split("\n\n")

  let parsed_stack = parse_stack(stack)
  let parsed_moves = parse_moves(moves)

  io.debug(parsed_stack)
  io.debug(part_a(parsed_stack, parsed_moves))
  io.debug(part_b(parsed_stack, parsed_moves))
}

fn parse_stack(stack: String) -> Map(Int, List(String)) {
  let stack_rows =
    stack
    |> string.split("\n")
    |> list.reverse()
    |> list.map(fn(x) {
      x
      |> string.split("")
      |> list.sized_chunk(4)
      |> list.fold(
        from: [],
        with: fn(acc, x) {
          case x {
            [" ", " ", " ", " "] -> list.append(acc, [""])
            [" ", " ", " "] -> list.append(acc, [""])
            [" ", _, " ", " "] -> acc
            [" ", _, " "] -> acc
            ["[", a, "]", " "] -> list.append(acc, [a])
            ["[", a, "]"] -> list.append(acc, [a])
          }
        },
      )
    })
    |> list.transpose()
    |> list.map(fn(x) {
      x
      |> list.filter(fn(x) { x != "" })
    })

  stack_rows
  |> list.index_fold(
    from: map.new(),
    with: fn(acc, x, i) { map.insert(acc, i + 1, x) },
  )
}

fn part_a(stack: Map(Int, List(String)), moves: List(List(Int))) {
  moves
  |> list.fold(
    from: stack,
    with: fn(acc, move) {
      let [nb, from, to] = move
      make_move(acc, #(nb, from, to), True)
    },
  )
}

fn part_b(stack: Map(Int, List(String)), moves: List(List(Int))) {
  moves
  |> list.fold(
    from: stack,
    with: fn(acc, move) {
      let [nb, from, to] = move
      make_move(acc, #(nb, from, to), False)
    },
  )
}

fn parse_moves(moves: String) -> List(List(Int)) {
  assert Ok(moves_regex) =
    regex.from_string("move (\\d+) from (\\d+) to (\\d+)")

  moves
  |> string.split("\n")
  |> list.map(fn(m) {
    assert [match] = regex.scan(with: moves_regex, content: m)
    assert [Some(Ok(nb)), Some(Ok(from)), Some(Ok(to))] =
      match.submatches
      |> list.map(option.map(_, int.parse))

    [nb, from, to]
  })
}

fn make_move(
  stack: Map(Int, List(String)),
  move: #(Int, Int, Int),
  reverse: Bool,
) -> Map(Int, List(String)) {
  let #(nb, from, to) = move

  case nb {
    0 -> stack
    i -> {
      assert Ok(from_val) = map.get(stack, from)
      assert Ok(to_val) = map.get(stack, to)
      let #(new_from_val, move_vals) =
        list.split(from_val, at: list.length(from_val) - i)
      let new_to_val = case reverse {
        True ->
          list.append(
            to_val,
            move_vals
            |> list.reverse(),
          )
        False -> list.append(to_val, move_vals)
      }
      let stack = map.insert(stack, from, new_from_val)
      map.insert(stack, to, new_to_val)
    }
  }
}
