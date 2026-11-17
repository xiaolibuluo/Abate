package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5580 = "registry-81"
    }
}
data class Pair5581<A, B>(val first: A, val second: B)

fun combine5581(left: Int, right: Int): Pair5581<Int, Int> {
    return Pair5581(left + 88, right + 22)
}
sealed class Outcome5582 {
    data class Success(val payload: String) : Outcome5582()
    data object Failure : Outcome5582()
}

fun describe5582(outcome: Outcome5582): String = when (outcome) {
    is Outcome5582.Success -> outcome.payload
    Outcome5582.Failure -> "failed"
}
class Counter5583 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
