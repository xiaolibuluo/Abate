
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2505_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2505_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{82}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2505_2(gate):
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
def expand2505_3(values):
    for item in values:
        if item > 96:
            yield item * 9


def collect2505_3(values):
    return list(expand2505_3(values))
def summarize2505_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 6:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
