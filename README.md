# 🚀 GoSortStack

A modern and blazingly fast **stack sorting algorithm** written in Go, featuring a stunning web-based visualizer

![Visualization](https://raw.githubusercontent.com/kwa0x2/GoSortStack/master/public/visualization.gif)

## 📋 Table of Contents

1. [What is GoSortStack?](#what-is-gosortstack)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Algorithm](#algorithm)
5. [Visualizer](#visualizer)
6. [Performance](#performance)
7. [Project Structure](#project-structure)

---

## 🎯 What is GoSortStack?

GoSortStack is a sorting algorithm that sorts a random list of integers using the **minimum number of moves** possible. It operates with two stacks (A and B) and a limited set of operations.

### 🔧 Available Operations

| Operation | Description |
|-----------|-------------|
| `SwapA` | Swap the first 2 elements at the top of stack A |
| `SwapB` | Swap the first 2 elements at the top of stack B |
| `SwapBoth` | `SwapA` and `SwapB` at the same time |
| `PushA` | Take the first element at the top of B and put it at the top of A |
| `PushB` | Take the first element at the top of A and put it at the top of B |
| `RotateA` | Shift up all elements of stack A by 1 |
| `RotateB` | Shift up all elements of stack B by 1 |
| `RotateBoth` | `RotateA` and `RotateB` at the same time |
| `ReverseRotateA` | Shift down all elements of stack A by 1 |
| `ReverseRotateB` | Shift down all elements of stack B by 1 |
| `ReverseRotateBoth` | `ReverseRotateA` and `ReverseRotateB` at the same time |

**Goal:** Stack B must be empty, and all integers must be in stack A, sorted in ascending order.

---

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/kwa0x2/GoSortStack.git
cd GoSortStack

# Build the project
go build -o sort_stack.exe cmd/main.go
```

---

## 🎮 Usage

### Normal Mode (Print operations to terminal)
```bash
./sort_stack.exe 5 2 8 1 9 3 7
```

### Visualization Mode (Watch in browser with stunning graphics)
```bash
./sort_stack.exe -v 5 2 8 1 9 3 7
```

or

```bash
./sort_stack.exe --visualize 5 2 8 1 9 3 7
```

---

## 🧠 Algorithm

This implementation uses the **Turk Algorithm**, a cost-based optimal sorting strategy that minimizes the number of operations.

### Algorithm Steps:

#### 1. **Initial Push**
Push the first 2 elements from Stack A to Stack B
- This creates a minimum and maximum value in Stack B
- Provides reference points for subsequent elements

![Initial Push](https://raw.githubusercontent.com/kwa0x2/GoSortStack/master/public/first_two.png)

#### 2. **Cost Calculation**
For each element in Stack A, calculate the minimum number of operations required to place it in the correct position in Stack B
- Compare 4 different rotation strategies:
  - `RotateA + RotateB` (Rotate both stacks upward)
  - `ReverseRotateA + ReverseRotateB` (Rotate both stacks downward)
  - `RotateA + ReverseRotateB` (A up, B down)
  - `ReverseRotateA + RotateB` (A down, B up)

#### 3. **Optimal Push**
Push the element with the lowest cost to Stack B
- This process continues until only 3 elements remain in Stack A

#### 4. **Sort Last Three**
Quickly sort the remaining 3 elements in Stack A

#### 5. **Push Back**
Push all elements from Stack B back to Stack A in optimal positions
- Each element is placed in its correct sorted position

#### 6. **Final Rotation**
Rotate Stack A to bring the smallest element to the top

![Final](https://raw.githubusercontent.com/kwa0x2/GoSortStack/master/public/final.png)

### Why This Algorithm?

- ✅ **Efficient**: Minimizes the number of operations
- ✅ **Optimal**: Cost-based approach selects the best rotation strategy
- ✅ **Scalable**: Works with 100, 500, even 1000+ elements
- ✅ **Fast**: Leverages Go's performance for lightning-fast execution

### Learn More

For a detailed explanation of the algorithm, check out this excellent article:
- [Push Swap — A journey to find the most efficient sorting algorithm](https://medium.com/@ayogun/push-swap-c1f5d2d41e97) by Ali Ogun
---

## 🎨 Visualizer

### Features

- ⚡ **Speed Control**: Adjust animation speed from 1ms to 2000ms
- 🎮 **Keyboard Shortcuts**:
  - `←` `→` Previous/Next step
  - `Space` Play/Pause
  - `Home` First step
  - `End` Last step
- 🎯 **Color-Coded Bars**: Blue → Green → Yellow → Red (small to large values)
- 📊 **Real-time Stats**: Current step and operation display

### How It Works

1. Run the program with the `-v` flag to record each operation
2. All steps are serialized to JSON and embedded in HTML
3. An HTML file is automatically generated and opened in your browser
4. Enjoy the Matrix rain effects in the background 🔥

---

## 📊 Performance

| Element Count | Average Moves | Max Moves |
|---------------|---------------|-----------|
| 3             | 2-3           | 3         |
| 5             | 8-12          | 12        |
| 100           | ~550          | <700      |
| 500           | ~3500         | <5500     |

---

## 🏗️ Project Structure

```
GoSortStack/
├── cmd/
│   └── main.go              # Main program entry point
├── internal/
│   ├── algorithm/           # Sorting algorithms
│   │   ├── sort.go         # Main sort function
│   │   ├── sort_utils.go   # Cost calculation utilities
│   │   ├── sort_three.go   # 3-element sort
│   │   └── sort_five.go    # 5-element sort
│   ├── checker/            # Sorted state checker
│   ├── operations/         # Stack operations
│   ├── parser/             # Input parsing
│   ├── stack/              # Stack data structure
│   └── visualizer/         # Visualization engine
├── visualizer.html         # Web visualizer template
└── README.md
```

### Call Graph

Visual representation of function calls and dependencies:

![Algorithm Call Graph](https://raw.githubusercontent.com/kwa0x2/GoSortStack/master/public/algorithm-graph.svg)

---

## 🔥 Cool Features

### 1. Silent Mode
Operations aren't printed to terminal when visualizer is active:
```go
if !o.Silent {
    fmt.Println("SwapA")
}
```

### 2. Embedded JSON
No CORS issues - JSON is directly embedded in HTML:
```javascript
const EMBEDDED_DATA = {...};
```

### 3. Matrix Rain Effect
Continuously flowing 0s and 1s in the background ⚡

---

## 🚀 Advanced Usage

### Large Tests
```bash
# 100 random numbers
./push_swap.exe -v $(shuf -i 1-100 -n 100 | tr '\n' ' ')

# 500 random numbers
./push_swap.exe $(shuf -i 1-500 -n 500 | tr '\n' ' ')
```

### Count Operations
```bash
./push_swap.exe 5 2 8 1 9 3 7 | wc -l
```

---

## 🛠️ Development

### Run Tests
```bash
go test ./...
```

### Build
```bash
go build -o push_swap.exe cmd/main.go
```

### Update Template
Edit `visualizer.html` - the program automatically uses it.

---

## 📝 Notes

- ⚠️ Duplicate numbers are not accepted
- ⚠️ Only integer values allowed
- ⚠️ Empty input returns an error
- ✅ Negative numbers are supported
- ✅ Maximum value: 2147483647 (INT_MAX)
- ✅ Minimum value: -2147483648 (INT_MIN)

---

## 🎓 What I Learned

- Implementing linked lists in Go
- Cost-based optimization algorithms
- Web-based visualization techniques
- Integrating Go with HTML/CSS/JavaScript
- Efficient sorting strategies

---

## 🙏 Credits

This project is a modern Go implementation of the classic C Push Swap project. The algorithm is based on the **Turk Algorithm**.

Special thanks to:
- [Ali Ogun](https://github.com/ayogun) for the excellent algorithm explanation and reference implementation
- [Medium Article](https://medium.com/@ayogun/push-swap-c1f5d2d41e97) - Comprehensive algorithm breakdown

---

## 📄 License

This project is open source and available under the MIT License.

---

**Made with ❤️ and lots of ☕**

*Sort with style - Matrix rain edition!* 🚀
