package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step120")
fun step120(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
sealed class Outcome121 {
    data class Success(val payload: String) : Outcome121()
    data object Failure : Outcome121()
}

fun describe121(outcome: Outcome121): String = when (outcome) {
    is Outcome121.Success -> outcome.payload
    Outcome121.Failure -> "failed"
}
class Counter122 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
