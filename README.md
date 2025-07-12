# Kafka Bridge Service (Go)

This is a high-performance Go service that consumes messages from a Kafka topic and produces them into another topic. It's designed with modular architecture, supports concurrent processing via goroutines, and includes full unit test coverage.

---

## 🚀 Features

- ✅ Consumes messages from a Kafka topic using [Confluent Kafka Go Client](https://github.com/confluentinc/confluent-kafka-go)
- ✅ Produces messages to a different Kafka topic
- ✅ Configurable number of goroutines (workers) for concurrent processing
- ✅ Graceful shutdown using `context.Context` and signal handling
- ✅ Environment-based configuration with fallback to `.env.local`
- ✅ Structured logging with [Uber Zap](https://github.com/uber-go/zap)
- ✅ Fully mockable architecture for unit testing
- ✅ Unit tests covering success and failure scenarios

---

## 🧱 Project Structure

```

kafka-bridge/
├── cmd/                # Application entrypoint
├── config/             # Configuration loader
├── consumer/           # Kafka consumer logic and interfaces
├── producer/           # Kafka producer logic and interfaces
├── service/            # Core processing service (consumes and produces)
├── mocks/              # Mocks for unit testing
├── tests/              # Unit tests
├── .env.local          # Local environment variables
├── go.mod / go.sum     # Module dependencies

````

---

## ⚙️ Configuration

The service uses environment variables for configuration. In local development, a `.env.local` file can be used:

### Example `.env.local`

```env
KAFKA_BROKERS=localhost:9092
KAFKA_GROUP_ID=my-group
KAFKA_INPUT_TOPIC=input-topic
KAFKA_OUTPUT_TOPIC=output-topic
NUM_WORKERS=4
````

---

## 🧪 Running Tests

```bash
go test ./...
```

---

## 🏁 Running the Application

```bash
go run cmd/main.go
```

Make sure your Kafka brokers are up and accessible based on the environment config.

---

## 📦 Dependencies

* [Confluent Kafka Go Client](https://github.com/confluentinc/confluent-kafka-go)
* [Zap Logger](https://github.com/uber-go/zap)
* [Godotenv](https://github.com/joho/godotenv)
* [Testify](https://github.com/stretchr/testify)

---

## 📌 License

MIT © Lucas Queiroz
