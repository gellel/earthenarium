# earthenarium

## Usage
`go get github.com/gellel/earthenarium`

### About 
Earthernarium is a small weather program intended to generate generalised climate seeds based on proximities to notable geographic ranges across a pseudo-earth planet. It uses a set of defined boundaries (Arctic Circle, Tropic of Cancer, Equator, Tropic of Capricorn, Antarctic Circle) to construct a normal distribution of random numbers to simulate the number of variable conditions for a seasonal range. Climates gradually shift from colder to warmer as the longtitudal reference slides closer towards the globes center band (minimum bounds Tropic of Cancer and maximum bounds Tropic of Capricon). 

#### Output
Running the program produces a tablised set of data, containing the following

| Location | Position | Local Time | Condition | Temperature | Pressure | Humidity | Season | Region | Hemisphere |
|----------|----------------------|------------------|-----------|-------------|------------|----------|--------|------------------|
| El Aauin | 27.1418,-13.18797,68 | 1903-3-24 0:10:0 | Sunny | +11.0 | 1005.10767 | 58.1 | Spring | Tropic of Cancer | Northern |
