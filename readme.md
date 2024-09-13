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
https://img.shields.io/badge/Golang?style=flat&logo=%3Csvg%20role%3D%22img%22%20viewBox%3D%220%200%2024%2024%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Ctitle%3EGo%3C%2Ftitle%3E%3Cpath%20d%3D%22M1.811%2010.231c-.047%200-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046%200%20.058.035.035.07l-.199.303c-.023.036-.082.07-.117.07zM.047%2011.306c-.047%200-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047%200%20.07.035.058.07l-.093.28c-.012.047-.058.07-.105.07zm2.828%201.075c-.047%200-.059-.035-.035-.07l.163-.292c.023-.035.07-.07.117-.07h2.337c.047%200%20.07.035.07.082l-.023.28c0%20.047-.047.082-.082.082zm12.129-2.36c-.736.187-1.239.327-1.963.514-.176.046-.187.058-.34-.117-.174-.199-.303-.327-.548-.444-.737-.362-1.45-.257-2.115.175-.795.514-1.204%201.274-1.192%202.22.011.935.654%201.706%201.577%201.835.795.105%201.46-.175%201.987-.77.105-.13.198-.27.315-.434H10.47c-.245%200-.304-.152-.222-.35.152-.362.432-.97.596-1.274a.315.315%200%2001.292-.187h4.253c-.023.316-.023.631-.07.947a4.983%204.983%200%2001-.958%202.29c-.841%201.11-1.94%201.8-3.33%201.986-1.145.152-2.209-.07-3.143-.77-.865-.655-1.356-1.52-1.484-2.595-.152-1.274.222-2.419.993-3.424.83-1.086%201.928-1.776%203.272-2.02%201.098-.2%202.15-.07%203.096.571.62.41%201.063.97%201.356%201.648.07.105.023.164-.117.2m3.868%206.461c-1.064-.024-2.034-.328-2.852-1.029a3.665%203.665%200%2001-1.262-2.255c-.21-1.32.152-2.489.947-3.529.853-1.122%201.881-1.706%203.272-1.95%201.192-.21%202.314-.095%203.33.595.923.63%201.496%201.484%201.648%202.605.198%201.578-.257%202.863-1.344%203.962-.771.783-1.718%201.273-2.805%201.495-.315.06-.63.07-.934.106zm2.78-4.72c-.011-.153-.011-.27-.034-.387-.21-1.157-1.274-1.81-2.384-1.554-1.087.245-1.788.935-2.045%202.033-.21.912.234%201.835%201.075%202.21.643.28%201.285.244%201.905-.07.923-.48%201.425-1.228%201.484-2.233z%22%2F%3E%3C%2Fsvg%3E&logoSize=auto&labelColor=rgb(%3Csvg%20role%3D%22img%22%20viewBox%3D%220%200%2024%2024%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Ctitle%3EGo%3C%2Ftitle%3E%3Cpath%20d%3D%22M1.811%2010.231c-.047%200-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046%200%20.058.035.035.07l-.199.303c-.023.036-.082.07-.117.07zM.047%2011.306c-.047%200-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047%200%20.07.035.058.07l-.093.28c-.012.047-.058.07-.105.07zm2.828%201.075c-.047%200-.059-.035-.035-.07l.163-.292c.023-.035.07-.07.117-.07h2.337c.047%200%20.07.035.07.082l-.023.28c0%20.047-.047.082-.082.082zm12.129-2.36c-.736.187-1.239.327-1.963.514-.176.046-.187.058-.34-.117-.174-.199-.303-.327-.548-.444-.737-.362-1.45-.257-2.115.175-.795.514-1.204%201.274-1.192%202.22.011.935.654%201.706%201.577%201.835.795.105%201.46-.175%201.987-.77.105-.13.198-.27.315-.434H10.47c-.245%200-.304-.152-.222-.35.152-.362.432-.97.596-1.274a.315.315%200%2001.292-.187h4.253c-.023.316-.023.631-.07.947a4.983%204.983%200%2001-.958%202.29c-.841%201.11-1.94%201.8-3.33%201.986-1.145.152-2.209-.07-3.143-.77-.865-.655-1.356-1.52-1.484-2.595-.152-1.274.222-2.419.993-3.424.83-1.086%201.928-1.776%203.272-2.02%201.098-.2%202.15-.07%203.096.571.62.41%201.063.97%201.356%201.648.07.105.023.164-.117.2m3.868%206.461c-1.064-.024-2.034-.328-2.852-1.029a3.665%203.665%200%2001-1.262-2.255c-.21-1.32.152-2.489.947-3.529.853-1.122%201.881-1.706%203.272-1.95%201.192-.21%202.314-.095%203.33.595.923.63%201.496%201.484%201.648%202.605.198%201.578-.257%202.863-1.344%203.962-.771.783-1.718%201.273-2.805%201.495-.315.06-.63.07-.934.106zm2.78-4.72c-.011-.153-.011-.27-.034-.387-.21-1.157-1.274-1.81-2.384-1.554-1.087.245-1.788.935-2.045%202.033-.21.912.234%201.835%201.075%202.21.643.28%201.285.244%201.905-.07.923-.48%201.425-1.228%201.484-2.233z%22%2F%3E%3C%2Fsvg%3E)&color=rgb(27%2C%2028%2C%2027)
