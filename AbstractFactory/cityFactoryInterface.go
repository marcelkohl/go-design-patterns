package main

type cityFactoryInterface interface {
  makeDemographics() demographicsInterface
  makeExpenses() expensesInterface
}
