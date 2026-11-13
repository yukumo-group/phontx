TEST_STRING="114514"

def main():
    byteData = str.encode(
        TEST_STRING,
        "utf-8"
    )
    print(bytes.hex(byteData))

main()
