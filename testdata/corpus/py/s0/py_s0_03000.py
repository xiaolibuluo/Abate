
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1373_0(gate):
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
class Slot1373_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{13}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1373_2(value: int) -> int:
    return value + 20


def decorate1373_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1373_2
def doubled1373_2(value: int) -> int:
    return value + value
