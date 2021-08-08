package main

type iCityFactory interface {
  makeDemographics() iDemographics
  makeExpenses() iExpenses
}
