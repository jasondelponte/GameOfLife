GameOfLife
==========

Playing around with Conways Game of Life

# Usage
```shell
./cgol -w <width> -h <height> -c <chance of life>
```

The implementation is very simple right now it builds two 2 dimentional arrays with their edges wrapped around. I used two buffers so i can just swap between them inbetween steps.