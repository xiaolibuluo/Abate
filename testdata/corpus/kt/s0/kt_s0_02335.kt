package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19180<A, B>(val first: A, val second: B)

fun combine19180(left: Int, right: Int): Pair19180<Int, Int> {
    return Pair19180(left + 10, right + 15)
}
class Counter19181 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
class Counter19182 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
sealed class Outcome19183 {
    data class Success(val payload: String) : Outcome19183()
    data object Failure : Outcome19183()
}

fun describe19183(outcome: Outcome19183): String = when (outcome) {
    is Outcome19183.Success -> outcome.payload
    Outcome19183.Failure -> "failed"
}
