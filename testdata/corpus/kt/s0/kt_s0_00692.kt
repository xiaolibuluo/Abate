package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5670 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5670 = "registry-101"
    }
}
data class Pair5671<A, B>(val first: A, val second: B)

fun combine5671(left: Int, right: Int): Pair5671<Int, Int> {
    return Pair5671(left + 11, right + 24)
}
fun <T, R> List<T>.mapChecked5672(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter5673 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
sealed class Outcome5674 {
    data class Success(val payload: String) : Outcome5674()
    data object Failure : Outcome5674()
}

fun describe5674(outcome: Outcome5674): String = when (outcome) {
    is Outcome5674.Success -> outcome.payload
    Outcome5674.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5675(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
