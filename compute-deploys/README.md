# Compute Deploys UI

This project is a single-page React application that allows you to view deployment information by querying the Datadog API.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:
- [Node.js](https://nodejs.org/en/) (which includes npm)
- `kubectl` configured with access to the required Kubernetes clusters.

## Getting Started

Follow these steps to get the application running on a new machine.

### 1. Install Dependencies

Clone the repository and navigate into the project directory. Then, install the required npm packages:

```bash
npm install
```

### 2. Configure Environment Variables

The application requires Datadog API and Application keys to fetch data.

1.  Create a new file named `.env` in the root of the project directory.
2.  Open the `.env` file and add your keys as follows. You will need to get these from your Datadog account. The Application Key requires at least read logs  permissions.

    ```env
    # .env
    DD_API_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    DD_APP_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    ```

### 3. Run the Application

Once the dependencies are installed and the environment variables are set, you can start the application. This command will run both the React frontend and the Node.js backend server concurrently.

```bash
npm run dev
```

The application will be available at [http://localhost:5173](http://localhost:5173). The backend server runs on port 3001.
