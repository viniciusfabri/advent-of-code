# aoc-skeleton
This repository serves as a skeleton for creating new Advent of Code (AoC) projects. You can use this template to quickly set up a new repository for solving AoC challenges.

## Getting Started

1. **Clone the repository:**
    ```sh
    git clone https://github.com/igorwulff/aoc-skeleton.git
    cd aoc-skeleton
    ```

2. **Change the module name:**
    Update the `go.mod` file with your new module name.
    ```go
    module github.com/yourusername/your-new-repo
    ```

3. **Install dependencies:**
    ```sh
    go mod tidy
    ```

4. **Run the application:**
    ```sh
    go run main.go
    ```

## Project Structure

- `main.go`: Entry point of the application.
- `2024/day1/part1.go`: Solution for Day 1, Part 1.
- `2024/day1/part1_test.go`: Tests for Day 1, Part 1.
- `2024/day1/input.txt`: Input data for Day 1.
- `2024/day1/sample.txt`: Sample input data for Day 1.

## Contributing

Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License.