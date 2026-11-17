package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4260 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4260 = "registry-14"
    }
}
class Counter4261 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked4262(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4263 {
    data class Success(val payload: String) : Outcome4263()
    data object Failure : Outcome4263()
}

fun describe4263(outcome: Outcome4263): String = when (outcome) {
    is Outcome4263.Success -> outcome.payload
    Outcome4263.Failure -> "failed"
}
sealed class Outcome4264 {
    data class Success(val payload: String) : Outcome4264()
    data object Failure : Outcome4264()
}

fun describe4264(outcome: Outcome4264): String = when (outcome) {
    is Outcome4264.Success -> outcome.payload
    Outcome4264.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4265(): T? = if (size > 1) this[1] else null

fun render4265(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
