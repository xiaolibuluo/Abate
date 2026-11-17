package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5060 {
    data class Success(val payload: String) : Outcome5060()
    data object Failure : Outcome5060()
}

fun describe5060(outcome: Outcome5060): String = when (outcome) {
    is Outcome5060.Success -> outcome.payload
    Outcome5060.Failure -> "failed"
}
class Counter5061 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
data class Pair5062<A, B>(val first: A, val second: B)

fun combine5062(left: Int, right: Int): Pair5062<Int, Int> {
    return Pair5062(left + 98, right + 35)
}
@JvmName("step5063")
fun step5063(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
