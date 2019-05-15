# earthenarium

## Binary
`go get github.com/gellel/earthenarium`
## Source
`git clone github.com/gellel/earthenarium`

### example 
`$ earthenarium`

### About 
Earthernarium is a small weather program intended to generate generalised climate seeds based on proximities to notable geographic ranges across a pseudo-earth planet. It uses a set of defined boundaries (Arctic Circle, Tropic of Cancer, Equator, Tropic of Capricorn, Antarctic Circle) to construct a normal distribution of random numbers to simulate the number of variable conditions for a seasonal range. Climates gradually shift from colder to warmer as the longtitudal reference slides closer towards the globes center band (minimum bounds Tropic of Cancer and maximum bounds Tropic of Capricon). 

#### Output
Running the program produces a tablised set of data, containing the following

| Location | Position | Local Time | Condition | Temperature | Pressure | Humidity | Season | Region | Hemisphere |
|----------|----------|------------|-----------|-------------|----------|----------|--------|--------|------------|
| El Aauin | 27.1418,-13.18797,68 | 1903-3-24 0:10:0 | Sunny | +11.0 | 1005.10767 | 58.1 | Spring | Tropic of Cancer | Northern |

#### Table contents
Location is city where the measurement is conducted.
Position is a comma-separated triple containing latitude, longitude, and elevation in metres above sea level,
Local time is an ISO8601 date time,
Conditions is either Snow, Rain, Sunny.
Temperature is in °C.
Pressure is in hPa
Relative humidity is a %.
Season is the sampling season collected from the generated timestamp.
Region is the temperate band of the planet.
Hemisphere is the positional reference.

#### Calculations
Temperatures are generate using a pseudo randomised normal distribution, taking the defined minimum and maximum bounds for a hemisphere and season. Each temperate range is build upon the span of days between the season starting time and season ending time, averaging out the temperature range across the month.
Humidity is generated from a random unix seed, with the minimum and maximum bounds being adjusted using the hemispheric bounds maximum and minimum ranges to scale desired change for higher or lower pressures.
Pressure is computed using an expanded barometric formula, modifying the potential minimum and maximum bounds based on the argument elevation and argument temperature to construct a pseudo accurate representation of the pressure range for a given condition.
