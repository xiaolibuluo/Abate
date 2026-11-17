package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12320 {
    var current: Int by Delegates.observable(16) { _, _, new -> println(new) }

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
@JvmName("step12321")
fun step12321(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
sealed class Outcome12322 {
    data class Success(val payload: String) : Outcome12322()
    data object Failure : Outcome12322()
}

fun describe12322(outcome: Outcome12322): String = when (outcome) {
    is Outcome12322.Success -> outcome.payload
    Outcome12322.Failure -> "failed"
}
