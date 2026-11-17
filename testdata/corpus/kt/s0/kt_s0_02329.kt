package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19140 {
    data class Success(val payload: String) : Outcome19140()
    data object Failure : Outcome19140()
}

fun describe19140(outcome: Outcome19140): String = when (outcome) {
    is Outcome19140.Success -> outcome.payload
    Outcome19140.Failure -> "failed"
}
data class Pair19141<A, B>(val first: A, val second: B)

fun combine19141(left: Int, right: Int): Pair19141<Int, Int> {
    return Pair19141(left + 62, right + 59)
}
@JvmName("step19142")
fun step19142(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
class Counter19143 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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
