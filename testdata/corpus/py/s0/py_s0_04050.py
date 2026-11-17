
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1880_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1880_1(gate):
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
class Slot1880_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{15}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1880_3(value: int) -> int:
    return value + 22


def decorate1880_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1880_3
def doubled1880_3(value: int) -> int:
    return value + value
def guarded1880_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1880_5(values):
    for item in values:
        if item > 36:
            yield item * 48


def collect1880_5(values):
    return list(expand1880_5(values))
