package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15960(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15961 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15961 = "registry-99"
    }
}
fun <T> List<T>.secondOrNull15962(): T? = if (size > 1) this[1] else null

fun render15962(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15963<A, B>(val first: A, val second: B)

fun combine15963(left: Int, right: Int): Pair15963<Int, Int> {
    return Pair15963(left + 16, right + 51)
}
class Counter15964 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
sealed class Outcome15965 {
    data class Success(val payload: String) : Outcome15965()
    data object Failure : Outcome15965()
}

fun describe15965(outcome: Outcome15965): String = when (outcome) {
    is Outcome15965.Success -> outcome.payload
    Outcome15965.Failure -> "failed"
}
class Counter15966 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
