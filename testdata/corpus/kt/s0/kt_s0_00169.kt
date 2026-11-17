package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked1380(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome1381 {
    data class Success(val payload: String) : Outcome1381()
    data object Failure : Outcome1381()
}

fun describe1381(outcome: Outcome1381): String = when (outcome) {
    is Outcome1381.Success -> outcome.payload
    Outcome1381.Failure -> "failed"
}
sealed class Outcome1382 {
    data class Success(val payload: String) : Outcome1382()
    data object Failure : Outcome1382()
}

fun describe1382(outcome: Outcome1382): String = when (outcome) {
    is Outcome1382.Success -> outcome.payload
    Outcome1382.Failure -> "failed"
}
@JvmName("step1383")
fun step1383(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
@JvmName("step1384")
fun step1384(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
class Counter1385 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
class Counter1386 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
