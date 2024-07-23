# gin-stream-middleware
`gin` 框架的流式中间件

- 对于常规的中间件，只能在接口的函数处理完成以后，在 `c.Next()` 后续中一次性处理。因此，对于流式的返回，这样的中间件可能并不满足你的需求。
- 本库提供了流式的中间件，其实现方式是重写了将返回内容写入 buffer 的 write 函数。这样，对于流式每次写入返回内容时，都可以触发自定义的处理函数，以充当中间件的功能。
- 要使用本库，你需要重写 `CallbackInstance` 结构体的 `Call` 方法，以实现自定义功能。
- 具体使用实例请参见 `example/main.go`。

# gin-stream-middleware
Stream middleware for the `gin` framework

- For regular middleware, it can only be processed once after the interface function is completed, in the subsequent `c.Next()`. Therefore, such middleware may not meet your requirements for stream returns.
- This library provides stream middleware, which is implemented by rewriting the write function that writes the return content into the buffer. In this way, each time the stream writes return content, it can trigger a custom processing function to act as middleware.
- To use this library, you need to rewrite the `Call` method of the `CallbackInstance` struct to implement custom functionality.
- For specific usage examples, please see `example/main.go`.