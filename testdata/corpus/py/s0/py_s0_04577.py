
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2125_0(gate):
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
def expand2125_1(values):
    for item in values:
        if item > 89:
            yield item * 11


def collect2125_1(values):
    return list(expand2125_1(values))
def expand2125_2(values):
    for item in values:
        if item > 96:
            yield item * 22


def collect2125_2(values):
    return list(expand2125_2(values))
def traced2125_3(value: int) -> int:
    return value + 6


def decorate2125_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2125_3
def doubled2125_3(value: int) -> int:
    return value + value
class Slot2125_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2125_5(gate):
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
def expand2125_6(values):
    for item in values:
        if item > 27:
            yield item * 66


def collect2125_6(values):
    return list(expand2125_6(values))
