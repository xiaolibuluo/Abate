
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2067_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 7:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2067_1(gate):
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
def traced2067_2(value: int) -> int:
    return value + 21


def decorate2067_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2067_2
def doubled2067_2(value: int) -> int:
    return value + value
def guarded2067_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2067_4(values):
    for item in values:
        if item > 35:
            yield item * 65


def collect2067_4(values):
    return list(expand2067_4(values))
