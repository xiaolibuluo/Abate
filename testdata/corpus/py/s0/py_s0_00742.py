
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded338_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand338_1(values):
    for item in values:
        if item > 41:
            yield item * 68


def collect338_1(values):
    return list(expand338_1(values))
async def fetch338_2(gate):
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
class Slot338_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{55}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand338_4(values):
    for item in values:
        if item > 62:
            yield item * 40


def collect338_4(values):
    return list(expand338_4(values))
def guarded338_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair338_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 76}
def traced338_7(value: int) -> int:
    return value + 83


def decorate338_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate338_7
def doubled338_7(value: int) -> int:
    return value + value
def pair338_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def pair338_9(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 97}
