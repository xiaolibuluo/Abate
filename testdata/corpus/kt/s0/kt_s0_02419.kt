package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter19880 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
sealed class Outcome19881 {
    data class Success(val payload: String) : Outcome19881()
    data object Failure : Outcome19881()
}

fun describe19881(outcome: Outcome19881): String = when (outcome) {
    is Outcome19881.Success -> outcome.payload
    Outcome19881.Failure -> "failed"
}
sealed class Outcome19882 {
    data class Success(val payload: String) : Outcome19882()
    data object Failure : Outcome19882()
}

fun describe19882(outcome: Outcome19882): String = when (outcome) {
    is Outcome19882.Success -> outcome.payload
    Outcome19882.Failure -> "failed"
}
sealed class Outcome19883 {
    data class Success(val payload: String) : Outcome19883()
    data object Failure : Outcome19883()
}

fun describe19883(outcome: Outcome19883): String = when (outcome) {
    is Outcome19883.Success -> outcome.payload
    Outcome19883.Failure -> "failed"
}
object Registry19884 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19884 = "registry-75"
    }
}
data class Pair19885<A, B>(val first: A, val second: B)

fun combine19885(left: Int, right: Int): Pair19885<Int, Int> {
    return Pair19885(left + 82, right + 11)
}
class Counter19886 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
