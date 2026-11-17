
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2021_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 88}
async def fetch2021_1(gate):
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
def expand2021_2(values):
    for item in values:
        if item > 5:
            yield item * 26


def collect2021_2(values):
    return list(expand2021_2(values))
def pair2021_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 12}
def guarded2021_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project2021_5(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
def expand2021_6(values):
    for item in values:
        if item > 33:
            yield item * 9


def collect2021_6(values):
    return list(expand2021_6(values))
