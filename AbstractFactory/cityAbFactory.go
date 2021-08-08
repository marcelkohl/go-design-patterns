package main

type cityAbFactory struct {
}

func (c *cityAbFactory) makeDemographics() iDemographics {
  return &cityAbDemographics{
    demographics: demographics{
      population: 331,
      majorityGender: "male",
    },
  }
}

func (c *cityAbFactory) makeExpenses() iExpenses {
  return &cityAbExpenses{
    expenses: expenses{
      forecast: 1832.88,
      spent: 113.55,
    },
  }
}
