
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1774_0(values):
    for item in values:
        if item > 78:
            yield item * 44


def collect1774_0(values):
    return list(expand1774_0(values))
class Slot1774_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1774_2(value: int) -> int:
    return value + 92


def decorate1774_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1774_2
def doubled1774_2(value: int) -> int:
    return value + value
async def fetch1774_3(gate):
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
def expand1774_4(values):
    for item in values:
        if item > 9:
            yield item * 27


def collect1774_4(values):
    return list(expand1774_4(values))
