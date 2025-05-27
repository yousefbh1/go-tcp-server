Phase 1: Single-Client Echo Server (Days 1–2)

    Skeleton server

        net.Listen("tcp", ":4000")

        listener.Accept() loop

    Echo logic

        Read from conn.Read() into a buffer

        Immediately write back via conn.Write()

    Manual test

        telnet localhost 4000 → type; verify you see your own input


Phase 2: Simple Client (Day 3)

    Dial in Go

        net.Dial("tcp", "localhost:4000")

    Std-in to server

        Read lines from bufio.NewReader(os.Stdin)

        Send each line to the server

    Print replies

        Spawn a goroutine that reads from the socket and prints to stdout


Phase 3: Concurrency & Connection Handling (Days 4–5)

    Spawn per-client goroutine

        On each Accept(), go handle(conn)

    Graceful shutdown

        Catch os.Signal (SIGINT) to call listener.Close()

    Limit max clients (stretch)

        Use a buffered channel semaphore or sync.WaitGroup


Phase 4: Message Framing (Days 6–7)

    Length-prefix protocol

        Client: send [4-byte length][payload]

        Server: io.ReadFull() the length, then payload

    Binary safety

        Support arbitrary bytes (no “line” boundaries)

    Error cases

        Handle partial reads, unexpected EOFs


Phase 5: Robustness & Logging (Day 8)

    Integrate a simple logger (standard log or zap)

    Emit connection open/close, bytes sent/received, errors

    Recover from panics inside handler

Phase 6: Testing & Benchmarking (Days 9–10)

    Unit tests

        Mock net.Conn with net.Pipe() to verify echo behavior

    Benchmark

        Use testing.B to flood-write/read in a loop

        Profile with go test -bench . -cpuprofile=cpu.prof + go tool pprof


Phase 7: Documentation & Packaging (Day 11)

    Write a clear README.md with usage examples:

    go run ./cmd/echoserver  --port 4000
    go run ./cmd/echoclient --addr localhost:4000

    Add Go module tags (go:generate, versioning)

Phase 8: Stretch Goals (Days 12–14)

    TLS support with crypto/tls (self-signed certs)

    CLI flags for timeouts, max-conns, log verbosity

    Metrics endpoint: embed a tiny HTTP server exposing prometheus‐style counters

    Dockerize server/client with a simple Dockerfile
