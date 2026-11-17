
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot513_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{78}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch513_1(gate):
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
async def fetch513_2(gate):
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
def pair513_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 99}
def traced513_4(value: int) -> int:
    return value + 9


def decorate513_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate513_4
def doubled513_4(value: int) -> int:
    return value + value
def traced513_5(value: int) -> int:
    return value + 16


def decorate513_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate513_5
def doubled513_5(value: int) -> int:
    return value + value
