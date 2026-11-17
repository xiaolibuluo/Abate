
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2409_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2409_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2409_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{5}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2409_3(gate):
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
def expand2409_4(values):
    for item in values:
        if item > 19:
            yield item * 19


def collect2409_4(values):
    return list(expand2409_4(values))
def summarize2409_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
