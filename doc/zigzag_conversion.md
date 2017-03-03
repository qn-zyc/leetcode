# zigzag_conversion

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R

```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string text, int nRows);
```

`convert("PAYPALISHIRING", 3)` should return `"PAHNAPLSIIGYIR"`



结构是这样的：

| 1    |      |      |      | 9    |      |      |      | 17   |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 2    |      |      | 8    | 10   |      |      | 16   | 18   |
| 3    |      | 7    |      | 11   |      | 15   |      | 19   |
| 4    | 6    |      |      | 12   | 14   |      |      | 20   |
| 5    |      |      |      | 13   |      |      |      | 21   |

