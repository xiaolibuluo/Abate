
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1746_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1746_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 12:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1746_2(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def expand1746_3(values):
    for item in values:
        if item > 26:
            yield item * 64


def collect1746_3(values):
    return list(expand1746_3(values))
def guarded1746_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1746_5(value: int) -> int:
    return value + 40


def decorate1746_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1746_5
def doubled1746_5(value: int) -> int:
    return value + value
async def fetch1746_6(gate):
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
class Slot1746_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{54}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1746_8(values):
    for item in values:
        if item > 61:
            yield item * 58


def collect1746_8(values):
    return list(expand1746_8(values))
