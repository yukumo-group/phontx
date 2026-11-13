import struct

MEMORY_NAME="testdata"

def main() -> None:
    param = struct.pack(
        "<10i",
        10,
        1,
        100,
        0,
        0,
        100,
        100,
        100,
        100,
        100,
    )
    print(bytes.hex(param))

main()
