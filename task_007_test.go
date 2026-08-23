package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry007PerErrorBudget(t *testing.T){ got:=func() uint { e:=errors.New("x"); return New(AttemptsForError(2,e)).attemptsForError[e] }(); if got != uint(2){t.Fatalf("got %v want %v",got,uint(2))} }

func TestTaskRetry007PerErrorBudgetAdjacent(t *testing.T){
 got:=func() uint { e:=errors.New("y"); return New(AttemptsForError(1,e)).attemptsForError[e] }()
 if got != uint(1){ t.Fatalf("adjacent got %v want %v",got,uint(1)) }
}
