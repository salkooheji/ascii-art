#!/bin/bash
Continue() {
  echo "Press any key to continue"
  while [ true ]; do
    read -t 3 -n 1
    if [ $? = 0 ]; then
      clear
      return
    fi
  done
}
clear

printf "test 1 'hello':\n" | cat -e
go run . "hello" | cat -e
Continue

printf "test 2 'HELLO':\n" | cat -e
go run . "HELLO" | cat -e
Continue

printf "test 3 'HeLlo HuMaN':\n" | cat -e
go run . "HeLlo HuMaN" | cat -e
Continue

printf "test 4 '1Hello 2There':\n" | cat -e
go run . "1Hello 2There" | cat -e
Continue

printf "test 5 'Hello\\\\nThere':\n" | cat -e
go run . "Hello\nThere" | cat -e
Continue

printf "test 6 'Hello\\\\n\\\\nThere':\n" | cat -e
go run . "Hello\n\nThere" | cat -e
Continue

printf "test 7 '{Hello & There #}':\n" | cat -e
go run . "{Hello & There #}" | cat -e
Continue

printf "test 8 'hello There 1 to 2!':\n" | cat -e
go run . 'hello There 1 to 2!' | cat -e
Continue

printf "test 9 'MaD3IrA&LiSboN':\n" | cat -e
go run . "MaD3IrA&LiSboN" | cat -e
Continue

printf "test 10 '1a\"#FdwHywR&/()=':\n" | cat -e
go run . "1a\"#FdwHywR&/()=" | cat -e
Continue