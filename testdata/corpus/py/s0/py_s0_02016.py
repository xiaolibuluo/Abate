
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded918_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch918_1(gate):
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
def guarded918_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize918_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 29:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair918_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 36}
class Slot918_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project918_6(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def guarded918_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced918_8(value: int) -> int:
    return value + 64


def decorate918_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate918_8
def doubled918_8(value: int) -> int:
    return value + value
