let parse =
  try
    let channel = open_in "input.txt" in
    let content = really_input_string channel (in_channel_length channel) in
    close_in channel;
    content
  with
  | Sys_error _ ->
    Printf.printf "Error: Could not open file\n";
    exit 1
;;

let count_differences row_1 row_2 =
  List.fold_left2 (fun acc c1 c2 -> if c1 <> c2 then acc + 1 else acc) 0 row_1 row_2
;;

let transpose matrix =
  List.map List.rev (List.rev matrix)
;;
  
let find_reflection rows use_smudge =
  let len = List.length rows in
  let rec check_reflection i =
    if i >= len then 0
    else
      let smudge_used = ref (not use_smudge) in
      let reflective = ref false in
      let rec check_rows j =
        if !reflective then ()
        else if j > i then ()
        else
          match List.nth_opt rows (i - j), List.nth_opt rows (i + j + 1) with
          | Some row_1, Some row_2 ->
            let differences = count_differences row_1 row_2 in
            reflective := differences = 0;
            if not !smudge_used then (
              reflective := differences <= 1;
              smudge_used := differences = 1
            );
            if not !reflective then ()
            else check_rows (j + 1)
          | _, _ -> ()
      in
      check_rows 0;
      if !reflective && !smudge_used then i + 1
      else check_reflection (i + 1)
  in
  check_reflection 0;;

let solve input use_smudge =
  let result = ref 0 in
  String.split_on_char '\n' input |> List.iter (fun pattern ->
    let rows = String.split_on_char '\n' pattern |> List.map (fun l -> List.init (String.length l) (fun i -> String.get l i)) in
    let h1 = find_reflection rows use_smudge in
    let chars_transposed = transpose rows in
    let h2 = find_reflection chars_transposed use_smudge in
    result := !result + h1 * 100 + h2
  );
  !result
;;

let solve_part1 input = solve input false
;;

let solve_part2 input = solve input true
;;

(* #trace find_reflection;; *)
(* main *)
let () =
  let input = parse in
  let p1 = solve_part1 input in
  let p2 = solve_part2 input in
  Printf.printf "part 1: %d\npart 2: %d\n" p1 p2
;;