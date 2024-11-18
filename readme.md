# Raptor Solar Analysis

## Objectives
- To provide authenticated data on global resource consumption and production from verified sources:
    - Global Resource Consumption: [https://worldpopulationreview.com/](https://worldpopulationreview.com/)
    - Global Resource Production: [https://ourworldindata.org/fossil-fuels](https://ourworldindata.org/fossil-fuels)
- To provide a tool to analyze a suitable plan to setup your own residential solar panel based on your requirements.

## Subprojects
The Subprojects within this project are:
- ResourceRadar
- RaptorSolar

## Working
The working of these projects is as follows:
- ResourceRadar
    - Global Resource Consumption
        - The web application displays a form that asks you to enter the country for which you want to view the data.
        - The data for each country is from 2023.
        - About 210 countries and international bodies have been included in the database.
        - Example: 
        The user enters "India" into the form. The form then processes and displays the resource consumption data for India. It includes various resources like Coal, Electricity, Oil, Natural Gas and Biofuel.
        - NOTE: The form is not case sensitive.

    - Global Resource Production
        - The web application displays a form that asks the user for the desired country and year, for which you want to retrieve the data.
        - The data for each country is available from the late 1900s.
        - Example:
        The user enters "India" in the country field and "2010" in the year field. The user will then recieve data about the resources produced by India in the year 2010.
        - NOTE: The form is not case sensitive.

- RaptorSolar
  - The web application displays a form that includes the following fields:
      - Land area the user is willing to devote to the solar panel setup.
      - The minimum power required by the user, to be generated from the solar panel setup.
      - The current annual electricity bill being paid by the user.
  - Based on the inputs form the user, the form processes the data and then displays all the suitable solar setup plans that suit the user's needs, which has been retrieved from the form.
  - All the plans are from [TataSolar](https://www.tatapowersolar.com/rooftops/residential/)
  - For each plan, a profit calculator has been included.
  - As solar setups are one time investments, and electricity bills are paid regularly, there will be a point in time or  "break even" when the user will start gaining profits from his/her investments.
  - The profit calculator, hence calculates the approximate number of years after which the investment in Tata Solar becomes a profit.
  - NOTE: The solar setup plans include both: off-grid and on-grid.

## Use Cases
- ResourceRadar
  - One place to get authentic data on global resource consumption and production from a span of many years.

- RaptorSolar
  - One place to analyze and a find out a solar panel setup plan for you.
  - To get an approximate break even using a profit calculator custom designed by me.

## Tools Used
- Golang
- PostgresSQL
- Go Templates
- CSS
- Bootstrap v5.3

## External Libraries
- go-chi
- jackc/pgconn
- crypto
- 