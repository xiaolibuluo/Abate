package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21140")
fun step21140(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
@JvmName("step21141")
fun step21141(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
class Counter21142 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
sealed class Outcome21143 {
    data class Success(val payload: String) : Outcome21143()
    data object Failure : Outcome21143()
}

fun describe21143(outcome: Outcome21143): String = when (outcome) {
    is Outcome21143.Success -> outcome.payload
    Outcome21143.Failure -> "failed"
}
