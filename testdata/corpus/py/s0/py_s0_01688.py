
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot773_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{63}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch773_1(gate):
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
async def fetch773_2(gate):
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
def expand773_3(values):
    for item in values:
        if item > 84:
            yield item * 24


def collect773_3(values):
    return list(expand773_3(values))
def project773_4(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def traced773_5(value: int) -> int:
    return value + 98


def decorate773_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate773_5
def doubled773_5(value: int) -> int:
    return value + value
def guarded773_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot773_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{15}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded773_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
