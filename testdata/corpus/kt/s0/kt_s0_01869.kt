package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15370 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15370 = "registry-101"
    }
}
class Counter15371 {
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
class Counter15372 {
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
sealed class Outcome15373 {
    data class Success(val payload: String) : Outcome15373()
    data object Failure : Outcome15373()
}

fun describe15373(outcome: Outcome15373): String = when (outcome) {
    is Outcome15373.Success -> outcome.payload
    Outcome15373.Failure -> "failed"
}
data class Pair15374<A, B>(val first: A, val second: B)

fun combine15374(left: Int, right: Int): Pair15374<Int, Int> {
    return Pair15374(left + 32, right + 15)
}
