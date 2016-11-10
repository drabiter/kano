# Kano

stupid simple [Hummingbird](https://hummingbird.me/) CLI client.

## Download

Debian package `wget https://github.com/drabiter/kano/raw/master/build/kano_1.0-1.deb`

Raw executable `wget https://github.com/drabiter/kano/raw/master/build/kano`

## Usage

**> list**

  List current watching series

**> history**
  
  List completed series

**> search <keywords>**               
  
  Search series based on `keywords` - eg `search macross delta`

**> add <hummingbird_anime_id>**
  
  Add series to watching list. This will reset watched episode to 0 - eg `add 9992`

**> bump <id> [count]**

  Increase watched episode count by `count` (optional, default is one) - eg `bump 2`, `bump 2 10`

**> delete <id>**
  
  Remove a series from watching list - eg `remove 1`

**> finish <id>**
  
  Bump episode count to total and mark it as finished - eg `finish 1`

**> rate <id> <rating>**
  
  Rate a series with rating one of 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, or 5 - eg `rate 7 4.5`

**> info**
  
  Show current logged in username and token
