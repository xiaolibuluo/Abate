
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced454_0(value: int) -> int:
    return value + 87


def decorate454_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate454_0
def doubled454_0(value: int) -> int:
    return value + value
class Slot454_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{94}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot454_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{101}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch454_3(gate):
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
