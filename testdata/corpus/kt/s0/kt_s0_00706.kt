package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5810 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5810 = "registry-89"
    }
}
class Counter5811 {
    var current: Int by Delegates.observable(96) { _, _, new -> println(new) }

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
sealed class Outcome5812 {
    data class Success(val payload: String) : Outcome5812()
    data object Failure : Outcome5812()
}

fun describe5812(outcome: Outcome5812): String = when (outcome) {
    is Outcome5812.Success -> outcome.payload
    Outcome5812.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5813(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked5814(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair5815<A, B>(val first: A, val second: B)

fun combine5815(left: Int, right: Int): Pair5815<Int, Int> {
    return Pair5815(left + 27, right + 44)
}
fun <T> List<T>.secondOrNull5816(): T? = if (size > 1) this[1] else null

fun render5816(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5817 {
    data class Success(val payload: String) : Outcome5817()
    data object Failure : Outcome5817()
}

fun describe5817(outcome: Outcome5817): String = when (outcome) {
    is Outcome5817.Success -> outcome.payload
    Outcome5817.Failure -> "failed"
}
