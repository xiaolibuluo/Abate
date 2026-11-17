
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2480_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2480_1(value: int) -> int:
    return value + 48


def decorate2480_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2480_1
def doubled2480_1(value: int) -> int:
    return value + value
async def fetch2480_2(gate):
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
def expand2480_3(values):
    for item in values:
        if item > 62:
            yield item * 17


def collect2480_3(values):
    return list(expand2480_3(values))
def pair2480_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 69}
async def fetch2480_5(gate):
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
def expand2480_6(values):
    for item in values:
        if item > 83:
            yield item * 50


def collect2480_6(values):
    return list(expand2480_6(values))
def traced2480_7(value: int) -> int:
    return value + 90


def decorate2480_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2480_7
def doubled2480_7(value: int) -> int:
    return value + value
def guarded2480_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2480_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
