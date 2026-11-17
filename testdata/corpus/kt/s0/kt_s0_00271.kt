package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2220<A, B>(val first: A, val second: B)

fun combine2220(left: Int, right: Int): Pair2220<Int, Int> {
    return Pair2220(left + 78, right + 38)
}
@JvmName("step2221")
fun step2221(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
sealed class Outcome2222 {
    data class Success(val payload: String) : Outcome2222()
    data object Failure : Outcome2222()
}

fun describe2222(outcome: Outcome2222): String = when (outcome) {
    is Outcome2222.Success -> outcome.payload
    Outcome2222.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked2223(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter2224 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
