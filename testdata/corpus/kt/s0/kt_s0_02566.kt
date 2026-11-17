package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21130 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
@JvmName("step21131")
fun step21131(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
@JvmName("step21132")
fun step21132(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
sealed class Outcome21133 {
    data class Success(val payload: String) : Outcome21133()
    data object Failure : Outcome21133()
}

fun describe21133(outcome: Outcome21133): String = when (outcome) {
    is Outcome21133.Success -> outcome.payload
    Outcome21133.Failure -> "failed"
}
