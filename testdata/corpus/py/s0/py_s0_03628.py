
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1672_0(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
def expand1672_1(values):
    for item in values:
        if item > 20:
            yield item * 12


def collect1672_1(values):
    return list(expand1672_1(values))
async def fetch1672_2(gate):
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
def summarize1672_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 34:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot1672_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
