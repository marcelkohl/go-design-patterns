package main

type cityXyFactory struct {
}

func (c *cityXyFactory) makeDemographics() demographicsInterface {
  return &cityXyDemographics{
    demographics: demographics{
      population: 426,
      majorityGender: "female",
    },
  }
}

func (c *cityXyFactory) makeExpenses() expensesInterface {
  return &cityXyExpenses{
    expenses: expenses{
      forecast: 2410.21,
      spent: 1005.72,
    },
  }
}
