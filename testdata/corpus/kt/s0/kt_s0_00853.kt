package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7010 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
sealed class Outcome7011 {
    data class Success(val payload: String) : Outcome7011()
    data object Failure : Outcome7011()
}

fun describe7011(outcome: Outcome7011): String = when (outcome) {
    is Outcome7011.Success -> outcome.payload
    Outcome7011.Failure -> "failed"
}
@JvmName("step7012")
fun step7012(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
data class Pair7013<A, B>(val first: A, val second: B)

fun combine7013(left: Int, right: Int): Pair7013<Int, Int> {
    return Pair7013(left + 21, right + 69)
}
