
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1602_0(values):
    select = lambda item: item + 73
    return [select(item) for item in values if item != 0]
def traced1602_1(value: int) -> int:
    return value + 80


def decorate1602_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1602_1
def doubled1602_1(value: int) -> int:
    return value + value
def expand1602_2(values):
    for item in values:
        if item > 87:
            yield item * 21


def collect1602_2(values):
    return list(expand1602_2(values))
def guarded1602_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1602_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{101}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1602_5(value: int) -> int:
    return value + 11


def decorate1602_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1602_5
def doubled1602_5(value: int) -> int:
    return value + value
async def fetch1602_6(gate):
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
def traced1602_7(value: int) -> int:
    return value + 25


def decorate1602_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1602_7
def doubled1602_7(value: int) -> int:
    return value + value
async def fetch1602_8(gate):
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
