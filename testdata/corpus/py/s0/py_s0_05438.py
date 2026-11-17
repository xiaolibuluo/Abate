
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2511_0(gate):
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
def summarize2511_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2511_2(values):
    select = lambda item: item + 70
    return [select(item) for item in values if item != 0]
def traced2511_3(value: int) -> int:
    return value + 77


def decorate2511_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2511_3
def doubled2511_3(value: int) -> int:
    return value + value
def expand2511_4(values):
    for item in values:
        if item > 84:
            yield item * 62


def collect2511_4(values):
    return list(expand2511_4(values))
def project2511_5(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def guarded2511_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
