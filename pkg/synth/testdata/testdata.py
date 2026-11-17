import struct

MEMORY_NAME="testdata"

def main() -> None:
    param = struct.pack(
        "<10i",
        10, # aq version
        1, # Preset
        60, # Speed
        0, # phont_data
        0, # aq10base
        100, # volume
        70, # pitch, accent...
        30,
        10,
        100,
    )
    print(bytes.hex(param))

main()
