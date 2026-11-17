package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11470 {
    data class Success(val payload: String) : Outcome11470()
    data object Failure : Outcome11470()
}

fun describe11470(outcome: Outcome11470): String = when (outcome) {
    is Outcome11470.Success -> outcome.payload
    Outcome11470.Failure -> "failed"
}
class Counter11471 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked11472(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step11473")
fun step11473(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
data class Pair11474<A, B>(val first: A, val second: B)

fun combine11474(left: Int, right: Int): Pair11474<Int, Int> {
    return Pair11474(left + 6, right + 30)
}
class Counter11475 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
