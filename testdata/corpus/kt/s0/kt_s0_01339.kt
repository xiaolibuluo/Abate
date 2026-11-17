package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11050 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
object Registry11051 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11051 = "registry-21"
    }
}
data class Pair11052<A, B>(val first: A, val second: B)

fun combine11052(left: Int, right: Int): Pair11052<Int, Int> {
    return Pair11052(left + 28, right + 19)
}
class Counter11053 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
object Registry11054 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11054 = "registry-42"
    }
}
sealed class Outcome11055 {
    data class Success(val payload: String) : Outcome11055()
    data object Failure : Outcome11055()
}

fun describe11055(outcome: Outcome11055): String = when (outcome) {
    is Outcome11055.Success -> outcome.payload
    Outcome11055.Failure -> "failed"
}
