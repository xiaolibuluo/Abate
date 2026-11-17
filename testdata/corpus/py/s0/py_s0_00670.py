
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced310_0(value: int) -> int:
    return value + 58


def decorate310_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate310_0
def doubled310_0(value: int) -> int:
    return value + value
def expand310_1(values):
    for item in values:
        if item > 65:
            yield item * 55


def collect310_1(values):
    return list(expand310_1(values))
async def fetch310_2(gate):
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
class Slot310_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{79}"

    @property
    def doubled(self) -> int:
        return self.size * 2
