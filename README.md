# Kano

stupid simple [Hummingbird](https://hummingbird.me/) CLI client.

[![asciicast](https://asciinema.org/a/2x7aoxgyiqp021cpjzr6agswx.png)](https://asciinema.org/a/2x7aoxgyiqp021cpjzr6agswx)

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

**> add [id]**
  
  Add series to watching list. This will reset watched episode to 0 - eg `add 9992`

**> bump [id] &lt;count&gt;**

  Increase watched episode count by `count` (optional, default is one) - eg `bump 2`, `bump 2 10`

**> delete [id]**
  
  Remove a series from watching list - eg `remove 1`

**> finish [id]**
  
  Bump episode count to total and mark it as finished - eg `finish 1`

**> rate [id] &lt;rating&gt;**
  
  Rate a series with rating one of 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, or 5 - eg `rate 7 4.5`

**> info**
  
  Show current logged in username and token

**> quit**
**> exit**

  Quit and get a life ;)

## Changelog

**1.1**

- [x] show user and community rating on watching list
- [x] use table number to add title to watching list

**1.0**

- [x] disaster

**TODO**

- [ ] give history table some color - it's hard to read
- [ ] one letter command alias - 's' for 'search'
