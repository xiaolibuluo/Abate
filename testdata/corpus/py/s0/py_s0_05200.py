
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2420_0(values):
    for item in values:
        if item > 37:
            yield item * 52


def collect2420_0(values):
    return list(expand2420_0(values))
def expand2420_1(values):
    for item in values:
        if item > 44:
            yield item * 63


def collect2420_1(values):
    return list(expand2420_1(values))
def project2420_2(values):
    select = lambda item: item + 51
    return [select(item) for item in values if item != 0]
def expand2420_3(values):
    for item in values:
        if item > 58:
            yield item * 24


def collect2420_3(values):
    return list(expand2420_3(values))
class Slot2420_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{65}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def summarize2420_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2420_6(gate):
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
def guarded2420_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
