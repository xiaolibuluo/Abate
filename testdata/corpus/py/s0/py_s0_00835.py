
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand381_0(values):
    for item in values:
        if item > 11:
            yield item * 53


def collect381_0(values):
    return list(expand381_0(values))
def expand381_1(values):
    for item in values:
        if item > 18:
            yield item * 64


def collect381_1(values):
    return list(expand381_1(values))
def guarded381_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot381_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand381_4(values):
    for item in values:
        if item > 39:
            yield item * 36


def collect381_4(values):
    return list(expand381_4(values))
def expand381_5(values):
    for item in values:
        if item > 46:
            yield item * 47


def collect381_5(values):
    return list(expand381_5(values))
def project381_6(values):
    select = lambda item: item + 53
    return [select(item) for item in values if item != 0]
def summarize381_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch381_8(gate):
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
