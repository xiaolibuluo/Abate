
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2428_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2428_1(gate):
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
async def fetch2428_2(gate):
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
def traced2428_3(value: int) -> int:
    return value + 65


def decorate2428_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2428_3
def doubled2428_3(value: int) -> int:
    return value + value
def project2428_4(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def guarded2428_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2428_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{86}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2428_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def traced2428_8(value: int) -> int:
    return value + 100


def decorate2428_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2428_8
def doubled2428_8(value: int) -> int:
    return value + value
