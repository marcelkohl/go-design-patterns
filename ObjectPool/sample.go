package main

import "github.com/google/uuid"

type sample struct {
  id string
}

func (s sample) new() sample {
  return sample{
    id: uuid.New().String(),
  }
}

func (s sample) getId() string {
  return s.id
}
