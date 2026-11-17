
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2060_0(value: int) -> int:
    return value + 13


def decorate2060_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2060_0
def doubled2060_0(value: int) -> int:
    return value + value
async def fetch2060_1(gate):
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
async def fetch2060_2(gate):
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
class Slot2060_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
