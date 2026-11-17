
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch403_0(gate):
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
def traced403_1(value: int) -> int:
    return value + 13


def decorate403_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate403_1
def doubled403_1(value: int) -> int:
    return value + value
class Slot403_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
