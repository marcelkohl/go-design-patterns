package main

type cityXyFactory struct {
}

func (c *cityXyFactory) makeDemographics() iDemographics {
  return &cityXyDemographics{
    demographics: demographics{
      population: 426,
      majorityGender: "female",
    },
  }
}

func (c *cityXyFactory) makeExpenses() iExpenses {
  return &cityXyExpenses{
    expenses: expenses{
      forecast: 2410.21,
      spent: 1005.72,
    },
  }
}
