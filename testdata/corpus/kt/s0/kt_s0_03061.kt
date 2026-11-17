package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25110")
fun step25110(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
class Counter25111 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
class Counter25112 {
    var current: Int by Delegates.observable(70) { _, _, new -> println(new) }

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
data class Pair25113<A, B>(val first: A, val second: B)

fun combine25113(left: Int, right: Int): Pair25113<Int, Int> {
    return Pair25113(left + 77, right + 51)
}
sealed class Outcome25114 {
    data class Success(val payload: String) : Outcome25114()
    data object Failure : Outcome25114()
}

fun describe25114(outcome: Outcome25114): String = when (outcome) {
    is Outcome25114.Success -> outcome.payload
    Outcome25114.Failure -> "failed"
}
