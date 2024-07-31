# 创建mkfifo的管道, 创建完会在文件系统中新建一个文件操作符
mkfifo myPipe

# 管道写入内容, 会阻塞!
cat "hello" > myPipe

# 管道读取内容
cat < myPipe