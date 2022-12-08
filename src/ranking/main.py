#!/usr/bin/env python3

import json
import sys

class Handle:
    def args_handling(self):
        if len(sys.argv) != 2:
            print("ERROR: One argument is provided", file=sys.stderr)
            sys.exit(84)


class Parser:
    def parse_file(self):
        with open(sys.argv[1]) as f:
            data = json.load(f)
        print(data)


def main():
    hdl = Handle()
    psr = Parser()
    try:
        hdl.args_handling()
        psr.parse_file()
    except FileNotFoundError:
        print("ERROR: File doesn't exist", file=sys.stderr)
        sys.exit(84)
    except IOError:
        print("ERROR: IO Error", file=sys.stderr)
        sys.exit(84)


if (__name__ == "__main__"):
    main()
