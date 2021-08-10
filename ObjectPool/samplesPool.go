package main

import (
  "fmt"
  "reflect"
)

type samplesPool struct {
  available   map[string]sampleInterface
	idle        map[string]sampleInterface
  maxPoolSize int
}

func (s samplesPool) new(maxPoolSize int) samplesPool {
  available := make(map[string]sampleInterface, maxPoolSize)
  idle := make(map[string]sampleInterface, maxPoolSize)

  sp := samplesPool{
    available: available,
    idle: idle,
    maxPoolSize: maxPoolSize,
  }

	return sp
}

func (s *samplesPool) acquire() (sampleInterface, error) {
  var obj sampleInterface = sample{}

  if len(s.available) > 0 {
    objKey := reflect.ValueOf(s.available).MapKeys()[0].String()

    obj = s.available[objKey]
    s.idle[objKey] = obj

    delete(s.available, obj.getId())
  } else if len(s.idle) < s.maxPoolSize  {
    obj = (sample{}).new()
    s.idle[obj.getId()] = obj
  } else {
    return nil, fmt.Errorf("Pool is full")
  }

  return s.idle[obj.getId()], nil
}

func (s *samplesPool) release(obj sampleInterface) {
  s.available[obj.getId()] = obj
  delete(s.idle, obj.getId())
}
