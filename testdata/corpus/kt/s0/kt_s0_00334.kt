package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter2770 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
data class Pair2771<A, B>(val first: A, val second: B)

fun combine2771(left: Int, right: Int): Pair2771<Int, Int> {
    return Pair2771(left + 24, right + 68)
}
@JvmName("step2772")
fun step2772(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T, R> List<T>.mapChecked2773(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step2774")
fun step2774(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
sealed class Outcome2775 {
    data class Success(val payload: String) : Outcome2775()
    data object Failure : Outcome2775()
}

fun describe2775(outcome: Outcome2775): String = when (outcome) {
    is Outcome2775.Success -> outcome.payload
    Outcome2775.Failure -> "failed"
}
@JvmName("step2776")
fun step2776(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
