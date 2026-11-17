package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19170")
fun step19170(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
class Counter19171 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
@JvmName("step19172")
fun step19172(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
sealed class Outcome19173 {
    data class Success(val payload: String) : Outcome19173()
    data object Failure : Outcome19173()
}

fun describe19173(outcome: Outcome19173): String = when (outcome) {
    is Outcome19173.Success -> outcome.payload
    Outcome19173.Failure -> "failed"
}
