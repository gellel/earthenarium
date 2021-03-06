# earthenarium

## Binary
`go get github.com/gellel/earthenarium`
## Source
`git clone github.com/gellel/earthenarium`

### example 
`$ earthenarium`

### About 
Earthernarium is a small weather program intended to generate generalised climate seeds based on proximities to notable geographic ranges across a pseudo-earth planet. It uses a set of defined boundaries (Arctic Circle, Tropic of Cancer, Equator, Tropic of Capricorn, Antarctic Circle) to construct a normal distribution of random numbers to simulate the number of variable conditions for a seasonal range. Climates gradually shift from colder to warmer as the longtitudal reference slides closer towards the globes center band (minimum bounds Tropic of Cancer and maximum bounds Tropic of Capricon). 

![screenshot](https://raw.githubusercontent.com/gellel/earthenarium/master/capture.PNG)

## Output
Running the program produces a tablised set of data, containing the following

| Location | Position | Local Time | Condition | Temperature | Pressure | Humidity | Season | Region | Hemisphere |
|----------|----------|------------|-----------|-------------|----------|----------|--------|--------|------------|
| El Aauin | 27.1418,-13.18797,68 | 1903-3-24 0:10:0 | Sunny | +11.0 | 1005.10767 | 58.1 | Spring | Tropic of Cancer | Northern |

#### Generated table contents
Location is city where the measurement is conducted.
Position is a comma-separated triple containing latitude, longitude, and elevation in metres above sea level.
Local time is an ISO8601 date time,
Conditions are either Snow, Rain or Sunny. 
Temperature is in °C.
Pressure is in hPa
Relative humidity is a %.
Season is the sampling season collected from the generated timestamp.
Region is the temperate band of the planet.
Hemisphere is the positional reference.

## Support
###### Locations
Locations that can be calculated are based on their availability within the Timezone database. These can be found at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. For the purposes of this implementation, main.go runs a series of predifined locations but can accept any valid location.

## Measurements
###### Latitude
Latitudes are given in degrees, using positive and negative floating point numbers to indicate North (+) or South(-) across the globe. A supported latitude for the program can only exist within the defined boundaries of +90/-90. Each *Latitude* is a unique data structure that allows its relative position to and from other latitudes across the Earth to be deduced. For future iterations, this would allow the program to intelligently modify climate conditions based on percieved proximity to important geological entities.
###### Longtitude
Longtitude are also given in degrees, using positive and negative floating point numbers to indicate East (+) or West(-) across the globe. A supported longtitude for the program can only exist within the defined boundaries of +180/-180. Each *Longtitude* is a unique data structure that allows its relative position to and from other longtitudes across the Earth to be deduced. For future iterations, this would allow the program to intelligently modify climate conditions based on percieved proximity to important geological entities.
###### Elevation
Elevations are calculated in meters and are always non-signed integers. For a valid elevation to exist in the context of this program, a floating point number must be created within the defined boundaries of 8848-0. These measurements are used to indicate sea-floor(0m) or Everest(8848m). Like Longtitude or Latitude pointers, Elevation addresses also have the ability to calculate proximities, allowing for a series of elevation points to be potentally used to normalise areas of varying topography (ie, Hong Kong).
###### Temperature
Temperatures are not clamped to specific ranges due to requirements to potentially throw a number out of bounds for a specific calculation. Temperature pointers, to aid in the calculation of other systems can be cast to all modern forms of temperature units, such as Kelvin, Fahrenheit and are default in Celcius.


## Calculations
###### Hemispheres
Hemispheres are generated from valid Latitude and Longtitude pointers which have been sanitised by the program. These measurements are then checked against the defined geographic boundaries and then are clamped to the approximate range the latitude and longtitude most accurately falls within. For simplicity, subsets and supersets of geographic hemispheres (such as *Tropic of Cancer, Sub-Tropical*, *Tropic of Cancer, Temperate-Arid*) are not factored into clipping to avoid large conditional checks for temperature and climate conditions, relying on approximates to satisfy the simulation output.
###### Seasons
Seasons are selected using quartile bounds, selected at runtime based on the argument chronograph.Time struct and hemipshere. This allows the program to dynamically fetch and generate timestamps for the appropriate seasonal range and climate, while factoring time appropriate conditions such as *leap years*, transitional timestamp boundaries (one year rolling over to the next) and whether or not the season system is inverted for the provided hemisphere. 
###### Temperatures
Temperatures are generated from a pseudo random normalised distribution. For a temperature to be generated, the program selects the appropriate climate conditions based on hemispheric and chronologic data. From these data points, the program generates a variying range of temperates for *n* number of days (that are appropriate to the selected hemisphere and geographic band). To populate *n*, the program uses the in memory chronograph.Time.Day as the lower bounds to calculate the required set of distrubtions needed to populate the number of required days for the season (which has been selected dynamically based on regional conditions). Once these numbers are in place, the program averages the temperatures across the month to produce the anticipated temperature for the current date. 
###### Humidity
Humidity is generated from a random unix seed, with the minimum and maximum bounds being adjusted using the hemispheric bounds maximum and minimum ranges to scale desired change for higher or lower pressures. Bands that exist in closer proximity to the edge of the Tropic of Cancer or Tropic of Capricon, the rise in expected humidity is reflected in the possible range of numbers that can be generated. For simplicity, factors such as carbon dixoide, proximity to jungle-like conditions have not been strict evaluated.
###### Pressure.
Pressure is computed using an expanded barometric formula. This formula takes into account elevations and temperatures, relative to their atmopsheric pressure bands, modifying the potential minimum and maximum bounds used in the final calculation. Based on the argument elevation and argument temperature, pressure systems will fluctuate to produce a pseudo accurate representation of the pressure range for a given set of conditions, leaning towards lower pressure systems for higher altitudes.
