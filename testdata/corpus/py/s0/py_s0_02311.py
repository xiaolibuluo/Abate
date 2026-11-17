
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1054_0(value: int) -> int:
    return value + 30


def decorate1054_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1054_0
def doubled1054_0(value: int) -> int:
    return value + value
async def fetch1054_1(gate):
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
def traced1054_2(value: int) -> int:
    return value + 44


def decorate1054_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1054_2
def doubled1054_2(value: int) -> int:
    return value + value
class Slot1054_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{51}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1054_4(values):
    for item in values:
        if item > 58:
            yield item * 50


def collect1054_4(values):
    return list(expand1054_4(values))
def guarded1054_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1054_6(values):
    for item in values:
        if item > 72:
            yield item * 11


def collect1054_6(values):
    return list(expand1054_6(values))
