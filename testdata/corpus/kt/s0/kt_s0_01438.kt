package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11850 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
sealed class Outcome11851 {
    data class Success(val payload: String) : Outcome11851()
    data object Failure : Outcome11851()
}

fun describe11851(outcome: Outcome11851): String = when (outcome) {
    is Outcome11851.Success -> outcome.payload
    Outcome11851.Failure -> "failed"
}
@JvmName("step11852")
fun step11852(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
