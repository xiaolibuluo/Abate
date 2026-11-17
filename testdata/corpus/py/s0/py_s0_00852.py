
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project388_0(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]
def guarded388_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded388_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project388_3(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
async def fetch388_4(gate):
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
def expand388_5(values):
    for item in values:
        if item > 40:
            yield item * 35


def collect388_5(values):
    return list(expand388_5(values))
async def fetch388_6(gate):
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
class Slot388_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{54}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch388_8(gate):
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
def project388_9(values):
    select = lambda item: item + 68
    return [select(item) for item in values if item != 0]
async def fetch388_10(gate):
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
