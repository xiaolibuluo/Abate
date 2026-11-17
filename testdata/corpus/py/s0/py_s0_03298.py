
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1526_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{55}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1526_1(value: int) -> int:
    return value + 62


def decorate1526_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1526_1
def doubled1526_1(value: int) -> int:
    return value + value
def expand1526_2(values):
    for item in values:
        if item > 69:
            yield item * 38


def collect1526_2(values):
    return list(expand1526_2(values))
async def fetch1526_3(gate):
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
