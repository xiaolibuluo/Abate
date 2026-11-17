package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15500 {
    data class Success(val payload: String) : Outcome15500()
    data object Failure : Outcome15500()
}

fun describe15500(outcome: Outcome15500): String = when (outcome) {
    is Outcome15500.Success -> outcome.payload
    Outcome15500.Failure -> "failed"
}
sealed class Outcome15501 {
    data class Success(val payload: String) : Outcome15501()
    data object Failure : Outcome15501()
}

fun describe15501(outcome: Outcome15501): String = when (outcome) {
    is Outcome15501.Success -> outcome.payload
    Outcome15501.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15502(): T? = if (size > 1) this[1] else null

fun render15502(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry15503 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15503 = "registry-97"
    }
}
data class Pair15504<A, B>(val first: A, val second: B)

fun combine15504(left: Int, right: Int): Pair15504<Int, Int> {
    return Pair15504(left + 7, right + 45)
}
class Counter15505 {
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
fun <T, R> List<T>.mapChecked15506(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
