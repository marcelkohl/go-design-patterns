package main

type iDemographics interface {
  getPopulation() int
  getMajorityGender() string
}

type demographics struct {
  population int
  majorityGender string
}

func (d *demographics) getMajorityGender() string {
    return d.majorityGender
}

func (d *demographics) getPopulation() int {
    return d.population
}
