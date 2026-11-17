
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2282_0(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
async def fetch2282_1(gate):
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
class Slot2282_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{100}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2282_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 10:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot2282_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2282_5(values):
    for item in values:
        if item > 24:
            yield item * 56


def collect2282_5(values):
    return list(expand2282_5(values))
class Slot2282_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2282_7(gate):
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
async def fetch2282_8(gate):
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
