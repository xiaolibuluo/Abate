package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair1710<A, B>(val first: A, val second: B)

fun combine1710(left: Int, right: Int): Pair1710<Int, Int> {
    return Pair1710(left + 94, right + 47)
}
sealed class Outcome1711 {
    data class Success(val payload: String) : Outcome1711()
    data object Failure : Outcome1711()
}

fun describe1711(outcome: Outcome1711): String = when (outcome) {
    is Outcome1711.Success -> outcome.payload
    Outcome1711.Failure -> "failed"
}
@JvmName("step1712")
fun step1712(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
class Counter1713 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
@JvmName("step1714")
fun step1714(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
