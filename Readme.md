# AI-Driven Golang Project Generator

This project is a Golang-based tool that utilizes OpenAI's GPT-3.5-turbo model to generate code and folder structures for a given topic. The project automates the process of creating a new Golang project by taking user input, generating the corresponding code and folder structure, and setting up the project directory.

## Features

- **Dynamic Code Generation:** Automatically generates Golang code based on the provided topic.
- **Folder Structure Creation:** Builds a project folder structure tailored to the given topic.
- **Integration with OpenAI API:** Uses the GPT-3.5-turbo model for generating code and folder structures.
- **Environment Variable Management:** Loads the OpenAI API key from a `.env` file.

## Prerequisites

- Golang installed on your system.
- A valid OpenAI API key.
- A `.env` file in the project root directory with the following content:

  ```plaintext
  API_KEY=your_openai_api_key

## How It Works
- nput Topic: The user provides a topic for which the project needs to be generated.
- API Request: The topic is sent to OpenAI’s GPT-3.5-turbo model via an API request.
- Response Handling: The model returns generated code and folder structure as a response.
- File Writing: The code is saved in code/code.go, and the folder structure is saved in code/output.txt.
- Folder Structure Creation: The folder structure is created in the project/ directory.

## Usage

1. **Clone the Repository:**

    ```sh
    git clone https://github.com/yourusername/ai-golang-generator.git
    cd ai-golang-generator
    ```

2. **Set Up Environment Variables:**

    Create a `.env` file in the root directory:

    ```plaintext
    API_KEY=your_openai_api_key
    ```

3. **Run the Application:**

    Execute the following command to start the application:

    ```sh
    go run main.go
    ```

4. **Enter a Topic:**

    When prompted, enter the topic for which you want to generate the project:

    ```plaintext
    Enter a topic: ecommerce platform
    ```

5. **View Generated Files:**

    - The generated code will be in `code/code.go`.
    - The folder structure will be detailed in `code/output.txt`.
    - The folder structure will created under the `project/ directory`.
