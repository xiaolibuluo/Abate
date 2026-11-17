package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter24570 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
@JvmName("step24571")
fun step24571(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
@JvmName("step24572")
fun step24572(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
@JvmName("step24573")
fun step24573(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
@JvmName("step24574")
fun step24574(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
sealed class Outcome24575 {
    data class Success(val payload: String) : Outcome24575()
    data object Failure : Outcome24575()
}

fun describe24575(outcome: Outcome24575): String = when (outcome) {
    is Outcome24575.Success -> outcome.payload
    Outcome24575.Failure -> "failed"
}
