package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter26120 {
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
sealed class Outcome26121 {
    data class Success(val payload: String) : Outcome26121()
    data object Failure : Outcome26121()
}

fun describe26121(outcome: Outcome26121): String = when (outcome) {
    is Outcome26121.Success -> outcome.payload
    Outcome26121.Failure -> "failed"
}
data class Pair26122<A, B>(val first: A, val second: B)

fun combine26122(left: Int, right: Int): Pair26122<Int, Int> {
    return Pair26122(left + 25, right + 15)
}
data class Pair26123<A, B>(val first: A, val second: B)

fun combine26123(left: Int, right: Int): Pair26123<Int, Int> {
    return Pair26123(left + 32, right + 26)
}
sealed class Outcome26124 {
    data class Success(val payload: String) : Outcome26124()
    data object Failure : Outcome26124()
}

fun describe26124(outcome: Outcome26124): String = when (outcome) {
    is Outcome26124.Success -> outcome.payload
    Outcome26124.Failure -> "failed"
}
class Counter26125 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
