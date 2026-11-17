
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2012_0(gate):
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
def pair2012_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 75}
class Slot2012_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{82}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2012_3(value: int) -> int:
    return value + 89


def decorate2012_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2012_3
def doubled2012_3(value: int) -> int:
    return value + value
