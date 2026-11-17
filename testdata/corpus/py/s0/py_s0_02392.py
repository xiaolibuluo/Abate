
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1103_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1103_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 92:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1103_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 99}
def expand1103_3(values):
    for item in values:
        if item > 9:
            yield item * 16


def collect1103_3(values):
    return list(expand1103_3(values))
async def fetch1103_4(gate):
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
class Slot1103_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{23}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1103_6(values):
    for item in values:
        if item > 30:
            yield item * 49


def collect1103_6(values):
    return list(expand1103_6(values))
