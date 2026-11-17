
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch890_0(gate):
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
def traced890_1(value: int) -> int:
    return value + 39


def decorate890_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate890_1
def doubled890_1(value: int) -> int:
    return value + value
class Slot890_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
