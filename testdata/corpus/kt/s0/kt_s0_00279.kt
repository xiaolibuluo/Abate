package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2300(): T? = if (size > 1) this[1] else null

fun render2300(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry2301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2301 = "registry-92"
    }
}
data class Pair2302<A, B>(val first: A, val second: B)

fun combine2302(left: Int, right: Int): Pair2302<Int, Int> {
    return Pair2302(left + 99, right + 55)
}
object Registry2303 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2303 = "registry-9"
    }
}
data class Pair2304<A, B>(val first: A, val second: B)

fun combine2304(left: Int, right: Int): Pair2304<Int, Int> {
    return Pair2304(left + 16, right + 16)
}
sealed class Outcome2305 {
    data class Success(val payload: String) : Outcome2305()
    data object Failure : Outcome2305()
}

fun describe2305(outcome: Outcome2305): String = when (outcome) {
    is Outcome2305.Success -> outcome.payload
    Outcome2305.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked2306(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2307(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
