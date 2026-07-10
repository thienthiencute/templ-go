# Templ setup

### Bước 1: Cài đặt Templ CLI
Cài đặt công cụ dòng lệnh `templ` vào máy:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Bước 2: Thêm thư viện vào dự án
Tải thư viện `templ` về thư mục dự án của bạn:
```bash
go get github.com/a-h/templ
```

### Bước 3: Tạo file giao diện
Tạo một thư mục và thêm file `.templ` (ví dụ: `test.templ`).

### Bước 4: Chạy lệnh generate
Để biên dịch file `.templ` ra file `.go` tương ứng, bạn chạy lệnh dưới. 
*(Lưu ý: Không chỉnh sửa và không cần push file `.go` được sinh ra này lên git. Mỗi khi sửa file `.templ` bạn phải chạy lại lệnh này để cập nhật.)*
```bash
templ generate
```

### Bước 5: Sử dụng component trong Go
Đây là ví dụ một cấu trúc đơn giản nhất để render component ra terminal:
```go
package main

import (
	"context"
	"os"
)

func main() {
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
```

### Bước 6: Chạy thử
Cuối cùng, chạy chương trình:
```bash
go run .
```

# Air-verse

Dưới đây là các bước cơ bản để thiết lập `air` chạy tự động `templ generate` và server Go:

### 1. Khởi tạo cấu hình Air
Nếu dự án của bạn chưa có file `.air.toml`, hãy khởi tạo bằng lệnh:
```bash
air init
```

### 2. Cấu hình để Air theo dõi file `.templ`
Mở file `.air.toml` và tìm đến phần `[build]`. Bạn cần thêm đuôi `templ` vào danh sách `include_ext` để Air biết và theo dõi khi có sự thay đổi:

```toml
[build]
  # Thêm "templ" vào danh sách này:
  include_ext = ["go", "tpl", "tmpl", "html", "templ"]
```

### 3. Cập nhật câu lệnh Build
Mặc định `air` chỉ chạy lệnh `go build`. Bạn cần đổi lệnh này để nó chạy cả `templ generate` trước khi build Go.

Trong `.air.toml`, cập nhật dòng `cmd`:

**Cho Windows:** (Sửa trong phần `[build.windows]`)
```toml
[build.windows]
  cmd = "cmd /c \"templ generate && go build -o ./tmp/main.exe .\""
```

**Cho Mac/Linux:** (Sửa trong phần `[build]`)
```toml
[build]
  cmd = "templ generate && go build -o ./tmp/main ."
```

### 4. Chạy dự án
Bây giờ mọi thứ đã xong, bạn chỉ cần gõ lệnh:
```bash
air
```

Mỗi khi bạn lưu file `.templ` hoặc file `.go`, `air` sẽ tự động chạy `templ generate`, build lại mã nguồn và khởi động lại server.
