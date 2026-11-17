
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1530_0(value: int) -> int:
    return value + 10


def decorate1530_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1530_0
def doubled1530_0(value: int) -> int:
    return value + value
async def fetch1530_1(gate):
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
def expand1530_2(values):
    for item in values:
        if item > 24:
            yield item * 66


def collect1530_2(values):
    return list(expand1530_2(values))
def traced1530_3(value: int) -> int:
    return value + 31


def decorate1530_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1530_3
def doubled1530_3(value: int) -> int:
    return value + value
class Slot1530_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{38}"

    @property
    def doubled(self) -> int:
        return self.size * 2
