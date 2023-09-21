# Contact Form Serverless Function

This project is designed to handle contact form submissions. Upon receiving a submission, it performs the following tasks:

1. Stores the contact form data into a database.
2. Sends an email notification to the team.
3. Sends a confirmation email to the user.

## Prerequisites

- **Go SDK Version**: `1.20`
- **Set Environment variables**: Copy sample.env to .env and update the values 
- **Platform**: DigitalOcean Serverless Functions. For more details, refer to the [official documentation](https://docs.digitalocean.com/reference/doctl/reference/serverless/).

## Configuration

The project uses a `project.yml` file for configuration. This file defines the structure and settings for the serverless function. For more information about the structure and details of `project.yml`, see the [official guide](https://docs.digitalocean.com/products/functions/reference/project-configuration/).

## Deployment Steps

1. **Connect to the Functions Namespace**: Before deploying, you need to connect to the appropriate functions namespace. Use the following command:
   ```bash
   doctl serverless connect <name-space>
   ```

2. **Check Connection Status**: To verify which namespace you're connected to, run:
   ```bash
   doctl serverless status
   ```

3. **Deploy the Project**: To deploy the serverless function, navigate to the project's root directory and run:
   ```bash
   doctl serverless deploy .
   ```
   Ensure that the directory in the command above points to the root of the project.

---

