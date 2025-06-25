# Big O Notation Playground

## 📚 Introduction

This playground is my personal space to explore and understand Big O notation. Here, I experiment with how Big O works and learn to apply it in real-world programming scenarios.

## 🧮 What is Big O Notation?

Big O notation characterizes functions according to their growth rates. Different algorithms may be more efficient in different situations. Understanding Big O helps you choose the right algorithm for your specific needs.

| Notation   | Name         | Description                                            |
| ---------- | ------------ | ------------------------------------------------------ |
| O(1)       | Constant     | Execution time stays the same regardless of input size |
| O(log n)   | Logarithmic  | Execution time grows logarithmically with input size   |
| O(n)       | Linear       | Execution time grows linearly with input size          |
| O(n log n) | Linearithmic | Combination of linear and logarithmic growth           |
| O(n²)      | Quadratic    | Execution time grows quadratically with input size     |

## 🗂️ Project Structure

```
├── golang/
│   ├── main.go                     # Main entry point for Go examples
│   ├── 1-o1/                       # O(1) - Constant time
│   │   └── 1-O1_constant.go
│   ├── 2-ologn/                    # O(log n) - Logarithmic time
│   │   └── 2-OlogN_binary_search.go
│   ├── 3-on/                       # O(n) - Linear time
│   │   └── 3-On_linear_scan.go
│   ├── 4-onlogn/                   # O(n log n) - Linearithmic time
│   │   └── 4-OnLogN_merge_sort.go
│   └── 5-on2/                      # O(n²) - Quadratic time
│       └── 5-On2_naive_pairs.go
├── javascript/
│   ├── 1-O1_constant.js            # O(1) - Constant time example
│   ├── 2-OlogN_binary_search.js    # O(log n) - Logarithmic time example
│   ├── 3-On_linear_scan.js         # O(n) - Linear time example
│   ├── 4-OnLogN_merge_sort.js      # O(n log n) - Linearithmic time example
│   └── 5-On2_naive_pairs.js        # O(n²) - Quadratic time example
├── php/
│   ├── 1-O1_constant.php           # O(1) - Constant time example
│   ├── 2-OlogN_binary_search.php   # O(log n) - Logarithmic time example
│   ├── 3-On_linear_scan.php        # O(n) - Linear time example
│   ├── 4-OnLogN_merge_sort.php     # O(n log n) - Linearithmic time example
│   └── 5-On2_naive_pairs.php       # O(n²) - Quadratic time example
├── info/
│   ├── big-o-notation-graph.png
│   ├── big-o-notation-table_data_structure.png
│   └── big-o-notation-array_sorting.png
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Node.js (version 20 or higher) for JavaScript examples
- PHP (version 8.1 or higher) for PHP examples
- Go (version 1.20 or higher) for Go examples

## 🔍 How to Use the Playground

Run each example to observe how different algorithms perform with the same input size:

```bash
# Javascript
node javascript/1-O1_constant.js
node javascript/2-OlogN_binary_search.js
node javascript/3-On_linear_scan.js
node javascript/4-OnLogN_merge_sort.js
node javascript/5-On2_naive_pairs.js

# Golang
GO111MODULE=off go run golang/main.go o1
GO111MODULE=off go run golang/main.go ologn
GO111MODULE=off go run golang/main.go on
GO111MODULE=off go run golang/main.go onlogn
GO111MODULE=off go run golang/main.go on2
GO111MODULE=off go run golang/main.go all

# PHP ( Coming Soon )
php php/1-O1_constant.php
php php/2-OlogN_binary_search.php
php php/3-On_linear_scan.php
php php/4-OnLogN_merge_sort.php
php php/5-On2_naive_pairs.php

```

### Experiment with Different Input Sizes

To truly understand Big O notation, try modifying the N value in each script (make it 10× bigger or smaller) and observe how the execution time changes:

- O(1) : Execution time barely changes regardless of input size
- O(log n) : Execution time grows very slowly (doubling N only adds one more iteration)
- O(n) : Execution time doubles when input size doubles
- O(n log n) : Execution time grows slightly faster than linear (watch what happens when you compare 100k vs 200k elements)
- O(n²) : Execution time grows dramatically (doubling N makes it approximately 4× slower)

## 🧠 Why Big O Matters

Understanding Big O notation helps you:

1. Predict Performance : Anticipate how your code will perform with large inputs
2. Compare Algorithms : Choose the most efficient solution for your specific problem
3. Optimize Code : Identify and eliminate inefficient patterns in your code
4. Technical Interviews : Discuss algorithm efficiency in a standardized way

## 📖 Reference & Further Learning

- Big O Cheat Sheet [ENG] : [Big O Notation - Know Thy Complexities!](https://www.bigocheatsheet.com/)
- Freecodecamp [ENG] : [Big O Notation: Why It Matters and Why It Doesn't](https://www.freecodecamp.org/news/big-o-notation-why-it-matters-and-why-it-doesnt-1674cfa8a23c/)
- Backend Magang [IND] : [Big O Notation - Memahami Kompleksitas Algoritma dalam Pemrograman](https://www.youtube.com/watch?v=XgKfcZctwA8)

## 👨‍💻 Author

Created with 💻 by Ignata

- 📂 GitHub: [Aldo Ignata Chandra](https://github.com/aldoignatachandra)
- 💼 LinkedIn: [Aldo Ignata Chandra](https://linkedin.com/in/aldoignatachandra)
