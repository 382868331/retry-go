package retry

import("context";"errors";"testing";"time")
var(_=context.Background;_=errors.New;_=time.Second)

func TestTaskRetry018ContextOption(t *testing.T){ got:=func() bool { c,cancel:=context.WithCancel(context.Background()); cancel(); return context.Cause(New(Context(c)).context)!=nil }(); if got != true{t.Fatalf("got %v want %v",got,true)} }

func TestTaskRetry018ContextOptionAdjacent(t *testing.T){
 got:=func() bool { c,_:=context.WithTimeout(context.Background(),time.Hour); return New(Context(c)).context==c }()
 if got != true{ t.Fatalf("adjacent got %v want %v",got,true) }
}
