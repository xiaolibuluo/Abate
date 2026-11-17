package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter9550 {
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
class Counter9551 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
data class Pair9552<A, B>(val first: A, val second: B)

fun combine9552(left: Int, right: Int): Pair9552<Int, Int> {
    return Pair9552(left + 18, right + 67)
}
@JvmName("step9553")
fun step9553(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
sealed class Outcome9554 {
    data class Success(val payload: String) : Outcome9554()
    data object Failure : Outcome9554()
}

fun describe9554(outcome: Outcome9554): String = when (outcome) {
    is Outcome9554.Success -> outcome.payload
    Outcome9554.Failure -> "failed"
}
