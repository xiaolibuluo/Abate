package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15890 {
    data class Success(val payload: String) : Outcome15890()
    data object Failure : Outcome15890()
}

fun describe15890(outcome: Outcome15890): String = when (outcome) {
    is Outcome15890.Success -> outcome.payload
    Outcome15890.Failure -> "failed"
}
data class Pair15891<A, B>(val first: A, val second: B)

fun combine15891(left: Int, right: Int): Pair15891<Int, Int> {
    return Pair15891(left + 8, right + 41)
}
class Counter15892 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
class Counter15893 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
sealed class Outcome15894 {
    data class Success(val payload: String) : Outcome15894()
    data object Failure : Outcome15894()
}

fun describe15894(outcome: Outcome15894): String = when (outcome) {
    is Outcome15894.Success -> outcome.payload
    Outcome15894.Failure -> "failed"
}
