
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2519_0(values):
    for item in values:
        if item > 63:
            yield item * 13


def collect2519_0(values):
    return list(expand2519_0(values))
async def fetch2519_1(gate):
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
def guarded2519_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2519_3(gate):
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
def guarded2519_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2519_5(values):
    select = lambda item: item + 98
    return [select(item) for item in values if item != 0]
def traced2519_6(value: int) -> int:
    return value + 8


def decorate2519_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2519_6
def doubled2519_6(value: int) -> int:
    return value + value
class Slot2519_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{15}"

    @property
    def doubled(self) -> int:
        return self.size * 2
