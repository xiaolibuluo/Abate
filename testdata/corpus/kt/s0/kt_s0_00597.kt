package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4850 {
    data class Success(val payload: String) : Outcome4850()
    data object Failure : Outcome4850()
}

fun describe4850(outcome: Outcome4850): String = when (outcome) {
    is Outcome4850.Success -> outcome.payload
    Outcome4850.Failure -> "failed"
}
class Counter4851 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
object Registry4852 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4852 = "registry-19"
    }
}
fun <T, R> List<T>.mapChecked4853(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull4854(): T? = if (size > 1) this[1] else null

fun render4854(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4855 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4855 = "registry-40"
    }
}
data class Pair4856<A, B>(val first: A, val second: B)

fun combine4856(left: Int, right: Int): Pair4856<Int, Int> {
    return Pair4856(left + 47, right + 54)
}
class Counter4857 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
