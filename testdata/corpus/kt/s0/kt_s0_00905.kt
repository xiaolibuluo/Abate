package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7420")
fun step7420(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome7421 {
    data class Success(val payload: String) : Outcome7421()
    data object Failure : Outcome7421()
}

fun describe7421(outcome: Outcome7421): String = when (outcome) {
    is Outcome7421.Success -> outcome.payload
    Outcome7421.Failure -> "failed"
}
class Counter7422 {
    var current: Int by Delegates.observable(62) { _, _, new -> println(new) }

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
