package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14310 {
    data class Success(val payload: String) : Outcome14310()
    data object Failure : Outcome14310()
}

fun describe14310(outcome: Outcome14310): String = when (outcome) {
    is Outcome14310.Success -> outcome.payload
    Outcome14310.Failure -> "failed"
}
class Counter14311 {
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
@JvmName("step14312")
fun step14312(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
data class Pair14313<A, B>(val first: A, val second: B)

fun combine14313(left: Int, right: Int): Pair14313<Int, Int> {
    return Pair14313(left + 5, right + 55)
}
