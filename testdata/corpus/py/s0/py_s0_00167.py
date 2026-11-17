
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand76_0(values):
    for item in values:
        if item > 23:
            yield item * 53


def collect76_0(values):
    return list(expand76_0(values))
def summarize76_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot76_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand76_3(values):
    for item in values:
        if item > 44:
            yield item * 25


def collect76_3(values):
    return list(expand76_3(values))
def guarded76_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project76_5(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
async def fetch76_6(gate):
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
