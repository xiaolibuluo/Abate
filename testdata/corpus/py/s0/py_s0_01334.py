
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize604_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch604_1(gate):
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
class Slot604_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{14}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot604_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{21}"

    @property
    def doubled(self) -> int:
        return self.size * 2
