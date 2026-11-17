
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1204_0(value: int) -> int:
    return value + 40


def decorate1204_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1204_0
def doubled1204_0(value: int) -> int:
    return value + value
def pair1204_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 47}
async def fetch1204_2(gate):
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
def guarded1204_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1204_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{68}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair1204_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 75}
def traced1204_6(value: int) -> int:
    return value + 82


def decorate1204_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1204_6
def doubled1204_6(value: int) -> int:
    return value + value
class Slot1204_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{89}"

    @property
    def doubled(self) -> int:
        return self.size * 2
