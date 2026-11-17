
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1537_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1537_1(gate):
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
async def fetch1537_2(gate):
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
def project1537_3(values):
    select = lambda item: item + 25
    return [select(item) for item in values if item != 0]
class Slot1537_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{32}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1537_5(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
class Slot1537_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{46}"

    @property
    def doubled(self) -> int:
        return self.size * 2
