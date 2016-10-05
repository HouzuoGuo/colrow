package main

import (
	"fmt"
	"testing"
)

func TestGetColumns(t *testing.T) {
	in := `0123456789
0123456789
0123456789
0123456789`
	if out := getColumns(in, 0, 0, 2, 2); out != `012
012
012` {
		t.Fatal(out)
	}
	if out := getColumns(in, 8, 1, 1000, 1000); out != `89
89
89` {
		t.Fatal(out)
	}

	in = `
howard@howie ~/g/s/g/H/colrow> git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   .gitignore

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        LICENSE
        main.go

no changes added to commit (use "git add" and/or "git commit -a")
howard@howie ~/g/s/g/H/colrow> git add .
howard@howie ~/g/s/g/H/colrow> l
total 2.5M
-rwx------ 1 howard users 2.5M Oct  5 09:37 colrow*
-rw------- 1 howard users 1.3K Oct  5 09:14 LICENSE
-rw------- 1 howard users 1.8K Oct  5 09:37 main.go
-rw------- 1 howard users   74 Oct  5 09:13 README.md
`
	expected := `
in/master'.

 what will be c
o discard chang




e in what will




 add" and/or "g
dd .


5 09:37 colrow*`
	if out := getColumns(in, 36, 2, 50, 20); out != expected {
		fmt.Println(out)
		t.Fatalf("\n%v\n%v", []byte(out), []byte(expected))
	}
}
