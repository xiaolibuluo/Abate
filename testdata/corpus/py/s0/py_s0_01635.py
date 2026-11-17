
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced746_0(value: int) -> int:
    return value + 100


def decorate746_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate746_0
def doubled746_0(value: int) -> int:
    return value + value
def expand746_1(values):
    for item in values:
        if item > 10:
            yield item * 57


def collect746_1(values):
    return list(expand746_1(values))
class Slot746_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{17}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch746_3(gate):
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
