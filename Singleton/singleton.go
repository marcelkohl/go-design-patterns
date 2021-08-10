package main

import (
  "fmt"
  "sync"
)

type singleton struct {
  lock *sync.Mutex
  instance *singleton
  counter int
}

func (s singleton) new() singleton {
  if s.lock == nil {
    return singleton{
      lock: &sync.Mutex{},
    }
  } else {
    return s
  }
}

func (s *singleton) getInstance() singleton {
  s.counter++

  if s.instance == nil {
    fmt.Println("Trying to create instance...")
    s.lock.Lock()
    defer s.lock.Unlock()

    if s.instance == nil {
      fmt.Println("Creating Single Instance Now!")
      s.instance = &singleton{
        instance: s,
      }
    } else {
      fmt.Println("Tried to create instance but another process created it first. Attempt n.: ", s.counter)
    }
  } else {
    fmt.Println("Instance already exists. Attempt n.: ", s.counter)
  }

  return *s.instance
}
