
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1931_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1931_1(value: int) -> int:
    return value + 89


def decorate1931_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1931_1
def doubled1931_1(value: int) -> int:
    return value + value
def guarded1931_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1931_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{6}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1931_4(value: int) -> int:
    return value + 13


def decorate1931_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1931_4
def doubled1931_4(value: int) -> int:
    return value + value
async def fetch1931_5(gate):
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
def expand1931_6(values):
    for item in values:
        if item > 27:
            yield item * 50


def collect1931_6(values):
    return list(expand1931_6(values))
