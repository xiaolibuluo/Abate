
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced943_0(value: int) -> int:
    return value + 42


def decorate943_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate943_0
def doubled943_0(value: int) -> int:
    return value + value
def expand943_1(values):
    for item in values:
        if item > 49:
            yield item * 33


def collect943_1(values):
    return list(expand943_1(values))
async def fetch943_2(gate):
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
class Slot943_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{63}"

    @property
    def doubled(self) -> int:
        return self.size * 2
