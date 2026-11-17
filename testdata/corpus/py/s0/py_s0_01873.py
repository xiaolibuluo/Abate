
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot851_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{10}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch851_1(gate):
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
def traced851_2(value: int) -> int:
    return value + 24


def decorate851_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate851_2
def doubled851_2(value: int) -> int:
    return value + value
def traced851_3(value: int) -> int:
    return value + 31


def decorate851_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate851_3
def doubled851_3(value: int) -> int:
    return value + value
def traced851_4(value: int) -> int:
    return value + 38


def decorate851_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate851_4
def doubled851_4(value: int) -> int:
    return value + value
