
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced465_0(value: int) -> int:
    return value + 36


def decorate465_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate465_0
def doubled465_0(value: int) -> int:
    return value + value
async def fetch465_1(gate):
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
class Slot465_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{50}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair465_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 57}
async def fetch465_4(gate):
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
def traced465_5(value: int) -> int:
    return value + 71


def decorate465_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate465_5
def doubled465_5(value: int) -> int:
    return value + value
