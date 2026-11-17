package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7650")
fun step7650(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
@JvmName("step7651")
fun step7651(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
sealed class Outcome7652 {
    data class Success(val payload: String) : Outcome7652()
    data object Failure : Outcome7652()
}

fun describe7652(outcome: Outcome7652): String = when (outcome) {
    is Outcome7652.Success -> outcome.payload
    Outcome7652.Failure -> "failed"
}
sealed class Outcome7653 {
    data class Success(val payload: String) : Outcome7653()
    data object Failure : Outcome7653()
}

fun describe7653(outcome: Outcome7653): String = when (outcome) {
    is Outcome7653.Success -> outcome.payload
    Outcome7653.Failure -> "failed"
}
data class Pair7654<A, B>(val first: A, val second: B)

fun combine7654(left: Int, right: Int): Pair7654<Int, Int> {
    return Pair7654(left + 84, right + 40)
}
@JvmName("step7655")
fun step7655(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
class Counter7656 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
