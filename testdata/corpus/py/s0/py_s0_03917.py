
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1815_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{29}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1815_1(value: int) -> int:
    return value + 36


def decorate1815_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1815_1
def doubled1815_1(value: int) -> int:
    return value + value
def project1815_2(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
async def fetch1815_3(gate):
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
