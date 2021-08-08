package main

type iExpenses interface {
  getForecast() float32
  getSpent() float32
}

type expenses struct {
  forecast float32
  spent float32
}

func (e *expenses) getForecast() float32 {
    return e.forecast
}

func (e *expenses) getSpent() float32 {
    return e.spent
}
