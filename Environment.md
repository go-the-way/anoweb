Environment Tables
---

| Name                       | Default     | Description                                                                                            |
|:---------------------------|:------------|:-------------------------------------------------------------------------------------------------------|
| SERVER_MAX_HEADER_SIZE     | 1 << 20     | Controls the maximum number of bytes the request header's keys and values, including the request line. |
| SERVER_READ_TIMEOUT        | time.Minute | ReadTimeout is the maximum duration for reading the entire request, including the body.                |
| SERVER_READ_HEADER_TIMEOUT | time.Minute | ReadHeaderTimeout is the amount of time allowed to read request headers.                               |
| SERVER_WRITE_TIMEOUT       | time.Minute | WriteTimeout is the maximum duration before timing out writes of the response.                         |
| SERVER_IDLE_TIMEOUT        | time.Second | A Duration represents the elapsed time between two instants as an int64 nanosecond count.              |
| CONFIG_FILE                | app.yml     | The YAML Configuration file.                                                                           |
| SERVER_HOST                | 0.0.0.0     | The Server Host.                                                                                       |
| SERVER_PORT                | 9494        | The Server Port .                                                                                      |
| SERVER_TLS_ENABLE          | False       | Enable TLS Support.                                                                                    |
| SERVER_TLS_CERT_FILE       | cart.pem    | TLS Cert File.                                                                                         |
| SERVER_TLS_KEY_FILE        | key.pem     | TLS Key File.                                                                                          |
| BANNER_ENABLE              | True        | Enable print banner.                                                                                   |
| BANNER_TYPE                | default     | Type of banner(Options: default, text, file).                                                          |
| BANNER_TEXT                | FLy GO GO   | Text type of banner.                                                                                   |
| BANNER_FILE                | banner.txt  | File type of banner.                                                                                   |
| TEMPLATE_CACHE             | True        | Enable template cache.                                                                                 |
| TEMPLATE_ROOT              | ./          | The template root path.                                                                                |
| TEMPLATE_SUFFIX            | .html       | The template file suffix.                                                                              |