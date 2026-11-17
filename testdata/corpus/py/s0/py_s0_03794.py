
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1750_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1750_1(value: int) -> int:
    return value + 64


def decorate1750_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1750_1
def doubled1750_1(value: int) -> int:
    return value + value
async def fetch1750_2(gate):
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
async def fetch1750_3(gate):
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
def guarded1750_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1750_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{92}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1750_6(value: int) -> int:
    return value + 99


def decorate1750_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1750_6
def doubled1750_6(value: int) -> int:
    return value + value
def pair1750_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 9}
def traced1750_8(value: int) -> int:
    return value + 16


def decorate1750_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1750_8
def doubled1750_8(value: int) -> int:
    return value + value
