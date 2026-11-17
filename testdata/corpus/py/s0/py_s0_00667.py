
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch307_0(gate):
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
def summarize307_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 26:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair307_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 33}
class Slot307_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
