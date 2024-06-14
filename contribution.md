# Contributing to API Authenticator Proxy

We welcome contributions to the API Authenticator Proxy project! Whether you want to report a bug, request a new feature, or submit a code change, please follow the guidelines outlined below.

## Bug Reports and Feature Requests

If you encounter a bug or have a feature request, please open an issue on the [issue tracker](https://github.com/Dhouib-Mohamed/api-authenticator-proxy/issues). Provide as much detail as possible, including steps to reproduce the bug or a clear description of the new feature.

## Code Contributions

### Getting Started

1. Fork the repository and clone it to your local machine:
   ```sh
   git clone https://github.com/yourusername/api-authenticator-proxy.git
   cd api-authenticator-proxy
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

3. Set up the configuration file:
   ```sh
   cp config.example.yaml config.yaml
   ```
   
4. Set up the environment variables:
    ```sh
    cp .env.example .env
    ```


### Making Changes

1. Create a new branch for your feature or bug fix:
   ```sh
   git checkout -b feature-name
   ```

2. Make your changes and ensure they follow the project's coding standards.

3. Run the test suite to ensure your changes didn't break existing functionality:
    ```bash
      bash scripts/test.sh
    ```

4. Commit your changes and push them to your fork:
   ```sh
   git add .
   git commit -m "Your commit message"
   git push origin feature-name
   ```

5. Submit a pull request to the `main` branch of the main repository.

### Code Review Process

Once you've submitted a pull request, the project maintainers will review your code. They may suggest changes or improvements before merging your code into the main branch.

### Code of Conduct

Please adhere to the project's [code of conduct](CODE_OF_CONDUCT.md) in all interactions.

### License

By contributing to the API Authenticator Proxy project, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).

Thank you for your interest in contributing to the API Authenticator Proxy project! Your contributions help improve the project for everyone.