
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced599_0(value: int) -> int:
    return value + 32


def decorate599_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate599_0
def doubled599_0(value: int) -> int:
    return value + value
def traced599_1(value: int) -> int:
    return value + 39


def decorate599_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate599_1
def doubled599_1(value: int) -> int:
    return value + value
async def fetch599_2(gate):
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
class Slot599_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{53}"

    @property
    def doubled(self) -> int:
        return self.size * 2
