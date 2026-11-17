
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2321_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 11:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2321_1(gate):
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
def pair2321_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 25}
class Slot2321_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2321_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize2321_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2321_6(values):
    for item in values:
        if item > 53:
            yield item * 35


def collect2321_6(values):
    return list(expand2321_6(values))
