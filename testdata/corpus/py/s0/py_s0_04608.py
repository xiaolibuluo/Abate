
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2140_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2140_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{90}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2140_2(values):
    for item in values:
        if item > 97:
            yield item * 66


def collect2140_2(values):
    return list(expand2140_2(values))
def project2140_3(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
async def fetch2140_4(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def pair2140_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 21}
def summarize2140_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 28:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2140_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2140_8(values):
    for item in values:
        if item > 42:
            yield item * 10


def collect2140_8(values):
    return list(expand2140_8(values))
class Slot2140_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
