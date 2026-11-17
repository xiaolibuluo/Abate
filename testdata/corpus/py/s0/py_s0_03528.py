
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1627_0(gate):
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
def summarize1627_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1627_2(values):
    for item in values:
        if item > 24:
            yield item * 13


def collect1627_2(values):
    return list(expand1627_2(values))
def traced1627_3(value: int) -> int:
    return value + 31


def decorate1627_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1627_3
def doubled1627_3(value: int) -> int:
    return value + value
def guarded1627_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
