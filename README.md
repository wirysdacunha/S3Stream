# S3Stream

🚀 **S3Stream** is a lightweight microservice written in Go, designed for uploading files to Amazon S3 upon receiving messages from RabbitMQ. This project leverages modern cloud technologies to provide a fast, scalable, and secure file upload solution.

---

## 🌟 Features

- **Efficient File Uploads**: Supports large file uploads to Amazon S3.
- **RabbitMQ Integration**: Seamlessly consumes messages containing file data.
- **AWS SDK for Go**: Utilizes the latest AWS SDK for high-performance operations.
- **Environment Configuration**: Easy configuration through environment variables.
- **Scalable Design**: Built to integrate into distributed systems.

---

## 📋 Prerequisites

1. **Go**: Version 1.20 or higher.
2. **AWS Account**: With access to an S3 bucket.
3. **RabbitMQ**: A running RabbitMQ server.

---

## 🛠️ Installation

### Clone the Repository
```bash
git clone https://github.com/wirysdacunha/S3Stream.git
cd S3Stream
```

### Initialize the Project
```bash
go mod tidy
```

### Set Up Environment Variables
Create a `.env` file in the project root:

```env
S3_BUCKET=your-s3-bucket-name
RABBITMQ_URL=amqp://user:password@host:5672/
```

Alternatively, export variables directly:

```bash
export S3_BUCKET=your-s3-bucket-name
export RABBITMQ_URL=amqp://user:password@host:5672/
```

---

## 🚀 Usage

### Start the Service
Run the application:

```bash
go run main.go
```

The service will connect to RabbitMQ, listen for messages, and upload files to the specified S3 bucket.

---

## 📂 Project Structure

```plaintext
S3Stream/
├── main.go         # Entry point of the application
├── config/         # Configuration logic
├── services/       # Core service logic
└── utils/          # Utility functions
```

---

## 🧪 Testing

### Send a Test Message to RabbitMQ
Use a RabbitMQ client or CLI to send a message:

```json
{
  "fileName": "example.txt",
  "fileData": "SGVsbG8gd29ybGQ="
}
```

Ensure the message contains a Base64-encoded file in the `fileData` field.

---

## 📖 Documentation

- **AWS SDK for Go**: [Documentation](https://aws.github.io/aws-sdk-go-v2/)
- **RabbitMQ**: [Documentation](https://www.rabbitmq.com/documentation.html)

---

## 🤝 Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/my-feature`).
5. Open a Pull Request.

---

## 🛡️ License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## 🌐 Connect

For questions or suggestions, feel free to reach out:
- **Email**: w.cunha@bitsti.com.br
- **GitHub**: [wirysdacunha](https://github.com/wirysdacunha)

---

Happy coding! 🎉
