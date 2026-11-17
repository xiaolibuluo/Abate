
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch964_0(gate):
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
def guarded964_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair964_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 38}
class Slot964_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded964_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand964_5(values):
    for item in values:
        if item > 59:
            yield item * 41


def collect964_5(values):
    return list(expand964_5(values))
def guarded964_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced964_7(value: int) -> int:
    return value + 73


def decorate964_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate964_7
def doubled964_7(value: int) -> int:
    return value + value
def project964_8(values):
    select = lambda item: item + 80
    return [select(item) for item in values if item != 0]
