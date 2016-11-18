# Kano

stupid simple [Hummingbird](https://hummingbird.me/) CLI client.

[![asciicast](https://asciinema.org/a/2x7aoxgyiqp021cpjzr6agswx.png)](https://asciinema.org/a/2x7aoxgyiqp021cpjzr6agswx)

## Download

[Debian package](https://github.com/drabiter/kano/raw/master/build/kano-latest.deb)

[Raw executable](https://github.com/drabiter/kano/raw/master/build/kano)

## Usage

```
$ kano
+----+---------------------------------+------+----------+-------+-------------+------------------+
| ID | Title                           | Type | Progress | Total | Your Rating | Community Rating |
+----+---------------------------------+------+----------+-------+-------------+------------------+
| 0  | Brave Witches                   | TV   | 2        | 12    |             | 3.58             |
| 1  | Love Live! Sunshine!!           | TV   | 12       | 13    |             | 3.79             |
+----+---------------------------------+------+----------+-------+-------------+------------------+

> search girlish
+----+-------------------+------+-------+------------------+--------+
| ID | Title             | Type | Total | Status           | Rating |
+----+-------------------+------+-------+------------------+--------+
| 0  | Gi(a)rlish Number | TV   | 12    | Currently Airing | 3.87   |
+----+-------------------+------+-------+------------------+--------+

> add 0
+----+---------------------------------+------+----------+-------+-------------+------------------+
| ID | Title                           | Type | Progress | Total | Your Rating | Community Rating |
+----+---------------------------------+------+----------+-------+-------------+------------------+
| 0  | Gi(a)rlish Number               | TV   | 0        | 12    |             | 3.87             |
| 1  | Brave Witches                   | TV   | 2        | 12    |             | 3.58             |
| 2  | Love Live! Sunshine!!           | TV   | 12       | 13    |             | 3.79             |
+----+---------------------------------+------+----------+-------+-------------+------------------+

> bump 0 2
+----+---------------------------------+------+----------+-------+-------------+------------------+
| ID | Title                           | Type | Progress | Total | Your Rating | Community Rating |
+----+---------------------------------+------+----------+-------+-------------+------------------+
| 0  | Gi(a)rlish Number               | TV   | 2        | 12    |             | 3.87             |
| 1  | Brave Witches                   | TV   | 2        | 12    |             | 3.58             |
| 2  | Love Live! Sunshine!!           | TV   | 12       | 13    |             | 3.79             |
+----+---------------------------------+------+----------+-------+-------------+------------------+

> bump 1
+----+---------------------------------+------+----------+-------+-------------+------------------+
| ID | Title                           | Type | Progress | Total | Your Rating | Community Rating |
+----+---------------------------------+------+----------+-------+-------------+------------------+
| 0  | Gi(a)rlish Number               | TV   | 2        | 12    |             | 3.87             |
| 1  | Brave Witches                   | TV   | 3        | 12    |             | 3.58             |
| 2  | Love Live! Sunshine!!           | TV   | 12       | 13    |             | 3.79             |
+----+---------------------------------+------+----------+-------+-------------+------------------+

> record
// ... bunch of finished series ...

> list
+----+---------------------------------+------+----------+-------+-------------+------------------+
| ID | Title                           | Type | Progress | Total | Your Rating | Community Rating |
+----+---------------------------------+------+----------+-------+-------------+------------------+
| 0  | Gi(a)rlish Number               | TV   | 2        | 12    |             | 3.87             |
| 1  | Brave Witches                   | TV   | 3        | 12    |             | 3.58             |
| 2  | Love Live! Sunshine!!           | TV   | 12       | 13    |             | 3.79             |
+----+---------------------------------+------+----------+-------+-------------+------------------+

> quit
```

## Changelog

**1.2**

- [x] fix deleting title from watching list

**1.1**

- [x] show user and community rating on watching list
- [x] use table number to add title to watching list

**1.0**

- [x] disaster

**TODO**

- [ ] give history table some color - it's hard to read
- [ ] one letter command alias - 's' for 'search'
