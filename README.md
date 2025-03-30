# R2PO (Revealed Resource & Performance Overview)

## Purpose

The **R2PO** tool is designed to benchmark and analyze the performance of API calls made to Azure resources across multiple programming languages. The goal is to help developers determine which language is most efficient for interacting with Azure resources based on key performance metrics such as execution time, CPU and memory usage, and disk space consumption.

This tool performs a series of tests across the following languages:

- **PowerShell**
- **Python**
- **.NET**
- **Go**
- **Rust**

By collecting and comparing these performance metrics, **R2PO** provides developers with insights into which programming language is best suited for their Azure-related tasks.

## Purpose of This Repository

1. **Performance Benchmarking**: Test and compare the speed, resource consumption, and overall efficiency of Azure-related API calls across different languages.
2. **Security Assessment**: Ensure that the implementation is secure and adheres to best practices when interacting with Azure services.
3. **Cross-Language Comparison**: Provide developers with concrete data that can help decide which language is most appropriate for their specific Azure task or use case.
4. **Multi-Language Proficiency**: Improve development skills by refactoring and coding the same task across various languages.

## Current Implementation

At present, the repository includes an implementation in **PowerShell**, with additional languages like **.NET**, **Go**, **Python**, and **Rust** planned for future additions.

### PowerShell Example

The current implementation uses **PowerShell** with the `Az` module to authenticate and retrieve Azure subscriptions. It measures the time taken for authentication and API requests to provide a benchmark of the process.

## Planned Features

- Extend the current test suite to include additional Azure-related tasks, such as retrieving resource groups and managing resources.
- Expand the tool to support **.NET**, **Go**, **Python**, and **Rust**, and implement the same benchmarked test.
- Include CPU, memory, and disk space usage during benchmarking.
- Add more specific tests for Azure Identity, resource group management, and more.

## Usage

Once the necessary dependencies for each language are installed, you can run the tests for each language separately. Each test measures the time taken to authenticate and retrieve Azure subscriptions.

### Example usage for PowerShell:

```powershell
# Run the Measure-R2PO function
Measure-R2PO
