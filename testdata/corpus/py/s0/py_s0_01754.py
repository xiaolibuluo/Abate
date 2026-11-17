
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand803_0(values):
    for item in values:
        if item > 65:
            yield item * 18


def collect803_0(values):
    return list(expand803_0(values))
async def fetch803_1(gate):
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
def expand803_2(values):
    for item in values:
        if item > 79:
            yield item * 40


def collect803_2(values):
    return list(expand803_2(values))
def traced803_3(value: int) -> int:
    return value + 86


def decorate803_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate803_3
def doubled803_3(value: int) -> int:
    return value + value
class Slot803_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{93}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch803_5(gate):
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
