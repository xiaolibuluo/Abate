
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot1700_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{86}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1700_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1700_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1700_3(values):
    for item in values:
        if item > 10:
            yield item * 47


def collect1700_3(values):
    return list(expand1700_3(values))
def guarded1700_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1700_5(gate):
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
def traced1700_6(value: int) -> int:
    return value + 31


def decorate1700_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1700_6
def doubled1700_6(value: int) -> int:
    return value + value
def project1700_7(values):
    select = lambda item: item + 38
    return [select(item) for item in values if item != 0]
class Slot1700_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{45}"

    @property
    def doubled(self) -> int:
        return self.size * 2
